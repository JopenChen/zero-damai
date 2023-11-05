package performanceSeat

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSeatRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformanceSeatRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSeatRemoveLogic {
	return &PerformanceSeatRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformanceSeatRemoveLogic) PerformanceSeatRemove(req *types.GeneralPathId) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
