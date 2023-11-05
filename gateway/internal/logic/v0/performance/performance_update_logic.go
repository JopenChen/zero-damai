package performance

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformanceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceUpdateLogic {
	return &PerformanceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformanceUpdateLogic) PerformanceUpdate(req *types.PerformanceUpdateReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
