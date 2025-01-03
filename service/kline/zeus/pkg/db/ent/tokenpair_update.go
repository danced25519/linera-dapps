// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/linera-hacker/linera-dapps/service/kline/zeus/pkg/db/ent/predicate"
	"github.com/linera-hacker/linera-dapps/service/kline/zeus/pkg/db/ent/tokenpair"
)

// TokenPairUpdate is the builder for updating TokenPair entities.
type TokenPairUpdate struct {
	config
	hooks     []Hook
	mutation  *TokenPairMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TokenPairUpdate builder.
func (tpu *TokenPairUpdate) Where(ps ...predicate.TokenPair) *TokenPairUpdate {
	tpu.mutation.Where(ps...)
	return tpu
}

// SetCreatedAt sets the "created_at" field.
func (tpu *TokenPairUpdate) SetCreatedAt(u uint32) *TokenPairUpdate {
	tpu.mutation.ResetCreatedAt()
	tpu.mutation.SetCreatedAt(u)
	return tpu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tpu *TokenPairUpdate) SetNillableCreatedAt(u *uint32) *TokenPairUpdate {
	if u != nil {
		tpu.SetCreatedAt(*u)
	}
	return tpu
}

// AddCreatedAt adds u to the "created_at" field.
func (tpu *TokenPairUpdate) AddCreatedAt(u int32) *TokenPairUpdate {
	tpu.mutation.AddCreatedAt(u)
	return tpu
}

// SetUpdatedAt sets the "updated_at" field.
func (tpu *TokenPairUpdate) SetUpdatedAt(u uint32) *TokenPairUpdate {
	tpu.mutation.ResetUpdatedAt()
	tpu.mutation.SetUpdatedAt(u)
	return tpu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tpu *TokenPairUpdate) AddUpdatedAt(u int32) *TokenPairUpdate {
	tpu.mutation.AddUpdatedAt(u)
	return tpu
}

// SetDeletedAt sets the "deleted_at" field.
func (tpu *TokenPairUpdate) SetDeletedAt(u uint32) *TokenPairUpdate {
	tpu.mutation.ResetDeletedAt()
	tpu.mutation.SetDeletedAt(u)
	return tpu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tpu *TokenPairUpdate) SetNillableDeletedAt(u *uint32) *TokenPairUpdate {
	if u != nil {
		tpu.SetDeletedAt(*u)
	}
	return tpu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tpu *TokenPairUpdate) AddDeletedAt(u int32) *TokenPairUpdate {
	tpu.mutation.AddDeletedAt(u)
	return tpu
}

// SetPoolID sets the "pool_id" field.
func (tpu *TokenPairUpdate) SetPoolID(u uint64) *TokenPairUpdate {
	tpu.mutation.ResetPoolID()
	tpu.mutation.SetPoolID(u)
	return tpu
}

// AddPoolID adds u to the "pool_id" field.
func (tpu *TokenPairUpdate) AddPoolID(u int64) *TokenPairUpdate {
	tpu.mutation.AddPoolID(u)
	return tpu
}

// SetTokenZeroID sets the "token_zero_id" field.
func (tpu *TokenPairUpdate) SetTokenZeroID(u uint32) *TokenPairUpdate {
	tpu.mutation.ResetTokenZeroID()
	tpu.mutation.SetTokenZeroID(u)
	return tpu
}

// AddTokenZeroID adds u to the "token_zero_id" field.
func (tpu *TokenPairUpdate) AddTokenZeroID(u int32) *TokenPairUpdate {
	tpu.mutation.AddTokenZeroID(u)
	return tpu
}

// SetTokenOneID sets the "token_one_id" field.
func (tpu *TokenPairUpdate) SetTokenOneID(u uint32) *TokenPairUpdate {
	tpu.mutation.ResetTokenOneID()
	tpu.mutation.SetTokenOneID(u)
	return tpu
}

// AddTokenOneID adds u to the "token_one_id" field.
func (tpu *TokenPairUpdate) AddTokenOneID(u int32) *TokenPairUpdate {
	tpu.mutation.AddTokenOneID(u)
	return tpu
}

// SetRemark sets the "remark" field.
func (tpu *TokenPairUpdate) SetRemark(s string) *TokenPairUpdate {
	tpu.mutation.SetRemark(s)
	return tpu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (tpu *TokenPairUpdate) SetNillableRemark(s *string) *TokenPairUpdate {
	if s != nil {
		tpu.SetRemark(*s)
	}
	return tpu
}

// ClearRemark clears the value of the "remark" field.
func (tpu *TokenPairUpdate) ClearRemark() *TokenPairUpdate {
	tpu.mutation.ClearRemark()
	return tpu
}

// Mutation returns the TokenPairMutation object of the builder.
func (tpu *TokenPairUpdate) Mutation() *TokenPairMutation {
	return tpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tpu *TokenPairUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := tpu.defaults(); err != nil {
		return 0, err
	}
	if len(tpu.hooks) == 0 {
		affected, err = tpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenPairMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tpu.mutation = mutation
			affected, err = tpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tpu.hooks) - 1; i >= 0; i-- {
			if tpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tpu *TokenPairUpdate) SaveX(ctx context.Context) int {
	affected, err := tpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tpu *TokenPairUpdate) Exec(ctx context.Context) error {
	_, err := tpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpu *TokenPairUpdate) ExecX(ctx context.Context) {
	if err := tpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tpu *TokenPairUpdate) defaults() error {
	if _, ok := tpu.mutation.UpdatedAt(); !ok {
		if tokenpair.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized tokenpair.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := tokenpair.UpdateDefaultUpdatedAt()
		tpu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tpu *TokenPairUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TokenPairUpdate {
	tpu.modifiers = append(tpu.modifiers, modifiers...)
	return tpu
}

func (tpu *TokenPairUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tokenpair.Table,
			Columns: tokenpair.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: tokenpair.FieldID,
			},
		},
	}
	if ps := tpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tpu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldCreatedAt,
		})
	}
	if value, ok := tpu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldCreatedAt,
		})
	}
	if value, ok := tpu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldUpdatedAt,
		})
	}
	if value, ok := tpu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldUpdatedAt,
		})
	}
	if value, ok := tpu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldDeletedAt,
		})
	}
	if value, ok := tpu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldDeletedAt,
		})
	}
	if value, ok := tpu.mutation.PoolID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: tokenpair.FieldPoolID,
		})
	}
	if value, ok := tpu.mutation.AddedPoolID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: tokenpair.FieldPoolID,
		})
	}
	if value, ok := tpu.mutation.TokenZeroID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldTokenZeroID,
		})
	}
	if value, ok := tpu.mutation.AddedTokenZeroID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldTokenZeroID,
		})
	}
	if value, ok := tpu.mutation.TokenOneID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldTokenOneID,
		})
	}
	if value, ok := tpu.mutation.AddedTokenOneID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldTokenOneID,
		})
	}
	if value, ok := tpu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tokenpair.FieldRemark,
		})
	}
	if tpu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tokenpair.FieldRemark,
		})
	}
	_spec.Modifiers = tpu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, tpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tokenpair.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TokenPairUpdateOne is the builder for updating a single TokenPair entity.
type TokenPairUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TokenPairMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (tpuo *TokenPairUpdateOne) SetCreatedAt(u uint32) *TokenPairUpdateOne {
	tpuo.mutation.ResetCreatedAt()
	tpuo.mutation.SetCreatedAt(u)
	return tpuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tpuo *TokenPairUpdateOne) SetNillableCreatedAt(u *uint32) *TokenPairUpdateOne {
	if u != nil {
		tpuo.SetCreatedAt(*u)
	}
	return tpuo
}

// AddCreatedAt adds u to the "created_at" field.
func (tpuo *TokenPairUpdateOne) AddCreatedAt(u int32) *TokenPairUpdateOne {
	tpuo.mutation.AddCreatedAt(u)
	return tpuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tpuo *TokenPairUpdateOne) SetUpdatedAt(u uint32) *TokenPairUpdateOne {
	tpuo.mutation.ResetUpdatedAt()
	tpuo.mutation.SetUpdatedAt(u)
	return tpuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tpuo *TokenPairUpdateOne) AddUpdatedAt(u int32) *TokenPairUpdateOne {
	tpuo.mutation.AddUpdatedAt(u)
	return tpuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tpuo *TokenPairUpdateOne) SetDeletedAt(u uint32) *TokenPairUpdateOne {
	tpuo.mutation.ResetDeletedAt()
	tpuo.mutation.SetDeletedAt(u)
	return tpuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tpuo *TokenPairUpdateOne) SetNillableDeletedAt(u *uint32) *TokenPairUpdateOne {
	if u != nil {
		tpuo.SetDeletedAt(*u)
	}
	return tpuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tpuo *TokenPairUpdateOne) AddDeletedAt(u int32) *TokenPairUpdateOne {
	tpuo.mutation.AddDeletedAt(u)
	return tpuo
}

// SetPoolID sets the "pool_id" field.
func (tpuo *TokenPairUpdateOne) SetPoolID(u uint64) *TokenPairUpdateOne {
	tpuo.mutation.ResetPoolID()
	tpuo.mutation.SetPoolID(u)
	return tpuo
}

// AddPoolID adds u to the "pool_id" field.
func (tpuo *TokenPairUpdateOne) AddPoolID(u int64) *TokenPairUpdateOne {
	tpuo.mutation.AddPoolID(u)
	return tpuo
}

// SetTokenZeroID sets the "token_zero_id" field.
func (tpuo *TokenPairUpdateOne) SetTokenZeroID(u uint32) *TokenPairUpdateOne {
	tpuo.mutation.ResetTokenZeroID()
	tpuo.mutation.SetTokenZeroID(u)
	return tpuo
}

// AddTokenZeroID adds u to the "token_zero_id" field.
func (tpuo *TokenPairUpdateOne) AddTokenZeroID(u int32) *TokenPairUpdateOne {
	tpuo.mutation.AddTokenZeroID(u)
	return tpuo
}

// SetTokenOneID sets the "token_one_id" field.
func (tpuo *TokenPairUpdateOne) SetTokenOneID(u uint32) *TokenPairUpdateOne {
	tpuo.mutation.ResetTokenOneID()
	tpuo.mutation.SetTokenOneID(u)
	return tpuo
}

// AddTokenOneID adds u to the "token_one_id" field.
func (tpuo *TokenPairUpdateOne) AddTokenOneID(u int32) *TokenPairUpdateOne {
	tpuo.mutation.AddTokenOneID(u)
	return tpuo
}

// SetRemark sets the "remark" field.
func (tpuo *TokenPairUpdateOne) SetRemark(s string) *TokenPairUpdateOne {
	tpuo.mutation.SetRemark(s)
	return tpuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (tpuo *TokenPairUpdateOne) SetNillableRemark(s *string) *TokenPairUpdateOne {
	if s != nil {
		tpuo.SetRemark(*s)
	}
	return tpuo
}

// ClearRemark clears the value of the "remark" field.
func (tpuo *TokenPairUpdateOne) ClearRemark() *TokenPairUpdateOne {
	tpuo.mutation.ClearRemark()
	return tpuo
}

// Mutation returns the TokenPairMutation object of the builder.
func (tpuo *TokenPairUpdateOne) Mutation() *TokenPairMutation {
	return tpuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tpuo *TokenPairUpdateOne) Select(field string, fields ...string) *TokenPairUpdateOne {
	tpuo.fields = append([]string{field}, fields...)
	return tpuo
}

// Save executes the query and returns the updated TokenPair entity.
func (tpuo *TokenPairUpdateOne) Save(ctx context.Context) (*TokenPair, error) {
	var (
		err  error
		node *TokenPair
	)
	if err := tpuo.defaults(); err != nil {
		return nil, err
	}
	if len(tpuo.hooks) == 0 {
		node, err = tpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenPairMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tpuo.mutation = mutation
			node, err = tpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tpuo.hooks) - 1; i >= 0; i-- {
			if tpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tpuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tpuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TokenPair)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TokenPairMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tpuo *TokenPairUpdateOne) SaveX(ctx context.Context) *TokenPair {
	node, err := tpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tpuo *TokenPairUpdateOne) Exec(ctx context.Context) error {
	_, err := tpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpuo *TokenPairUpdateOne) ExecX(ctx context.Context) {
	if err := tpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tpuo *TokenPairUpdateOne) defaults() error {
	if _, ok := tpuo.mutation.UpdatedAt(); !ok {
		if tokenpair.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized tokenpair.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := tokenpair.UpdateDefaultUpdatedAt()
		tpuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tpuo *TokenPairUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TokenPairUpdateOne {
	tpuo.modifiers = append(tpuo.modifiers, modifiers...)
	return tpuo
}

func (tpuo *TokenPairUpdateOne) sqlSave(ctx context.Context) (_node *TokenPair, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tokenpair.Table,
			Columns: tokenpair.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: tokenpair.FieldID,
			},
		},
	}
	id, ok := tpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TokenPair.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tokenpair.FieldID)
		for _, f := range fields {
			if !tokenpair.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tokenpair.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tpuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldCreatedAt,
		})
	}
	if value, ok := tpuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldCreatedAt,
		})
	}
	if value, ok := tpuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldUpdatedAt,
		})
	}
	if value, ok := tpuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldUpdatedAt,
		})
	}
	if value, ok := tpuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldDeletedAt,
		})
	}
	if value, ok := tpuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldDeletedAt,
		})
	}
	if value, ok := tpuo.mutation.PoolID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: tokenpair.FieldPoolID,
		})
	}
	if value, ok := tpuo.mutation.AddedPoolID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: tokenpair.FieldPoolID,
		})
	}
	if value, ok := tpuo.mutation.TokenZeroID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldTokenZeroID,
		})
	}
	if value, ok := tpuo.mutation.AddedTokenZeroID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldTokenZeroID,
		})
	}
	if value, ok := tpuo.mutation.TokenOneID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldTokenOneID,
		})
	}
	if value, ok := tpuo.mutation.AddedTokenOneID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldTokenOneID,
		})
	}
	if value, ok := tpuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tokenpair.FieldRemark,
		})
	}
	if tpuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tokenpair.FieldRemark,
		})
	}
	_spec.Modifiers = tpuo.modifiers
	_node = &TokenPair{config: tpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tokenpair.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
