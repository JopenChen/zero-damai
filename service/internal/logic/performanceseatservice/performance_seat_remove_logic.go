package performanceseatservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSeatRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformanceSeatRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSeatRemoveLogic {
	return &PerformanceSeatRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformanceSeatRemove 演出活动场次座位表 删除
func (l *PerformanceSeatRemoveLogic) PerformanceSeatRemove(in *service_pb.PerformanceSeatRemoveReq) (resp *service_pb.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
