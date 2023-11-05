package performanceSeat

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSeatAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformanceSeatAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSeatAddLogic {
	return &PerformanceSeatAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformanceSeatAddLogic) PerformanceSeatAdd(req *types.PerformanceSeatAddReq) (resp *types.PerformanceSeatAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
