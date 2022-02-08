// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/attribute"
	"bongo/ent/predicate"
	"bongo/ent/sellerproductvariation"
	"bongo/ent/sellerproductvariationvalues"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SellerProductVariationValuesUpdate is the builder for updating SellerProductVariationValues entities.
type SellerProductVariationValuesUpdate struct {
	config
	hooks    []Hook
	mutation *SellerProductVariationValuesMutation
}

// Where appends a list predicates to the SellerProductVariationValuesUpdate builder.
func (spvvu *SellerProductVariationValuesUpdate) Where(ps ...predicate.SellerProductVariationValues) *SellerProductVariationValuesUpdate {
	spvvu.mutation.Where(ps...)
	return spvvu
}

// SetName sets the "name" field.
func (spvvu *SellerProductVariationValuesUpdate) SetName(s string) *SellerProductVariationValuesUpdate {
	spvvu.mutation.SetName(s)
	return spvvu
}

// SetDescription sets the "description" field.
func (spvvu *SellerProductVariationValuesUpdate) SetDescription(s string) *SellerProductVariationValuesUpdate {
	spvvu.mutation.SetDescription(s)
	return spvvu
}

// SetUpdatedAt sets the "updated_at" field.
func (spvvu *SellerProductVariationValuesUpdate) SetUpdatedAt(t time.Time) *SellerProductVariationValuesUpdate {
	spvvu.mutation.SetUpdatedAt(t)
	return spvvu
}

// SetDeletedAt sets the "deleted_at" field.
func (spvvu *SellerProductVariationValuesUpdate) SetDeletedAt(t time.Time) *SellerProductVariationValuesUpdate {
	spvvu.mutation.SetDeletedAt(t)
	return spvvu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (spvvu *SellerProductVariationValuesUpdate) SetNillableDeletedAt(t *time.Time) *SellerProductVariationValuesUpdate {
	if t != nil {
		spvvu.SetDeletedAt(*t)
	}
	return spvvu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (spvvu *SellerProductVariationValuesUpdate) ClearDeletedAt() *SellerProductVariationValuesUpdate {
	spvvu.mutation.ClearDeletedAt()
	return spvvu
}

// SetAttributeID sets the "attribute" edge to the Attribute entity by ID.
func (spvvu *SellerProductVariationValuesUpdate) SetAttributeID(id int) *SellerProductVariationValuesUpdate {
	spvvu.mutation.SetAttributeID(id)
	return spvvu
}

// SetNillableAttributeID sets the "attribute" edge to the Attribute entity by ID if the given value is not nil.
func (spvvu *SellerProductVariationValuesUpdate) SetNillableAttributeID(id *int) *SellerProductVariationValuesUpdate {
	if id != nil {
		spvvu = spvvu.SetAttributeID(*id)
	}
	return spvvu
}

// SetAttribute sets the "attribute" edge to the Attribute entity.
func (spvvu *SellerProductVariationValuesUpdate) SetAttribute(a *Attribute) *SellerProductVariationValuesUpdate {
	return spvvu.SetAttributeID(a.ID)
}

// SetSellerProductVariationID sets the "seller_product_variation" edge to the SellerProductVariation entity by ID.
func (spvvu *SellerProductVariationValuesUpdate) SetSellerProductVariationID(id int) *SellerProductVariationValuesUpdate {
	spvvu.mutation.SetSellerProductVariationID(id)
	return spvvu
}

// SetNillableSellerProductVariationID sets the "seller_product_variation" edge to the SellerProductVariation entity by ID if the given value is not nil.
func (spvvu *SellerProductVariationValuesUpdate) SetNillableSellerProductVariationID(id *int) *SellerProductVariationValuesUpdate {
	if id != nil {
		spvvu = spvvu.SetSellerProductVariationID(*id)
	}
	return spvvu
}

// SetSellerProductVariation sets the "seller_product_variation" edge to the SellerProductVariation entity.
func (spvvu *SellerProductVariationValuesUpdate) SetSellerProductVariation(s *SellerProductVariation) *SellerProductVariationValuesUpdate {
	return spvvu.SetSellerProductVariationID(s.ID)
}

// Mutation returns the SellerProductVariationValuesMutation object of the builder.
func (spvvu *SellerProductVariationValuesUpdate) Mutation() *SellerProductVariationValuesMutation {
	return spvvu.mutation
}

// ClearAttribute clears the "attribute" edge to the Attribute entity.
func (spvvu *SellerProductVariationValuesUpdate) ClearAttribute() *SellerProductVariationValuesUpdate {
	spvvu.mutation.ClearAttribute()
	return spvvu
}

// ClearSellerProductVariation clears the "seller_product_variation" edge to the SellerProductVariation entity.
func (spvvu *SellerProductVariationValuesUpdate) ClearSellerProductVariation() *SellerProductVariationValuesUpdate {
	spvvu.mutation.ClearSellerProductVariation()
	return spvvu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (spvvu *SellerProductVariationValuesUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	spvvu.defaults()
	if len(spvvu.hooks) == 0 {
		affected, err = spvvu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SellerProductVariationValuesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spvvu.mutation = mutation
			affected, err = spvvu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spvvu.hooks) - 1; i >= 0; i-- {
			if spvvu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spvvu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spvvu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (spvvu *SellerProductVariationValuesUpdate) SaveX(ctx context.Context) int {
	affected, err := spvvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (spvvu *SellerProductVariationValuesUpdate) Exec(ctx context.Context) error {
	_, err := spvvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spvvu *SellerProductVariationValuesUpdate) ExecX(ctx context.Context) {
	if err := spvvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spvvu *SellerProductVariationValuesUpdate) defaults() {
	if _, ok := spvvu.mutation.UpdatedAt(); !ok {
		v := sellerproductvariationvalues.UpdateDefaultUpdatedAt()
		spvvu.mutation.SetUpdatedAt(v)
	}
}

func (spvvu *SellerProductVariationValuesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sellerproductvariationvalues.Table,
			Columns: sellerproductvariationvalues.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sellerproductvariationvalues.FieldID,
			},
		},
	}
	if ps := spvvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spvvu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerproductvariationvalues.FieldName,
		})
	}
	if value, ok := spvvu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerproductvariationvalues.FieldDescription,
		})
	}
	if value, ok := spvvu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductvariationvalues.FieldUpdatedAt,
		})
	}
	if value, ok := spvvu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductvariationvalues.FieldDeletedAt,
		})
	}
	if spvvu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sellerproductvariationvalues.FieldDeletedAt,
		})
	}
	if spvvu.mutation.AttributeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductvariationvalues.AttributeTable,
			Columns: []string{sellerproductvariationvalues.AttributeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attribute.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spvvu.mutation.AttributeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductvariationvalues.AttributeTable,
			Columns: []string{sellerproductvariationvalues.AttributeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attribute.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if spvvu.mutation.SellerProductVariationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductvariationvalues.SellerProductVariationTable,
			Columns: []string{sellerproductvariationvalues.SellerProductVariationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sellerproductvariation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spvvu.mutation.SellerProductVariationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductvariationvalues.SellerProductVariationTable,
			Columns: []string{sellerproductvariationvalues.SellerProductVariationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sellerproductvariation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, spvvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sellerproductvariationvalues.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SellerProductVariationValuesUpdateOne is the builder for updating a single SellerProductVariationValues entity.
type SellerProductVariationValuesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SellerProductVariationValuesMutation
}

// SetName sets the "name" field.
func (spvvuo *SellerProductVariationValuesUpdateOne) SetName(s string) *SellerProductVariationValuesUpdateOne {
	spvvuo.mutation.SetName(s)
	return spvvuo
}

// SetDescription sets the "description" field.
func (spvvuo *SellerProductVariationValuesUpdateOne) SetDescription(s string) *SellerProductVariationValuesUpdateOne {
	spvvuo.mutation.SetDescription(s)
	return spvvuo
}

// SetUpdatedAt sets the "updated_at" field.
func (spvvuo *SellerProductVariationValuesUpdateOne) SetUpdatedAt(t time.Time) *SellerProductVariationValuesUpdateOne {
	spvvuo.mutation.SetUpdatedAt(t)
	return spvvuo
}

// SetDeletedAt sets the "deleted_at" field.
func (spvvuo *SellerProductVariationValuesUpdateOne) SetDeletedAt(t time.Time) *SellerProductVariationValuesUpdateOne {
	spvvuo.mutation.SetDeletedAt(t)
	return spvvuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (spvvuo *SellerProductVariationValuesUpdateOne) SetNillableDeletedAt(t *time.Time) *SellerProductVariationValuesUpdateOne {
	if t != nil {
		spvvuo.SetDeletedAt(*t)
	}
	return spvvuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (spvvuo *SellerProductVariationValuesUpdateOne) ClearDeletedAt() *SellerProductVariationValuesUpdateOne {
	spvvuo.mutation.ClearDeletedAt()
	return spvvuo
}

// SetAttributeID sets the "attribute" edge to the Attribute entity by ID.
func (spvvuo *SellerProductVariationValuesUpdateOne) SetAttributeID(id int) *SellerProductVariationValuesUpdateOne {
	spvvuo.mutation.SetAttributeID(id)
	return spvvuo
}

// SetNillableAttributeID sets the "attribute" edge to the Attribute entity by ID if the given value is not nil.
func (spvvuo *SellerProductVariationValuesUpdateOne) SetNillableAttributeID(id *int) *SellerProductVariationValuesUpdateOne {
	if id != nil {
		spvvuo = spvvuo.SetAttributeID(*id)
	}
	return spvvuo
}

// SetAttribute sets the "attribute" edge to the Attribute entity.
func (spvvuo *SellerProductVariationValuesUpdateOne) SetAttribute(a *Attribute) *SellerProductVariationValuesUpdateOne {
	return spvvuo.SetAttributeID(a.ID)
}

// SetSellerProductVariationID sets the "seller_product_variation" edge to the SellerProductVariation entity by ID.
func (spvvuo *SellerProductVariationValuesUpdateOne) SetSellerProductVariationID(id int) *SellerProductVariationValuesUpdateOne {
	spvvuo.mutation.SetSellerProductVariationID(id)
	return spvvuo
}

// SetNillableSellerProductVariationID sets the "seller_product_variation" edge to the SellerProductVariation entity by ID if the given value is not nil.
func (spvvuo *SellerProductVariationValuesUpdateOne) SetNillableSellerProductVariationID(id *int) *SellerProductVariationValuesUpdateOne {
	if id != nil {
		spvvuo = spvvuo.SetSellerProductVariationID(*id)
	}
	return spvvuo
}

// SetSellerProductVariation sets the "seller_product_variation" edge to the SellerProductVariation entity.
func (spvvuo *SellerProductVariationValuesUpdateOne) SetSellerProductVariation(s *SellerProductVariation) *SellerProductVariationValuesUpdateOne {
	return spvvuo.SetSellerProductVariationID(s.ID)
}

// Mutation returns the SellerProductVariationValuesMutation object of the builder.
func (spvvuo *SellerProductVariationValuesUpdateOne) Mutation() *SellerProductVariationValuesMutation {
	return spvvuo.mutation
}

// ClearAttribute clears the "attribute" edge to the Attribute entity.
func (spvvuo *SellerProductVariationValuesUpdateOne) ClearAttribute() *SellerProductVariationValuesUpdateOne {
	spvvuo.mutation.ClearAttribute()
	return spvvuo
}

// ClearSellerProductVariation clears the "seller_product_variation" edge to the SellerProductVariation entity.
func (spvvuo *SellerProductVariationValuesUpdateOne) ClearSellerProductVariation() *SellerProductVariationValuesUpdateOne {
	spvvuo.mutation.ClearSellerProductVariation()
	return spvvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (spvvuo *SellerProductVariationValuesUpdateOne) Select(field string, fields ...string) *SellerProductVariationValuesUpdateOne {
	spvvuo.fields = append([]string{field}, fields...)
	return spvvuo
}

// Save executes the query and returns the updated SellerProductVariationValues entity.
func (spvvuo *SellerProductVariationValuesUpdateOne) Save(ctx context.Context) (*SellerProductVariationValues, error) {
	var (
		err  error
		node *SellerProductVariationValues
	)
	spvvuo.defaults()
	if len(spvvuo.hooks) == 0 {
		node, err = spvvuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SellerProductVariationValuesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spvvuo.mutation = mutation
			node, err = spvvuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(spvvuo.hooks) - 1; i >= 0; i-- {
			if spvvuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spvvuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spvvuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (spvvuo *SellerProductVariationValuesUpdateOne) SaveX(ctx context.Context) *SellerProductVariationValues {
	node, err := spvvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (spvvuo *SellerProductVariationValuesUpdateOne) Exec(ctx context.Context) error {
	_, err := spvvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spvvuo *SellerProductVariationValuesUpdateOne) ExecX(ctx context.Context) {
	if err := spvvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spvvuo *SellerProductVariationValuesUpdateOne) defaults() {
	if _, ok := spvvuo.mutation.UpdatedAt(); !ok {
		v := sellerproductvariationvalues.UpdateDefaultUpdatedAt()
		spvvuo.mutation.SetUpdatedAt(v)
	}
}

func (spvvuo *SellerProductVariationValuesUpdateOne) sqlSave(ctx context.Context) (_node *SellerProductVariationValues, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sellerproductvariationvalues.Table,
			Columns: sellerproductvariationvalues.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sellerproductvariationvalues.FieldID,
			},
		},
	}
	id, ok := spvvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SellerProductVariationValues.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := spvvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sellerproductvariationvalues.FieldID)
		for _, f := range fields {
			if !sellerproductvariationvalues.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sellerproductvariationvalues.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := spvvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spvvuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerproductvariationvalues.FieldName,
		})
	}
	if value, ok := spvvuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sellerproductvariationvalues.FieldDescription,
		})
	}
	if value, ok := spvvuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductvariationvalues.FieldUpdatedAt,
		})
	}
	if value, ok := spvvuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sellerproductvariationvalues.FieldDeletedAt,
		})
	}
	if spvvuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sellerproductvariationvalues.FieldDeletedAt,
		})
	}
	if spvvuo.mutation.AttributeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductvariationvalues.AttributeTable,
			Columns: []string{sellerproductvariationvalues.AttributeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attribute.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spvvuo.mutation.AttributeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductvariationvalues.AttributeTable,
			Columns: []string{sellerproductvariationvalues.AttributeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attribute.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if spvvuo.mutation.SellerProductVariationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductvariationvalues.SellerProductVariationTable,
			Columns: []string{sellerproductvariationvalues.SellerProductVariationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sellerproductvariation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spvvuo.mutation.SellerProductVariationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sellerproductvariationvalues.SellerProductVariationTable,
			Columns: []string{sellerproductvariationvalues.SellerProductVariationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sellerproductvariation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SellerProductVariationValues{config: spvvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, spvvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sellerproductvariationvalues.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
