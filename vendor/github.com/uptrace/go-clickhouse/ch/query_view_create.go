package ch

import (
	"context"
	"database/sql"

	"github.com/uptrace/go-clickhouse/ch/chschema"
	"github.com/uptrace/go-clickhouse/ch/internal"
)

type CreateViewQuery struct {
	whereBaseQuery

	materialized bool
	ifNotExists  bool
	view         chschema.QueryWithArgs
	cluster      chschema.QueryWithArgs
	to           chschema.QueryWithArgs
	group        []chschema.QueryWithArgs
	order        chschema.QueryWithArgs
}

var _ Query = (*CreateViewQuery)(nil)

func NewCreateViewQuery(db *DB) *CreateViewQuery {
	return &CreateViewQuery{
		whereBaseQuery: whereBaseQuery{
			baseQuery: baseQuery{
				db: db,
			},
		},
	}
}

func (q *CreateViewQuery) Model(model any) *CreateViewQuery {
	q.setTableModel(model)
	return q
}

func (q *CreateViewQuery) WithQuery(fn func(*CreateViewQuery) *CreateViewQuery) *CreateViewQuery {
	return fn(q)
}

//------------------------------------------------------------------------------

func (q *CreateViewQuery) View(view string) *CreateViewQuery {
	q.view = chschema.UnsafeIdent(view)
	return q
}

func (q *CreateViewQuery) ViewExpr(query string, args ...any) *CreateViewQuery {
	q.view = chschema.SafeQuery(query, args)
	return q
}

func (q *CreateViewQuery) OnCluster(cluster string) *CreateViewQuery {
	q.cluster = chschema.UnsafeIdent(cluster)
	return q
}

func (q *CreateViewQuery) OnClusterExpr(query string, args ...any) *CreateViewQuery {
	q.cluster = chschema.SafeQuery(query, args)
	return q
}

func (q *CreateViewQuery) To(to string) *CreateViewQuery {
	q.to = chschema.UnsafeIdent(to)
	return q
}

func (q *CreateViewQuery) ToExpr(query string, args ...any) *CreateViewQuery {
	q.to = chschema.SafeQuery(query, args)
	return q
}

func (q *CreateViewQuery) Table(tables ...string) *CreateViewQuery {
	for _, table := range tables {
		q.addTable(chschema.UnsafeIdent(table))
	}
	return q
}

func (q *CreateViewQuery) TableExpr(query string, args ...any) *CreateViewQuery {
	q.addTable(chschema.SafeQuery(query, args))
	return q
}

func (q *CreateViewQuery) ModelTableExpr(query string, args ...any) *CreateViewQuery {
	q.modelTableName = chschema.SafeQuery(query, args)
	return q
}

//------------------------------------------------------------------------------

func (q *CreateViewQuery) Column(columns ...string) *CreateViewQuery {
	for _, column := range columns {
		q.addColumn(chschema.UnsafeIdent(column))
	}
	return q
}

func (q *CreateViewQuery) ColumnExpr(query string, args ...any) *CreateViewQuery {
	q.addColumn(chschema.SafeQuery(query, args))
	return q
}

func (q *CreateViewQuery) ExcludeColumn(columns ...string) *CreateViewQuery {
	q.excludeColumn(columns)
	return q
}

//------------------------------------------------------------------------------

func (q *CreateViewQuery) Materialized() *CreateViewQuery {
	q.materialized = true
	return q
}

func (q *CreateViewQuery) IfNotExists() *CreateViewQuery {
	q.ifNotExists = true
	return q
}

//------------------------------------------------------------------------------

func (q *CreateViewQuery) Where(query string, args ...any) *CreateViewQuery {
	q.addWhere(chschema.SafeQueryWithSep(query, args, " AND "))
	return q
}

func (q *CreateViewQuery) WhereOr(query string, args ...any) *CreateViewQuery {
	q.addWhere(chschema.SafeQueryWithSep(query, args, " OR "))
	return q
}

func (q *CreateViewQuery) WhereGroup(sep string, fn func(*CreateViewQuery) *CreateViewQuery) *CreateViewQuery {
	saved := q.where
	q.where = nil

	q = fn(q)

	where := q.where
	q.where = saved

	q.addWhereGroup(sep, where)

	return q
}

func (q *CreateViewQuery) Group(columns ...string) *CreateViewQuery {
	for _, column := range columns {
		q.group = append(q.group, chschema.UnsafeIdent(column))
	}
	return q
}

func (q *CreateViewQuery) GroupExpr(group string, args ...any) *CreateViewQuery {
	q.group = append(q.group, chschema.SafeQuery(group, args))
	return q
}

func (q *CreateViewQuery) OrderExpr(query string, args ...any) *CreateViewQuery {
	q.order = chschema.SafeQuery(query, args)
	return q
}

func (q *CreateViewQuery) Setting(query string, args ...any) *CreateViewQuery {
	q.settings = append(q.settings, chschema.SafeQuery(query, args))
	return q
}

//------------------------------------------------------------------------------

func (q *CreateViewQuery) Operation() string {
	return "CREATE VIEW"
}

var _ chschema.QueryAppender = (*CreateViewQuery)(nil)

func (q *CreateViewQuery) AppendQuery(fmter chschema.Formatter, b []byte) (_ []byte, err error) {
	if q.err != nil {
		return nil, q.err
	}

	b = append(b, "CREATE "...)
	if q.materialized {
		b = append(b, "MATERIALIZED "...)
	}
	b = append(b, "VIEW "...)
	if q.ifNotExists {
		b = append(b, "IF NOT EXISTS "...)
	}

	b, err = q.view.AppendQuery(fmter, b)
	if err != nil {
		return nil, err
	}

	if !q.cluster.IsEmpty() {
		b = append(b, " ON CLUSTER "...)
		b, err = q.cluster.AppendQuery(fmter, b)
		if err != nil {
			return nil, err
		}
	}

	b = append(b, " TO "...)
	b, err = q.to.AppendQuery(fmter, b)
	if err != nil {
		return nil, err
	}
	b = append(b, " AS "...)

	b = append(b, "SELECT "...)

	b, err = q.appendColumns(fmter, b)
	if err != nil {
		return nil, err
	}

	b = append(b, " FROM "...)
	b, err = q.appendTablesWithAlias(fmter, b)
	if err != nil {
		return nil, err
	}

	b, err = q.appendWhere(fmter, b)
	if err != nil {
		return nil, err
	}

	if len(q.group) > 0 {
		b = append(b, " GROUP BY "...)
		for i, f := range q.group {
			if i > 0 {
				b = append(b, ", "...)
			}
			b, err = f.AppendQuery(fmter, b)
			if err != nil {
				return nil, err
			}
		}
	}

	if !q.order.IsZero() {
		b = append(b, " ORDER BY "...)
		b, err = q.order.AppendQuery(fmter, b)
		if err != nil {
			return nil, err
		}
	}

	b, err = q.appendSettings(fmter, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

//------------------------------------------------------------------------------

func (q *CreateViewQuery) Exec(ctx context.Context, dest ...any) (sql.Result, error) {
	queryBytes, err := q.AppendQuery(q.db.fmter, q.db.makeQueryBytes())
	if err != nil {
		return nil, err
	}
	query := internal.String(queryBytes)

	return q.exec(ctx, q, query)
}
