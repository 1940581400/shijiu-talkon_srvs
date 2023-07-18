package initialize

import (
	"go.uber.org/zap"
	"log"
	"strconv"
	"talkon_srvs/user_srv/config"
	"talkon_srvs/user_srv/global"
	"talkon_srvs/user_srv/utils"

	"github.com/spf13/viper"
)

func GetEnvInfo(key string) string {
	viper.AutomaticEnv()
	return viper.GetString(key)
}

func InitConfig() {
	info := GetEnvInfo("TALKON_DEV")
	configName := "./conf/config_pro.yaml"
	if flag, _ := strconv.ParseBool(info); flag {
		//configName = "user_srv/config_dev.yaml"
		configName = "./conf/config_dev.yaml"
	}
	logger := log.Default()
	logger.Printf("[配置文件] 读取中 %s", configName)
	v := viper.New()
	v.SetConfigFile(configName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	err = v.Unmarshal(&serverConfig)
	if err != nil {
		panic(err)
	}
	global.ServerConfig = &serverConfig
	logger.Printf("[配置文件] 读取完成")
	// 初始化consul配置中心
	addr := global.ServerConfig.ConsulConfig.Host + ":" + global.ServerConfig.ConsulConfig.Port
	global.ServerConfig.ConsulConfig.ConsulAddr = addr
	consulConfig, err := utils.GetConsulConfig(addr, global.ServerConfig.ConsulConfig.ConfigurationCenter.Prefix)
	if err != nil {
		logger.Panic("[consul] 配置中心获取失败", zap.Error(err))
	}
	global.ConfigSource = &consulConfig
}
