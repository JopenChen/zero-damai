package performanceSession

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSessionRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformanceSessionRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSessionRemoveLogic {
	return &PerformanceSessionRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformanceSessionRemoveLogic) PerformanceSessionRemove(req *types.GeneralPathId) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
