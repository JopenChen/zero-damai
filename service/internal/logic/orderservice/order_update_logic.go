package orderservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderUpdateLogic {
	return &OrderUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderUpdate 订单表 更新
func (l *OrderUpdateLogic) OrderUpdate(in *service_pb.OrderUpdateReq) (resp *service_pb.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
