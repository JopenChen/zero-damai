package user

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRetrieveLogic {
	return &UserRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRetrieveLogic) UserRetrieve(req *types.UserRetrieveReq) (resp *types.UserRetrieveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
