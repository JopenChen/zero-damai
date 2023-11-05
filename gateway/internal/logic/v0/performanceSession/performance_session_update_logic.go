package performanceSession

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSessionUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformanceSessionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSessionUpdateLogic {
	return &PerformanceSessionUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformanceSessionUpdateLogic) PerformanceSessionUpdate(req *types.PerformanceSessionUpdateReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
