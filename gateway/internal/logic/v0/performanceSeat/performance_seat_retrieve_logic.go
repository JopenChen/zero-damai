package performanceSeat

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSeatRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformanceSeatRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSeatRetrieveLogic {
	return &PerformanceSeatRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformanceSeatRetrieveLogic) PerformanceSeatRetrieve(req *types.PerformanceSeatRetrieveReq) (resp *types.PerformanceSeatRetrieveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
