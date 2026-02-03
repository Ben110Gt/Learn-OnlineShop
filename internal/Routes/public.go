package routes

import (
	Handler "OnlineShop/Api/internal/Handler"
	"OnlineShop/Api/internal/configs"
	"OnlineShop/Api/internal/repositories"
	Service "OnlineShop/Api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	configs.ConnectDatabase()
	db := configs.GetDB()
	//User
	userRepo := repositories.NewUserRepository(db)
	userService := Service.NewUserService(userRepo)
	userHandler := Handler.NewUserHandler(userService)
	//Product
	ProductRepo := repositories.NewProductRepository(db)
	productService := Service.NewProductService(ProductRepo)
	ProductHadler := Handler.NewProductHandler(productService)
	//Category
	CategoryRepo := repositories.NewCategoryRepository(db)
	CategoryService := Service.NewCategoryService(CategoryRepo)
	CategoryHandler := Handler.NewCategoryHandler(*CategoryService)

	public := app.Group("/")
	public.Post("/register", userHandler.RegisterUser)         //complete✅
	public.Post("/login", userHandler.Login)                   //complete✅
	public.Get("/products", ProductHadler.GetAllProductPublic) //complete✅
	public.Get("/categories", CategoryHandler.GetAllCategory)  //complete✅

}
