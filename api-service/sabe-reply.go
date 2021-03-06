package main

import (
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/moyrne/sabe-reply/api-service/internal/config"
	"github.com/moyrne/sabe-reply/api-service/internal/handler"
	"github.com/moyrne/sabe-reply/api-service/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/sabe-reply-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
