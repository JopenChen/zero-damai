package performanceservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformanceRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceRemoveLogic {
	return &PerformanceRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformanceRemove 演出活动表 删除
func (l *PerformanceRemoveLogic) PerformanceRemove(in *service_pb.PerformanceRemoveReq) (resp *service_pb.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
