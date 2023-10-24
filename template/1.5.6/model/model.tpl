package {{.pkg}}
{{if .withCache}}
import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)
{{else}}
import (
	"context"
	"github.com/JopenChen/zero-damai/common/global"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)
{{end}}
var _ {{.upperStartCamelObject}}Model = (*custom{{.upperStartCamelObject}}Model)(nil)

type (
	// {{.upperStartCamelObject}}Model is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.upperStartCamelObject}}Model.
	{{.upperStartCamelObject}}Model interface {
		{{.lowerStartCamelObject}}Model
		RowBuilder() squirrel.SelectBuilder

		// GetByCondition 根据条件获取单条记录
        GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*{{.upperStartCamelObject}}, error)
        // FindByCondition 根据条件获取所有符合的记录
        FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*{{.upperStartCamelObject}}, error)
	}

	custom{{.upperStartCamelObject}}Model struct {
		*default{{.upperStartCamelObject}}Model
	}
)

// New{{.upperStartCamelObject}}Model returns a model for the database table.
func New{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf, opts ...cache.Option{{end}}) {{.upperStartCamelObject}}Model {
	return &custom{{.upperStartCamelObject}}Model{
		default{{.upperStartCamelObject}}Model: new{{.upperStartCamelObject}}Model(conn{{if .withCache}}, c, opts...{{end}}),
	}
}

func (c custom{{.upperStartCamelObject}}Model) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select({{.lowerStartCamelObject}}Rows).From(c.table)
}

// GetByCondition 根据条件获取单条记录
func (c custom{{.upperStartCamelObject}}Model) GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (data *{{.upperStartCamelObject}}, err error) {
	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return
	}

	var resp {{.upperStartCamelObject}}
	err = c.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// FindByCondition 根据条件获取所有符合的记录
func (c custom{{.upperStartCamelObject}}Model) FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*{{.upperStartCamelObject}}, error) {
	if orderBy == global.EmptyString {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	err = c.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
