package initialize

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	microPlg "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"go.uber.org/zap"

	"talkon_srvs/user_srv/config"
	"talkon_srvs/user_srv/global"
	"talkon_srvs/user_srv/utils"
)

// InitConsul 注册服务
func InitConsul() {
	addr := global.ServerConfig.ConsulConfig.ConsulAddr
	zap.L().Info("[consul] 服务注册中 ", zap.String("地址", addr))
	newRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			addr,
		}
	})
	var jgr = &config.JaegerConfig{}
	err := jgr.GetJaegerConfig(*global.ConfigSource, "jaeger")
	zap.L().Info("[consul] 链路配置 ", zap.Reflect("jaeger", jgr))
	if err != nil {
		zap.L().Panic("[链路追踪] 配置获取失败", zap.Error(err))
	}
	globalTrancer, closer, err := utils.NewTracer(jgr.ServiceName, jgr.Addr)
	defer closer.Close()
	if err != nil {
		zap.L().Panic("[链路追踪] 创建链路失败", zap.Error(err))
	}
	microSrv := micro.NewService(
		micro.Name("user_srv"),
		micro.Version("1.0"),
		micro.Registry(newRegistry),
		micro.WrapHandler(microPlg.NewHandlerWrapper(globalTrancer)),
	)
	microSrv.Init()
	err = microSrv.Run()
	if err != nil {
		panic("failed to consul " + err.Error())
	}
}
