package orderservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderRetrieveLogic {
	return &OrderRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderRetrieve 订单表 获取
func (l *OrderRetrieveLogic) OrderRetrieve(in *service_pb.OrderRetrieveReq) (resp *service_pb.OrderRetrieveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
