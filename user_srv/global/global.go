package global

import (
	"github.com/asim/go-micro/v3/config"
	sjConf "github.com/shijiu-xf/go-base/config"
	"gorm.io/gorm"
	myConf "talkon_srvs/user_srv/config"
)

var (
	// DB 全局DB
	DB *gorm.DB
	// ServerConfig 全局配置
	ServerConfig *myConf.ServerConfig
	// ConfigSource 配置中心
	ConfigSource *config.Config

	LogFileConfig *sjConf.LogFileConfig
)
