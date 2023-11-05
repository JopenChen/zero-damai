package performanceservicelogic

import (
	"context"
	"fmt"
	"github.com/JopenChen/zero-damai/common/global"
	"github.com/zeromicro/go-zero/core/logx"

	"service/internal/svc"
	"service/service_pb"
)

type PerformanceRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformanceRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceRetrieveLogic {
	return &PerformanceRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformanceRetrieve 演出活动表 获取
func (l *PerformanceRetrieveLogic) PerformanceRetrieve(in *service_pb.PerformanceRetrieveReq) (resp *service_pb.PerformanceRetrieveResp, err error) {
	resp = new(service_pb.PerformanceRetrieveResp)
	resp.Data = make([]*service_pb.Performance, 0)
	performanceRowBuilder := l.svcCtx.PerformanceModel.RowBuilder()
	countBuilder := l.svcCtx.PerformanceModel.CountBuilder(global.IdString)

	// 过滤条件
	for _, v := range in.Paging.Filter {
		performanceRowBuilder = performanceRowBuilder.Where(fmt.Sprintf("%s = ?", v.Field), v.Value)
		countBuilder = countBuilder.Where(fmt.Sprintf("`%s` = ?", v.Field), v.Value)
	}
	// 排序条件
	for _, v := range in.Paging.Sort {
		performanceRowBuilder = performanceRowBuilder.OrderBy(fmt.Sprintf("%s %s", v.Field, v.Order))
		countBuilder = countBuilder.OrderBy(fmt.Sprintf("%s %s", v.Field, v.Order))
	}

	// 获取数据
	performanceList, err := l.svcCtx.PerformanceModel.FindPageByCondition(l.ctx, in.Paging.Page, in.Paging.PageSize, performanceRowBuilder, "sale_at DESC")
	if err != nil {
		l.Logger.Errorf("PerformanceModel.FindPageByCondition error: %v", err)
		return
	}

	// 获取总数
	count, err := l.svcCtx.PerformanceModel.CountByCondition(l.ctx, countBuilder)
	if err != nil {
		l.Logger.Errorf("PerformanceModel.CountByCondition error: %v", err)
		return
	}

	// Return
	resp.Total = count
	for _, v := range performanceList {
		resp.Data = append(resp.Data, &service_pb.Performance{
			Id: v.Id,
			//Title:          v.Title,
			//Description:    v.Description,
			//City:           v.City,
			Address: v.Address,
			//PrioritySaleAt: v.PrioritySaleAt,
			//SaleAt:         v.SaleAt,
			CreatedAt: v.CreatedAt.Unix(),
			UpdatedAt: v.UpdatedAt.Unix(),
		})
	}
	return
}
