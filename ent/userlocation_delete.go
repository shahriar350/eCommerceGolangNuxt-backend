// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/predicate"
	"bongo/ent/userlocation"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserLocationDelete is the builder for deleting a UserLocation entity.
type UserLocationDelete struct {
	config
	hooks    []Hook
	mutation *UserLocationMutation
}

// Where appends a list predicates to the UserLocationDelete builder.
func (uld *UserLocationDelete) Where(ps ...predicate.UserLocation) *UserLocationDelete {
	uld.mutation.Where(ps...)
	return uld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uld *UserLocationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uld.hooks) == 0 {
		affected, err = uld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserLocationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uld.mutation = mutation
			affected, err = uld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uld.hooks) - 1; i >= 0; i-- {
			if uld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (uld *UserLocationDelete) ExecX(ctx context.Context) int {
	n, err := uld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uld *UserLocationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userlocation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userlocation.FieldID,
			},
		},
	}
	if ps := uld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, uld.driver, _spec)
}

// UserLocationDeleteOne is the builder for deleting a single UserLocation entity.
type UserLocationDeleteOne struct {
	uld *UserLocationDelete
}

// Exec executes the deletion query.
func (uldo *UserLocationDeleteOne) Exec(ctx context.Context) error {
	n, err := uldo.uld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userlocation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uldo *UserLocationDeleteOne) ExecX(ctx context.Context) {
	uldo.uld.ExecX(ctx)
}
