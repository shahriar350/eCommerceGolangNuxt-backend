// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/predicate"
	"bongo/ent/shopcategory"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShopCategoryDelete is the builder for deleting a ShopCategory entity.
type ShopCategoryDelete struct {
	config
	hooks    []Hook
	mutation *ShopCategoryMutation
}

// Where appends a list predicates to the ShopCategoryDelete builder.
func (scd *ShopCategoryDelete) Where(ps ...predicate.ShopCategory) *ShopCategoryDelete {
	scd.mutation.Where(ps...)
	return scd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scd *ShopCategoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(scd.hooks) == 0 {
		affected, err = scd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShopCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			scd.mutation = mutation
			affected, err = scd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(scd.hooks) - 1; i >= 0; i-- {
			if scd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (scd *ShopCategoryDelete) ExecX(ctx context.Context) int {
	n, err := scd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scd *ShopCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: shopcategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shopcategory.FieldID,
			},
		},
	}
	if ps := scd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, scd.driver, _spec)
}

// ShopCategoryDeleteOne is the builder for deleting a single ShopCategory entity.
type ShopCategoryDeleteOne struct {
	scd *ShopCategoryDelete
}

// Exec executes the deletion query.
func (scdo *ShopCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := scdo.scd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{shopcategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scdo *ShopCategoryDeleteOne) ExecX(ctx context.Context) {
	scdo.scd.ExecX(ctx)
}
