// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/danced25519/linera-dapps/service/kline/zeus/pkg/db/ent/predicate"
	"github.com/danced25519/linera-dapps/service/kline/zeus/pkg/db/ent/tokenpair"
)

// TokenPairDelete is the builder for deleting a TokenPair entity.
type TokenPairDelete struct {
	config
	hooks    []Hook
	mutation *TokenPairMutation
}

// Where appends a list predicates to the TokenPairDelete builder.
func (tpd *TokenPairDelete) Where(ps ...predicate.TokenPair) *TokenPairDelete {
	tpd.mutation.Where(ps...)
	return tpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tpd *TokenPairDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tpd.hooks) == 0 {
		affected, err = tpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenPairMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tpd.mutation = mutation
			affected, err = tpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tpd.hooks) - 1; i >= 0; i-- {
			if tpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpd *TokenPairDelete) ExecX(ctx context.Context) int {
	n, err := tpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tpd *TokenPairDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tokenpair.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: tokenpair.FieldID,
			},
		},
	}
	if ps := tpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TokenPairDeleteOne is the builder for deleting a single TokenPair entity.
type TokenPairDeleteOne struct {
	tpd *TokenPairDelete
}

// Exec executes the deletion query.
func (tpdo *TokenPairDeleteOne) Exec(ctx context.Context) error {
	n, err := tpdo.tpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tokenpair.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tpdo *TokenPairDeleteOne) ExecX(ctx context.Context) {
	tpdo.tpd.ExecX(ctx)
}
