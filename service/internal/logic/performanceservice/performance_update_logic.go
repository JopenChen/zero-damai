package performanceservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformanceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceUpdateLogic {
	return &PerformanceUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformanceUpdate 演出活动表 更新
func (l *PerformanceUpdateLogic) PerformanceUpdate(in *service_pb.PerformanceUpdateReq) (resp *service_pb.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
