package performancesessionservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSessionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformanceSessionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSessionAddLogic {
	return &PerformanceSessionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformanceSessionAdd 演出活动场次表 创建
func (l *PerformanceSessionAddLogic) PerformanceSessionAdd(in *service_pb.PerformanceSessionAddReq) (resp *service_pb.PerformanceSessionAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
