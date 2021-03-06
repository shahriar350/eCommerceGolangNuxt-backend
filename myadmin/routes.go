package myadmin

import (
	"bongo/db"
	"bongo/mixin"
	"context"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func adminMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("bongoauth")
	err, data, _ := mixin.VerifyToken(cookie)
	if !err {
		c.Status(fiber.StatusForbidden)
	}
	id, _ := strconv.Atoi(data.Issuer)
	user, err2 := db.Client.User.Get(context.Background(),id)
	if err2 != nil {
		return c.Status(401).SendString("You are not allowed.")
	}
	//var user model.User
	//result := model.DB.First(&user, "id = ?", data.Issuer)
	//if result.Error != nil {
	//	return c.SendStatus(401)
	//}
	if !user.Admin {
		return c.Status(401).SendString("You are not allowed.")
	}
	c.Locals("AuthID", id)
	return c.Next()
}
func AdminRoutes(app *fiber.App) {
	// admin
	admin := app.Group("/api/admin", adminMiddleware)
	// shop category
	admin.Post("/logout", AdminLogout)
	admin.Get("/shop/categories/only/all", GetShopCategoryOnly)
	admin.Get("/shop/categories/all", GetShopCategory)
	admin.Get("/shop/categories/active", GetActiveShopCategory)
	admin.Get("/shop/categories/:id", GetTargetShopCategory)
	admin.Post("/shop/category/create", CreateShopCategory)
	admin.Delete("/shop/category/delete/soft/:id", ShopCategorySoftDelete)
	admin.Put("/shop/category/delete/recover/:id", ShopCategorySoftRecoverDelete)
	admin.Delete("/shop/category/delete/permanent/:id", ShopCategoryPermanentDelete)
	admin.Put("/shop/category/update/:id", UpdateShopCategory)
	// category
	admin.Get("/categories", GetCategory)
	admin.Get("/categories/deleted", GetDeletedCategory)
	admin.Post("/category/create", CreateCategory)
	admin.Delete("/category/delete/soft/:id", CategorySoftDelete)
	admin.Put("/category/delete/recover/:id", CategoryRecoverDelete)
	admin.Delete("/category/delete/permanent/:id", CategoryPermanentDelete)
	admin.Put("/category/update/:id", UpdateCategory)
	// Product
	//attributes
	admin.Get("/attributes", GetAttributes)
	admin.Post("/attributes/create", CreateAttributes)
	admin.Post("/attributes/single/:id", GetSingleAttributes)
	admin.Put("/attributes/edit/:id", EditSingleAttributes)
	admin.Delete("/attributes/delete/permanent/:id", PermanentDeleteSingleAttributes)
	// seller request
	admin.Get("/seller/request/all", AllSellerRequest)
	admin.Get("/seller/request", GetSellerRequest)
	admin.Post("/seller/request/accepted/:id", AcceptSellerRequest)
	admin.Delete("/seller/request/remove/:id", RemoveSellerRequest)
	admin.Delete("/seller/request/remove/permanent/:id", RemovePermanentSellerRequest)
	admin.Put("/seller/request/remove/recover/:id", RecoverSellerRequest)
	// seller shop start
	admin.Get("/seller/shops/non/activate", SellerShopsNonActivate)
	admin.Get("/seller/shops/activate", SellerShopsActivate)
	admin.Get("/seller/shops/delete", SellerShopsDeleted)
	admin.Get("/seller/shops/all", SellerShopsAll)
	admin.Post("/seller/shops/active/:id", ActiveSellerShops)
	admin.Put("/seller/shops/soft/delete/:id", SoftDeleteSellerShops)
	admin.Delete("/seller/shops/permanent/delete/:id", PermanentDeleteSellerShops)
	admin.Put("/seller/shops/recover/delete/:id", RecoverDeleteSellerShops)
	//	product brand
	admin.Get("/product/brand/all", BrandAll)
	admin.Post("/product/brand/create", BrandCreate)
	admin.Put("/product/brand/update/:id", BrandEdit)
	admin.Delete("/product/brand/soft/delete/:id", BrandSoftDelete)
	admin.Delete("/product/brand/delete/:id", BrandDelete)
	admin.Put("/product/brand/recover/:id", BrandRecoverDelete)

}
