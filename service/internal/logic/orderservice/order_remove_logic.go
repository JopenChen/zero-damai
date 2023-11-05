package orderservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderRemoveLogic {
	return &OrderRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderRemove 订单表 删除
func (l *OrderRemoveLogic) OrderRemove(in *service_pb.OrderRemoveReq) (resp *service_pb.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
