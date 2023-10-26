// Code generated by goctl. DO NOT EDIT.

package performance_session

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"

	"github.com/JopenChen/zero-damai/common/global"
	"github.com/Masterminds/squirrel"
)

var (
	performanceSessionFieldNames          = builder.RawFieldNames(&PerformanceSession{})
	performanceSessionRows                = strings.Join(performanceSessionFieldNames, ",")
	performanceSessionRowsExpectAutoSet   = strings.Join(stringx.Remove(performanceSessionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	performanceSessionRowsWithPlaceHolder = strings.Join(stringx.Remove(performanceSessionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	performanceSessionModel interface {
		Insert(ctx context.Context, data *PerformanceSession) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PerformanceSession, error)
		Update(ctx context.Context, data *PerformanceSession) error
		Delete(ctx context.Context, id int64) error

		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder

		// GetByCondition 根据条件获取单条记录
		GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*PerformanceSession, error)
		// FindByCondition 根据条件获取所有符合的记录
		FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*PerformanceSession, error)
		// FindPageByCondition 根据条件获取分页记录
		FindPageByCondition(ctx context.Context, page int64, pageSize int64, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*PerformanceSession, error)
		// CountByCondition 根据条件获取总数
		CountByCondition(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
	}

	defaultPerformanceSessionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PerformanceSession struct {
		Id            int64          `db:"id"`             // ID
		PerformanceId int64          `db:"performance_id"` // 关联的演出活动ID
		Tickets       sql.NullString `db:"tickets"`        // 规格票数，格式为：[{"seat": "vip","price": 2023,"quantity": 2000}]。seat: 席位、price: 票价、quantity: 数量。
		ShowAt        int64          `db:"show_at"`        // 演出时间
		CreatedAt     time.Time      `db:"created_at"`     // 创建时间
		UpdatedAt     time.Time      `db:"updated_at"`     // 更新时间
		IsDel         int64          `db:"is_del"`         // 是否删除: 0=未指定、1=是、2=否: 默认=2
	}
)

func (m *defaultPerformanceSessionModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(performanceSessionRows).From(m.table)
}

func (m *defaultPerformanceSessionModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ") as count").From(m.table)
}

// GetByCondition 根据条件获取单条记录
func (m *defaultPerformanceSessionModel) GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (data *PerformanceSession, err error) {
	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return
	}

	var resp PerformanceSession
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// FindByCondition 根据条件获取所有符合的记录
func (m *defaultPerformanceSessionModel) FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*PerformanceSession, error) {
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
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// FindPageByCondition 根据条件获取分页记录
func (m *defaultPerformanceSessionModel) FindPageByCondition(ctx context.Context, page int64, pageSize int64, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*PerformanceSession, error) {
	if orderBy == global.EmptyString {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if pageSize == global.Zero {
		pageSize = 10
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, paramList, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).
		Offset(uint64(offset)).
		Limit(uint64(pageSize)).
		ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*PerformanceSession
	err = m.conn.QueryRowsCtx(ctx, &resp, query, paramList...)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// CountByCondition 根据条件获取总数
func (m *defaultPerformanceSessionModel) CountByCondition(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {
	query, paramList, err := countBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.conn.QueryRowCtx(ctx, &resp, query, paramList...)
	if err != nil {
		return 0, err
	}
	return resp, nil
}

func newPerformanceSessionModel(conn sqlx.SqlConn) *defaultPerformanceSessionModel {
	return &defaultPerformanceSessionModel{
		conn:  conn,
		table: "`performance_session`",
	}
}

func (m *defaultPerformanceSessionModel) withSession(session sqlx.Session) *defaultPerformanceSessionModel {
	return &defaultPerformanceSessionModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`performance_session`",
	}
}

func (m *defaultPerformanceSessionModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ? and `is_del` = %s", m.table, global.DelStateNo)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPerformanceSessionModel) FindOne(ctx context.Context, id int64) (*PerformanceSession, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? and `is_del` = %s limit 1", performanceSessionRows, m.table, global.DelStateNo)
	var resp PerformanceSession
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPerformanceSessionModel) Insert(ctx context.Context, data *PerformanceSession) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, performanceSessionRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.PerformanceId, data.Tickets, data.ShowAt, data.IsDel)
	return ret, err
}

func (m *defaultPerformanceSessionModel) Update(ctx context.Context, data *PerformanceSession) error {
	query := fmt.Sprintf("update %s set %s where `id` = ? and `is_del` = %s", m.table, performanceSessionRowsWithPlaceHolder, global.DelStateNo)
	_, err := m.conn.ExecCtx(ctx, query, data.PerformanceId, data.Tickets, data.ShowAt, data.IsDel, data.Id)
	return err
}

func (m *defaultPerformanceSessionModel) tableName() string {
	return m.table
}
