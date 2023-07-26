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
	UserController := controllers.UserController{}
	KelasController := controllers.KelasController{}
	SiswaController := controllers.SiswaController{}
	NilaiController := controllers.NewNilaiController()
	TransformasiController := controllers.NewTransformasiController()
	AttributeController := controllers.NewAttributeController()
	DashboardController := controllers.NewsDashboardController()

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
			adminDashboard := admin.Group("dashboard").Use(middlewares.JwtAuthRolesMiddleware(models.RoleAdmin, models.RoleOperator))
			{
				adminDashboard.Get("donut", DashboardController.Donut)
				adminDashboard.Get("column", DashboardController.Column)
			}
			adminKelas := admin.Group("kelas").Use(middlewares.JwtAuthRolesMiddleware(models.RoleAdmin, models.RoleOperator))
			{
				adminKelas.Get("", KelasController.Index)
			}
			adminAttribute := admin.Group("attribute").Use(middlewares.JwtAuthRolesMiddleware(models.RoleAdmin, models.RoleOperator))
			{
				adminAttribute.Get("", AttributeController.GetPaginate)
				adminAttribute.Get("/:uuid", AttributeController.Show)
				adminAttribute.Put("/:uuid", AttributeController.Update)
			}
			adminSiswa := admin.Group("siswa").Use(middlewares.JwtAuthRolesMiddleware(models.RoleAdmin, models.RoleOperator))
			{
				adminSiswa.Get("", SiswaController.GetPaginate)
				adminSiswa.Post("", SiswaController.Store)
				adminSiswa.Get("/:uuid", SiswaController.Show)
				adminSiswa.Put("/:uuid", SiswaController.Update)
				adminSiswa.Delete("/:uuid", SiswaController.Delete)
			}
			adminNilai := admin.Group("nilai").Use(middlewares.JwtAuthRolesMiddleware(models.RoleAdmin, models.RoleOperator))
			{
				adminNilai.Get("", NilaiController.GetPaginate)
				adminNilai.Post("", NilaiController.Store)
				adminNilai.Get("/:uuid", NilaiController.Show)
			}
			adminTransformasi := admin.Group("transformasi").Use(middlewares.JwtAuthRolesMiddleware(models.RoleAdmin, models.RoleOperator))
			{
				adminTransformasi.Get("", TransformasiController.GetPaginate)
				adminTransformasi.Post("/bayes", TransformasiController.Bayes)
			}
			adminUser := admin.Group("user").Use(middlewares.JwtAuthRolesMiddleware(models.RoleAdmin))
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
