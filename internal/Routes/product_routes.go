package routes

import (
	Handler "OnlineShop/Api/internal/Handler"
	"OnlineShop/Api/internal/configs"
	"OnlineShop/Api/internal/middlewares"
	"OnlineShop/Api/internal/repositories"
	Service "OnlineShop/Api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func ProductrRoutes(app *fiber.App) {
	configs.ConnectDatabase()
	db := configs.GetDB()
	ProductRepo := repositories.NewProductRepository(db)
	productService := Service.NewProductService(ProductRepo)
	ProductHadler := Handler.NewProductHandler(productService)

	auth := app.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	admin := auth.Group("/admin")
	admin.Use(middlewares.AdminOnly())
	{
		admin.Get("/products", ProductHadler.GetAllProduct)               //complete✅
		admin.Get("/product/:product_id", ProductHadler.GetProductrByID)  //complete✅
		admin.Post("/products", ProductHadler.CreateProduct)              //complete✅
		admin.Delete("/product/:product_id", ProductHadler.DeleteProduct) //complete✅
		admin.Put("/product/:product_id", ProductHadler.UpdateProduct)    //complete✅
	}
}
