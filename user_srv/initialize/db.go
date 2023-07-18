package initialize

import (
	"fmt"
	"github.com/shijiu-xf/go-base/comm/zapsj"
	"go.uber.org/zap"
	"log"
	"os"
	"talkon_srvs/user_srv/config"
	"talkon_srvs/user_srv/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() {
	zapsj.ZapS().Infof("[数据库配置] 初始化")
	var mysqlConfig = &config.MySQLConfig{}
	err := mysqlConfig.GetMysqlConfig(*global.ConfigSource, "mysql")
	if err != nil {
		zap.L().Panic("get config field :", zap.Error(err))
	}
	zapsj.ZapS().Infof("[数据库配置] 连接 地址:%s 端口:%s", mysqlConfig.Host, mysqlConfig.Port)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Schema)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "t_",
		},
		Logger: newLogger,
	})
	global.DB = DB
	defer func() {
		if err != nil {
			panic(err)
		}
		zapsj.ZapS().Infof("[数据库配置] 初始化完成")
	}()
}
