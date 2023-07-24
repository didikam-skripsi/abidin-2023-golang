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
				auth.Put("/profile", AuthController.Profile)
			}
			admin := jwt.Group("admin")
			adminProduct := admin.Group("product").Use(middlewares.JwtAuthRolesMiddleware(models.RoleAdmin, models.RoleUser))
			{
				adminProduct.Get("", ProductController.GetPostPaginate)
				adminProduct.Post("", ProductController.Store)
				adminProduct.Get("/:uuid", ProductController.Show)
				adminProduct.Put("/:uuid", ProductController.Update)
				adminProduct.Delete("/:uuid", ProductController.Delete)
			}
			adminUser := admin.Group("user").Use(middlewares.JwtAuthRolesMiddleware(models.RoleOperator, models.RoleAdmin))
			{
				adminUser.Get("", UserController.GetPostPaginate)
				adminUser.Post("", UserController.Store)
				adminUser.Get("/:uuid", UserController.Show)
				adminUser.Put("/:uuid", UserController.Update)
				adminUser.Delete("/:uuid", UserController.Delete)
			}
		}
	}

	return app
}
