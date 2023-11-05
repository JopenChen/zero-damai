package userservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRetrieveLogic {
	return &UserRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserRetrieve 用户表 获取
func (l *UserRetrieveLogic) UserRetrieve(in *service_pb.UserRetrieveReq) (resp *service_pb.UserRetrieveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
