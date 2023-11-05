package userservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRemoveLogic {
	return &UserRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserRemove 用户表 删除
func (l *UserRemoveLogic) UserRemove(in *service_pb.UserRemoveReq) (resp *service_pb.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
