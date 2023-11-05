package order

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderRemoveLogic {
	return &OrderRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderRemoveLogic) OrderRemove(req *types.GeneralPathId) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
