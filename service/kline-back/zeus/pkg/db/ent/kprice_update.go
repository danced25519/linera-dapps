// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Geapefurit/kline-back/zeus/pkg/db/ent/kprice"
	"github.com/Geapefurit/kline-back/zeus/pkg/db/ent/predicate"
)

// KPriceUpdate is the builder for updating KPrice entities.
type KPriceUpdate struct {
	config
	hooks     []Hook
	mutation  *KPriceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the KPriceUpdate builder.
func (ku *KPriceUpdate) Where(ps ...predicate.KPrice) *KPriceUpdate {
	ku.mutation.Where(ps...)
	return ku
}

// SetCreatedAt sets the "created_at" field.
func (ku *KPriceUpdate) SetCreatedAt(u uint32) *KPriceUpdate {
	ku.mutation.ResetCreatedAt()
	ku.mutation.SetCreatedAt(u)
	return ku
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ku *KPriceUpdate) SetNillableCreatedAt(u *uint32) *KPriceUpdate {
	if u != nil {
		ku.SetCreatedAt(*u)
	}
	return ku
}

// AddCreatedAt adds u to the "created_at" field.
func (ku *KPriceUpdate) AddCreatedAt(u int32) *KPriceUpdate {
	ku.mutation.AddCreatedAt(u)
	return ku
}

// SetUpdatedAt sets the "updated_at" field.
func (ku *KPriceUpdate) SetUpdatedAt(u uint32) *KPriceUpdate {
	ku.mutation.ResetUpdatedAt()
	ku.mutation.SetUpdatedAt(u)
	return ku
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ku *KPriceUpdate) AddUpdatedAt(u int32) *KPriceUpdate {
	ku.mutation.AddUpdatedAt(u)
	return ku
}

// SetDeletedAt sets the "deleted_at" field.
func (ku *KPriceUpdate) SetDeletedAt(u uint32) *KPriceUpdate {
	ku.mutation.ResetDeletedAt()
	ku.mutation.SetDeletedAt(u)
	return ku
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ku *KPriceUpdate) SetNillableDeletedAt(u *uint32) *KPriceUpdate {
	if u != nil {
		ku.SetDeletedAt(*u)
	}
	return ku
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ku *KPriceUpdate) AddDeletedAt(u int32) *KPriceUpdate {
	ku.mutation.AddDeletedAt(u)
	return ku
}

// SetTokenPairID sets the "token_pair_id" field.
func (ku *KPriceUpdate) SetTokenPairID(u uint32) *KPriceUpdate {
	ku.mutation.ResetTokenPairID()
	ku.mutation.SetTokenPairID(u)
	return ku
}

// AddTokenPairID adds u to the "token_pair_id" field.
func (ku *KPriceUpdate) AddTokenPairID(u int32) *KPriceUpdate {
	ku.mutation.AddTokenPairID(u)
	return ku
}

// SetPrice sets the "price" field.
func (ku *KPriceUpdate) SetPrice(f float64) *KPriceUpdate {
	ku.mutation.ResetPrice()
	ku.mutation.SetPrice(f)
	return ku
}

// AddPrice adds f to the "price" field.
func (ku *KPriceUpdate) AddPrice(f float64) *KPriceUpdate {
	ku.mutation.AddPrice(f)
	return ku
}

// SetTimestamp sets the "timestamp" field.
func (ku *KPriceUpdate) SetTimestamp(u uint32) *KPriceUpdate {
	ku.mutation.ResetTimestamp()
	ku.mutation.SetTimestamp(u)
	return ku
}

// AddTimestamp adds u to the "timestamp" field.
func (ku *KPriceUpdate) AddTimestamp(u int32) *KPriceUpdate {
	ku.mutation.AddTimestamp(u)
	return ku
}

// Mutation returns the KPriceMutation object of the builder.
func (ku *KPriceUpdate) Mutation() *KPriceMutation {
	return ku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ku *KPriceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ku.defaults(); err != nil {
		return 0, err
	}
	if len(ku.hooks) == 0 {
		affected, err = ku.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KPriceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ku.mutation = mutation
			affected, err = ku.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ku.hooks) - 1; i >= 0; i-- {
			if ku.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ku.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ku.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ku *KPriceUpdate) SaveX(ctx context.Context) int {
	affected, err := ku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ku *KPriceUpdate) Exec(ctx context.Context) error {
	_, err := ku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ku *KPriceUpdate) ExecX(ctx context.Context) {
	if err := ku.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ku *KPriceUpdate) defaults() error {
	if _, ok := ku.mutation.UpdatedAt(); !ok {
		if kprice.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized kprice.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := kprice.UpdateDefaultUpdatedAt()
		ku.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ku *KPriceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *KPriceUpdate {
	ku.modifiers = append(ku.modifiers, modifiers...)
	return ku
}

func (ku *KPriceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kprice.Table,
			Columns: kprice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: kprice.FieldID,
			},
		},
	}
	if ps := ku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ku.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldCreatedAt,
		})
	}
	if value, ok := ku.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldCreatedAt,
		})
	}
	if value, ok := ku.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldUpdatedAt,
		})
	}
	if value, ok := ku.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldUpdatedAt,
		})
	}
	if value, ok := ku.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldDeletedAt,
		})
	}
	if value, ok := ku.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldDeletedAt,
		})
	}
	if value, ok := ku.mutation.TokenPairID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldTokenPairID,
		})
	}
	if value, ok := ku.mutation.AddedTokenPairID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldTokenPairID,
		})
	}
	if value, ok := ku.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: kprice.FieldPrice,
		})
	}
	if value, ok := ku.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: kprice.FieldPrice,
		})
	}
	if value, ok := ku.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldTimestamp,
		})
	}
	if value, ok := ku.mutation.AddedTimestamp(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldTimestamp,
		})
	}
	_spec.Modifiers = ku.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kprice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// KPriceUpdateOne is the builder for updating a single KPrice entity.
type KPriceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *KPriceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (kuo *KPriceUpdateOne) SetCreatedAt(u uint32) *KPriceUpdateOne {
	kuo.mutation.ResetCreatedAt()
	kuo.mutation.SetCreatedAt(u)
	return kuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kuo *KPriceUpdateOne) SetNillableCreatedAt(u *uint32) *KPriceUpdateOne {
	if u != nil {
		kuo.SetCreatedAt(*u)
	}
	return kuo
}

// AddCreatedAt adds u to the "created_at" field.
func (kuo *KPriceUpdateOne) AddCreatedAt(u int32) *KPriceUpdateOne {
	kuo.mutation.AddCreatedAt(u)
	return kuo
}

// SetUpdatedAt sets the "updated_at" field.
func (kuo *KPriceUpdateOne) SetUpdatedAt(u uint32) *KPriceUpdateOne {
	kuo.mutation.ResetUpdatedAt()
	kuo.mutation.SetUpdatedAt(u)
	return kuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (kuo *KPriceUpdateOne) AddUpdatedAt(u int32) *KPriceUpdateOne {
	kuo.mutation.AddUpdatedAt(u)
	return kuo
}

// SetDeletedAt sets the "deleted_at" field.
func (kuo *KPriceUpdateOne) SetDeletedAt(u uint32) *KPriceUpdateOne {
	kuo.mutation.ResetDeletedAt()
	kuo.mutation.SetDeletedAt(u)
	return kuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kuo *KPriceUpdateOne) SetNillableDeletedAt(u *uint32) *KPriceUpdateOne {
	if u != nil {
		kuo.SetDeletedAt(*u)
	}
	return kuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (kuo *KPriceUpdateOne) AddDeletedAt(u int32) *KPriceUpdateOne {
	kuo.mutation.AddDeletedAt(u)
	return kuo
}

// SetTokenPairID sets the "token_pair_id" field.
func (kuo *KPriceUpdateOne) SetTokenPairID(u uint32) *KPriceUpdateOne {
	kuo.mutation.ResetTokenPairID()
	kuo.mutation.SetTokenPairID(u)
	return kuo
}

// AddTokenPairID adds u to the "token_pair_id" field.
func (kuo *KPriceUpdateOne) AddTokenPairID(u int32) *KPriceUpdateOne {
	kuo.mutation.AddTokenPairID(u)
	return kuo
}

// SetPrice sets the "price" field.
func (kuo *KPriceUpdateOne) SetPrice(f float64) *KPriceUpdateOne {
	kuo.mutation.ResetPrice()
	kuo.mutation.SetPrice(f)
	return kuo
}

// AddPrice adds f to the "price" field.
func (kuo *KPriceUpdateOne) AddPrice(f float64) *KPriceUpdateOne {
	kuo.mutation.AddPrice(f)
	return kuo
}

// SetTimestamp sets the "timestamp" field.
func (kuo *KPriceUpdateOne) SetTimestamp(u uint32) *KPriceUpdateOne {
	kuo.mutation.ResetTimestamp()
	kuo.mutation.SetTimestamp(u)
	return kuo
}

// AddTimestamp adds u to the "timestamp" field.
func (kuo *KPriceUpdateOne) AddTimestamp(u int32) *KPriceUpdateOne {
	kuo.mutation.AddTimestamp(u)
	return kuo
}

// Mutation returns the KPriceMutation object of the builder.
func (kuo *KPriceUpdateOne) Mutation() *KPriceMutation {
	return kuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kuo *KPriceUpdateOne) Select(field string, fields ...string) *KPriceUpdateOne {
	kuo.fields = append([]string{field}, fields...)
	return kuo
}

// Save executes the query and returns the updated KPrice entity.
func (kuo *KPriceUpdateOne) Save(ctx context.Context) (*KPrice, error) {
	var (
		err  error
		node *KPrice
	)
	if err := kuo.defaults(); err != nil {
		return nil, err
	}
	if len(kuo.hooks) == 0 {
		node, err = kuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KPriceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kuo.mutation = mutation
			node, err = kuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kuo.hooks) - 1; i >= 0; i-- {
			if kuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, kuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*KPrice)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from KPriceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kuo *KPriceUpdateOne) SaveX(ctx context.Context) *KPrice {
	node, err := kuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kuo *KPriceUpdateOne) Exec(ctx context.Context) error {
	_, err := kuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kuo *KPriceUpdateOne) ExecX(ctx context.Context) {
	if err := kuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kuo *KPriceUpdateOne) defaults() error {
	if _, ok := kuo.mutation.UpdatedAt(); !ok {
		if kprice.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized kprice.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := kprice.UpdateDefaultUpdatedAt()
		kuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (kuo *KPriceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *KPriceUpdateOne {
	kuo.modifiers = append(kuo.modifiers, modifiers...)
	return kuo
}

func (kuo *KPriceUpdateOne) sqlSave(ctx context.Context) (_node *KPrice, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kprice.Table,
			Columns: kprice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: kprice.FieldID,
			},
		},
	}
	id, ok := kuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "KPrice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := kuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kprice.FieldID)
		for _, f := range fields {
			if !kprice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != kprice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldCreatedAt,
		})
	}
	if value, ok := kuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldCreatedAt,
		})
	}
	if value, ok := kuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldUpdatedAt,
		})
	}
	if value, ok := kuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldUpdatedAt,
		})
	}
	if value, ok := kuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldDeletedAt,
		})
	}
	if value, ok := kuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldDeletedAt,
		})
	}
	if value, ok := kuo.mutation.TokenPairID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldTokenPairID,
		})
	}
	if value, ok := kuo.mutation.AddedTokenPairID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldTokenPairID,
		})
	}
	if value, ok := kuo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: kprice.FieldPrice,
		})
	}
	if value, ok := kuo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: kprice.FieldPrice,
		})
	}
	if value, ok := kuo.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldTimestamp,
		})
	}
	if value, ok := kuo.mutation.AddedTimestamp(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kprice.FieldTimestamp,
		})
	}
	_spec.Modifiers = kuo.modifiers
	_node = &KPrice{config: kuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kprice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
