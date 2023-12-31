// Code generated by goctl. DO NOT EDIT.
// Source: service.proto

package service

import (
	"context"

	"service/service_pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	EmptyResp                      = service_pb.EmptyResp
	FilterItem                     = service_pb.FilterItem
	LoginReq                       = service_pb.LoginReq
	LoginResp                      = service_pb.LoginResp
	Order                          = service_pb.Order
	OrderAddReq                    = service_pb.OrderAddReq
	OrderAddResp                   = service_pb.OrderAddResp
	OrderRemoveReq                 = service_pb.OrderRemoveReq
	OrderRetrieveReq               = service_pb.OrderRetrieveReq
	OrderRetrieveResp              = service_pb.OrderRetrieveResp
	OrderUpdateReq                 = service_pb.OrderUpdateReq
	Paging                         = service_pb.Paging
	Performance                    = service_pb.Performance
	PerformanceAddReq              = service_pb.PerformanceAddReq
	PerformanceAddResp             = service_pb.PerformanceAddResp
	PerformanceRemoveReq           = service_pb.PerformanceRemoveReq
	PerformanceRetrieveReq         = service_pb.PerformanceRetrieveReq
	PerformanceRetrieveResp        = service_pb.PerformanceRetrieveResp
	PerformanceSeat                = service_pb.PerformanceSeat
	PerformanceSeatAddReq          = service_pb.PerformanceSeatAddReq
	PerformanceSeatAddResp         = service_pb.PerformanceSeatAddResp
	PerformanceSeatRemoveReq       = service_pb.PerformanceSeatRemoveReq
	PerformanceSeatRetrieveReq     = service_pb.PerformanceSeatRetrieveReq
	PerformanceSeatRetrieveResp    = service_pb.PerformanceSeatRetrieveResp
	PerformanceSeatUpdateReq       = service_pb.PerformanceSeatUpdateReq
	PerformanceSession             = service_pb.PerformanceSession
	PerformanceSessionAddReq       = service_pb.PerformanceSessionAddReq
	PerformanceSessionAddResp      = service_pb.PerformanceSessionAddResp
	PerformanceSessionRemoveReq    = service_pb.PerformanceSessionRemoveReq
	PerformanceSessionRetrieveReq  = service_pb.PerformanceSessionRetrieveReq
	PerformanceSessionRetrieveResp = service_pb.PerformanceSessionRetrieveResp
	PerformanceSessionUpdateReq    = service_pb.PerformanceSessionUpdateReq
	PerformanceUpdateReq           = service_pb.PerformanceUpdateReq
	SortItem                       = service_pb.SortItem
	User                           = service_pb.User
	UserAddReq                     = service_pb.UserAddReq
	UserAddResp                    = service_pb.UserAddResp
	UserRemoveReq                  = service_pb.UserRemoveReq
	UserRetrieveReq                = service_pb.UserRetrieveReq
	UserRetrieveResp               = service_pb.UserRetrieveResp
	UserUpdateReq                  = service_pb.UserUpdateReq

	Service interface {
		// Login 登录
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	}

	defaultService struct {
		cli zrpc.Client
	}
)

func NewService(cli zrpc.Client) Service {
	return &defaultService{
		cli: cli,
	}
}

// Login 登录
func (m *defaultService) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := service_pb.NewServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
