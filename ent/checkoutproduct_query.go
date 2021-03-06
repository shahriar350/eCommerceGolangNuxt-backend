// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/checkout"
	"bongo/ent/checkoutproduct"
	"bongo/ent/predicate"
	"bongo/ent/sellerproduct"
	"bongo/ent/sellerproductvariation"
	"bongo/ent/user"
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CheckoutProductQuery is the builder for querying CheckoutProduct entities.
type CheckoutProductQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CheckoutProduct
	// eager-loading edges.
	withUser                   *UserQuery
	withCheckout               *CheckoutQuery
	withSeller                 *UserQuery
	withSellerProduct          *SellerProductQuery
	withSellerProductVariation *SellerProductVariationQuery
	withFKs                    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CheckoutProductQuery builder.
func (cpq *CheckoutProductQuery) Where(ps ...predicate.CheckoutProduct) *CheckoutProductQuery {
	cpq.predicates = append(cpq.predicates, ps...)
	return cpq
}

// Limit adds a limit step to the query.
func (cpq *CheckoutProductQuery) Limit(limit int) *CheckoutProductQuery {
	cpq.limit = &limit
	return cpq
}

// Offset adds an offset step to the query.
func (cpq *CheckoutProductQuery) Offset(offset int) *CheckoutProductQuery {
	cpq.offset = &offset
	return cpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cpq *CheckoutProductQuery) Unique(unique bool) *CheckoutProductQuery {
	cpq.unique = &unique
	return cpq
}

// Order adds an order step to the query.
func (cpq *CheckoutProductQuery) Order(o ...OrderFunc) *CheckoutProductQuery {
	cpq.order = append(cpq.order, o...)
	return cpq
}

// QueryUser chains the current query on the "user" edge.
func (cpq *CheckoutProductQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checkoutproduct.Table, checkoutproduct.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checkoutproduct.UserTable, checkoutproduct.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCheckout chains the current query on the "checkout" edge.
func (cpq *CheckoutProductQuery) QueryCheckout() *CheckoutQuery {
	query := &CheckoutQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checkoutproduct.Table, checkoutproduct.FieldID, selector),
			sqlgraph.To(checkout.Table, checkout.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checkoutproduct.CheckoutTable, checkoutproduct.CheckoutColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySeller chains the current query on the "seller" edge.
func (cpq *CheckoutProductQuery) QuerySeller() *UserQuery {
	query := &UserQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checkoutproduct.Table, checkoutproduct.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checkoutproduct.SellerTable, checkoutproduct.SellerColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySellerProduct chains the current query on the "seller_product" edge.
func (cpq *CheckoutProductQuery) QuerySellerProduct() *SellerProductQuery {
	query := &SellerProductQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checkoutproduct.Table, checkoutproduct.FieldID, selector),
			sqlgraph.To(sellerproduct.Table, sellerproduct.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checkoutproduct.SellerProductTable, checkoutproduct.SellerProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySellerProductVariation chains the current query on the "seller_product_variation" edge.
func (cpq *CheckoutProductQuery) QuerySellerProductVariation() *SellerProductVariationQuery {
	query := &SellerProductVariationQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checkoutproduct.Table, checkoutproduct.FieldID, selector),
			sqlgraph.To(sellerproductvariation.Table, sellerproductvariation.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checkoutproduct.SellerProductVariationTable, checkoutproduct.SellerProductVariationColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CheckoutProduct entity from the query.
// Returns a *NotFoundError when no CheckoutProduct was found.
func (cpq *CheckoutProductQuery) First(ctx context.Context) (*CheckoutProduct, error) {
	nodes, err := cpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{checkoutproduct.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cpq *CheckoutProductQuery) FirstX(ctx context.Context) *CheckoutProduct {
	node, err := cpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CheckoutProduct ID from the query.
// Returns a *NotFoundError when no CheckoutProduct ID was found.
func (cpq *CheckoutProductQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{checkoutproduct.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cpq *CheckoutProductQuery) FirstIDX(ctx context.Context) int {
	id, err := cpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CheckoutProduct entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CheckoutProduct entity is not found.
// Returns a *NotFoundError when no CheckoutProduct entities are found.
func (cpq *CheckoutProductQuery) Only(ctx context.Context) (*CheckoutProduct, error) {
	nodes, err := cpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{checkoutproduct.Label}
	default:
		return nil, &NotSingularError{checkoutproduct.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cpq *CheckoutProductQuery) OnlyX(ctx context.Context) *CheckoutProduct {
	node, err := cpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CheckoutProduct ID in the query.
// Returns a *NotSingularError when exactly one CheckoutProduct ID is not found.
// Returns a *NotFoundError when no entities are found.
func (cpq *CheckoutProductQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{checkoutproduct.Label}
	default:
		err = &NotSingularError{checkoutproduct.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cpq *CheckoutProductQuery) OnlyIDX(ctx context.Context) int {
	id, err := cpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CheckoutProducts.
func (cpq *CheckoutProductQuery) All(ctx context.Context) ([]*CheckoutProduct, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cpq *CheckoutProductQuery) AllX(ctx context.Context) []*CheckoutProduct {
	nodes, err := cpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CheckoutProduct IDs.
func (cpq *CheckoutProductQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cpq.Select(checkoutproduct.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cpq *CheckoutProductQuery) IDsX(ctx context.Context) []int {
	ids, err := cpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cpq *CheckoutProductQuery) Count(ctx context.Context) (int, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cpq *CheckoutProductQuery) CountX(ctx context.Context) int {
	count, err := cpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cpq *CheckoutProductQuery) Exist(ctx context.Context) (bool, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cpq *CheckoutProductQuery) ExistX(ctx context.Context) bool {
	exist, err := cpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CheckoutProductQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cpq *CheckoutProductQuery) Clone() *CheckoutProductQuery {
	if cpq == nil {
		return nil
	}
	return &CheckoutProductQuery{
		config:                     cpq.config,
		limit:                      cpq.limit,
		offset:                     cpq.offset,
		order:                      append([]OrderFunc{}, cpq.order...),
		predicates:                 append([]predicate.CheckoutProduct{}, cpq.predicates...),
		withUser:                   cpq.withUser.Clone(),
		withCheckout:               cpq.withCheckout.Clone(),
		withSeller:                 cpq.withSeller.Clone(),
		withSellerProduct:          cpq.withSellerProduct.Clone(),
		withSellerProductVariation: cpq.withSellerProductVariation.Clone(),
		// clone intermediate query.
		sql:  cpq.sql.Clone(),
		path: cpq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *CheckoutProductQuery) WithUser(opts ...func(*UserQuery)) *CheckoutProductQuery {
	query := &UserQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withUser = query
	return cpq
}

// WithCheckout tells the query-builder to eager-load the nodes that are connected to
// the "checkout" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *CheckoutProductQuery) WithCheckout(opts ...func(*CheckoutQuery)) *CheckoutProductQuery {
	query := &CheckoutQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withCheckout = query
	return cpq
}

// WithSeller tells the query-builder to eager-load the nodes that are connected to
// the "seller" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *CheckoutProductQuery) WithSeller(opts ...func(*UserQuery)) *CheckoutProductQuery {
	query := &UserQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withSeller = query
	return cpq
}

// WithSellerProduct tells the query-builder to eager-load the nodes that are connected to
// the "seller_product" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *CheckoutProductQuery) WithSellerProduct(opts ...func(*SellerProductQuery)) *CheckoutProductQuery {
	query := &SellerProductQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withSellerProduct = query
	return cpq
}

// WithSellerProductVariation tells the query-builder to eager-load the nodes that are connected to
// the "seller_product_variation" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *CheckoutProductQuery) WithSellerProductVariation(opts ...func(*SellerProductVariationQuery)) *CheckoutProductQuery {
	query := &SellerProductVariationQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withSellerProductVariation = query
	return cpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Quantity int `json:"quantity,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CheckoutProduct.Query().
//		GroupBy(checkoutproduct.FieldQuantity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cpq *CheckoutProductQuery) GroupBy(field string, fields ...string) *CheckoutProductGroupBy {
	group := &CheckoutProductGroupBy{config: cpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Quantity int `json:"quantity,omitempty"`
//	}
//
//	client.CheckoutProduct.Query().
//		Select(checkoutproduct.FieldQuantity).
//		Scan(ctx, &v)
//
func (cpq *CheckoutProductQuery) Select(fields ...string) *CheckoutProductSelect {
	cpq.fields = append(cpq.fields, fields...)
	return &CheckoutProductSelect{CheckoutProductQuery: cpq}
}

func (cpq *CheckoutProductQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cpq.fields {
		if !checkoutproduct.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cpq.path != nil {
		prev, err := cpq.path(ctx)
		if err != nil {
			return err
		}
		cpq.sql = prev
	}
	return nil
}

func (cpq *CheckoutProductQuery) sqlAll(ctx context.Context) ([]*CheckoutProduct, error) {
	var (
		nodes       = []*CheckoutProduct{}
		withFKs     = cpq.withFKs
		_spec       = cpq.querySpec()
		loadedTypes = [5]bool{
			cpq.withUser != nil,
			cpq.withCheckout != nil,
			cpq.withSeller != nil,
			cpq.withSellerProduct != nil,
			cpq.withSellerProductVariation != nil,
		}
	)
	if cpq.withUser != nil || cpq.withCheckout != nil || cpq.withSeller != nil || cpq.withSellerProduct != nil || cpq.withSellerProductVariation != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, checkoutproduct.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CheckoutProduct{config: cpq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, cpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := cpq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CheckoutProduct)
		for i := range nodes {
			if nodes[i].user_checkout_products == nil {
				continue
			}
			fk := *nodes[i].user_checkout_products
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_checkout_products" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := cpq.withCheckout; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CheckoutProduct)
		for i := range nodes {
			if nodes[i].checkout_checkout_products == nil {
				continue
			}
			fk := *nodes[i].checkout_checkout_products
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(checkout.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "checkout_checkout_products" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Checkout = n
			}
		}
	}

	if query := cpq.withSeller; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CheckoutProduct)
		for i := range nodes {
			if nodes[i].user_seller_checkout_products == nil {
				continue
			}
			fk := *nodes[i].user_seller_checkout_products
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_seller_checkout_products" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Seller = n
			}
		}
	}

	if query := cpq.withSellerProduct; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CheckoutProduct)
		for i := range nodes {
			if nodes[i].seller_product_checkout_products == nil {
				continue
			}
			fk := *nodes[i].seller_product_checkout_products
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(sellerproduct.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "seller_product_checkout_products" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.SellerProduct = n
			}
		}
	}

	if query := cpq.withSellerProductVariation; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CheckoutProduct)
		for i := range nodes {
			if nodes[i].seller_product_variation_checkout_products == nil {
				continue
			}
			fk := *nodes[i].seller_product_variation_checkout_products
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(sellerproductvariation.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "seller_product_variation_checkout_products" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.SellerProductVariation = n
			}
		}
	}

	return nodes, nil
}

func (cpq *CheckoutProductQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cpq.querySpec()
	_spec.Node.Columns = cpq.fields
	if len(cpq.fields) > 0 {
		_spec.Unique = cpq.unique != nil && *cpq.unique
	}
	return sqlgraph.CountNodes(ctx, cpq.driver, _spec)
}

func (cpq *CheckoutProductQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cpq *CheckoutProductQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   checkoutproduct.Table,
			Columns: checkoutproduct.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checkoutproduct.FieldID,
			},
		},
		From:   cpq.sql,
		Unique: true,
	}
	if unique := cpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, checkoutproduct.FieldID)
		for i := range fields {
			if fields[i] != checkoutproduct.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cpq *CheckoutProductQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cpq.driver.Dialect())
	t1 := builder.Table(checkoutproduct.Table)
	columns := cpq.fields
	if len(columns) == 0 {
		columns = checkoutproduct.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cpq.sql != nil {
		selector = cpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cpq.unique != nil && *cpq.unique {
		selector.Distinct()
	}
	for _, p := range cpq.predicates {
		p(selector)
	}
	for _, p := range cpq.order {
		p(selector)
	}
	if offset := cpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CheckoutProductGroupBy is the group-by builder for CheckoutProduct entities.
type CheckoutProductGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cpgb *CheckoutProductGroupBy) Aggregate(fns ...AggregateFunc) *CheckoutProductGroupBy {
	cpgb.fns = append(cpgb.fns, fns...)
	return cpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cpgb *CheckoutProductGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cpgb.path(ctx)
	if err != nil {
		return err
	}
	cpgb.sql = query
	return cpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cpgb *CheckoutProductGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CheckoutProductGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CheckoutProductGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cpgb *CheckoutProductGroupBy) StringsX(ctx context.Context) []string {
	v, err := cpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CheckoutProductGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkoutproduct.Label}
	default:
		err = fmt.Errorf("ent: CheckoutProductGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cpgb *CheckoutProductGroupBy) StringX(ctx context.Context) string {
	v, err := cpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CheckoutProductGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CheckoutProductGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cpgb *CheckoutProductGroupBy) IntsX(ctx context.Context) []int {
	v, err := cpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CheckoutProductGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkoutproduct.Label}
	default:
		err = fmt.Errorf("ent: CheckoutProductGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cpgb *CheckoutProductGroupBy) IntX(ctx context.Context) int {
	v, err := cpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CheckoutProductGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CheckoutProductGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cpgb *CheckoutProductGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CheckoutProductGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkoutproduct.Label}
	default:
		err = fmt.Errorf("ent: CheckoutProductGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cpgb *CheckoutProductGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CheckoutProductGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CheckoutProductGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cpgb *CheckoutProductGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CheckoutProductGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkoutproduct.Label}
	default:
		err = fmt.Errorf("ent: CheckoutProductGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cpgb *CheckoutProductGroupBy) BoolX(ctx context.Context) bool {
	v, err := cpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cpgb *CheckoutProductGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cpgb.fields {
		if !checkoutproduct.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cpgb *CheckoutProductGroupBy) sqlQuery() *sql.Selector {
	selector := cpgb.sql.Select()
	aggregation := make([]string, 0, len(cpgb.fns))
	for _, fn := range cpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cpgb.fields)+len(cpgb.fns))
		for _, f := range cpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cpgb.fields...)...)
}

// CheckoutProductSelect is the builder for selecting fields of CheckoutProduct entities.
type CheckoutProductSelect struct {
	*CheckoutProductQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cps *CheckoutProductSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cps.prepareQuery(ctx); err != nil {
		return err
	}
	cps.sql = cps.CheckoutProductQuery.sqlQuery(ctx)
	return cps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cps *CheckoutProductSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cps *CheckoutProductSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CheckoutProductSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cps *CheckoutProductSelect) StringsX(ctx context.Context) []string {
	v, err := cps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cps *CheckoutProductSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkoutproduct.Label}
	default:
		err = fmt.Errorf("ent: CheckoutProductSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cps *CheckoutProductSelect) StringX(ctx context.Context) string {
	v, err := cps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cps *CheckoutProductSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CheckoutProductSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cps *CheckoutProductSelect) IntsX(ctx context.Context) []int {
	v, err := cps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cps *CheckoutProductSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkoutproduct.Label}
	default:
		err = fmt.Errorf("ent: CheckoutProductSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cps *CheckoutProductSelect) IntX(ctx context.Context) int {
	v, err := cps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cps *CheckoutProductSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CheckoutProductSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cps *CheckoutProductSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cps *CheckoutProductSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkoutproduct.Label}
	default:
		err = fmt.Errorf("ent: CheckoutProductSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cps *CheckoutProductSelect) Float64X(ctx context.Context) float64 {
	v, err := cps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cps *CheckoutProductSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CheckoutProductSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cps *CheckoutProductSelect) BoolsX(ctx context.Context) []bool {
	v, err := cps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cps *CheckoutProductSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkoutproduct.Label}
	default:
		err = fmt.Errorf("ent: CheckoutProductSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cps *CheckoutProductSelect) BoolX(ctx context.Context) bool {
	v, err := cps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cps *CheckoutProductSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cps.sql.Query()
	if err := cps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
