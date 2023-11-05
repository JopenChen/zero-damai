package performancesessionservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSessionRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformanceSessionRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSessionRemoveLogic {
	return &PerformanceSessionRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformanceSessionRemove 演出活动场次表 删除
func (l *PerformanceSessionRemoveLogic) PerformanceSessionRemove(in *service_pb.PerformanceSessionRemoveReq) (resp *service_pb.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
