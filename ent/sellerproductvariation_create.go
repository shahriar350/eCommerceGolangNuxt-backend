// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/cartproduct"
	"bongo/ent/checkoutproduct"
	"bongo/ent/sellerproduct"
	"bongo/ent/sellerproductvariation"
	"bongo/ent/sellerproductvariationvalues"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// SellerProductVariationCreate is the builder for creating a SellerProductVariation entity.
type SellerProductVariationCreate struct {
	config
	mutation *SellerProductVariationMutation
	hooks    []Hook
}

// SetProductPrice sets the "product_price" field.
func (spvc *SellerProductVariationCreate) SetProductPrice(d decimal.Decimal) *SellerProductVariationCreate {
	spvc.mutation.SetProductPrice(d)
	return spvc
}

// SetSellingPrice sets the "selling_price" field.
func (spvc *SellerProductVariationCreate) SetSellingPrice(d decimal.Decimal) *SellerProductVariationCreate {
	spvc.mutation.SetSellingPrice(d)
	return spvc
}

// SetQuantity sets the "quantity" field.
func (spvc *SellerProductVariationCreate) SetQuantity(i int) *SellerProductVariationCreate {
	spvc.mutation.SetQuantity(i)
	return spvc
}

// SetImage sets the "image" field.
func (spvc *SellerProductVariationCreate) SetImage(i int) *SellerProductVariationCreate {
	spvc.mutation.SetImage(i)
	return spvc
}

// SetCreatedAt sets the "created_at" field.
func (spvc *SellerProductVariationCreate) SetCreatedAt(t time.Time) *SellerProductVariationCreate {
	spvc.mutation.SetCreatedAt(t)
	return spvc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (spvc *SellerProductVariationCreate) SetNillableCreatedAt(t *time.Time) *SellerProductVariationCreate {
	if t != nil {
		spvc.SetCreatedAt(*t)
	}
	return spvc
}

// SetUpdatedAt sets the "updated_at" field.
func (spvc *SellerProductVariationCreate) SetUpdatedAt(t time.Time) *SellerProductVariationCreate {
	spvc.mutation.SetUpdatedAt(t)
	return spvc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (spvc *SellerProductVariationCreate) SetNillableUpdatedAt(t *time.Time) *SellerProductVariationCreate {
	if t != nil {
		spvc.SetUpdatedAt(*t)
	}
	return spvc
}

// SetDeletedAt sets the "deleted_at" field.
func (spvc *SellerProductVariationCreate) SetDeletedAt(t time.Time) *SellerProductVariationCreate {
	spvc.mutation.SetDeletedAt(t)
	return spvc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (spvc *SellerProductVariationCreate) SetNillableDeletedAt(t *time.Time) *SellerProductVariationCreate {
	if t != nil {
		spvc.SetDeletedAt(*t)
	}
	return spvc
}

// SetSellerProductID sets the "seller_product" edge to the SellerProduct entity by ID.
func (spvc *SellerProductVariationCreate) SetSellerProductID(id int) *SellerProductVariationCreate {
	spvc.mutation.SetSellerProductID(id)
	return spvc
}

// SetNillableSellerProductID sets the "seller_product" edge to the SellerProduct entity by ID if the given value is not nil.
func (spvc *SellerProductVariationCreate) SetNillableSellerProductID(id *int) *SellerProductVariationCreate {
	if id != nil {
		spvc = spvc.SetSellerProductID(*id)
	}
	return spvc
}

// SetSellerProduct sets the "seller_product" edge to the SellerProduct entity.
func (spvc *SellerProductVariationCreate) SetSellerProduct(s *SellerProduct) *SellerProductVariationCreate {
	return spvc.SetSellerProductID(s.ID)
}

// AddSellerProductVariationValueIDs adds the "seller_product_variation_values" edge to the SellerProductVariationValues entity by IDs.
func (spvc *SellerProductVariationCreate) AddSellerProductVariationValueIDs(ids ...int) *SellerProductVariationCreate {
	spvc.mutation.AddSellerProductVariationValueIDs(ids...)
	return spvc
}

// AddSellerProductVariationValues adds the "seller_product_variation_values" edges to the SellerProductVariationValues entity.
func (spvc *SellerProductVariationCreate) AddSellerProductVariationValues(s ...*SellerProductVariationValues) *SellerProductVariationCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return spvc.AddSellerProductVariationValueIDs(ids...)
}

// AddCartProductIDs adds the "cart_products" edge to the CartProduct entity by IDs.
func (spvc *SellerProductVariationCreate) AddCartProductIDs(ids ...int) *SellerProductVariationCreate {
	spvc.mutation.AddCartProductIDs(ids...)
	return spvc
}

// AddCartProducts adds the "cart_products" edges to the CartProduct entity.
func (spvc *SellerProductVariationCreate) AddCartProducts(c ...*CartProduct) *SellerProductVariationCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return spvc.AddCartProductIDs(ids...)
}

// AddCheckoutProductIDs adds the "checkout_products" edge to the CheckoutProduct entity by IDs.
func (spvc *SellerProductVariationCreate) AddCheckoutProductIDs(ids ...int) *SellerProductVariationCreate {
	spvc.mutation.AddCheckoutProductIDs(ids...)
	return spvc
}

// AddCheckoutProducts adds the "checkout_products" edges to the CheckoutProduct entity.
func (spvc *SellerProductVariationCreate) AddCheckoutProducts(c ...*CheckoutProduct) *SellerProductVariationCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return spvc.AddCheckoutProductIDs(ids...)
}

// Mutation returns the SellerProductVariationMutation object of the builder.
func (spvc *SellerProductVariationCreate) Mutation() *SellerProductVariationMutation {
	return spvc.mutation
}

// Save creates the SellerProductVariation in the database.
func (spvc *SellerProductVariationCreate) Save(ctx context.Context) (*SellerProductVariation, error) {
	var (
		err  error
		node *SellerProductVariation
	)
	spvc.defaults()
	if len(spvc.hooks) == 0 {
		if err = spvc.check(); err != nil {
			return nil, err
		}
		node, err = spvc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SellerProductVariationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = spvc.check(); err != nil {
				return nil, err
			}
			spvc.mutation = mutation
			if node, err = spvc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(spvc.hooks) - 1; i >= 0; i-- {
			if spvc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spvc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spvc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (spvc *SellerProductVariationCreate) SaveX(ctx context.Context) *SellerProductVariation {
	v, err := spvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spvc *SellerProductVariationCreate) Exec(ctx context.Context) error {
	_, err := spvc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spvc *SellerProductVariationCreate) ExecX(ctx context.Context) {
	if err := spvc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spvc *SellerProductVariationCreate) defaults() {
	if _, ok := spvc.mutation.CreatedAt(); !ok {
		v := sellerproductvariation.DefaultCreatedAt()
		spvc.mutation.SetCreatedAt(v)
	}
	if _, ok := spvc.mutation.UpdatedAt(); !ok {
		v := sellerproductvariation.DefaultUpdatedAt()
		spvc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spvc *SellerProductVariationCreate) check() error {
	if _, ok := spvc.mutation.ProductPrice(); !ok {
		return &ValidationError{Name: "product_price", err: errors.New(`ent: missing required field "product_price"`)}
	}
	if _, ok := spvc.mutation.SellingPrice(); !ok {
		return &ValidationError{Name: "selling_price", err: errors.New(`ent: missing required field "selling_price"`)}
	}
	if _, ok := spvc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "quantity"`)}
	}
	if _, ok := spvc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "image"`)}
	}
	if _, ok := spvc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := spvc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (spvc *SellerProductVariationCreate) sqlSave(ctx context.Context) (*SellerProductVariation, error) {
	_node, _spec := spvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spvc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (spvc *SellerProductVariationCreate) createSpec() (*SellerProductVariation, *sqlgraph.CreateSpec) {
	var (
		_node = &SellerProductVariation{config: spvc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sellerproductvariation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sellerproductvariation.FieldID,
			},
		}
	)
	if value, ok := spvc.mutation.ProductPrice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sellerproductvariation.FieldProductPrice,
		})
		_node.ProductPrice = value
	}
	if value, ok := spvc.mutation.SellingPrice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sellerproductvariation.FieldSellingPrice,
		})
		_node.SellingPrice = value
	}
	if value, ok := spvc.mutation.Quantity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sellerproductvariation.FieldQuantity,
		})
		_node.Quantity = value
	}
	if value, ok := spvc.mutation.Image(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sellerproductvariation.FieldImage,
		})
		_node.Image = value
	}
	if value, ok := spvc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductvariation.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := spvc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductvariation.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := spvc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductvariation.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := spvc.mutation.SellerProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductvariation.SellerProductTable,
			Columns: []string{sellerproductvariation.SellerProductColumn},
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
		_node.seller_product_seller_product_variations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := spvc.mutation.SellerProductVariationValuesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sellerproductvariation.SellerProductVariationValuesTable,
			Columns: []string{sellerproductvariation.SellerProductVariationValuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sellerproductvariationvalues.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := spvc.mutation.CartProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sellerproductvariation.CartProductsTable,
			Columns: []string{sellerproductvariation.CartProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cartproduct.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := spvc.mutation.CheckoutProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sellerproductvariation.CheckoutProductsTable,
			Columns: []string{sellerproductvariation.CheckoutProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: checkoutproduct.FieldID,
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

// SellerProductVariationCreateBulk is the builder for creating many SellerProductVariation entities in bulk.
type SellerProductVariationCreateBulk struct {
	config
	builders []*SellerProductVariationCreate
}

// Save creates the SellerProductVariation entities in the database.
func (spvcb *SellerProductVariationCreateBulk) Save(ctx context.Context) ([]*SellerProductVariation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(spvcb.builders))
	nodes := make([]*SellerProductVariation, len(spvcb.builders))
	mutators := make([]Mutator, len(spvcb.builders))
	for i := range spvcb.builders {
		func(i int, root context.Context) {
			builder := spvcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SellerProductVariationMutation)
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
					_, err = mutators[i+1].Mutate(root, spvcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spvcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, spvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spvcb *SellerProductVariationCreateBulk) SaveX(ctx context.Context) []*SellerProductVariation {
	v, err := spvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spvcb *SellerProductVariationCreateBulk) Exec(ctx context.Context) error {
	_, err := spvcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spvcb *SellerProductVariationCreateBulk) ExecX(ctx context.Context) {
	if err := spvcb.Exec(ctx); err != nil {
		panic(err)
	}
}
