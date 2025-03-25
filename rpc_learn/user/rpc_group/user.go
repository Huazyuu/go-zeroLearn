package main

import (
	"flag"
	"fmt"

	"gozero_learn/rpc_learn/user/rpc_group/internal/config"
	usercreateServer "gozero_learn/rpc_learn/user/rpc_group/internal/server/usercreate"
	userinfoServer "gozero_learn/rpc_learn/user/rpc_group/internal/server/userinfo"
	"gozero_learn/rpc_learn/user/rpc_group/internal/svc"
	"gozero_learn/rpc_learn/user/rpc_group/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserCreateServer(grpcServer, usercreateServer.NewUserCreateServer(ctx))
		user.RegisterUserInfoServer(grpcServer, userinfoServer.NewUserInfoServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
