package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strconv"
	"strings"
	"talkon_srvs/user_srv/global"
	"talkon_srvs/user_srv/global/zero"
)

func NewMyConfig() zap.Config {
	config := global.ServerConfig.LoggerInfo
	var development bool
	var err error
	// 开发者模式默认开启
	if config.Development == zero.String {
		development = true
	}
	development, err = strconv.ParseBool(config.Development)
	if err != nil {
		panic(err)
	}
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(GetZapLogLevelCode(config.Level)),
		Development:      development,
		Encoding:         config.Encoding,
		EncoderConfig:    GetEncoderConfig(config.EncoderConfig),
		OutputPaths:      GetOutPaths(config.OutputPaths),
		ErrorOutputPaths: GetOutPaths(config.ErrorOutputPaths),
	}
}

// GetZapLogLevelCode 将配置文件中的日志级别转换为zap的码值，便于配置
func GetZapLogLevelCode(level string) zapcore.Level {
	// 默认为Debug级别
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "dpanic":
		zapLevel = zapcore.DPanicLevel
	case "panic":
		zapLevel = zapcore.PanicLevel
	case "fatal":
		zapLevel = zapcore.FatalLevel
	default:
		zapLevel = zapcore.DebugLevel
	}
	return zapLevel
}

// GetEncoderConfig 获取编码配置文件
func GetEncoderConfig(encoderConfig string) zapcore.EncoderConfig {
	if encoderConfig == "dev" {
		return zap.NewDevelopmentEncoderConfig()
	}
	if encoderConfig == "pro" {
		return zap.NewProductionEncoderConfig()
	}
	return zap.NewDevelopmentEncoderConfig()
}

func GetOutPaths(paths string) []string {
	outPaths := make([]string, 0)
	split := strings.Split(paths, ",")
	if paths == zero.String {
		outPaths = append(outPaths, "stderr")
		return outPaths
	}
	if len(split) == 0 {
		outPaths = append(outPaths, paths)
		return outPaths
	}
	for _, s := range split {
		outPaths = append(outPaths, s)
	}
	return outPaths
}

func InitLogger() {
	lg := log.Default()
	lg.Printf("[日志配置] 初始化")
	config := NewMyConfig()
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	lg.Printf("[日志配置] 初始化完成")
	zap.L().Sync()
}
