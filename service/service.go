package main

import (
	"flag"
	"fmt"

	"service/internal/config"
	orderserviceServer "service/internal/server/orderservice"
	performanceseatserviceServer "service/internal/server/performanceseatservice"
	performanceserviceServer "service/internal/server/performanceservice"
	performancesessionserviceServer "service/internal/server/performancesessionservice"
	serviceServer "service/internal/server/service"
	userserviceServer "service/internal/server/userservice"
	"service/internal/svc"
	"service/service_pb"

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
		service_pb.RegisterPerformanceServiceServer(grpcServer, performanceserviceServer.NewPerformanceServiceServer(ctx))
		service_pb.RegisterOrderServiceServer(grpcServer, orderserviceServer.NewOrderServiceServer(ctx))
		service_pb.RegisterPerformanceSeatServiceServer(grpcServer, performanceseatserviceServer.NewPerformanceSeatServiceServer(ctx))
		service_pb.RegisterPerformanceSessionServiceServer(grpcServer, performancesessionserviceServer.NewPerformanceSessionServiceServer(ctx))
		service_pb.RegisterUserServiceServer(grpcServer, userserviceServer.NewUserServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
