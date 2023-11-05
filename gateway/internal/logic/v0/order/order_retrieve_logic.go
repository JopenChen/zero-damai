package order

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderRetrieveLogic {
	return &OrderRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderRetrieveLogic) OrderRetrieve(req *types.OrderRetrieveReq) (resp *types.OrderRetrieveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
