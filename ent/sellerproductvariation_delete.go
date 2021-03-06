// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/predicate"
	"bongo/ent/sellerproductvariation"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SellerProductVariationDelete is the builder for deleting a SellerProductVariation entity.
type SellerProductVariationDelete struct {
	config
	hooks    []Hook
	mutation *SellerProductVariationMutation
}

// Where appends a list predicates to the SellerProductVariationDelete builder.
func (spvd *SellerProductVariationDelete) Where(ps ...predicate.SellerProductVariation) *SellerProductVariationDelete {
	spvd.mutation.Where(ps...)
	return spvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spvd *SellerProductVariationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(spvd.hooks) == 0 {
		affected, err = spvd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SellerProductVariationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spvd.mutation = mutation
			affected, err = spvd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spvd.hooks) - 1; i >= 0; i-- {
			if spvd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spvd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spvd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (spvd *SellerProductVariationDelete) ExecX(ctx context.Context) int {
	n, err := spvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spvd *SellerProductVariationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sellerproductvariation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sellerproductvariation.FieldID,
			},
		},
	}
	if ps := spvd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, spvd.driver, _spec)
}

// SellerProductVariationDeleteOne is the builder for deleting a single SellerProductVariation entity.
type SellerProductVariationDeleteOne struct {
	spvd *SellerProductVariationDelete
}

// Exec executes the deletion query.
func (spvdo *SellerProductVariationDeleteOne) Exec(ctx context.Context) error {
	n, err := spvdo.spvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sellerproductvariation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spvdo *SellerProductVariationDeleteOne) ExecX(ctx context.Context) {
	spvdo.spvd.ExecX(ctx)
}
