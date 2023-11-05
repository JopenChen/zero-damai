package performanceSession

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceSessionRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformanceSessionRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceSessionRetrieveLogic {
	return &PerformanceSessionRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformanceSessionRetrieveLogic) PerformanceSessionRetrieve(req *types.PerformanceSessionRetrieveReq) (resp *types.PerformanceSessionRetrieveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
