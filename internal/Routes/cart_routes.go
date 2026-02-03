package routes

import (
	handler "OnlineShop/Api/internal/Handler"
	"OnlineShop/Api/internal/configs"
	"OnlineShop/Api/internal/middlewares"
	"OnlineShop/Api/internal/repositories"
	"OnlineShop/Api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func CartRoutes(app *fiber.App) {
	configs.ConnectDatabase()
	db := configs.GetDB()
	CartRepo := repositories.NewCartRepository(db)
	CartService := service.NewCartService(CartRepo)
	CartHandler := handler.NewCartHandler(*CartService)

	auth := app.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	admin := auth.Group("/user")
	admin.Use(middlewares.UserOnly())
	{
		admin.Post("/cart", CartHandler.AddCart)             //complete✅
		admin.Get("/cart/:user_id", CartHandler.GetCartbyID) //complete✅
	}
}
