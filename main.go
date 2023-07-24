package main

import (
	"fmt"
	"gostarter-backend/models"
	"gostarter-backend/routes"
	"os"
)

func main() {
	models.ConnectDataBase()
	// models.SeedProducts()
	// models.SeedUsers()
	app := routes.SetupRouter()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	err := app.Listen(":" + port)
	if err != nil {
		fmt.Printf("Error listening on port %s: %v", port, err)
	}
}
