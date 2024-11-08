// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billinginvoicemanuallineconfig"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// BillingInvoiceManualLineConfigQuery is the builder for querying BillingInvoiceManualLineConfig entities.
type BillingInvoiceManualLineConfigQuery struct {
	config
	ctx        *QueryContext
	order      []billinginvoicemanuallineconfig.OrderOption
	inters     []Interceptor
	predicates []predicate.BillingInvoiceManualLineConfig
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BillingInvoiceManualLineConfigQuery builder.
func (bimlcq *BillingInvoiceManualLineConfigQuery) Where(ps ...predicate.BillingInvoiceManualLineConfig) *BillingInvoiceManualLineConfigQuery {
	bimlcq.predicates = append(bimlcq.predicates, ps...)
	return bimlcq
}

// Limit the number of records to be returned by this query.
func (bimlcq *BillingInvoiceManualLineConfigQuery) Limit(limit int) *BillingInvoiceManualLineConfigQuery {
	bimlcq.ctx.Limit = &limit
	return bimlcq
}

// Offset to start from.
func (bimlcq *BillingInvoiceManualLineConfigQuery) Offset(offset int) *BillingInvoiceManualLineConfigQuery {
	bimlcq.ctx.Offset = &offset
	return bimlcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bimlcq *BillingInvoiceManualLineConfigQuery) Unique(unique bool) *BillingInvoiceManualLineConfigQuery {
	bimlcq.ctx.Unique = &unique
	return bimlcq
}

// Order specifies how the records should be ordered.
func (bimlcq *BillingInvoiceManualLineConfigQuery) Order(o ...billinginvoicemanuallineconfig.OrderOption) *BillingInvoiceManualLineConfigQuery {
	bimlcq.order = append(bimlcq.order, o...)
	return bimlcq
}

// First returns the first BillingInvoiceManualLineConfig entity from the query.
// Returns a *NotFoundError when no BillingInvoiceManualLineConfig was found.
func (bimlcq *BillingInvoiceManualLineConfigQuery) First(ctx context.Context) (*BillingInvoiceManualLineConfig, error) {
	nodes, err := bimlcq.Limit(1).All(setContextOp(ctx, bimlcq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{billinginvoicemanuallineconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bimlcq *BillingInvoiceManualLineConfigQuery) FirstX(ctx context.Context) *BillingInvoiceManualLineConfig {
	node, err := bimlcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BillingInvoiceManualLineConfig ID from the query.
// Returns a *NotFoundError when no BillingInvoiceManualLineConfig ID was found.
func (bimlcq *BillingInvoiceManualLineConfigQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bimlcq.Limit(1).IDs(setContextOp(ctx, bimlcq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{billinginvoicemanuallineconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bimlcq *BillingInvoiceManualLineConfigQuery) FirstIDX(ctx context.Context) string {
	id, err := bimlcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BillingInvoiceManualLineConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BillingInvoiceManualLineConfig entity is found.
// Returns a *NotFoundError when no BillingInvoiceManualLineConfig entities are found.
func (bimlcq *BillingInvoiceManualLineConfigQuery) Only(ctx context.Context) (*BillingInvoiceManualLineConfig, error) {
	nodes, err := bimlcq.Limit(2).All(setContextOp(ctx, bimlcq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{billinginvoicemanuallineconfig.Label}
	default:
		return nil, &NotSingularError{billinginvoicemanuallineconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bimlcq *BillingInvoiceManualLineConfigQuery) OnlyX(ctx context.Context) *BillingInvoiceManualLineConfig {
	node, err := bimlcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BillingInvoiceManualLineConfig ID in the query.
// Returns a *NotSingularError when more than one BillingInvoiceManualLineConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (bimlcq *BillingInvoiceManualLineConfigQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bimlcq.Limit(2).IDs(setContextOp(ctx, bimlcq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{billinginvoicemanuallineconfig.Label}
	default:
		err = &NotSingularError{billinginvoicemanuallineconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bimlcq *BillingInvoiceManualLineConfigQuery) OnlyIDX(ctx context.Context) string {
	id, err := bimlcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BillingInvoiceManualLineConfigs.
func (bimlcq *BillingInvoiceManualLineConfigQuery) All(ctx context.Context) ([]*BillingInvoiceManualLineConfig, error) {
	ctx = setContextOp(ctx, bimlcq.ctx, ent.OpQueryAll)
	if err := bimlcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BillingInvoiceManualLineConfig, *BillingInvoiceManualLineConfigQuery]()
	return withInterceptors[[]*BillingInvoiceManualLineConfig](ctx, bimlcq, qr, bimlcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bimlcq *BillingInvoiceManualLineConfigQuery) AllX(ctx context.Context) []*BillingInvoiceManualLineConfig {
	nodes, err := bimlcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BillingInvoiceManualLineConfig IDs.
func (bimlcq *BillingInvoiceManualLineConfigQuery) IDs(ctx context.Context) (ids []string, err error) {
	if bimlcq.ctx.Unique == nil && bimlcq.path != nil {
		bimlcq.Unique(true)
	}
	ctx = setContextOp(ctx, bimlcq.ctx, ent.OpQueryIDs)
	if err = bimlcq.Select(billinginvoicemanuallineconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bimlcq *BillingInvoiceManualLineConfigQuery) IDsX(ctx context.Context) []string {
	ids, err := bimlcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bimlcq *BillingInvoiceManualLineConfigQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bimlcq.ctx, ent.OpQueryCount)
	if err := bimlcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bimlcq, querierCount[*BillingInvoiceManualLineConfigQuery](), bimlcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bimlcq *BillingInvoiceManualLineConfigQuery) CountX(ctx context.Context) int {
	count, err := bimlcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bimlcq *BillingInvoiceManualLineConfigQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bimlcq.ctx, ent.OpQueryExist)
	switch _, err := bimlcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bimlcq *BillingInvoiceManualLineConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := bimlcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BillingInvoiceManualLineConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bimlcq *BillingInvoiceManualLineConfigQuery) Clone() *BillingInvoiceManualLineConfigQuery {
	if bimlcq == nil {
		return nil
	}
	return &BillingInvoiceManualLineConfigQuery{
		config:     bimlcq.config,
		ctx:        bimlcq.ctx.Clone(),
		order:      append([]billinginvoicemanuallineconfig.OrderOption{}, bimlcq.order...),
		inters:     append([]Interceptor{}, bimlcq.inters...),
		predicates: append([]predicate.BillingInvoiceManualLineConfig{}, bimlcq.predicates...),
		// clone intermediate query.
		sql:  bimlcq.sql.Clone(),
		path: bimlcq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BillingInvoiceManualLineConfig.Query().
//		GroupBy(billinginvoicemanuallineconfig.FieldNamespace).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (bimlcq *BillingInvoiceManualLineConfigQuery) GroupBy(field string, fields ...string) *BillingInvoiceManualLineConfigGroupBy {
	bimlcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BillingInvoiceManualLineConfigGroupBy{build: bimlcq}
	grbuild.flds = &bimlcq.ctx.Fields
	grbuild.label = billinginvoicemanuallineconfig.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//	}
//
//	client.BillingInvoiceManualLineConfig.Query().
//		Select(billinginvoicemanuallineconfig.FieldNamespace).
//		Scan(ctx, &v)
func (bimlcq *BillingInvoiceManualLineConfigQuery) Select(fields ...string) *BillingInvoiceManualLineConfigSelect {
	bimlcq.ctx.Fields = append(bimlcq.ctx.Fields, fields...)
	sbuild := &BillingInvoiceManualLineConfigSelect{BillingInvoiceManualLineConfigQuery: bimlcq}
	sbuild.label = billinginvoicemanuallineconfig.Label
	sbuild.flds, sbuild.scan = &bimlcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BillingInvoiceManualLineConfigSelect configured with the given aggregations.
func (bimlcq *BillingInvoiceManualLineConfigQuery) Aggregate(fns ...AggregateFunc) *BillingInvoiceManualLineConfigSelect {
	return bimlcq.Select().Aggregate(fns...)
}

func (bimlcq *BillingInvoiceManualLineConfigQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bimlcq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bimlcq); err != nil {
				return err
			}
		}
	}
	for _, f := range bimlcq.ctx.Fields {
		if !billinginvoicemanuallineconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if bimlcq.path != nil {
		prev, err := bimlcq.path(ctx)
		if err != nil {
			return err
		}
		bimlcq.sql = prev
	}
	return nil
}

func (bimlcq *BillingInvoiceManualLineConfigQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BillingInvoiceManualLineConfig, error) {
	var (
		nodes = []*BillingInvoiceManualLineConfig{}
		_spec = bimlcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BillingInvoiceManualLineConfig).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BillingInvoiceManualLineConfig{config: bimlcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(bimlcq.modifiers) > 0 {
		_spec.Modifiers = bimlcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bimlcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bimlcq *BillingInvoiceManualLineConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bimlcq.querySpec()
	if len(bimlcq.modifiers) > 0 {
		_spec.Modifiers = bimlcq.modifiers
	}
	_spec.Node.Columns = bimlcq.ctx.Fields
	if len(bimlcq.ctx.Fields) > 0 {
		_spec.Unique = bimlcq.ctx.Unique != nil && *bimlcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bimlcq.driver, _spec)
}

func (bimlcq *BillingInvoiceManualLineConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(billinginvoicemanuallineconfig.Table, billinginvoicemanuallineconfig.Columns, sqlgraph.NewFieldSpec(billinginvoicemanuallineconfig.FieldID, field.TypeString))
	_spec.From = bimlcq.sql
	if unique := bimlcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bimlcq.path != nil {
		_spec.Unique = true
	}
	if fields := bimlcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, billinginvoicemanuallineconfig.FieldID)
		for i := range fields {
			if fields[i] != billinginvoicemanuallineconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bimlcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bimlcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bimlcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bimlcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bimlcq *BillingInvoiceManualLineConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bimlcq.driver.Dialect())
	t1 := builder.Table(billinginvoicemanuallineconfig.Table)
	columns := bimlcq.ctx.Fields
	if len(columns) == 0 {
		columns = billinginvoicemanuallineconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bimlcq.sql != nil {
		selector = bimlcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bimlcq.ctx.Unique != nil && *bimlcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range bimlcq.modifiers {
		m(selector)
	}
	for _, p := range bimlcq.predicates {
		p(selector)
	}
	for _, p := range bimlcq.order {
		p(selector)
	}
	if offset := bimlcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bimlcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (bimlcq *BillingInvoiceManualLineConfigQuery) ForUpdate(opts ...sql.LockOption) *BillingInvoiceManualLineConfigQuery {
	if bimlcq.driver.Dialect() == dialect.Postgres {
		bimlcq.Unique(false)
	}
	bimlcq.modifiers = append(bimlcq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return bimlcq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (bimlcq *BillingInvoiceManualLineConfigQuery) ForShare(opts ...sql.LockOption) *BillingInvoiceManualLineConfigQuery {
	if bimlcq.driver.Dialect() == dialect.Postgres {
		bimlcq.Unique(false)
	}
	bimlcq.modifiers = append(bimlcq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return bimlcq
}

// BillingInvoiceManualLineConfigGroupBy is the group-by builder for BillingInvoiceManualLineConfig entities.
type BillingInvoiceManualLineConfigGroupBy struct {
	selector
	build *BillingInvoiceManualLineConfigQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bimlcgb *BillingInvoiceManualLineConfigGroupBy) Aggregate(fns ...AggregateFunc) *BillingInvoiceManualLineConfigGroupBy {
	bimlcgb.fns = append(bimlcgb.fns, fns...)
	return bimlcgb
}

// Scan applies the selector query and scans the result into the given value.
func (bimlcgb *BillingInvoiceManualLineConfigGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bimlcgb.build.ctx, ent.OpQueryGroupBy)
	if err := bimlcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingInvoiceManualLineConfigQuery, *BillingInvoiceManualLineConfigGroupBy](ctx, bimlcgb.build, bimlcgb, bimlcgb.build.inters, v)
}

func (bimlcgb *BillingInvoiceManualLineConfigGroupBy) sqlScan(ctx context.Context, root *BillingInvoiceManualLineConfigQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bimlcgb.fns))
	for _, fn := range bimlcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bimlcgb.flds)+len(bimlcgb.fns))
		for _, f := range *bimlcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bimlcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bimlcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BillingInvoiceManualLineConfigSelect is the builder for selecting fields of BillingInvoiceManualLineConfig entities.
type BillingInvoiceManualLineConfigSelect struct {
	*BillingInvoiceManualLineConfigQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bimlcs *BillingInvoiceManualLineConfigSelect) Aggregate(fns ...AggregateFunc) *BillingInvoiceManualLineConfigSelect {
	bimlcs.fns = append(bimlcs.fns, fns...)
	return bimlcs
}

// Scan applies the selector query and scans the result into the given value.
func (bimlcs *BillingInvoiceManualLineConfigSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bimlcs.ctx, ent.OpQuerySelect)
	if err := bimlcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingInvoiceManualLineConfigQuery, *BillingInvoiceManualLineConfigSelect](ctx, bimlcs.BillingInvoiceManualLineConfigQuery, bimlcs, bimlcs.inters, v)
}

func (bimlcs *BillingInvoiceManualLineConfigSelect) sqlScan(ctx context.Context, root *BillingInvoiceManualLineConfigQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bimlcs.fns))
	for _, fn := range bimlcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bimlcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bimlcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
