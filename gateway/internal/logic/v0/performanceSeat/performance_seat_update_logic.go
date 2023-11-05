package performanceSeat

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSeatUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformanceSeatUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSeatUpdateLogic {
	return &PerformanceSeatUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformanceSeatUpdateLogic) PerformanceSeatUpdate(req *types.PerformanceSeatUpdateReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
