// Code generated by entc, DO NOT EDIT.

package sellershopproduct

import (
	"time"
)

const (
	// Label holds the string label denoting the sellershopproduct type in the database.
	Label = "seller_shop_product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeSellerShop holds the string denoting the seller_shop edge name in mutations.
	EdgeSellerShop = "seller_shop"
	// EdgeSellerProduct holds the string denoting the seller_product edge name in mutations.
	EdgeSellerProduct = "seller_product"
	// Table holds the table name of the sellershopproduct in the database.
	Table = "seller_shop_products"
	// SellerShopTable is the table that holds the seller_shop relation/edge.
	SellerShopTable = "seller_shop_products"
	// SellerShopInverseTable is the table name for the SellerShop entity.
	// It exists in this package in order to avoid circular dependency with the "sellershop" package.
	SellerShopInverseTable = "seller_shops"
	// SellerShopColumn is the table column denoting the seller_shop relation/edge.
	SellerShopColumn = "seller_shop_seller_shop_products"
	// SellerProductTable is the table that holds the seller_product relation/edge.
	SellerProductTable = "seller_shop_products"
	// SellerProductInverseTable is the table name for the SellerProduct entity.
	// It exists in this package in order to avoid circular dependency with the "sellerproduct" package.
	SellerProductInverseTable = "seller_products"
	// SellerProductColumn is the table column denoting the seller_product relation/edge.
	SellerProductColumn = "seller_product_seller_shop_products"
)

// Columns holds all SQL columns for sellershopproduct fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "seller_shop_products"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"seller_product_seller_shop_products",
	"seller_shop_seller_shop_products",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)