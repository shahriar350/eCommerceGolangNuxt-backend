// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/predicate"
	"bongo/ent/sellerproductimage"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SellerProductImageDelete is the builder for deleting a SellerProductImage entity.
type SellerProductImageDelete struct {
	config
	hooks    []Hook
	mutation *SellerProductImageMutation
}

// Where appends a list predicates to the SellerProductImageDelete builder.
func (spid *SellerProductImageDelete) Where(ps ...predicate.SellerProductImage) *SellerProductImageDelete {
	spid.mutation.Where(ps...)
	return spid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spid *SellerProductImageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(spid.hooks) == 0 {
		affected, err = spid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SellerProductImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spid.mutation = mutation
			affected, err = spid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spid.hooks) - 1; i >= 0; i-- {
			if spid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (spid *SellerProductImageDelete) ExecX(ctx context.Context) int {
	n, err := spid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spid *SellerProductImageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sellerproductimage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sellerproductimage.FieldID,
			},
		},
	}
	if ps := spid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, spid.driver, _spec)
}

// SellerProductImageDeleteOne is the builder for deleting a single SellerProductImage entity.
type SellerProductImageDeleteOne struct {
	spid *SellerProductImageDelete
}

// Exec executes the deletion query.
func (spido *SellerProductImageDeleteOne) Exec(ctx context.Context) error {
	n, err := spido.spid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sellerproductimage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spido *SellerProductImageDeleteOne) ExecX(ctx context.Context) {
	spido.spid.ExecX(ctx)
}
