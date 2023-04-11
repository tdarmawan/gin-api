package main

import (
	"os"
	"swaggo-gin-api-basic/database"
	"swaggo-gin-api-basic/routes"
)

func main() {
	database.StartDb()

	var PORT = os.Getenv("PORT")

	routes.StartServer().Run(":" + PORT)
}