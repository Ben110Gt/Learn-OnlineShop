package routes

import (
	Handler "OnlineShop/Api/internal/Handler"
	"OnlineShop/Api/internal/configs"
	"OnlineShop/Api/internal/middlewares"
	"OnlineShop/Api/internal/repositories"
	Service "OnlineShop/Api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	configs.ConnectDatabase()
	db := configs.GetDB()

	userRepo := repositories.NewUserRepository(db)
	userService := Service.NewUserService(userRepo)
	userHandler := Handler.NewUserHandler(userService)

	auth := app.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	admin := auth.Group("/admin")
	admin.Use(middlewares.AdminOnly())
	{

		admin.Get("/users", userHandler.GetAllUser)            //complete✅
		admin.Get("/user/:user_id", userHandler.GetUser)       //complete✅
		admin.Post("/register/", userHandler.RegisterAdmin)    //complete✅
		admin.Delete("/user/:user_id", userHandler.DeleteUser) //complete✅
		admin.Put("/user/:user_id", userHandler.UpdateUser)    //complete✅

	}

}
