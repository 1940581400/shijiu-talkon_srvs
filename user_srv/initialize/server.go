package initialize

import (
	"fmt"
	"github.com/shijiu-xf/go-base/comm/zapsj"
	"net"
	"talkon_srvs/user_srv/global"

	"google.golang.org/grpc"

	"talkon_srvs/user_srv/handler"
	"talkon_srvs/user_srv/proto"
)

func InitServer() {
	config := global.ServerConfig.UserSrvInfo
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserService{})
	zapsj.ZapS().Infof("[用户服务] 启动：地址:%s 端口:%s", config.Host, config.Port)
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		panic("failed to listen " + err.Error())
	}
	err = server.Serve(listen)
	if err != nil {
		panic("failed to start grpc " + err.Error())
	}
}
