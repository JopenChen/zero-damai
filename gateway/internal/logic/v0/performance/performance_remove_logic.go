package performance

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformanceRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceRemoveLogic {
	return &PerformanceRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformanceRemoveLogic) PerformanceRemove(req *types.GeneralPathId) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
