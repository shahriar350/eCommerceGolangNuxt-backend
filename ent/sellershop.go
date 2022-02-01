// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/sellershop"
	"bongo/ent/shopcategory"
	"bongo/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// SellerShop is the model entity for the SellerShop schema.
type SellerShop struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// ContactNumber holds the value of the "contact_number" field.
	ContactNumber string `json:"contact_number,omitempty"`
	// Banner holds the value of the "banner" field.
	Banner string `json:"banner,omitempty"`
	// BusinessLocation holds the value of the "business_location" field.
	BusinessLocation string `json:"business_location,omitempty"`
	// TaxID holds the value of the "tax_id" field.
	TaxID string `json:"tax_id,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SellerShopQuery when eager-loading is set.
	Edges                      SellerShopEdges `json:"edges"`
	shop_category_seller_shops *int
	user_seller_shops          *int
	user_approved_shops        *int
}

// SellerShopEdges holds the relations/edges for other nodes in the graph.
type SellerShopEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Admin holds the value of the admin edge.
	Admin *User `json:"admin,omitempty"`
	// GetShopCategory holds the value of the get_shop_category edge.
	GetShopCategory *ShopCategory `json:"get_shop_category,omitempty"`
	// SellerProducts holds the value of the seller_products edge.
	SellerProducts []*SellerProduct `json:"seller_products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SellerShopEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// AdminOrErr returns the Admin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SellerShopEdges) AdminOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Admin == nil {
			// The edge admin was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Admin, nil
	}
	return nil, &NotLoadedError{edge: "admin"}
}

// GetShopCategoryOrErr returns the GetShopCategory value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SellerShopEdges) GetShopCategoryOrErr() (*ShopCategory, error) {
	if e.loadedTypes[2] {
		if e.GetShopCategory == nil {
			// The edge get_shop_category was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: shopcategory.Label}
		}
		return e.GetShopCategory, nil
	}
	return nil, &NotLoadedError{edge: "get_shop_category"}
}

// SellerProductsOrErr returns the SellerProducts value or an error if the edge
// was not loaded in eager-loading.
func (e SellerShopEdges) SellerProductsOrErr() ([]*SellerProduct, error) {
	if e.loadedTypes[3] {
		return e.SellerProducts, nil
	}
	return nil, &NotLoadedError{edge: "seller_products"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SellerShop) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sellershop.FieldActive:
			values[i] = new(sql.NullBool)
		case sellershop.FieldID:
			values[i] = new(sql.NullInt64)
		case sellershop.FieldName, sellershop.FieldSlug, sellershop.FieldContactNumber, sellershop.FieldBanner, sellershop.FieldBusinessLocation, sellershop.FieldTaxID:
			values[i] = new(sql.NullString)
		case sellershop.FieldCreatedAt, sellershop.FieldUpdatedAt, sellershop.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case sellershop.ForeignKeys[0]: // shop_category_seller_shops
			values[i] = new(sql.NullInt64)
		case sellershop.ForeignKeys[1]: // user_seller_shops
			values[i] = new(sql.NullInt64)
		case sellershop.ForeignKeys[2]: // user_approved_shops
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SellerShop", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SellerShop fields.
func (ss *SellerShop) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sellershop.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ss.ID = int(value.Int64)
		case sellershop.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ss.Name = value.String
			}
		case sellershop.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				ss.Slug = value.String
			}
		case sellershop.FieldContactNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact_number", values[i])
			} else if value.Valid {
				ss.ContactNumber = value.String
			}
		case sellershop.FieldBanner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner", values[i])
			} else if value.Valid {
				ss.Banner = value.String
			}
		case sellershop.FieldBusinessLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_location", values[i])
			} else if value.Valid {
				ss.BusinessLocation = value.String
			}
		case sellershop.FieldTaxID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tax_id", values[i])
			} else if value.Valid {
				ss.TaxID = value.String
			}
		case sellershop.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				ss.Active = value.Bool
			}
		case sellershop.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ss.CreatedAt = value.Time
			}
		case sellershop.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ss.UpdatedAt = value.Time
			}
		case sellershop.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ss.DeletedAt = new(time.Time)
				*ss.DeletedAt = value.Time
			}
		case sellershop.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field shop_category_seller_shops", value)
			} else if value.Valid {
				ss.shop_category_seller_shops = new(int)
				*ss.shop_category_seller_shops = int(value.Int64)
			}
		case sellershop.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_seller_shops", value)
			} else if value.Valid {
				ss.user_seller_shops = new(int)
				*ss.user_seller_shops = int(value.Int64)
			}
		case sellershop.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_approved_shops", value)
			} else if value.Valid {
				ss.user_approved_shops = new(int)
				*ss.user_approved_shops = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the SellerShop entity.
func (ss *SellerShop) QueryUser() *UserQuery {
	return (&SellerShopClient{config: ss.config}).QueryUser(ss)
}

// QueryAdmin queries the "admin" edge of the SellerShop entity.
func (ss *SellerShop) QueryAdmin() *UserQuery {
	return (&SellerShopClient{config: ss.config}).QueryAdmin(ss)
}

// QueryGetShopCategory queries the "get_shop_category" edge of the SellerShop entity.
func (ss *SellerShop) QueryGetShopCategory() *ShopCategoryQuery {
	return (&SellerShopClient{config: ss.config}).QueryGetShopCategory(ss)
}

// QuerySellerProducts queries the "seller_products" edge of the SellerShop entity.
func (ss *SellerShop) QuerySellerProducts() *SellerProductQuery {
	return (&SellerShopClient{config: ss.config}).QuerySellerProducts(ss)
}

// Update returns a builder for updating this SellerShop.
// Note that you need to call SellerShop.Unwrap() before calling this method if this SellerShop
// was returned from a transaction, and the transaction was committed or rolled back.
func (ss *SellerShop) Update() *SellerShopUpdateOne {
	return (&SellerShopClient{config: ss.config}).UpdateOne(ss)
}

// Unwrap unwraps the SellerShop entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ss *SellerShop) Unwrap() *SellerShop {
	tx, ok := ss.config.driver.(*txDriver)
	if !ok {
		panic("ent: SellerShop is not a transactional entity")
	}
	ss.config.driver = tx.drv
	return ss
}

// String implements the fmt.Stringer.
func (ss *SellerShop) String() string {
	var builder strings.Builder
	builder.WriteString("SellerShop(")
	builder.WriteString(fmt.Sprintf("id=%v", ss.ID))
	builder.WriteString(", name=")
	builder.WriteString(ss.Name)
	builder.WriteString(", slug=")
	builder.WriteString(ss.Slug)
	builder.WriteString(", contact_number=")
	builder.WriteString(ss.ContactNumber)
	builder.WriteString(", banner=")
	builder.WriteString(ss.Banner)
	builder.WriteString(", business_location=")
	builder.WriteString(ss.BusinessLocation)
	builder.WriteString(", tax_id=")
	builder.WriteString(ss.TaxID)
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", ss.Active))
	builder.WriteString(", created_at=")
	builder.WriteString(ss.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ss.UpdatedAt.Format(time.ANSIC))
	if v := ss.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// SellerShops is a parsable slice of SellerShop.
type SellerShops []*SellerShop

func (ss SellerShops) config(cfg config) {
	for _i := range ss {
		ss[_i].config = cfg
	}
}
