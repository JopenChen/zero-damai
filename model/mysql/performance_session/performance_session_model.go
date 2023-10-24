package performance_session

import (
	"context"
	"github.com/JopenChen/zero-damai/common/global"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PerformanceSessionModel = (*customPerformanceSessionModel)(nil)

type (
	// PerformanceSessionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPerformanceSessionModel.
	PerformanceSessionModel interface {
		performanceSessionModel
		RowBuilder() squirrel.SelectBuilder

		// GetByCondition 根据条件获取单条记录
		GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*PerformanceSession, error)
		// FindByCondition 根据条件获取所有符合的记录
		FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*PerformanceSession, error)
	}

	customPerformanceSessionModel struct {
		*defaultPerformanceSessionModel
	}
)

// NewPerformanceSessionModel returns a model for the database table.
func NewPerformanceSessionModel(conn sqlx.SqlConn) PerformanceSessionModel {
	return &customPerformanceSessionModel{
		defaultPerformanceSessionModel: newPerformanceSessionModel(conn),
	}
}

func (c customPerformanceSessionModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(performanceSessionRows).From(c.table)
}

// GetByCondition 根据条件获取单条记录
func (c customPerformanceSessionModel) GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (data *PerformanceSession, err error) {
	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return
	}

	var resp PerformanceSession
	err = c.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// FindByCondition 根据条件获取所有符合的记录
func (c customPerformanceSessionModel) FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*PerformanceSession, error) {
	if orderBy == global.EmptyString {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*PerformanceSession
	err = c.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
