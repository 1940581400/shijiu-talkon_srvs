package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"talkon_srvs/user_srv/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() {
	zap.S().Infof("[数据库配置] 初始化")
	config := global.ServerConfig.MySQLInfo
	zap.S().Infof("[数据库配置] 初始化 :%s", global.ServerConfig.MySQLInfo)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Schema)
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
		zap.S().Infof("[数据库配置] 初始化完成")
	}()
}
