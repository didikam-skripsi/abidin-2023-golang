package main

import (
	"gostarter-backend/models"
	"gostarter-backend/routes"
)

func main() {
	models.ConnectDataBase()
	// models.SeedProducts()
	r := routes.SetupRouter()
	r.Listen(":8080")

}
