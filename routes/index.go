package routes

import (
	"gostarter-backend/controllers"
	"gostarter-backend/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRouter() *fiber.App {
	HomeController := controllers.HomeController{}
	AuthController := controllers.AuthController{}
	ProductController := controllers.ProductController{}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/home", HomeController.Index)

	public := app.Group("/api")
	{
		public.Post("/register", AuthController.Register)
		public.Post("/login", AuthController.Login)

		auth := app.Group("/api/admin")
		auth.Use(middlewares.JwtAuthMiddleware())
		auth.Get("/user", AuthController.CurrentUser)
		auth.Get("/product", ProductController.GetPostPaginate)
		auth.Post("/product", ProductController.Store)
		auth.Get("/product/:id", ProductController.Show)
		auth.Put("/product/:id", ProductController.Update)
		auth.Delete("/product/:id", ProductController.Delete)
	}

	return app
}
