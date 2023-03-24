package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"talkon_srvs/user_srv/model"
	"time"
)

func main() {
	dsn := "shijiu:Qq1633841065@tcp(127.0.0.1:8888)/talkon_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "t_",
		},
		Logger: newLogger,
	})
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	// 建表
	err = db.AutoMigrate(&model.User{})
}
