package global

import (
	"gorm.io/gorm"
	"talkon_srvs/user_srv/config"
)

var (
	// DB 全局DB
	DB *gorm.DB
	// ServerConfig 全局配置
	ServerConfig *config.ServerConfig
)
