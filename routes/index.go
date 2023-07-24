package routes

import (
	"gostarter-backend/controllers"
	"gostarter-backend/middlewares"
	"gostarter-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRouter() *fiber.App {
	// HomeController := controllers.HomeController{}
	AuthController := controllers.AuthController{}
	ProductController := controllers.ProductController{}
	UserController := controllers.UserController{}

	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "./public")
	// app.Get("/", HomeController.Index)

	api := app.Group("/api")
	{
		api.Post("/register", AuthController.Register)
		api.Post("/login", AuthController.Login)
		jwt := api.Use(middlewares.JwtAuthMiddleware())
		{
			auth := jwt.Group("/auth")
			{
				auth.Get("/user", AuthController.CurrentUser)
			}
			admin := jwt.Group("admin")
			adminProduct := admin.Use(middlewares.JwtAuthRolesMiddleware(models.RoleAdmin, models.RoleUser))
			{
				adminProduct.Get("/product", ProductController.GetPostPaginate)
				adminProduct.Post("/product", ProductController.Store)
				adminProduct.Get("/product/:uuid", ProductController.Show)
				adminProduct.Put("/product/:uuid", ProductController.Update)
				adminProduct.Delete("/product/:uuid", ProductController.Delete)
			}
			adminUser := admin.Use(middlewares.JwtAuthRolesMiddleware(models.RoleAdmin))
			{
				adminUser.Get("/user", UserController.GetPostPaginate)
				adminUser.Post("/user", UserController.Store)
				adminUser.Get("/user/:uuid", UserController.Show)
				adminUser.Put("/user/:uuid", UserController.Update)
				adminUser.Delete("/user/:uuid", UserController.Delete)
			}
		}
	}

	return app
}
