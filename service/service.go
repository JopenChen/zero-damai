package main

import (
	"flag"
	"fmt"

	"github.com/JopenChen/zero-damai/service/internal/config"
	serviceServer "github.com/JopenChen/zero-damai/service/internal/server/service"
	testserviceServer "github.com/JopenChen/zero-damai/service/internal/server/testservice"
	"github.com/JopenChen/zero-damai/service/internal/svc"
	"github.com/JopenChen/zero-damai/service/service_pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/service.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		service_pb.RegisterServiceServer(grpcServer, serviceServer.NewServiceServer(ctx))
		service_pb.RegisterTestServiceServer(grpcServer, testserviceServer.NewTestServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
