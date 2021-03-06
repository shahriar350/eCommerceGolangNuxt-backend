// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/category"
	"bongo/ent/sellerrequest"
	"bongo/ent/sellershop"
	"bongo/ent/shopcategory"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShopCategoryCreate is the builder for creating a ShopCategory entity.
type ShopCategoryCreate struct {
	config
	mutation *ShopCategoryMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (scc *ShopCategoryCreate) SetName(s string) *ShopCategoryCreate {
	scc.mutation.SetName(s)
	return scc
}

// SetSlug sets the "slug" field.
func (scc *ShopCategoryCreate) SetSlug(s string) *ShopCategoryCreate {
	scc.mutation.SetSlug(s)
	return scc
}

// SetImage sets the "image" field.
func (scc *ShopCategoryCreate) SetImage(s string) *ShopCategoryCreate {
	scc.mutation.SetImage(s)
	return scc
}

// SetCreatedAt sets the "created_at" field.
func (scc *ShopCategoryCreate) SetCreatedAt(t time.Time) *ShopCategoryCreate {
	scc.mutation.SetCreatedAt(t)
	return scc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (scc *ShopCategoryCreate) SetNillableCreatedAt(t *time.Time) *ShopCategoryCreate {
	if t != nil {
		scc.SetCreatedAt(*t)
	}
	return scc
}

// SetUpdatedAt sets the "updated_at" field.
func (scc *ShopCategoryCreate) SetUpdatedAt(t time.Time) *ShopCategoryCreate {
	scc.mutation.SetUpdatedAt(t)
	return scc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (scc *ShopCategoryCreate) SetNillableUpdatedAt(t *time.Time) *ShopCategoryCreate {
	if t != nil {
		scc.SetUpdatedAt(*t)
	}
	return scc
}

// SetDeletedAt sets the "deleted_at" field.
func (scc *ShopCategoryCreate) SetDeletedAt(t time.Time) *ShopCategoryCreate {
	scc.mutation.SetDeletedAt(t)
	return scc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (scc *ShopCategoryCreate) SetNillableDeletedAt(t *time.Time) *ShopCategoryCreate {
	if t != nil {
		scc.SetDeletedAt(*t)
	}
	return scc
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (scc *ShopCategoryCreate) AddCategoryIDs(ids ...int) *ShopCategoryCreate {
	scc.mutation.AddCategoryIDs(ids...)
	return scc
}

// AddCategories adds the "categories" edges to the Category entity.
func (scc *ShopCategoryCreate) AddCategories(c ...*Category) *ShopCategoryCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return scc.AddCategoryIDs(ids...)
}

// AddSellerRequestIDs adds the "seller_requests" edge to the SellerRequest entity by IDs.
func (scc *ShopCategoryCreate) AddSellerRequestIDs(ids ...int) *ShopCategoryCreate {
	scc.mutation.AddSellerRequestIDs(ids...)
	return scc
}

// AddSellerRequests adds the "seller_requests" edges to the SellerRequest entity.
func (scc *ShopCategoryCreate) AddSellerRequests(s ...*SellerRequest) *ShopCategoryCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scc.AddSellerRequestIDs(ids...)
}

// AddSellerShopIDs adds the "seller_shops" edge to the SellerShop entity by IDs.
func (scc *ShopCategoryCreate) AddSellerShopIDs(ids ...int) *ShopCategoryCreate {
	scc.mutation.AddSellerShopIDs(ids...)
	return scc
}

// AddSellerShops adds the "seller_shops" edges to the SellerShop entity.
func (scc *ShopCategoryCreate) AddSellerShops(s ...*SellerShop) *ShopCategoryCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scc.AddSellerShopIDs(ids...)
}

// Mutation returns the ShopCategoryMutation object of the builder.
func (scc *ShopCategoryCreate) Mutation() *ShopCategoryMutation {
	return scc.mutation
}

// Save creates the ShopCategory in the database.
func (scc *ShopCategoryCreate) Save(ctx context.Context) (*ShopCategory, error) {
	var (
		err  error
		node *ShopCategory
	)
	scc.defaults()
	if len(scc.hooks) == 0 {
		if err = scc.check(); err != nil {
			return nil, err
		}
		node, err = scc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShopCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = scc.check(); err != nil {
				return nil, err
			}
			scc.mutation = mutation
			if node, err = scc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(scc.hooks) - 1; i >= 0; i-- {
			if scc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (scc *ShopCategoryCreate) SaveX(ctx context.Context) *ShopCategory {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *ShopCategoryCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *ShopCategoryCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scc *ShopCategoryCreate) defaults() {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		v := shopcategory.DefaultCreatedAt()
		scc.mutation.SetCreatedAt(v)
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		v := shopcategory.DefaultUpdatedAt()
		scc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *ShopCategoryCreate) check() error {
	if _, ok := scc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ShopCategory.name"`)}
	}
	if _, ok := scc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "ShopCategory.slug"`)}
	}
	if _, ok := scc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "ShopCategory.image"`)}
	}
	if _, ok := scc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ShopCategory.created_at"`)}
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ShopCategory.updated_at"`)}
	}
	return nil
}

func (scc *ShopCategoryCreate) sqlSave(ctx context.Context) (*ShopCategory, error) {
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (scc *ShopCategoryCreate) createSpec() (*ShopCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &ShopCategory{config: scc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: shopcategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shopcategory.FieldID,
			},
		}
	)
	if value, ok := scc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shopcategory.FieldName,
		})
		_node.Name = value
	}
	if value, ok := scc.mutation.Slug(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shopcategory.FieldSlug,
		})
		_node.Slug = value
	}
	if value, ok := scc.mutation.Image(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shopcategory.FieldImage,
		})
		_node.Image = value
	}
	if value, ok := scc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shopcategory.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := scc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shopcategory.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := scc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shopcategory.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := scc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shopcategory.CategoriesTable,
			Columns: []string{shopcategory.CategoriesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := scc.mutation.SellerRequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shopcategory.SellerRequestsTable,
			Columns: []string{shopcategory.SellerRequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sellerrequest.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := scc.mutation.SellerShopsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shopcategory.SellerShopsTable,
			Columns: []string{shopcategory.SellerShopsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sellershop.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ShopCategoryCreateBulk is the builder for creating many ShopCategory entities in bulk.
type ShopCategoryCreateBulk struct {
	config
	builders []*ShopCategoryCreate
}

// Save creates the ShopCategory entities in the database.
func (sccb *ShopCategoryCreateBulk) Save(ctx context.Context) ([]*ShopCategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*ShopCategory, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShopCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *ShopCategoryCreateBulk) SaveX(ctx context.Context) []*ShopCategory {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *ShopCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *ShopCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}
