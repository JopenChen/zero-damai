package performanceseatservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSeatUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformanceSeatUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSeatUpdateLogic {
	return &PerformanceSeatUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformanceSeatUpdate 演出活动场次座位表 更新
func (l *PerformanceSeatUpdateLogic) PerformanceSeatUpdate(in *service_pb.PerformanceSeatUpdateReq) (resp *service_pb.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
