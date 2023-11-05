package performanceservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformanceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceAddLogic {
	return &PerformanceAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformanceAdd 演出活动表 创建
func (l *PerformanceAddLogic) PerformanceAdd(in *service_pb.PerformanceAddReq) (resp *service_pb.PerformanceAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
