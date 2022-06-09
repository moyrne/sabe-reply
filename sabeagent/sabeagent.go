package main

import (
	"flag"
	"fmt"

	"github.com/moyrne/sabe-reply/sabeagent/internal/config"
	"github.com/moyrne/sabe-reply/sabeagent/internal/server"
	"github.com/moyrne/sabe-reply/sabeagent/internal/svc"
	"github.com/moyrne/sabe-reply/sabeagent/sabeagent"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/sabeagent.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewSabeAgentServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sabeagent.RegisterSabeAgentServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
