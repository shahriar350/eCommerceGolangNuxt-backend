// Code generated by entc, DO NOT EDIT.

package checkout

import (
	"bongo/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TotalPrice applies equality check predicate on the "total_price" field. It's identical to TotalPriceEQ.
func TotalPrice(v decimal.Decimal) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalPrice), v))
	})
}

// Completed applies equality check predicate on the "completed" field. It's identical to CompletedEQ.
func Completed(v bool) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompleted), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// TotalPriceEQ applies the EQ predicate on the "total_price" field.
func TotalPriceEQ(v decimal.Decimal) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalPrice), v))
	})
}

// TotalPriceNEQ applies the NEQ predicate on the "total_price" field.
func TotalPriceNEQ(v decimal.Decimal) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalPrice), v))
	})
}

// TotalPriceIn applies the In predicate on the "total_price" field.
func TotalPriceIn(vs ...decimal.Decimal) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTotalPrice), v...))
	})
}

// TotalPriceNotIn applies the NotIn predicate on the "total_price" field.
func TotalPriceNotIn(vs ...decimal.Decimal) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTotalPrice), v...))
	})
}

// TotalPriceGT applies the GT predicate on the "total_price" field.
func TotalPriceGT(v decimal.Decimal) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalPrice), v))
	})
}

// TotalPriceGTE applies the GTE predicate on the "total_price" field.
func TotalPriceGTE(v decimal.Decimal) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalPrice), v))
	})
}

// TotalPriceLT applies the LT predicate on the "total_price" field.
func TotalPriceLT(v decimal.Decimal) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalPrice), v))
	})
}

// TotalPriceLTE applies the LTE predicate on the "total_price" field.
func TotalPriceLTE(v decimal.Decimal) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalPrice), v))
	})
}

// CompletedEQ applies the EQ predicate on the "completed" field.
func CompletedEQ(v bool) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompleted), v))
	})
}

// CompletedNEQ applies the NEQ predicate on the "completed" field.
func CompletedNEQ(v bool) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCompleted), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLocation applies the HasEdge predicate on the "location" edge.
func HasLocation() predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LocationTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LocationTable, LocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationWith applies the HasEdge predicate on the "location" edge with a given conditions (other predicates).
func HasLocationWith(preds ...predicate.UserLocation) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LocationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LocationTable, LocationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCart applies the HasEdge predicate on the "cart" edge.
func HasCart() predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CartTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, CartTable, CartColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCartWith applies the HasEdge predicate on the "cart" edge with a given conditions (other predicates).
func HasCartWith(preds ...predicate.Cart) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CartInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, CartTable, CartColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCheckoutProducts applies the HasEdge predicate on the "checkout_products" edge.
func HasCheckoutProducts() predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CheckoutProductsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CheckoutProductsTable, CheckoutProductsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCheckoutProductsWith applies the HasEdge predicate on the "checkout_products" edge with a given conditions (other predicates).
func HasCheckoutProductsWith(preds ...predicate.CheckoutProduct) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CheckoutProductsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CheckoutProductsTable, CheckoutProductsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Checkout) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Checkout) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Checkout) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		p(s.Not())
	})
}
