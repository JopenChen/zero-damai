package performancesessionservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSessionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformanceSessionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSessionUpdateLogic {
	return &PerformanceSessionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformanceSessionUpdate 演出活动场次表 更新
func (l *PerformanceSessionUpdateLogic) PerformanceSessionUpdate(in *service_pb.PerformanceSessionUpdateReq) (resp *service_pb.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
