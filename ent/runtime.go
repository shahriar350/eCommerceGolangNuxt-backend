// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/attribute"
	"bongo/ent/brand"
	"bongo/ent/cart"
	"bongo/ent/cartproduct"
	"bongo/ent/category"
	"bongo/ent/checkout"
	"bongo/ent/checkoutproduct"
	"bongo/ent/schema"
	"bongo/ent/sellerproduct"
	"bongo/ent/sellerproductcategory"
	"bongo/ent/sellerproductimage"
	"bongo/ent/sellerproductvariation"
	"bongo/ent/sellerproductvariationvalues"
	"bongo/ent/sellerrequest"
	"bongo/ent/sellershop"
	"bongo/ent/sellershopproduct"
	"bongo/ent/shopcategory"
	"bongo/ent/user"
	"bongo/ent/userlocation"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	attributeFields := schema.Attribute{}.Fields()
	_ = attributeFields
	// attributeDescCreatedAt is the schema descriptor for created_at field.
	attributeDescCreatedAt := attributeFields[1].Descriptor()
	// attribute.DefaultCreatedAt holds the default value on creation for the created_at field.
	attribute.DefaultCreatedAt = attributeDescCreatedAt.Default.(func() time.Time)
	// attributeDescUpdatedAt is the schema descriptor for updated_at field.
	attributeDescUpdatedAt := attributeFields[2].Descriptor()
	// attribute.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	attribute.DefaultUpdatedAt = attributeDescUpdatedAt.Default.(func() time.Time)
	// attribute.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	attribute.UpdateDefaultUpdatedAt = attributeDescUpdatedAt.UpdateDefault.(func() time.Time)
	brandFields := schema.Brand{}.Fields()
	_ = brandFields
	// brandDescCreatedAt is the schema descriptor for created_at field.
	brandDescCreatedAt := brandFields[1].Descriptor()
	// brand.DefaultCreatedAt holds the default value on creation for the created_at field.
	brand.DefaultCreatedAt = brandDescCreatedAt.Default.(func() time.Time)
	// brandDescUpdatedAt is the schema descriptor for updated_at field.
	brandDescUpdatedAt := brandFields[2].Descriptor()
	// brand.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	brand.DefaultUpdatedAt = brandDescUpdatedAt.Default.(func() time.Time)
	// brand.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	brand.UpdateDefaultUpdatedAt = brandDescUpdatedAt.UpdateDefault.(func() time.Time)
	cartFields := schema.Cart{}.Fields()
	_ = cartFields
	// cartDescCreatedAt is the schema descriptor for created_at field.
	cartDescCreatedAt := cartFields[1].Descriptor()
	// cart.DefaultCreatedAt holds the default value on creation for the created_at field.
	cart.DefaultCreatedAt = cartDescCreatedAt.Default.(func() time.Time)
	// cartDescUpdatedAt is the schema descriptor for updated_at field.
	cartDescUpdatedAt := cartFields[2].Descriptor()
	// cart.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cart.DefaultUpdatedAt = cartDescUpdatedAt.Default.(func() time.Time)
	// cart.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	cart.UpdateDefaultUpdatedAt = cartDescUpdatedAt.UpdateDefault.(func() time.Time)
	cartproductFields := schema.CartProduct{}.Fields()
	_ = cartproductFields
	// cartproductDescCreatedAt is the schema descriptor for created_at field.
	cartproductDescCreatedAt := cartproductFields[1].Descriptor()
	// cartproduct.DefaultCreatedAt holds the default value on creation for the created_at field.
	cartproduct.DefaultCreatedAt = cartproductDescCreatedAt.Default.(func() time.Time)
	// cartproductDescUpdatedAt is the schema descriptor for updated_at field.
	cartproductDescUpdatedAt := cartproductFields[2].Descriptor()
	// cartproduct.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cartproduct.DefaultUpdatedAt = cartproductDescUpdatedAt.Default.(func() time.Time)
	// cartproduct.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	cartproduct.UpdateDefaultUpdatedAt = cartproductDescUpdatedAt.UpdateDefault.(func() time.Time)
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescCreatedAt is the schema descriptor for created_at field.
	categoryDescCreatedAt := categoryFields[2].Descriptor()
	// category.DefaultCreatedAt holds the default value on creation for the created_at field.
	category.DefaultCreatedAt = categoryDescCreatedAt.Default.(func() time.Time)
	// categoryDescUpdatedAt is the schema descriptor for updated_at field.
	categoryDescUpdatedAt := categoryFields[3].Descriptor()
	// category.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	category.DefaultUpdatedAt = categoryDescUpdatedAt.Default.(func() time.Time)
	// category.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	category.UpdateDefaultUpdatedAt = categoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	checkoutFields := schema.Checkout{}.Fields()
	_ = checkoutFields
	// checkoutDescCompleted is the schema descriptor for completed field.
	checkoutDescCompleted := checkoutFields[1].Descriptor()
	// checkout.DefaultCompleted holds the default value on creation for the completed field.
	checkout.DefaultCompleted = checkoutDescCompleted.Default.(bool)
	// checkoutDescCreatedAt is the schema descriptor for created_at field.
	checkoutDescCreatedAt := checkoutFields[2].Descriptor()
	// checkout.DefaultCreatedAt holds the default value on creation for the created_at field.
	checkout.DefaultCreatedAt = checkoutDescCreatedAt.Default.(func() time.Time)
	// checkoutDescUpdatedAt is the schema descriptor for updated_at field.
	checkoutDescUpdatedAt := checkoutFields[3].Descriptor()
	// checkout.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	checkout.DefaultUpdatedAt = checkoutDescUpdatedAt.Default.(func() time.Time)
	// checkout.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	checkout.UpdateDefaultUpdatedAt = checkoutDescUpdatedAt.UpdateDefault.(func() time.Time)
	checkoutproductFields := schema.CheckoutProduct{}.Fields()
	_ = checkoutproductFields
	// checkoutproductDescReceived is the schema descriptor for received field.
	checkoutproductDescReceived := checkoutproductFields[3].Descriptor()
	// checkoutproduct.DefaultReceived holds the default value on creation for the received field.
	checkoutproduct.DefaultReceived = checkoutproductDescReceived.Default.(bool)
	// checkoutproductDescStatus is the schema descriptor for status field.
	checkoutproductDescStatus := checkoutproductFields[4].Descriptor()
	// checkoutproduct.DefaultStatus holds the default value on creation for the status field.
	checkoutproduct.DefaultStatus = checkoutproductDescStatus.Default.(int)
	// checkoutproductDescCreatedAt is the schema descriptor for created_at field.
	checkoutproductDescCreatedAt := checkoutproductFields[5].Descriptor()
	// checkoutproduct.DefaultCreatedAt holds the default value on creation for the created_at field.
	checkoutproduct.DefaultCreatedAt = checkoutproductDescCreatedAt.Default.(func() time.Time)
	// checkoutproductDescUpdatedAt is the schema descriptor for updated_at field.
	checkoutproductDescUpdatedAt := checkoutproductFields[6].Descriptor()
	// checkoutproduct.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	checkoutproduct.DefaultUpdatedAt = checkoutproductDescUpdatedAt.Default.(func() time.Time)
	// checkoutproduct.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	checkoutproduct.UpdateDefaultUpdatedAt = checkoutproductDescUpdatedAt.UpdateDefault.(func() time.Time)
	sellerproductFields := schema.SellerProduct{}.Fields()
	_ = sellerproductFields
	// sellerproductDescActive is the schema descriptor for active field.
	sellerproductDescActive := sellerproductFields[5].Descriptor()
	// sellerproduct.DefaultActive holds the default value on creation for the active field.
	sellerproduct.DefaultActive = sellerproductDescActive.Default.(bool)
	// sellerproductDescCreatedAt is the schema descriptor for created_at field.
	sellerproductDescCreatedAt := sellerproductFields[11].Descriptor()
	// sellerproduct.DefaultCreatedAt holds the default value on creation for the created_at field.
	sellerproduct.DefaultCreatedAt = sellerproductDescCreatedAt.Default.(func() time.Time)
	// sellerproductDescUpdatedAt is the schema descriptor for updated_at field.
	sellerproductDescUpdatedAt := sellerproductFields[12].Descriptor()
	// sellerproduct.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sellerproduct.DefaultUpdatedAt = sellerproductDescUpdatedAt.Default.(func() time.Time)
	// sellerproduct.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sellerproduct.UpdateDefaultUpdatedAt = sellerproductDescUpdatedAt.UpdateDefault.(func() time.Time)
	sellerproductcategoryFields := schema.SellerProductCategory{}.Fields()
	_ = sellerproductcategoryFields
	// sellerproductcategoryDescCreatedAt is the schema descriptor for created_at field.
	sellerproductcategoryDescCreatedAt := sellerproductcategoryFields[0].Descriptor()
	// sellerproductcategory.DefaultCreatedAt holds the default value on creation for the created_at field.
	sellerproductcategory.DefaultCreatedAt = sellerproductcategoryDescCreatedAt.Default.(func() time.Time)
	// sellerproductcategoryDescUpdatedAt is the schema descriptor for updated_at field.
	sellerproductcategoryDescUpdatedAt := sellerproductcategoryFields[1].Descriptor()
	// sellerproductcategory.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sellerproductcategory.DefaultUpdatedAt = sellerproductcategoryDescUpdatedAt.Default.(func() time.Time)
	// sellerproductcategory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sellerproductcategory.UpdateDefaultUpdatedAt = sellerproductcategoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	sellerproductimageFields := schema.SellerProductImage{}.Fields()
	_ = sellerproductimageFields
	// sellerproductimageDescDisplay is the schema descriptor for display field.
	sellerproductimageDescDisplay := sellerproductimageFields[0].Descriptor()
	// sellerproductimage.DefaultDisplay holds the default value on creation for the display field.
	sellerproductimage.DefaultDisplay = sellerproductimageDescDisplay.Default.(bool)
	// sellerproductimageDescCreatedAt is the schema descriptor for created_at field.
	sellerproductimageDescCreatedAt := sellerproductimageFields[2].Descriptor()
	// sellerproductimage.DefaultCreatedAt holds the default value on creation for the created_at field.
	sellerproductimage.DefaultCreatedAt = sellerproductimageDescCreatedAt.Default.(func() time.Time)
	// sellerproductimageDescUpdatedAt is the schema descriptor for updated_at field.
	sellerproductimageDescUpdatedAt := sellerproductimageFields[3].Descriptor()
	// sellerproductimage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sellerproductimage.DefaultUpdatedAt = sellerproductimageDescUpdatedAt.Default.(func() time.Time)
	// sellerproductimage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sellerproductimage.UpdateDefaultUpdatedAt = sellerproductimageDescUpdatedAt.UpdateDefault.(func() time.Time)
	sellerproductvariationFields := schema.SellerProductVariation{}.Fields()
	_ = sellerproductvariationFields
	// sellerproductvariationDescCreatedAt is the schema descriptor for created_at field.
	sellerproductvariationDescCreatedAt := sellerproductvariationFields[4].Descriptor()
	// sellerproductvariation.DefaultCreatedAt holds the default value on creation for the created_at field.
	sellerproductvariation.DefaultCreatedAt = sellerproductvariationDescCreatedAt.Default.(func() time.Time)
	// sellerproductvariationDescUpdatedAt is the schema descriptor for updated_at field.
	sellerproductvariationDescUpdatedAt := sellerproductvariationFields[5].Descriptor()
	// sellerproductvariation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sellerproductvariation.DefaultUpdatedAt = sellerproductvariationDescUpdatedAt.Default.(func() time.Time)
	// sellerproductvariation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sellerproductvariation.UpdateDefaultUpdatedAt = sellerproductvariationDescUpdatedAt.UpdateDefault.(func() time.Time)
	sellerproductvariationvaluesFields := schema.SellerProductVariationValues{}.Fields()
	_ = sellerproductvariationvaluesFields
	// sellerproductvariationvaluesDescCreatedAt is the schema descriptor for created_at field.
	sellerproductvariationvaluesDescCreatedAt := sellerproductvariationvaluesFields[2].Descriptor()
	// sellerproductvariationvalues.DefaultCreatedAt holds the default value on creation for the created_at field.
	sellerproductvariationvalues.DefaultCreatedAt = sellerproductvariationvaluesDescCreatedAt.Default.(func() time.Time)
	// sellerproductvariationvaluesDescUpdatedAt is the schema descriptor for updated_at field.
	sellerproductvariationvaluesDescUpdatedAt := sellerproductvariationvaluesFields[3].Descriptor()
	// sellerproductvariationvalues.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sellerproductvariationvalues.DefaultUpdatedAt = sellerproductvariationvaluesDescUpdatedAt.Default.(func() time.Time)
	// sellerproductvariationvalues.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sellerproductvariationvalues.UpdateDefaultUpdatedAt = sellerproductvariationvaluesDescUpdatedAt.UpdateDefault.(func() time.Time)
	sellerrequestFields := schema.SellerRequest{}.Fields()
	_ = sellerrequestFields
	// sellerrequestDescContactNumber is the schema descriptor for contact_number field.
	sellerrequestDescContactNumber := sellerrequestFields[2].Descriptor()
	// sellerrequest.ContactNumberValidator is a validator for the "contact_number" field. It is called by the builders before save.
	sellerrequest.ContactNumberValidator = sellerrequestDescContactNumber.Validators[0].(func(string) error)
	// sellerrequestDescAccepted is the schema descriptor for accepted field.
	sellerrequestDescAccepted := sellerrequestFields[5].Descriptor()
	// sellerrequest.DefaultAccepted holds the default value on creation for the accepted field.
	sellerrequest.DefaultAccepted = sellerrequestDescAccepted.Default.(bool)
	// sellerrequestDescCreatedAt is the schema descriptor for created_at field.
	sellerrequestDescCreatedAt := sellerrequestFields[6].Descriptor()
	// sellerrequest.DefaultCreatedAt holds the default value on creation for the created_at field.
	sellerrequest.DefaultCreatedAt = sellerrequestDescCreatedAt.Default.(func() time.Time)
	// sellerrequestDescUpdatedAt is the schema descriptor for updated_at field.
	sellerrequestDescUpdatedAt := sellerrequestFields[7].Descriptor()
	// sellerrequest.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sellerrequest.DefaultUpdatedAt = sellerrequestDescUpdatedAt.Default.(func() time.Time)
	// sellerrequest.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sellerrequest.UpdateDefaultUpdatedAt = sellerrequestDescUpdatedAt.UpdateDefault.(func() time.Time)
	sellershopFields := schema.SellerShop{}.Fields()
	_ = sellershopFields
	// sellershopDescContactNumber is the schema descriptor for contact_number field.
	sellershopDescContactNumber := sellershopFields[2].Descriptor()
	// sellershop.ContactNumberValidator is a validator for the "contact_number" field. It is called by the builders before save.
	sellershop.ContactNumberValidator = sellershopDescContactNumber.Validators[0].(func(string) error)
	// sellershopDescCreatedAt is the schema descriptor for created_at field.
	sellershopDescCreatedAt := sellershopFields[9].Descriptor()
	// sellershop.DefaultCreatedAt holds the default value on creation for the created_at field.
	sellershop.DefaultCreatedAt = sellershopDescCreatedAt.Default.(func() time.Time)
	// sellershopDescUpdatedAt is the schema descriptor for updated_at field.
	sellershopDescUpdatedAt := sellershopFields[10].Descriptor()
	// sellershop.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sellershop.DefaultUpdatedAt = sellershopDescUpdatedAt.Default.(func() time.Time)
	// sellershop.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sellershop.UpdateDefaultUpdatedAt = sellershopDescUpdatedAt.UpdateDefault.(func() time.Time)
	sellershopproductFields := schema.SellerShopProduct{}.Fields()
	_ = sellershopproductFields
	// sellershopproductDescCreatedAt is the schema descriptor for created_at field.
	sellershopproductDescCreatedAt := sellershopproductFields[0].Descriptor()
	// sellershopproduct.DefaultCreatedAt holds the default value on creation for the created_at field.
	sellershopproduct.DefaultCreatedAt = sellershopproductDescCreatedAt.Default.(func() time.Time)
	// sellershopproductDescUpdatedAt is the schema descriptor for updated_at field.
	sellershopproductDescUpdatedAt := sellershopproductFields[1].Descriptor()
	// sellershopproduct.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sellershopproduct.DefaultUpdatedAt = sellershopproductDescUpdatedAt.Default.(func() time.Time)
	// sellershopproduct.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sellershopproduct.UpdateDefaultUpdatedAt = sellershopproductDescUpdatedAt.UpdateDefault.(func() time.Time)
	shopcategoryFields := schema.ShopCategory{}.Fields()
	_ = shopcategoryFields
	// shopcategoryDescCreatedAt is the schema descriptor for created_at field.
	shopcategoryDescCreatedAt := shopcategoryFields[3].Descriptor()
	// shopcategory.DefaultCreatedAt holds the default value on creation for the created_at field.
	shopcategory.DefaultCreatedAt = shopcategoryDescCreatedAt.Default.(func() time.Time)
	// shopcategoryDescUpdatedAt is the schema descriptor for updated_at field.
	shopcategoryDescUpdatedAt := shopcategoryFields[4].Descriptor()
	// shopcategory.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	shopcategory.DefaultUpdatedAt = shopcategoryDescUpdatedAt.Default.(func() time.Time)
	// shopcategory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	shopcategory.UpdateDefaultUpdatedAt = shopcategoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescPhoneNumber is the schema descriptor for phone_number field.
	userDescPhoneNumber := userFields[1].Descriptor()
	// user.PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	user.PhoneNumberValidator = userDescPhoneNumber.Validators[0].(func(string) error)
	// userDescAdmin is the schema descriptor for admin field.
	userDescAdmin := userFields[3].Descriptor()
	// user.DefaultAdmin holds the default value on creation for the admin field.
	user.DefaultAdmin = userDescAdmin.Default.(bool)
	// userDescStaff is the schema descriptor for staff field.
	userDescStaff := userFields[4].Descriptor()
	// user.DefaultStaff holds the default value on creation for the staff field.
	user.DefaultStaff = userDescStaff.Default.(bool)
	// userDescSeller is the schema descriptor for seller field.
	userDescSeller := userFields[5].Descriptor()
	// user.DefaultSeller holds the default value on creation for the seller field.
	user.DefaultSeller = userDescSeller.Default.(bool)
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[6].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[9].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[10].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	userlocationFields := schema.UserLocation{}.Fields()
	_ = userlocationFields
	// userlocationDescCreatedAt is the schema descriptor for created_at field.
	userlocationDescCreatedAt := userlocationFields[7].Descriptor()
	// userlocation.DefaultCreatedAt holds the default value on creation for the created_at field.
	userlocation.DefaultCreatedAt = userlocationDescCreatedAt.Default.(func() time.Time)
	// userlocationDescUpdatedAt is the schema descriptor for updated_at field.
	userlocationDescUpdatedAt := userlocationFields[8].Descriptor()
	// userlocation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	userlocation.DefaultUpdatedAt = userlocationDescUpdatedAt.Default.(func() time.Time)
	// userlocation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	userlocation.UpdateDefaultUpdatedAt = userlocationDescUpdatedAt.UpdateDefault.(func() time.Time)
}
