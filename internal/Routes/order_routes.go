package routes

import (
	Handler "OnlineShop/Api/internal/Handler"
	"OnlineShop/Api/internal/configs"
	"OnlineShop/Api/internal/middlewares"
	"OnlineShop/Api/internal/repositories"
	Service "OnlineShop/Api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(app *fiber.App) {
	configs.ConnectDatabase()
	db := configs.GetDB()

	OrderRepo := repositories.NewOrderRepository(db)
	CartRepo := repositories.NewCartRepository(db)
	OrderService := Service.NewOrderService(OrderRepo, CartRepo)
	OrderHandler := Handler.NewOrderHandler(*OrderService)

	auth := app.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	User := auth.Group("/user")
	User.Use(middlewares.UserOnly())
	{
		User.Post("/order", OrderHandler.CreateOrder) //complete✅
	}
	admin := auth.Group("/admin")
	admin.Use(middlewares.AdminOnly())
	{
		admin.Post("/order/:order_id", OrderHandler.ConfirmOrder) //complete✅
		admin.Get("/orders/", OrderHandler.GetAllOrder)           //complete✅
	}
}
