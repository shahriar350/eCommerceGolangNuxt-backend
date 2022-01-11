// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/category"
	"bongo/ent/sellerproduct"
	"bongo/ent/sellerproductcategory"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SellerProductCategoryCreate is the builder for creating a SellerProductCategory entity.
type SellerProductCategoryCreate struct {
	config
	mutation *SellerProductCategoryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (spcc *SellerProductCategoryCreate) SetCreatedAt(t time.Time) *SellerProductCategoryCreate {
	spcc.mutation.SetCreatedAt(t)
	return spcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (spcc *SellerProductCategoryCreate) SetNillableCreatedAt(t *time.Time) *SellerProductCategoryCreate {
	if t != nil {
		spcc.SetCreatedAt(*t)
	}
	return spcc
}

// SetUpdatedAt sets the "updated_at" field.
func (spcc *SellerProductCategoryCreate) SetUpdatedAt(t time.Time) *SellerProductCategoryCreate {
	spcc.mutation.SetUpdatedAt(t)
	return spcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (spcc *SellerProductCategoryCreate) SetNillableUpdatedAt(t *time.Time) *SellerProductCategoryCreate {
	if t != nil {
		spcc.SetUpdatedAt(*t)
	}
	return spcc
}

// SetDeletedAt sets the "deleted_at" field.
func (spcc *SellerProductCategoryCreate) SetDeletedAt(t time.Time) *SellerProductCategoryCreate {
	spcc.mutation.SetDeletedAt(t)
	return spcc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (spcc *SellerProductCategoryCreate) SetNillableDeletedAt(t *time.Time) *SellerProductCategoryCreate {
	if t != nil {
		spcc.SetDeletedAt(*t)
	}
	return spcc
}

// SetSellerProductID sets the "seller_product" edge to the SellerProduct entity by ID.
func (spcc *SellerProductCategoryCreate) SetSellerProductID(id int) *SellerProductCategoryCreate {
	spcc.mutation.SetSellerProductID(id)
	return spcc
}

// SetNillableSellerProductID sets the "seller_product" edge to the SellerProduct entity by ID if the given value is not nil.
func (spcc *SellerProductCategoryCreate) SetNillableSellerProductID(id *int) *SellerProductCategoryCreate {
	if id != nil {
		spcc = spcc.SetSellerProductID(*id)
	}
	return spcc
}

// SetSellerProduct sets the "seller_product" edge to the SellerProduct entity.
func (spcc *SellerProductCategoryCreate) SetSellerProduct(s *SellerProduct) *SellerProductCategoryCreate {
	return spcc.SetSellerProductID(s.ID)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (spcc *SellerProductCategoryCreate) SetCategoryID(id int) *SellerProductCategoryCreate {
	spcc.mutation.SetCategoryID(id)
	return spcc
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (spcc *SellerProductCategoryCreate) SetNillableCategoryID(id *int) *SellerProductCategoryCreate {
	if id != nil {
		spcc = spcc.SetCategoryID(*id)
	}
	return spcc
}

// SetCategory sets the "category" edge to the Category entity.
func (spcc *SellerProductCategoryCreate) SetCategory(c *Category) *SellerProductCategoryCreate {
	return spcc.SetCategoryID(c.ID)
}

// Mutation returns the SellerProductCategoryMutation object of the builder.
func (spcc *SellerProductCategoryCreate) Mutation() *SellerProductCategoryMutation {
	return spcc.mutation
}

// Save creates the SellerProductCategory in the database.
func (spcc *SellerProductCategoryCreate) Save(ctx context.Context) (*SellerProductCategory, error) {
	var (
		err  error
		node *SellerProductCategory
	)
	spcc.defaults()
	if len(spcc.hooks) == 0 {
		if err = spcc.check(); err != nil {
			return nil, err
		}
		node, err = spcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SellerProductCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = spcc.check(); err != nil {
				return nil, err
			}
			spcc.mutation = mutation
			if node, err = spcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(spcc.hooks) - 1; i >= 0; i-- {
			if spcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (spcc *SellerProductCategoryCreate) SaveX(ctx context.Context) *SellerProductCategory {
	v, err := spcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spcc *SellerProductCategoryCreate) Exec(ctx context.Context) error {
	_, err := spcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spcc *SellerProductCategoryCreate) ExecX(ctx context.Context) {
	if err := spcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spcc *SellerProductCategoryCreate) defaults() {
	if _, ok := spcc.mutation.CreatedAt(); !ok {
		v := sellerproductcategory.DefaultCreatedAt()
		spcc.mutation.SetCreatedAt(v)
	}
	if _, ok := spcc.mutation.UpdatedAt(); !ok {
		v := sellerproductcategory.DefaultUpdatedAt()
		spcc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spcc *SellerProductCategoryCreate) check() error {
	if _, ok := spcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := spcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (spcc *SellerProductCategoryCreate) sqlSave(ctx context.Context) (*SellerProductCategory, error) {
	_node, _spec := spcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (spcc *SellerProductCategoryCreate) createSpec() (*SellerProductCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &SellerProductCategory{config: spcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sellerproductcategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sellerproductcategory.FieldID,
			},
		}
	)
	if value, ok := spcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductcategory.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := spcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductcategory.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := spcc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductcategory.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := spcc.mutation.SellerProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductcategory.SellerProductTable,
			Columns: []string{sellerproductcategory.SellerProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sellerproduct.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.seller_product_seller_product_categories = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := spcc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductcategory.CategoryTable,
			Columns: []string{sellerproductcategory.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.category_product_categories = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SellerProductCategoryCreateBulk is the builder for creating many SellerProductCategory entities in bulk.
type SellerProductCategoryCreateBulk struct {
	config
	builders []*SellerProductCategoryCreate
}

// Save creates the SellerProductCategory entities in the database.
func (spccb *SellerProductCategoryCreateBulk) Save(ctx context.Context) ([]*SellerProductCategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(spccb.builders))
	nodes := make([]*SellerProductCategory, len(spccb.builders))
	mutators := make([]Mutator, len(spccb.builders))
	for i := range spccb.builders {
		func(i int, root context.Context) {
			builder := spccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SellerProductCategoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, spccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, spccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spccb *SellerProductCategoryCreateBulk) SaveX(ctx context.Context) []*SellerProductCategory {
	v, err := spccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spccb *SellerProductCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := spccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spccb *SellerProductCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := spccb.Exec(ctx); err != nil {
		panic(err)
	}
}
