package main

import "talkon_srvs/user_srv/initialize"

func main() {

	// 1.初始化配置，注意 此操作一定要在第一，不然后面初始化读不到配置
	initialize.InitConfig()
	// 2.初始化日志，注意 此操作一定要在第二，不然初始化文件当中的 日志 无法打印
	initialize.InitLogger()

	// 下面顺序可随意
	// 3.初始化全局数据连接
	initialize.InitDB()
	// 4.初始化服务
	initialize.InitServer()
}
