package svc

import (
	"gateway/internal/config"
	"github.com/JopenChen/zero-damai/service/client/service"
	"github.com/JopenChen/zero-damai/service/client/testservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	ServiceRpc     service.Service
	TestServiceRpc testservice.TestService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		ServiceRpc:     service.NewService(zrpc.MustNewClient(c.ServiceRpc)),
		TestServiceRpc: testservice.NewTestService(zrpc.MustNewClient(c.ServiceRpc)),
	}
}
