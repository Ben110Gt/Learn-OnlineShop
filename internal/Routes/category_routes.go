package routes

import (
	Handler "OnlineShop/Api/internal/Handler"
	"OnlineShop/Api/internal/configs"
	"OnlineShop/Api/internal/middlewares"
	"OnlineShop/Api/internal/repositories"
	Service "OnlineShop/Api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	configs.ConnectDatabase()
	db := configs.GetDB()
	CategoryRepo := repositories.NewCategoryRepository(db)
	CategoryService := Service.NewCategoryService(CategoryRepo)
	CategoryHandler := Handler.NewCategoryHandler(*CategoryService)

	auth := app.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	admin := auth.Group("/admin")
	admin.Use(middlewares.AdminOnly())
	{
		admin.Post("/categories", CategoryHandler.CreateCategory)                    //complete✅
		admin.Get("/categories", CategoryHandler.GetAllCategory)                     //complete✅
		admin.Delete("/categories/:category_id", CategoryHandler.DeleteCategoryByid) //complete✅
	}
}
