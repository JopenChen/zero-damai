package performance_seat

import (
	"context"
	"github.com/JopenChen/zero-damai/common/global"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PerformanceSeatModel = (*customPerformanceSeatModel)(nil)

type (
	// PerformanceSeatModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPerformanceSeatModel.
	PerformanceSeatModel interface {
		performanceSeatModel
		RowBuilder() squirrel.SelectBuilder

		// GetByCondition 根据条件获取单条记录
		GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*PerformanceSeat, error)
		// FindByCondition 根据条件获取所有符合的记录
		FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*PerformanceSeat, error)
	}

	customPerformanceSeatModel struct {
		*defaultPerformanceSeatModel
	}
)

// NewPerformanceSeatModel returns a model for the database table.
func NewPerformanceSeatModel(conn sqlx.SqlConn) PerformanceSeatModel {
	return &customPerformanceSeatModel{
		defaultPerformanceSeatModel: newPerformanceSeatModel(conn),
	}
}

func (c customPerformanceSeatModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(performanceSeatRows).From(c.table)
}

// GetByCondition 根据条件获取单条记录
func (c customPerformanceSeatModel) GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (data *PerformanceSeat, err error) {
	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return
	}

	var resp PerformanceSeat
	err = c.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// FindByCondition 根据条件获取所有符合的记录
func (c customPerformanceSeatModel) FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*PerformanceSeat, error) {
	if orderBy == global.EmptyString {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*PerformanceSeat
	err = c.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
