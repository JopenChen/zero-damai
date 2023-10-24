package performance

import (
	"context"
	"github.com/JopenChen/zero-damai/common/global"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PerformanceModel = (*customPerformanceModel)(nil)

type (
	// PerformanceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPerformanceModel.
	PerformanceModel interface {
		performanceModel
		RowBuilder() squirrel.SelectBuilder

		// GetByCondition 根据条件获取单条记录
		GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Performance, error)
		// FindByCondition 根据条件获取所有符合的记录
		FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Performance, error)
	}

	customPerformanceModel struct {
		*defaultPerformanceModel
	}
)

// NewPerformanceModel returns a model for the database table.
func NewPerformanceModel(conn sqlx.SqlConn) PerformanceModel {
	return &customPerformanceModel{
		defaultPerformanceModel: newPerformanceModel(conn),
	}
}

func (c customPerformanceModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(performanceRows).From(c.table)
}

// GetByCondition 根据条件获取单条记录
func (c customPerformanceModel) GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (data *Performance, err error) {
	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return
	}

	var resp Performance
	err = c.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// FindByCondition 根据条件获取所有符合的记录
func (c customPerformanceModel) FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Performance, error) {
	if orderBy == global.EmptyString {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Performance
	err = c.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
