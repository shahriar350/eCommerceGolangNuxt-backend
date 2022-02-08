// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/cart"
	"bongo/ent/checkout"
	"bongo/ent/checkoutproduct"
	"bongo/ent/predicate"
	"bongo/ent/user"
	"bongo/ent/userlocation"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// CheckoutUpdate is the builder for updating Checkout entities.
type CheckoutUpdate struct {
	config
	hooks    []Hook
	mutation *CheckoutMutation
}

// Where appends a list predicates to the CheckoutUpdate builder.
func (cu *CheckoutUpdate) Where(ps ...predicate.Checkout) *CheckoutUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetTotalPrice sets the "total_price" field.
func (cu *CheckoutUpdate) SetTotalPrice(d decimal.Decimal) *CheckoutUpdate {
	cu.mutation.ResetTotalPrice()
	cu.mutation.SetTotalPrice(d)
	return cu
}

// AddTotalPrice adds d to the "total_price" field.
func (cu *CheckoutUpdate) AddTotalPrice(d decimal.Decimal) *CheckoutUpdate {
	cu.mutation.AddTotalPrice(d)
	return cu
}

// SetCompleted sets the "completed" field.
func (cu *CheckoutUpdate) SetCompleted(b bool) *CheckoutUpdate {
	cu.mutation.SetCompleted(b)
	return cu
}

// SetNillableCompleted sets the "completed" field if the given value is not nil.
func (cu *CheckoutUpdate) SetNillableCompleted(b *bool) *CheckoutUpdate {
	if b != nil {
		cu.SetCompleted(*b)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CheckoutUpdate) SetUpdatedAt(t time.Time) *CheckoutUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CheckoutUpdate) SetDeletedAt(t time.Time) *CheckoutUpdate {
	cu.mutation.SetDeletedAt(t)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CheckoutUpdate) SetNillableDeletedAt(t *time.Time) *CheckoutUpdate {
	if t != nil {
		cu.SetDeletedAt(*t)
	}
	return cu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cu *CheckoutUpdate) ClearDeletedAt() *CheckoutUpdate {
	cu.mutation.ClearDeletedAt()
	return cu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cu *CheckoutUpdate) SetUserID(id int) *CheckoutUpdate {
	cu.mutation.SetUserID(id)
	return cu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cu *CheckoutUpdate) SetNillableUserID(id *int) *CheckoutUpdate {
	if id != nil {
		cu = cu.SetUserID(*id)
	}
	return cu
}

// SetUser sets the "user" edge to the User entity.
func (cu *CheckoutUpdate) SetUser(u *User) *CheckoutUpdate {
	return cu.SetUserID(u.ID)
}

// SetLocationID sets the "location" edge to the UserLocation entity by ID.
func (cu *CheckoutUpdate) SetLocationID(id int) *CheckoutUpdate {
	cu.mutation.SetLocationID(id)
	return cu
}

// SetNillableLocationID sets the "location" edge to the UserLocation entity by ID if the given value is not nil.
func (cu *CheckoutUpdate) SetNillableLocationID(id *int) *CheckoutUpdate {
	if id != nil {
		cu = cu.SetLocationID(*id)
	}
	return cu
}

// SetLocation sets the "location" edge to the UserLocation entity.
func (cu *CheckoutUpdate) SetLocation(u *UserLocation) *CheckoutUpdate {
	return cu.SetLocationID(u.ID)
}

// SetCartID sets the "cart" edge to the Cart entity by ID.
func (cu *CheckoutUpdate) SetCartID(id int) *CheckoutUpdate {
	cu.mutation.SetCartID(id)
	return cu
}

// SetNillableCartID sets the "cart" edge to the Cart entity by ID if the given value is not nil.
func (cu *CheckoutUpdate) SetNillableCartID(id *int) *CheckoutUpdate {
	if id != nil {
		cu = cu.SetCartID(*id)
	}
	return cu
}

// SetCart sets the "cart" edge to the Cart entity.
func (cu *CheckoutUpdate) SetCart(c *Cart) *CheckoutUpdate {
	return cu.SetCartID(c.ID)
}

// AddCheckoutProductIDs adds the "checkout_products" edge to the CheckoutProduct entity by IDs.
func (cu *CheckoutUpdate) AddCheckoutProductIDs(ids ...int) *CheckoutUpdate {
	cu.mutation.AddCheckoutProductIDs(ids...)
	return cu
}

// AddCheckoutProducts adds the "checkout_products" edges to the CheckoutProduct entity.
func (cu *CheckoutUpdate) AddCheckoutProducts(c ...*CheckoutProduct) *CheckoutUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCheckoutProductIDs(ids...)
}

// Mutation returns the CheckoutMutation object of the builder.
func (cu *CheckoutUpdate) Mutation() *CheckoutMutation {
	return cu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cu *CheckoutUpdate) ClearUser() *CheckoutUpdate {
	cu.mutation.ClearUser()
	return cu
}

// ClearLocation clears the "location" edge to the UserLocation entity.
func (cu *CheckoutUpdate) ClearLocation() *CheckoutUpdate {
	cu.mutation.ClearLocation()
	return cu
}

// ClearCart clears the "cart" edge to the Cart entity.
func (cu *CheckoutUpdate) ClearCart() *CheckoutUpdate {
	cu.mutation.ClearCart()
	return cu
}

// ClearCheckoutProducts clears all "checkout_products" edges to the CheckoutProduct entity.
func (cu *CheckoutUpdate) ClearCheckoutProducts() *CheckoutUpdate {
	cu.mutation.ClearCheckoutProducts()
	return cu
}

// RemoveCheckoutProductIDs removes the "checkout_products" edge to CheckoutProduct entities by IDs.
func (cu *CheckoutUpdate) RemoveCheckoutProductIDs(ids ...int) *CheckoutUpdate {
	cu.mutation.RemoveCheckoutProductIDs(ids...)
	return cu
}

// RemoveCheckoutProducts removes "checkout_products" edges to CheckoutProduct entities.
func (cu *CheckoutUpdate) RemoveCheckoutProducts(c ...*CheckoutProduct) *CheckoutUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCheckoutProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CheckoutUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CheckoutMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CheckoutUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CheckoutUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CheckoutUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CheckoutUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := checkout.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CheckoutUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   checkout.Table,
			Columns: checkout.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checkout.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.TotalPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: checkout.FieldTotalPrice,
		})
	}
	if value, ok := cu.mutation.AddedTotalPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: checkout.FieldTotalPrice,
		})
	}
	if value, ok := cu.mutation.Completed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: checkout.FieldCompleted,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: checkout.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: checkout.FieldDeletedAt,
		})
	}
	if cu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: checkout.FieldDeletedAt,
		})
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checkout.UserTable,
			Columns: []string{checkout.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checkout.UserTable,
			Columns: []string{checkout.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checkout.LocationTable,
			Columns: []string{checkout.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlocation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checkout.LocationTable,
			Columns: []string{checkout.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlocation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   checkout.CartTable,
			Columns: []string{checkout.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cart.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   checkout.CartTable,
			Columns: []string{checkout.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cart.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CheckoutProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   checkout.CheckoutProductsTable,
			Columns: []string{checkout.CheckoutProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: checkoutproduct.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCheckoutProductsIDs(); len(nodes) > 0 && !cu.mutation.CheckoutProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   checkout.CheckoutProductsTable,
			Columns: []string{checkout.CheckoutProductsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CheckoutProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   checkout.CheckoutProductsTable,
			Columns: []string{checkout.CheckoutProductsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{checkout.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CheckoutUpdateOne is the builder for updating a single Checkout entity.
type CheckoutUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CheckoutMutation
}

// SetTotalPrice sets the "total_price" field.
func (cuo *CheckoutUpdateOne) SetTotalPrice(d decimal.Decimal) *CheckoutUpdateOne {
	cuo.mutation.ResetTotalPrice()
	cuo.mutation.SetTotalPrice(d)
	return cuo
}

// AddTotalPrice adds d to the "total_price" field.
func (cuo *CheckoutUpdateOne) AddTotalPrice(d decimal.Decimal) *CheckoutUpdateOne {
	cuo.mutation.AddTotalPrice(d)
	return cuo
}

// SetCompleted sets the "completed" field.
func (cuo *CheckoutUpdateOne) SetCompleted(b bool) *CheckoutUpdateOne {
	cuo.mutation.SetCompleted(b)
	return cuo
}

// SetNillableCompleted sets the "completed" field if the given value is not nil.
func (cuo *CheckoutUpdateOne) SetNillableCompleted(b *bool) *CheckoutUpdateOne {
	if b != nil {
		cuo.SetCompleted(*b)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CheckoutUpdateOne) SetUpdatedAt(t time.Time) *CheckoutUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CheckoutUpdateOne) SetDeletedAt(t time.Time) *CheckoutUpdateOne {
	cuo.mutation.SetDeletedAt(t)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CheckoutUpdateOne) SetNillableDeletedAt(t *time.Time) *CheckoutUpdateOne {
	if t != nil {
		cuo.SetDeletedAt(*t)
	}
	return cuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cuo *CheckoutUpdateOne) ClearDeletedAt() *CheckoutUpdateOne {
	cuo.mutation.ClearDeletedAt()
	return cuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cuo *CheckoutUpdateOne) SetUserID(id int) *CheckoutUpdateOne {
	cuo.mutation.SetUserID(id)
	return cuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cuo *CheckoutUpdateOne) SetNillableUserID(id *int) *CheckoutUpdateOne {
	if id != nil {
		cuo = cuo.SetUserID(*id)
	}
	return cuo
}

// SetUser sets the "user" edge to the User entity.
func (cuo *CheckoutUpdateOne) SetUser(u *User) *CheckoutUpdateOne {
	return cuo.SetUserID(u.ID)
}

// SetLocationID sets the "location" edge to the UserLocation entity by ID.
func (cuo *CheckoutUpdateOne) SetLocationID(id int) *CheckoutUpdateOne {
	cuo.mutation.SetLocationID(id)
	return cuo
}

// SetNillableLocationID sets the "location" edge to the UserLocation entity by ID if the given value is not nil.
func (cuo *CheckoutUpdateOne) SetNillableLocationID(id *int) *CheckoutUpdateOne {
	if id != nil {
		cuo = cuo.SetLocationID(*id)
	}
	return cuo
}

// SetLocation sets the "location" edge to the UserLocation entity.
func (cuo *CheckoutUpdateOne) SetLocation(u *UserLocation) *CheckoutUpdateOne {
	return cuo.SetLocationID(u.ID)
}

// SetCartID sets the "cart" edge to the Cart entity by ID.
func (cuo *CheckoutUpdateOne) SetCartID(id int) *CheckoutUpdateOne {
	cuo.mutation.SetCartID(id)
	return cuo
}

// SetNillableCartID sets the "cart" edge to the Cart entity by ID if the given value is not nil.
func (cuo *CheckoutUpdateOne) SetNillableCartID(id *int) *CheckoutUpdateOne {
	if id != nil {
		cuo = cuo.SetCartID(*id)
	}
	return cuo
}

// SetCart sets the "cart" edge to the Cart entity.
func (cuo *CheckoutUpdateOne) SetCart(c *Cart) *CheckoutUpdateOne {
	return cuo.SetCartID(c.ID)
}

// AddCheckoutProductIDs adds the "checkout_products" edge to the CheckoutProduct entity by IDs.
func (cuo *CheckoutUpdateOne) AddCheckoutProductIDs(ids ...int) *CheckoutUpdateOne {
	cuo.mutation.AddCheckoutProductIDs(ids...)
	return cuo
}

// AddCheckoutProducts adds the "checkout_products" edges to the CheckoutProduct entity.
func (cuo *CheckoutUpdateOne) AddCheckoutProducts(c ...*CheckoutProduct) *CheckoutUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCheckoutProductIDs(ids...)
}

// Mutation returns the CheckoutMutation object of the builder.
func (cuo *CheckoutUpdateOne) Mutation() *CheckoutMutation {
	return cuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cuo *CheckoutUpdateOne) ClearUser() *CheckoutUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// ClearLocation clears the "location" edge to the UserLocation entity.
func (cuo *CheckoutUpdateOne) ClearLocation() *CheckoutUpdateOne {
	cuo.mutation.ClearLocation()
	return cuo
}

// ClearCart clears the "cart" edge to the Cart entity.
func (cuo *CheckoutUpdateOne) ClearCart() *CheckoutUpdateOne {
	cuo.mutation.ClearCart()
	return cuo
}

// ClearCheckoutProducts clears all "checkout_products" edges to the CheckoutProduct entity.
func (cuo *CheckoutUpdateOne) ClearCheckoutProducts() *CheckoutUpdateOne {
	cuo.mutation.ClearCheckoutProducts()
	return cuo
}

// RemoveCheckoutProductIDs removes the "checkout_products" edge to CheckoutProduct entities by IDs.
func (cuo *CheckoutUpdateOne) RemoveCheckoutProductIDs(ids ...int) *CheckoutUpdateOne {
	cuo.mutation.RemoveCheckoutProductIDs(ids...)
	return cuo
}

// RemoveCheckoutProducts removes "checkout_products" edges to CheckoutProduct entities.
func (cuo *CheckoutUpdateOne) RemoveCheckoutProducts(c ...*CheckoutProduct) *CheckoutUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCheckoutProductIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CheckoutUpdateOne) Select(field string, fields ...string) *CheckoutUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Checkout entity.
func (cuo *CheckoutUpdateOne) Save(ctx context.Context) (*Checkout, error) {
	var (
		err  error
		node *Checkout
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CheckoutMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CheckoutUpdateOne) SaveX(ctx context.Context) *Checkout {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CheckoutUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CheckoutUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CheckoutUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := checkout.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CheckoutUpdateOne) sqlSave(ctx context.Context) (_node *Checkout, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   checkout.Table,
			Columns: checkout.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checkout.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Checkout.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, checkout.FieldID)
		for _, f := range fields {
			if !checkout.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != checkout.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.TotalPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: checkout.FieldTotalPrice,
		})
	}
	if value, ok := cuo.mutation.AddedTotalPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: checkout.FieldTotalPrice,
		})
	}
	if value, ok := cuo.mutation.Completed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: checkout.FieldCompleted,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: checkout.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: checkout.FieldDeletedAt,
		})
	}
	if cuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: checkout.FieldDeletedAt,
		})
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checkout.UserTable,
			Columns: []string{checkout.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checkout.UserTable,
			Columns: []string{checkout.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checkout.LocationTable,
			Columns: []string{checkout.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlocation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checkout.LocationTable,
			Columns: []string{checkout.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlocation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   checkout.CartTable,
			Columns: []string{checkout.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cart.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   checkout.CartTable,
			Columns: []string{checkout.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cart.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CheckoutProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   checkout.CheckoutProductsTable,
			Columns: []string{checkout.CheckoutProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: checkoutproduct.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCheckoutProductsIDs(); len(nodes) > 0 && !cuo.mutation.CheckoutProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   checkout.CheckoutProductsTable,
			Columns: []string{checkout.CheckoutProductsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CheckoutProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   checkout.CheckoutProductsTable,
			Columns: []string{checkout.CheckoutProductsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Checkout{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{checkout.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
