package performanceseatservicelogic

import (
	"context"

	"service/internal/svc"
	"service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSeatRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformanceSeatRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSeatRetrieveLogic {
	return &PerformanceSeatRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformanceSeatRetrieve 演出活动场次座位表 获取
func (l *PerformanceSeatRetrieveLogic) PerformanceSeatRetrieve(in *service_pb.PerformanceSeatRetrieveReq) (resp *service_pb.PerformanceSeatRetrieveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
