// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/brand"
	"bongo/ent/sellerproduct"
	"bongo/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shopspring/decimal"
)

// SellerProduct is the model entity for the SellerProduct schema.
type SellerProduct struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// SellingPrice holds the value of the "selling_price" field.
	SellingPrice decimal.Decimal `json:"selling_price,omitempty"`
	// ProductPrice holds the value of the "product_price" field.
	ProductPrice decimal.Decimal `json:"product_price,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// OfferPrice holds the value of the "offer_price" field.
	OfferPrice *int `json:"offer_price,omitempty"`
	// OfferPriceStart holds the value of the "offer_price_start" field.
	OfferPriceStart *time.Time `json:"offer_price_start,omitempty"`
	// OfferPriceEnd holds the value of the "offer_price_end" field.
	OfferPriceEnd *time.Time `json:"offer_price_end,omitempty"`
	// NextStock holds the value of the "next_stock" field.
	NextStock *string `json:"next_stock,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SellerProductQuery when eager-loading is set.
	Edges                       SellerProductEdges `json:"edges"`
	brand_brand                 *int
	seller_shop_seller_products *int
	user_seller_products        *int
}

// SellerProductEdges holds the relations/edges for other nodes in the graph.
type SellerProductEdges struct {
	// Brand holds the value of the brand edge.
	Brand *Brand `json:"brand,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// SellerProductImages holds the value of the seller_product_images edge.
	SellerProductImages []*SellerProductImage `json:"seller_product_images,omitempty"`
	// SellerProductCategories holds the value of the seller_product_categories edge.
	SellerProductCategories []*SellerProductCategory `json:"seller_product_categories,omitempty"`
	// CartProducts holds the value of the cart_products edge.
	CartProducts []*CartProduct `json:"cart_products,omitempty"`
	// CheckoutProducts holds the value of the checkout_products edge.
	CheckoutProducts []*CheckoutProduct `json:"checkout_products,omitempty"`
	// SellerProductVariations holds the value of the seller_product_variations edge.
	SellerProductVariations []*SellerProductVariation `json:"seller_product_variations,omitempty"`
	// SellerShopProducts holds the value of the seller_shop_products edge.
	SellerShopProducts []*SellerShopProduct `json:"seller_shop_products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// BrandOrErr returns the Brand value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SellerProductEdges) BrandOrErr() (*Brand, error) {
	if e.loadedTypes[0] {
		if e.Brand == nil {
			// The edge brand was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: brand.Label}
		}
		return e.Brand, nil
	}
	return nil, &NotLoadedError{edge: "brand"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SellerProductEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// SellerProductImagesOrErr returns the SellerProductImages value or an error if the edge
// was not loaded in eager-loading.
func (e SellerProductEdges) SellerProductImagesOrErr() ([]*SellerProductImage, error) {
	if e.loadedTypes[2] {
		return e.SellerProductImages, nil
	}
	return nil, &NotLoadedError{edge: "seller_product_images"}
}

// SellerProductCategoriesOrErr returns the SellerProductCategories value or an error if the edge
// was not loaded in eager-loading.
func (e SellerProductEdges) SellerProductCategoriesOrErr() ([]*SellerProductCategory, error) {
	if e.loadedTypes[3] {
		return e.SellerProductCategories, nil
	}
	return nil, &NotLoadedError{edge: "seller_product_categories"}
}

// CartProductsOrErr returns the CartProducts value or an error if the edge
// was not loaded in eager-loading.
func (e SellerProductEdges) CartProductsOrErr() ([]*CartProduct, error) {
	if e.loadedTypes[4] {
		return e.CartProducts, nil
	}
	return nil, &NotLoadedError{edge: "cart_products"}
}

// CheckoutProductsOrErr returns the CheckoutProducts value or an error if the edge
// was not loaded in eager-loading.
func (e SellerProductEdges) CheckoutProductsOrErr() ([]*CheckoutProduct, error) {
	if e.loadedTypes[5] {
		return e.CheckoutProducts, nil
	}
	return nil, &NotLoadedError{edge: "checkout_products"}
}

// SellerProductVariationsOrErr returns the SellerProductVariations value or an error if the edge
// was not loaded in eager-loading.
func (e SellerProductEdges) SellerProductVariationsOrErr() ([]*SellerProductVariation, error) {
	if e.loadedTypes[6] {
		return e.SellerProductVariations, nil
	}
	return nil, &NotLoadedError{edge: "seller_product_variations"}
}

// SellerShopProductsOrErr returns the SellerShopProducts value or an error if the edge
// was not loaded in eager-loading.
func (e SellerProductEdges) SellerShopProductsOrErr() ([]*SellerShopProduct, error) {
	if e.loadedTypes[7] {
		return e.SellerShopProducts, nil
	}
	return nil, &NotLoadedError{edge: "seller_shop_products"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SellerProduct) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sellerproduct.FieldSellingPrice, sellerproduct.FieldProductPrice:
			values[i] = new(decimal.Decimal)
		case sellerproduct.FieldActive:
			values[i] = new(sql.NullBool)
		case sellerproduct.FieldID, sellerproduct.FieldQuantity, sellerproduct.FieldOfferPrice:
			values[i] = new(sql.NullInt64)
		case sellerproduct.FieldName, sellerproduct.FieldSlug, sellerproduct.FieldDescription, sellerproduct.FieldNextStock:
			values[i] = new(sql.NullString)
		case sellerproduct.FieldOfferPriceStart, sellerproduct.FieldOfferPriceEnd, sellerproduct.FieldCreatedAt, sellerproduct.FieldUpdatedAt, sellerproduct.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case sellerproduct.ForeignKeys[0]: // brand_brand
			values[i] = new(sql.NullInt64)
		case sellerproduct.ForeignKeys[1]: // seller_shop_seller_products
			values[i] = new(sql.NullInt64)
		case sellerproduct.ForeignKeys[2]: // user_seller_products
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SellerProduct", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SellerProduct fields.
func (sp *SellerProduct) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sellerproduct.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sp.ID = int(value.Int64)
		case sellerproduct.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sp.Name = value.String
			}
		case sellerproduct.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				sp.Slug = value.String
			}
		case sellerproduct.FieldSellingPrice:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field selling_price", values[i])
			} else if value != nil {
				sp.SellingPrice = *value
			}
		case sellerproduct.FieldProductPrice:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field product_price", values[i])
			} else if value != nil {
				sp.ProductPrice = *value
			}
		case sellerproduct.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				sp.Quantity = int(value.Int64)
			}
		case sellerproduct.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				sp.Active = value.Bool
			}
		case sellerproduct.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				sp.Description = new(string)
				*sp.Description = value.String
			}
		case sellerproduct.FieldOfferPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field offer_price", values[i])
			} else if value.Valid {
				sp.OfferPrice = new(int)
				*sp.OfferPrice = int(value.Int64)
			}
		case sellerproduct.FieldOfferPriceStart:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field offer_price_start", values[i])
			} else if value.Valid {
				sp.OfferPriceStart = new(time.Time)
				*sp.OfferPriceStart = value.Time
			}
		case sellerproduct.FieldOfferPriceEnd:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field offer_price_end", values[i])
			} else if value.Valid {
				sp.OfferPriceEnd = new(time.Time)
				*sp.OfferPriceEnd = value.Time
			}
		case sellerproduct.FieldNextStock:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field next_stock", values[i])
			} else if value.Valid {
				sp.NextStock = new(string)
				*sp.NextStock = value.String
			}
		case sellerproduct.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sp.CreatedAt = value.Time
			}
		case sellerproduct.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sp.UpdatedAt = value.Time
			}
		case sellerproduct.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sp.DeletedAt = new(time.Time)
				*sp.DeletedAt = value.Time
			}
		case sellerproduct.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field brand_brand", value)
			} else if value.Valid {
				sp.brand_brand = new(int)
				*sp.brand_brand = int(value.Int64)
			}
		case sellerproduct.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field seller_shop_seller_products", value)
			} else if value.Valid {
				sp.seller_shop_seller_products = new(int)
				*sp.seller_shop_seller_products = int(value.Int64)
			}
		case sellerproduct.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_seller_products", value)
			} else if value.Valid {
				sp.user_seller_products = new(int)
				*sp.user_seller_products = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryBrand queries the "brand" edge of the SellerProduct entity.
func (sp *SellerProduct) QueryBrand() *BrandQuery {
	return (&SellerProductClient{config: sp.config}).QueryBrand(sp)
}

// QueryUser queries the "user" edge of the SellerProduct entity.
func (sp *SellerProduct) QueryUser() *UserQuery {
	return (&SellerProductClient{config: sp.config}).QueryUser(sp)
}

// QuerySellerProductImages queries the "seller_product_images" edge of the SellerProduct entity.
func (sp *SellerProduct) QuerySellerProductImages() *SellerProductImageQuery {
	return (&SellerProductClient{config: sp.config}).QuerySellerProductImages(sp)
}

// QuerySellerProductCategories queries the "seller_product_categories" edge of the SellerProduct entity.
func (sp *SellerProduct) QuerySellerProductCategories() *SellerProductCategoryQuery {
	return (&SellerProductClient{config: sp.config}).QuerySellerProductCategories(sp)
}

// QueryCartProducts queries the "cart_products" edge of the SellerProduct entity.
func (sp *SellerProduct) QueryCartProducts() *CartProductQuery {
	return (&SellerProductClient{config: sp.config}).QueryCartProducts(sp)
}

// QueryCheckoutProducts queries the "checkout_products" edge of the SellerProduct entity.
func (sp *SellerProduct) QueryCheckoutProducts() *CheckoutProductQuery {
	return (&SellerProductClient{config: sp.config}).QueryCheckoutProducts(sp)
}

// QuerySellerProductVariations queries the "seller_product_variations" edge of the SellerProduct entity.
func (sp *SellerProduct) QuerySellerProductVariations() *SellerProductVariationQuery {
	return (&SellerProductClient{config: sp.config}).QuerySellerProductVariations(sp)
}

// QuerySellerShopProducts queries the "seller_shop_products" edge of the SellerProduct entity.
func (sp *SellerProduct) QuerySellerShopProducts() *SellerShopProductQuery {
	return (&SellerProductClient{config: sp.config}).QuerySellerShopProducts(sp)
}

// Update returns a builder for updating this SellerProduct.
// Note that you need to call SellerProduct.Unwrap() before calling this method if this SellerProduct
// was returned from a transaction, and the transaction was committed or rolled back.
func (sp *SellerProduct) Update() *SellerProductUpdateOne {
	return (&SellerProductClient{config: sp.config}).UpdateOne(sp)
}

// Unwrap unwraps the SellerProduct entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sp *SellerProduct) Unwrap() *SellerProduct {
	tx, ok := sp.config.driver.(*txDriver)
	if !ok {
		panic("ent: SellerProduct is not a transactional entity")
	}
	sp.config.driver = tx.drv
	return sp
}

// String implements the fmt.Stringer.
func (sp *SellerProduct) String() string {
	var builder strings.Builder
	builder.WriteString("SellerProduct(")
	builder.WriteString(fmt.Sprintf("id=%v", sp.ID))
	builder.WriteString(", name=")
	builder.WriteString(sp.Name)
	builder.WriteString(", slug=")
	builder.WriteString(sp.Slug)
	builder.WriteString(", selling_price=")
	builder.WriteString(fmt.Sprintf("%v", sp.SellingPrice))
	builder.WriteString(", product_price=")
	builder.WriteString(fmt.Sprintf("%v", sp.ProductPrice))
	builder.WriteString(", quantity=")
	builder.WriteString(fmt.Sprintf("%v", sp.Quantity))
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", sp.Active))
	if v := sp.Description; v != nil {
		builder.WriteString(", description=")
		builder.WriteString(*v)
	}
	if v := sp.OfferPrice; v != nil {
		builder.WriteString(", offer_price=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := sp.OfferPriceStart; v != nil {
		builder.WriteString(", offer_price_start=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := sp.OfferPriceEnd; v != nil {
		builder.WriteString(", offer_price_end=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := sp.NextStock; v != nil {
		builder.WriteString(", next_stock=")
		builder.WriteString(*v)
	}
	builder.WriteString(", created_at=")
	builder.WriteString(sp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(sp.UpdatedAt.Format(time.ANSIC))
	if v := sp.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// SellerProducts is a parsable slice of SellerProduct.
type SellerProducts []*SellerProduct

func (sp SellerProducts) config(cfg config) {
	for _i := range sp {
		sp[_i].config = cfg
	}
}
