// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Geapefurit/kline-back/zeus/pkg/db/ent/kpoint"
	"github.com/Geapefurit/kline-back/zeus/pkg/db/ent/predicate"
)

// KPointQuery is the builder for querying KPoint entities.
type KPointQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.KPoint
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KPointQuery builder.
func (kq *KPointQuery) Where(ps ...predicate.KPoint) *KPointQuery {
	kq.predicates = append(kq.predicates, ps...)
	return kq
}

// Limit adds a limit step to the query.
func (kq *KPointQuery) Limit(limit int) *KPointQuery {
	kq.limit = &limit
	return kq
}

// Offset adds an offset step to the query.
func (kq *KPointQuery) Offset(offset int) *KPointQuery {
	kq.offset = &offset
	return kq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kq *KPointQuery) Unique(unique bool) *KPointQuery {
	kq.unique = &unique
	return kq
}

// Order adds an order step to the query.
func (kq *KPointQuery) Order(o ...OrderFunc) *KPointQuery {
	kq.order = append(kq.order, o...)
	return kq
}

// First returns the first KPoint entity from the query.
// Returns a *NotFoundError when no KPoint was found.
func (kq *KPointQuery) First(ctx context.Context) (*KPoint, error) {
	nodes, err := kq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{kpoint.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kq *KPointQuery) FirstX(ctx context.Context) *KPoint {
	node, err := kq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first KPoint ID from the query.
// Returns a *NotFoundError when no KPoint ID was found.
func (kq *KPointQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = kq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{kpoint.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kq *KPointQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := kq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single KPoint entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one KPoint entity is found.
// Returns a *NotFoundError when no KPoint entities are found.
func (kq *KPointQuery) Only(ctx context.Context) (*KPoint, error) {
	nodes, err := kq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{kpoint.Label}
	default:
		return nil, &NotSingularError{kpoint.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kq *KPointQuery) OnlyX(ctx context.Context) *KPoint {
	node, err := kq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only KPoint ID in the query.
// Returns a *NotSingularError when more than one KPoint ID is found.
// Returns a *NotFoundError when no entities are found.
func (kq *KPointQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = kq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{kpoint.Label}
	default:
		err = &NotSingularError{kpoint.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kq *KPointQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := kq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KPoints.
func (kq *KPointQuery) All(ctx context.Context) ([]*KPoint, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return kq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (kq *KPointQuery) AllX(ctx context.Context) []*KPoint {
	nodes, err := kq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of KPoint IDs.
func (kq *KPointQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := kq.Select(kpoint.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kq *KPointQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := kq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kq *KPointQuery) Count(ctx context.Context) (int, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return kq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (kq *KPointQuery) CountX(ctx context.Context) int {
	count, err := kq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kq *KPointQuery) Exist(ctx context.Context) (bool, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return kq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (kq *KPointQuery) ExistX(ctx context.Context) bool {
	exist, err := kq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KPointQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kq *KPointQuery) Clone() *KPointQuery {
	if kq == nil {
		return nil
	}
	return &KPointQuery{
		config:     kq.config,
		limit:      kq.limit,
		offset:     kq.offset,
		order:      append([]OrderFunc{}, kq.order...),
		predicates: append([]predicate.KPoint{}, kq.predicates...),
		// clone intermediate query.
		sql:    kq.sql.Clone(),
		path:   kq.path,
		unique: kq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.KPoint.Query().
//		GroupBy(kpoint.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (kq *KPointQuery) GroupBy(field string, fields ...string) *KPointGroupBy {
	grbuild := &KPointGroupBy{config: kq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kq.sqlQuery(ctx), nil
	}
	grbuild.label = kpoint.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.KPoint.Query().
//		Select(kpoint.FieldCreatedAt).
//		Scan(ctx, &v)
func (kq *KPointQuery) Select(fields ...string) *KPointSelect {
	kq.fields = append(kq.fields, fields...)
	selbuild := &KPointSelect{KPointQuery: kq}
	selbuild.label = kpoint.Label
	selbuild.flds, selbuild.scan = &kq.fields, selbuild.Scan
	return selbuild
}

func (kq *KPointQuery) prepareQuery(ctx context.Context) error {
	for _, f := range kq.fields {
		if !kpoint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if kq.path != nil {
		prev, err := kq.path(ctx)
		if err != nil {
			return err
		}
		kq.sql = prev
	}
	if kpoint.Policy == nil {
		return errors.New("ent: uninitialized kpoint.Policy (forgotten import ent/runtime?)")
	}
	if err := kpoint.Policy.EvalQuery(ctx, kq); err != nil {
		return err
	}
	return nil
}

func (kq *KPointQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*KPoint, error) {
	var (
		nodes = []*KPoint{}
		_spec = kq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*KPoint).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &KPoint{config: kq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(kq.modifiers) > 0 {
		_spec.Modifiers = kq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, kq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (kq *KPointQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kq.querySpec()
	if len(kq.modifiers) > 0 {
		_spec.Modifiers = kq.modifiers
	}
	_spec.Node.Columns = kq.fields
	if len(kq.fields) > 0 {
		_spec.Unique = kq.unique != nil && *kq.unique
	}
	return sqlgraph.CountNodes(ctx, kq.driver, _spec)
}

func (kq *KPointQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := kq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (kq *KPointQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kpoint.Table,
			Columns: kpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: kpoint.FieldID,
			},
		},
		From:   kq.sql,
		Unique: true,
	}
	if unique := kq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := kq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kpoint.FieldID)
		for i := range fields {
			if fields[i] != kpoint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kq *KPointQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kq.driver.Dialect())
	t1 := builder.Table(kpoint.Table)
	columns := kq.fields
	if len(columns) == 0 {
		columns = kpoint.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if kq.sql != nil {
		selector = kq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if kq.unique != nil && *kq.unique {
		selector.Distinct()
	}
	for _, m := range kq.modifiers {
		m(selector)
	}
	for _, p := range kq.predicates {
		p(selector)
	}
	for _, p := range kq.order {
		p(selector)
	}
	if offset := kq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (kq *KPointQuery) ForUpdate(opts ...sql.LockOption) *KPointQuery {
	if kq.driver.Dialect() == dialect.Postgres {
		kq.Unique(false)
	}
	kq.modifiers = append(kq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return kq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (kq *KPointQuery) ForShare(opts ...sql.LockOption) *KPointQuery {
	if kq.driver.Dialect() == dialect.Postgres {
		kq.Unique(false)
	}
	kq.modifiers = append(kq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return kq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (kq *KPointQuery) Modify(modifiers ...func(s *sql.Selector)) *KPointSelect {
	kq.modifiers = append(kq.modifiers, modifiers...)
	return kq.Select()
}

// KPointGroupBy is the group-by builder for KPoint entities.
type KPointGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kgb *KPointGroupBy) Aggregate(fns ...AggregateFunc) *KPointGroupBy {
	kgb.fns = append(kgb.fns, fns...)
	return kgb
}

// Scan applies the group-by query and scans the result into the given value.
func (kgb *KPointGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := kgb.path(ctx)
	if err != nil {
		return err
	}
	kgb.sql = query
	return kgb.sqlScan(ctx, v)
}

func (kgb *KPointGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range kgb.fields {
		if !kpoint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := kgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kgb *KPointGroupBy) sqlQuery() *sql.Selector {
	selector := kgb.sql.Select()
	aggregation := make([]string, 0, len(kgb.fns))
	for _, fn := range kgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(kgb.fields)+len(kgb.fns))
		for _, f := range kgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(kgb.fields...)...)
}

// KPointSelect is the builder for selecting fields of KPoint entities.
type KPointSelect struct {
	*KPointQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ks *KPointSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ks.prepareQuery(ctx); err != nil {
		return err
	}
	ks.sql = ks.KPointQuery.sqlQuery(ctx)
	return ks.sqlScan(ctx, v)
}

func (ks *KPointSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ks.sql.Query()
	if err := ks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ks *KPointSelect) Modify(modifiers ...func(s *sql.Selector)) *KPointSelect {
	ks.modifiers = append(ks.modifiers, modifiers...)
	return ks
}