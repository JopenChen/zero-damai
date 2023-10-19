package testservicelogic

import (
	"context"

	"github.com/JopenChen/zero-damai/service/internal/svc"
	"github.com/JopenChen/zero-damai/service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *service_pb.LoginReq) (*service_pb.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &service_pb.LoginResp{}, nil
}
