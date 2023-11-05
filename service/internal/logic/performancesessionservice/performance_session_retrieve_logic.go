package performancesessionservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSessionRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformanceSessionRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSessionRetrieveLogic {
	return &PerformanceSessionRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformanceSessionRetrieve 演出活动场次表 获取
func (l *PerformanceSessionRetrieveLogic) PerformanceSessionRetrieve(in *service_pb.PerformanceSessionRetrieveReq) (resp *service_pb.PerformanceSessionRetrieveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
