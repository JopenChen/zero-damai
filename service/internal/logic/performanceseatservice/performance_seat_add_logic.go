package performanceseatservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSeatAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformanceSeatAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSeatAddLogic {
	return &PerformanceSeatAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformanceSeatAdd 演出活动场次座位表 创建
func (l *PerformanceSeatAddLogic) PerformanceSeatAdd(in *service_pb.PerformanceSeatAddReq) (resp *service_pb.PerformanceSeatAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
