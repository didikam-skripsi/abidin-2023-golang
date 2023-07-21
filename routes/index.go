package routes

import (
	"gostarter-backend/controllers"
	"gostarter-backend/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	HomeController := controllers.HomeController{}
	AuthController := controllers.AuthController{}
	ProductController := controllers.ProductController{}

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/home", HomeController.Index)

	public := router.Group("/api")
	{
		public.POST("/register", AuthController.Register)
		public.POST("/login", AuthController.Login)

		auth := router.Group("/api/admin")
		{
			auth.Use(middlewares.JwtAuthMiddleware())
			auth.GET("/user", AuthController.CurrentUser)

			auth.GET("/product", ProductController.GetPostPaginate)
			auth.POST("/product", ProductController.Store)
			auth.GET("/product/:id", ProductController.Show)
			auth.PUT("/product/:id", ProductController.Update)
			auth.DELETE("/product/:id", ProductController.Delete)
		}
	}

	return router
}
