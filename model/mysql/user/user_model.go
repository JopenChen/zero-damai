package user

import (
	"context"
	"github.com/JopenChen/zero-damai/common/global"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		RowBuilder() squirrel.SelectBuilder

		// GetByCondition 根据条件获取单条记录
		GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*User, error)
		// FindByCondition 根据条件获取所有符合的记录
		FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (c customUserModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(userRows).From(c.table)
}

// GetByCondition 根据条件获取单条记录
func (c customUserModel) GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (data *User, err error) {
	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return
	}

	var resp User
	err = c.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// FindByCondition 根据条件获取所有符合的记录
func (c customUserModel) FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*User, error) {
	if orderBy == global.EmptyString {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User
	err = c.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
