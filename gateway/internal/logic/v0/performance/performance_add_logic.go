package performance

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformanceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceAddLogic {
	return &PerformanceAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformanceAddLogic) PerformanceAdd(req *types.PerformanceAddReq) (resp *types.PerformanceAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
