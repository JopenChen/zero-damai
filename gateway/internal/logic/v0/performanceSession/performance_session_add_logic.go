package performanceSession

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSessionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformanceSessionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSessionAddLogic {
	return &PerformanceSessionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformanceSessionAddLogic) PerformanceSessionAdd(req *types.PerformanceSessionAddReq) (resp *types.PerformanceSessionAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
