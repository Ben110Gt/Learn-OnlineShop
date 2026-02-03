package server

import (
	routes "OnlineShop/Api/internal/Routes"
	"OnlineShop/Api/internal/configs"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer() {
	// Connect to the database
	configs.ConnectDatabase()
	// Initialize Fiber app
	app := fiber.New()
	//cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://foo.com, https://github.com",
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           int((12 * time.Hour).Seconds()),
	}))

	// Setup routes
	routes.PublicRoutes(app)
	routes.UserRoutes(app)
	routes.ProductrRoutes(app)
	routes.CategoryRoutes(app)
	routes.CartRoutes(app)
	routes.OrderRoutes(app)

	// Start server
	app.Listen("localhost:8080")
}
