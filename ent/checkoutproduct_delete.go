// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/checkoutproduct"
	"bongo/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CheckoutProductDelete is the builder for deleting a CheckoutProduct entity.
type CheckoutProductDelete struct {
	config
	hooks    []Hook
	mutation *CheckoutProductMutation
}

// Where appends a list predicates to the CheckoutProductDelete builder.
func (cpd *CheckoutProductDelete) Where(ps ...predicate.CheckoutProduct) *CheckoutProductDelete {
	cpd.mutation.Where(ps...)
	return cpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cpd *CheckoutProductDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cpd.hooks) == 0 {
		affected, err = cpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CheckoutProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cpd.mutation = mutation
			affected, err = cpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cpd.hooks) - 1; i >= 0; i-- {
			if cpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpd *CheckoutProductDelete) ExecX(ctx context.Context) int {
	n, err := cpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cpd *CheckoutProductDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: checkoutproduct.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checkoutproduct.FieldID,
			},
		},
	}
	if ps := cpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cpd.driver, _spec)
}

// CheckoutProductDeleteOne is the builder for deleting a single CheckoutProduct entity.
type CheckoutProductDeleteOne struct {
	cpd *CheckoutProductDelete
}

// Exec executes the deletion query.
func (cpdo *CheckoutProductDeleteOne) Exec(ctx context.Context) error {
	n, err := cpdo.cpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{checkoutproduct.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cpdo *CheckoutProductDeleteOne) ExecX(ctx context.Context) {
	cpdo.cpd.ExecX(ctx)
}
