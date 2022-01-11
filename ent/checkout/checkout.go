// Code generated by entc, DO NOT EDIT.

package checkout

import (
	"time"
)

const (
	// Label holds the string label denoting the checkout type in the database.
	Label = "checkout"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTotalPrice holds the string denoting the total_price field in the database.
	FieldTotalPrice = "total_price"
	// FieldCompleted holds the string denoting the completed field in the database.
	FieldCompleted = "completed"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeLocation holds the string denoting the location edge name in mutations.
	EdgeLocation = "location"
	// EdgeCart holds the string denoting the cart edge name in mutations.
	EdgeCart = "cart"
	// EdgeCheckoutProducts holds the string denoting the checkout_products edge name in mutations.
	EdgeCheckoutProducts = "checkout_products"
	// Table holds the table name of the checkout in the database.
	Table = "checkouts"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "checkouts"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_checkouts"
	// LocationTable is the table that holds the location relation/edge.
	LocationTable = "checkouts"
	// LocationInverseTable is the table name for the UserLocation entity.
	// It exists in this package in order to avoid circular dependency with the "userlocation" package.
	LocationInverseTable = "user_locations"
	// LocationColumn is the table column denoting the location relation/edge.
	LocationColumn = "user_location_user_locations"
	// CartTable is the table that holds the cart relation/edge.
	CartTable = "checkouts"
	// CartInverseTable is the table name for the Cart entity.
	// It exists in this package in order to avoid circular dependency with the "cart" package.
	CartInverseTable = "carts"
	// CartColumn is the table column denoting the cart relation/edge.
	CartColumn = "cart_checkout"
	// CheckoutProductsTable is the table that holds the checkout_products relation/edge.
	CheckoutProductsTable = "checkout_products"
	// CheckoutProductsInverseTable is the table name for the CheckoutProduct entity.
	// It exists in this package in order to avoid circular dependency with the "checkoutproduct" package.
	CheckoutProductsInverseTable = "checkout_products"
	// CheckoutProductsColumn is the table column denoting the checkout_products relation/edge.
	CheckoutProductsColumn = "checkout_checkout_products"
)

// Columns holds all SQL columns for checkout fields.
var Columns = []string{
	FieldID,
	FieldTotalPrice,
	FieldCompleted,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "checkouts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"cart_checkout",
	"user_checkouts",
	"user_location_user_locations",
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
	// DefaultCompleted holds the default value on creation for the "completed" field.
	DefaultCompleted bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)