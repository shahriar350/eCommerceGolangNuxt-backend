// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/category"
	"bongo/ent/predicate"
	"bongo/ent/sellerrequest"
	"bongo/ent/sellershop"
	"bongo/ent/shopcategory"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShopCategoryUpdate is the builder for updating ShopCategory entities.
type ShopCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *ShopCategoryMutation
}

// Where appends a list predicates to the ShopCategoryUpdate builder.
func (scu *ShopCategoryUpdate) Where(ps ...predicate.ShopCategory) *ShopCategoryUpdate {
	scu.mutation.Where(ps...)
	return scu
}

// SetName sets the "name" field.
func (scu *ShopCategoryUpdate) SetName(s string) *ShopCategoryUpdate {
	scu.mutation.SetName(s)
	return scu
}

// SetSlug sets the "slug" field.
func (scu *ShopCategoryUpdate) SetSlug(s string) *ShopCategoryUpdate {
	scu.mutation.SetSlug(s)
	return scu
}

// SetImage sets the "image" field.
func (scu *ShopCategoryUpdate) SetImage(s string) *ShopCategoryUpdate {
	scu.mutation.SetImage(s)
	return scu
}

// SetUpdatedAt sets the "updated_at" field.
func (scu *ShopCategoryUpdate) SetUpdatedAt(t time.Time) *ShopCategoryUpdate {
	scu.mutation.SetUpdatedAt(t)
	return scu
}

// SetDeletedAt sets the "deleted_at" field.
func (scu *ShopCategoryUpdate) SetDeletedAt(t time.Time) *ShopCategoryUpdate {
	scu.mutation.SetDeletedAt(t)
	return scu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (scu *ShopCategoryUpdate) SetNillableDeletedAt(t *time.Time) *ShopCategoryUpdate {
	if t != nil {
		scu.SetDeletedAt(*t)
	}
	return scu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (scu *ShopCategoryUpdate) ClearDeletedAt() *ShopCategoryUpdate {
	scu.mutation.ClearDeletedAt()
	return scu
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (scu *ShopCategoryUpdate) AddCategoryIDs(ids ...int) *ShopCategoryUpdate {
	scu.mutation.AddCategoryIDs(ids...)
	return scu
}

// AddCategories adds the "categories" edges to the Category entity.
func (scu *ShopCategoryUpdate) AddCategories(c ...*Category) *ShopCategoryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return scu.AddCategoryIDs(ids...)
}

// AddSellerRequestIDs adds the "seller_requests" edge to the SellerRequest entity by IDs.
func (scu *ShopCategoryUpdate) AddSellerRequestIDs(ids ...int) *ShopCategoryUpdate {
	scu.mutation.AddSellerRequestIDs(ids...)
	return scu
}

// AddSellerRequests adds the "seller_requests" edges to the SellerRequest entity.
func (scu *ShopCategoryUpdate) AddSellerRequests(s ...*SellerRequest) *ShopCategoryUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scu.AddSellerRequestIDs(ids...)
}

// AddSellerShopIDs adds the "seller_shops" edge to the SellerShop entity by IDs.
func (scu *ShopCategoryUpdate) AddSellerShopIDs(ids ...int) *ShopCategoryUpdate {
	scu.mutation.AddSellerShopIDs(ids...)
	return scu
}

// AddSellerShops adds the "seller_shops" edges to the SellerShop entity.
func (scu *ShopCategoryUpdate) AddSellerShops(s ...*SellerShop) *ShopCategoryUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scu.AddSellerShopIDs(ids...)
}

// Mutation returns the ShopCategoryMutation object of the builder.
func (scu *ShopCategoryUpdate) Mutation() *ShopCategoryMutation {
	return scu.mutation
}

// ClearCategories clears all "categories" edges to the Category entity.
func (scu *ShopCategoryUpdate) ClearCategories() *ShopCategoryUpdate {
	scu.mutation.ClearCategories()
	return scu
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (scu *ShopCategoryUpdate) RemoveCategoryIDs(ids ...int) *ShopCategoryUpdate {
	scu.mutation.RemoveCategoryIDs(ids...)
	return scu
}

// RemoveCategories removes "categories" edges to Category entities.
func (scu *ShopCategoryUpdate) RemoveCategories(c ...*Category) *ShopCategoryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return scu.RemoveCategoryIDs(ids...)
}

// ClearSellerRequests clears all "seller_requests" edges to the SellerRequest entity.
func (scu *ShopCategoryUpdate) ClearSellerRequests() *ShopCategoryUpdate {
	scu.mutation.ClearSellerRequests()
	return scu
}

// RemoveSellerRequestIDs removes the "seller_requests" edge to SellerRequest entities by IDs.
func (scu *ShopCategoryUpdate) RemoveSellerRequestIDs(ids ...int) *ShopCategoryUpdate {
	scu.mutation.RemoveSellerRequestIDs(ids...)
	return scu
}

// RemoveSellerRequests removes "seller_requests" edges to SellerRequest entities.
func (scu *ShopCategoryUpdate) RemoveSellerRequests(s ...*SellerRequest) *ShopCategoryUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scu.RemoveSellerRequestIDs(ids...)
}

// ClearSellerShops clears all "seller_shops" edges to the SellerShop entity.
func (scu *ShopCategoryUpdate) ClearSellerShops() *ShopCategoryUpdate {
	scu.mutation.ClearSellerShops()
	return scu
}

// RemoveSellerShopIDs removes the "seller_shops" edge to SellerShop entities by IDs.
func (scu *ShopCategoryUpdate) RemoveSellerShopIDs(ids ...int) *ShopCategoryUpdate {
	scu.mutation.RemoveSellerShopIDs(ids...)
	return scu
}

// RemoveSellerShops removes "seller_shops" edges to SellerShop entities.
func (scu *ShopCategoryUpdate) RemoveSellerShops(s ...*SellerShop) *ShopCategoryUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scu.RemoveSellerShopIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (scu *ShopCategoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	scu.defaults()
	if len(scu.hooks) == 0 {
		affected, err = scu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShopCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			scu.mutation = mutation
			affected, err = scu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(scu.hooks) - 1; i >= 0; i-- {
			if scu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (scu *ShopCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := scu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (scu *ShopCategoryUpdate) Exec(ctx context.Context) error {
	_, err := scu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scu *ShopCategoryUpdate) ExecX(ctx context.Context) {
	if err := scu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scu *ShopCategoryUpdate) defaults() {
	if _, ok := scu.mutation.UpdatedAt(); !ok {
		v := shopcategory.UpdateDefaultUpdatedAt()
		scu.mutation.SetUpdatedAt(v)
	}
}

func (scu *ShopCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shopcategory.Table,
			Columns: shopcategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shopcategory.FieldID,
			},
		},
	}
	if ps := scu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shopcategory.FieldName,
		})
	}
	if value, ok := scu.mutation.Slug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shopcategory.FieldSlug,
		})
	}
	if value, ok := scu.mutation.Image(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shopcategory.FieldImage,
		})
	}
	if value, ok := scu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shopcategory.FieldUpdatedAt,
		})
	}
	if value, ok := scu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shopcategory.FieldDeletedAt,
		})
	}
	if scu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: shopcategory.FieldDeletedAt,
		})
	}
	if scu.mutation.CategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !scu.mutation.CategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.CategoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if scu.mutation.SellerRequestsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.RemovedSellerRequestsIDs(); len(nodes) > 0 && !scu.mutation.SellerRequestsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.SellerRequestsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if scu.mutation.SellerShopsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.RemovedSellerShopsIDs(); len(nodes) > 0 && !scu.mutation.SellerShopsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.SellerShopsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, scu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shopcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ShopCategoryUpdateOne is the builder for updating a single ShopCategory entity.
type ShopCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShopCategoryMutation
}

// SetName sets the "name" field.
func (scuo *ShopCategoryUpdateOne) SetName(s string) *ShopCategoryUpdateOne {
	scuo.mutation.SetName(s)
	return scuo
}

// SetSlug sets the "slug" field.
func (scuo *ShopCategoryUpdateOne) SetSlug(s string) *ShopCategoryUpdateOne {
	scuo.mutation.SetSlug(s)
	return scuo
}

// SetImage sets the "image" field.
func (scuo *ShopCategoryUpdateOne) SetImage(s string) *ShopCategoryUpdateOne {
	scuo.mutation.SetImage(s)
	return scuo
}

// SetUpdatedAt sets the "updated_at" field.
func (scuo *ShopCategoryUpdateOne) SetUpdatedAt(t time.Time) *ShopCategoryUpdateOne {
	scuo.mutation.SetUpdatedAt(t)
	return scuo
}

// SetDeletedAt sets the "deleted_at" field.
func (scuo *ShopCategoryUpdateOne) SetDeletedAt(t time.Time) *ShopCategoryUpdateOne {
	scuo.mutation.SetDeletedAt(t)
	return scuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (scuo *ShopCategoryUpdateOne) SetNillableDeletedAt(t *time.Time) *ShopCategoryUpdateOne {
	if t != nil {
		scuo.SetDeletedAt(*t)
	}
	return scuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (scuo *ShopCategoryUpdateOne) ClearDeletedAt() *ShopCategoryUpdateOne {
	scuo.mutation.ClearDeletedAt()
	return scuo
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (scuo *ShopCategoryUpdateOne) AddCategoryIDs(ids ...int) *ShopCategoryUpdateOne {
	scuo.mutation.AddCategoryIDs(ids...)
	return scuo
}

// AddCategories adds the "categories" edges to the Category entity.
func (scuo *ShopCategoryUpdateOne) AddCategories(c ...*Category) *ShopCategoryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return scuo.AddCategoryIDs(ids...)
}

// AddSellerRequestIDs adds the "seller_requests" edge to the SellerRequest entity by IDs.
func (scuo *ShopCategoryUpdateOne) AddSellerRequestIDs(ids ...int) *ShopCategoryUpdateOne {
	scuo.mutation.AddSellerRequestIDs(ids...)
	return scuo
}

// AddSellerRequests adds the "seller_requests" edges to the SellerRequest entity.
func (scuo *ShopCategoryUpdateOne) AddSellerRequests(s ...*SellerRequest) *ShopCategoryUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scuo.AddSellerRequestIDs(ids...)
}

// AddSellerShopIDs adds the "seller_shops" edge to the SellerShop entity by IDs.
func (scuo *ShopCategoryUpdateOne) AddSellerShopIDs(ids ...int) *ShopCategoryUpdateOne {
	scuo.mutation.AddSellerShopIDs(ids...)
	return scuo
}

// AddSellerShops adds the "seller_shops" edges to the SellerShop entity.
func (scuo *ShopCategoryUpdateOne) AddSellerShops(s ...*SellerShop) *ShopCategoryUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scuo.AddSellerShopIDs(ids...)
}

// Mutation returns the ShopCategoryMutation object of the builder.
func (scuo *ShopCategoryUpdateOne) Mutation() *ShopCategoryMutation {
	return scuo.mutation
}

// ClearCategories clears all "categories" edges to the Category entity.
func (scuo *ShopCategoryUpdateOne) ClearCategories() *ShopCategoryUpdateOne {
	scuo.mutation.ClearCategories()
	return scuo
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (scuo *ShopCategoryUpdateOne) RemoveCategoryIDs(ids ...int) *ShopCategoryUpdateOne {
	scuo.mutation.RemoveCategoryIDs(ids...)
	return scuo
}

// RemoveCategories removes "categories" edges to Category entities.
func (scuo *ShopCategoryUpdateOne) RemoveCategories(c ...*Category) *ShopCategoryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return scuo.RemoveCategoryIDs(ids...)
}

// ClearSellerRequests clears all "seller_requests" edges to the SellerRequest entity.
func (scuo *ShopCategoryUpdateOne) ClearSellerRequests() *ShopCategoryUpdateOne {
	scuo.mutation.ClearSellerRequests()
	return scuo
}

// RemoveSellerRequestIDs removes the "seller_requests" edge to SellerRequest entities by IDs.
func (scuo *ShopCategoryUpdateOne) RemoveSellerRequestIDs(ids ...int) *ShopCategoryUpdateOne {
	scuo.mutation.RemoveSellerRequestIDs(ids...)
	return scuo
}

// RemoveSellerRequests removes "seller_requests" edges to SellerRequest entities.
func (scuo *ShopCategoryUpdateOne) RemoveSellerRequests(s ...*SellerRequest) *ShopCategoryUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scuo.RemoveSellerRequestIDs(ids...)
}

// ClearSellerShops clears all "seller_shops" edges to the SellerShop entity.
func (scuo *ShopCategoryUpdateOne) ClearSellerShops() *ShopCategoryUpdateOne {
	scuo.mutation.ClearSellerShops()
	return scuo
}

// RemoveSellerShopIDs removes the "seller_shops" edge to SellerShop entities by IDs.
func (scuo *ShopCategoryUpdateOne) RemoveSellerShopIDs(ids ...int) *ShopCategoryUpdateOne {
	scuo.mutation.RemoveSellerShopIDs(ids...)
	return scuo
}

// RemoveSellerShops removes "seller_shops" edges to SellerShop entities.
func (scuo *ShopCategoryUpdateOne) RemoveSellerShops(s ...*SellerShop) *ShopCategoryUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scuo.RemoveSellerShopIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (scuo *ShopCategoryUpdateOne) Select(field string, fields ...string) *ShopCategoryUpdateOne {
	scuo.fields = append([]string{field}, fields...)
	return scuo
}

// Save executes the query and returns the updated ShopCategory entity.
func (scuo *ShopCategoryUpdateOne) Save(ctx context.Context) (*ShopCategory, error) {
	var (
		err  error
		node *ShopCategory
	)
	scuo.defaults()
	if len(scuo.hooks) == 0 {
		node, err = scuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShopCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			scuo.mutation = mutation
			node, err = scuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(scuo.hooks) - 1; i >= 0; i-- {
			if scuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (scuo *ShopCategoryUpdateOne) SaveX(ctx context.Context) *ShopCategory {
	node, err := scuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (scuo *ShopCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := scuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scuo *ShopCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := scuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scuo *ShopCategoryUpdateOne) defaults() {
	if _, ok := scuo.mutation.UpdatedAt(); !ok {
		v := shopcategory.UpdateDefaultUpdatedAt()
		scuo.mutation.SetUpdatedAt(v)
	}
}

func (scuo *ShopCategoryUpdateOne) sqlSave(ctx context.Context) (_node *ShopCategory, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shopcategory.Table,
			Columns: shopcategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shopcategory.FieldID,
			},
		},
	}
	id, ok := scuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ShopCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := scuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shopcategory.FieldID)
		for _, f := range fields {
			if !shopcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != shopcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := scuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shopcategory.FieldName,
		})
	}
	if value, ok := scuo.mutation.Slug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shopcategory.FieldSlug,
		})
	}
	if value, ok := scuo.mutation.Image(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shopcategory.FieldImage,
		})
	}
	if value, ok := scuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shopcategory.FieldUpdatedAt,
		})
	}
	if value, ok := scuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shopcategory.FieldDeletedAt,
		})
	}
	if scuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: shopcategory.FieldDeletedAt,
		})
	}
	if scuo.mutation.CategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !scuo.mutation.CategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.CategoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if scuo.mutation.SellerRequestsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.RemovedSellerRequestsIDs(); len(nodes) > 0 && !scuo.mutation.SellerRequestsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.SellerRequestsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if scuo.mutation.SellerShopsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.RemovedSellerShopsIDs(); len(nodes) > 0 && !scuo.mutation.SellerShopsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.SellerShopsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ShopCategory{config: scuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, scuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shopcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
