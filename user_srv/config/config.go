package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Schema   string `mapstructure:"schema"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// LoggerConfig struct 的配置字段含义请参考 zap.Config
type LoggerConfig struct {
	Level            string `mapstructure:"level"`
	Development      string `mapstructure:"development"`
	Encoding         string `mapstructure:"encoding"`
	EncoderConfig    string `mapstructure:"encoderConfig"`
	OutputPaths      string `mapstructure:"outputPaths"`
	ErrorOutputPaths string `mapstructure:"errorOutputPaths"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	UserSrvInfo UserSrvConfig `mapstructure:"server"`
	MySQLInfo   MySQLConfig   `mapstructure:"mysql"`
	LoggerInfo  LoggerConfig  `mapstructure:"logger"`
}
