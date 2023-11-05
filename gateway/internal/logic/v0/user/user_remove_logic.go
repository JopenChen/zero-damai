package user

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRemoveLogic {
	return &UserRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRemoveLogic) UserRemove(req *types.GeneralPathId) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
