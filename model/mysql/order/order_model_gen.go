// Code generated by goctl. DO NOT EDIT.

package order

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
	orderFieldNames          = builder.RawFieldNames(&Order{})
	orderRows                = strings.Join(orderFieldNames, ",")
	orderRowsExpectAutoSet   = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	orderRowsWithPlaceHolder = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	orderModel interface {
		Insert(ctx context.Context, data *Order) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Order, error)
		Update(ctx context.Context, data *Order) error
		Delete(ctx context.Context, id int64) error

		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder

		// GetByCondition 根据条件获取单条记录
		GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Order, error)
		// FindByCondition 根据条件获取所有符合的记录
		FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Order, error)
		// FindPageByCondition 根据条件获取分页记录
		FindPageByCondition(ctx context.Context, page int64, pageSize int64, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Order, error)
		// CountByCondition 根据条件获取总数
		CountByCondition(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
	}

	defaultOrderModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Order struct {
		Id         int64     `db:"id"`         // 用户ID
		Name       string    `db:"name"`       // 姓名
		Nickname   string    `db:"nickname"`   // 昵称
		Avatar     string    `db:"avatar"`     // 头像
		Background string    `db:"background"` // 个人中心背景图
		Mobile     string    `db:"mobile"`     // 手机号
		Password   string    `db:"password"`   // 密码
		Mail       string    `db:"mail"`       // 邮箱地址
		Identity   string    `db:"identity"`   // 身份证号
		Gender     int64     `db:"gender"`     // 性别: 0=未指定、1=男、2=女、3=第三性别、4=保密: 默认=4
		Nation     string    `db:"nation"`     // 民族
		Birthday   int64     `db:"birthday"`   // 出生日期
		Address    string    `db:"address"`    // 收货地址
		Audience   string    `db:"audience"`   // 观影人信息,格式：[{"name":"xxx","identity_type": 1,"identity_number"}],identity 可选值：1=身份证、2=港澳台居民居住证、3=港澳居民来往内地通行证、4=台湾居民来往大陆通行证、5=护照、6=外国人永久居留身份证
		Status     int64     `db:"status"`     // 是否禁用: 0=未指定、1=开启、2=禁用: 默认=1
		LoginAt    time.Time `db:"login_at"`   // 登录时间
		CreatedAt  time.Time `db:"created_at"` // 创建时间
		UpdatedAt  time.Time `db:"updated_at"` // 更新时间
		IsDel      int64     `db:"is_del"`     // 是否删除: 0=未指定、1=是、2=否: 默认=2
	}
)

func (m *defaultOrderModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(orderRows).From(m.table)
}

func (m *defaultOrderModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ") as count").From(m.table)
}

// GetByCondition 根据条件获取单条记录
func (m *defaultOrderModel) GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (data *Order, err error) {
	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return
	}

	var resp Order
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// FindByCondition 根据条件获取所有符合的记录
func (m *defaultOrderModel) FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Order, error) {
	if orderBy == global.EmptyString {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Order
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// FindPageByCondition 根据条件获取分页记录
func (m *defaultOrderModel) FindPageByCondition(ctx context.Context, page int64, pageSize int64, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Order, error) {
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

	var resp []*Order
	err = m.conn.QueryRowsCtx(ctx, &resp, query, paramList...)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// CountByCondition 根据条件获取总数
func (m *defaultOrderModel) CountByCondition(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {
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

func newOrderModel(conn sqlx.SqlConn) *defaultOrderModel {
	return &defaultOrderModel{
		conn:  conn,
		table: "`order`",
	}
}

func (m *defaultOrderModel) withSession(session sqlx.Session) *defaultOrderModel {
	return &defaultOrderModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`order`",
	}
}

func (m *defaultOrderModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ? and `is_del` = %s", m.table, global.DelStateNo)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOrderModel) FindOne(ctx context.Context, id int64) (*Order, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? and `is_del` = %s limit 1", orderRows, m.table, global.DelStateNo)
	var resp Order
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

func (m *defaultOrderModel) Insert(ctx context.Context, data *Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Nickname, data.Avatar, data.Background, data.Mobile, data.Password, data.Mail, data.Identity, data.Gender, data.Nation, data.Birthday, data.Address, data.Audience, data.Status, data.LoginAt, data.IsDel)
	return ret, err
}

func (m *defaultOrderModel) Update(ctx context.Context, data *Order) error {
	query := fmt.Sprintf("update %s set %s where `id` = ? and `is_del` = %s", m.table, orderRowsWithPlaceHolder, global.DelStateNo)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Nickname, data.Avatar, data.Background, data.Mobile, data.Password, data.Mail, data.Identity, data.Gender, data.Nation, data.Birthday, data.Address, data.Audience, data.Status, data.LoginAt, data.IsDel, data.Id)
	return err
}

func (m *defaultOrderModel) tableName() string {
	return m.table
}
