type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}

		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder

        // GetByCondition 根据条件获取单条记录
        GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*{{.upperStartCamelObject}}, error)
        // FindByCondition 根据条件获取所有符合的记录
        FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*{{.upperStartCamelObject}}, error)
        // FindPageByCondition 根据条件获取分页记录
        FindPageByCondition(ctx context.Context, page int64, pageSize int64, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*{{.upperStartCamelObject}}, error)
	}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}sqlc.CachedConn{{else}}conn sqlx.SqlConn{{end}}
		table string
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)

func (m *default{{.upperStartCamelObject}}Model) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select({{.lowerStartCamelObject}}Rows).From(m.table)
}

func (m *default{{.upperStartCamelObject}}Model) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ") as count").From(m.table)
}

// GetByCondition 根据条件获取单条记录
func (m *default{{.upperStartCamelObject}}Model) GetByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder) (data *{{.upperStartCamelObject}}, err error) {
	query, values, err := rowBuilder.Where(squirrel.Eq{"is_del": global.DelStateNo}).ToSql()
	if err != nil {
		return
	}

	var resp {{.upperStartCamelObject}}
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// FindByCondition 根据条件获取所有符合的记录
func (m *default{{.upperStartCamelObject}}Model) FindByCondition(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*{{.upperStartCamelObject}}, error) {
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
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}


// FindPageByCondition 根据条件获取分页记录
func (m *default{{.upperStartCamelObject}}Model) FindPageByCondition(ctx context.Context, page int64, pageSize int64, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*{{.upperStartCamelObject}}, error) {
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

	var resp []*{{.upperStartCamelObject}}
	err = m.conn.QueryRowsCtx(ctx, &resp, query, paramList...)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// CountByCondition 根据条件获取总数
func (m *default{{.upperStartCamelObject}}Model) CountByCondition(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {
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

