package performance

import (
	"context"
	"github.com/JopenChen/zero-damai/common/global"
	"github.com/JopenChen/zero-damai/service/client/service"
	"github.com/bytedance/sonic"

	"github.com/JopenChen/zero-damai/gateway/internal/svc"
	"github.com/JopenChen/zero-damai/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformanceRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformanceRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceRetrieveLogic {
	return &PerformanceRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformanceRetrieveLogic) PerformanceRetrieve(req *types.PerformanceRetrieveReq) (resp *types.PerformanceRetrieveResp, err error) {
	resp = new(types.PerformanceRetrieveResp)
	resp.Data = make([]*types.Performance, 0)

	// 过滤条件解析
	var filterArray []*service.FilterItem
	if req.Filter != global.EmptyString {
		err = sonic.UnmarshalString(req.Filter, &filterArray)
		if err != nil {
			l.Logger.Errorf("sonic.UnmarshalString(req.Filter, &filterArray) error: %v, filter: %+v", err, filterArray)
			return
		}
	}
	// 排序条件解析
	var sortArray []*service.SortItem
	if req.Sort != global.EmptyString {
		err = sonic.UnmarshalString(req.Sort, &sortArray)
		if err != nil {
			l.Logger.Errorf("sonic.UnmarshalString(req.Sort, &sortArray) error: %v, sort: %+v", err, sortArray)
			return
		}
	}

	rpcResp, err := l.svcCtx.ServiceRpc.PerformanceRetrieve(l.ctx, &service.PerformanceRetrieveReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Filter:   filterArray,
		Sort:     sortArray,
	})
	if err != nil {
		return
	}

	// Return
	resp.Total = rpcResp.Total
	for _, v := range rpcResp.Data {
		resp.Data = append(resp.Data, &types.Performance{
			Id:             v.Id,
			Title:          v.Title,
			Description:    v.Description,
			City:           v.City,
			Address:        v.Address,
			PrioritySaleAt: v.PrioritySaleAt,
			SaleAt:         v.SaleAt,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
		})
	}
	return
}
