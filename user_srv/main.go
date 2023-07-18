package main

import (
	"github.com/shijiu-xf/go-base/comm/prometheussj"
	"github.com/shijiu-xf/go-base/comm/zapsj"
	"github.com/shijiu-xf/go-base/config"
	"log"
	"talkon_srvs/user_srv/global"
	"talkon_srvs/user_srv/initialize"
)

func main() {
	lg := log.Default()
	// 1.初始化配置，注意 此操作一定要在第一，不然后面初始化读不到配置
	initialize.InitConfig()
	// 2.初始化日志，注意 此操作一定要在第二，不然初始化文件当中的 日志 无法打印
	//initialize.InitLogger()
	logFileCfg := &config.LogFileConfig{}
	err := logFileCfg.GetLogFileConfig(*global.ConfigSource, "logfile")
	if err != nil {
		lg.Panicf("[GetLogFileConfig] 日志配置获取失败%s", err.Error())
	}
	err = zapsj.InitSJZap(*logFileCfg)
	if err != nil {
		lg.Panicf("[InitSJZap] 日志初始化失败%s", err.Error())
	}
	zapsj.ZapL().Info("日志初始化成功")
	defer zapsj.ZapL().Sync()
	// 下面顺序可随意
	// 3.初始化全局数据连接
	initialize.InitDB()

	// 4.注册服务
	go initialize.InitConsul()

	// 添加监控
	prometheussj.Monitoring(9015)

	// 初始化服务
	initialize.InitServer()
}
