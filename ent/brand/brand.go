// Code generated by entc, DO NOT EDIT.

package brand

import (
	"time"
)

const (
	// Label holds the string label denoting the brand type in the database.
	Label = "brand"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeBrand holds the string denoting the brand edge name in mutations.
	EdgeBrand = "brand"
	// Table holds the table name of the brand in the database.
	Table = "brands"
	// BrandTable is the table that holds the brand relation/edge.
	BrandTable = "seller_products"
	// BrandInverseTable is the table name for the SellerProduct entity.
	// It exists in this package in order to avoid circular dependency with the "sellerproduct" package.
	BrandInverseTable = "seller_products"
	// BrandColumn is the table column denoting the brand relation/edge.
	BrandColumn = "brand_brand"
)

// Columns holds all SQL columns for brand fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
