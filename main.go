package main

import (
	"gostarter-backend/models"
	"gostarter-backend/routes"
)

func main() {
	models.ConnectDataBase()
	r := routes.SetupRouter()
	r.Run(":8080")

}