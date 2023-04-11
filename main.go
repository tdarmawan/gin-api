package main

import (
	"swaggo-gin-api-basic/database"
	"swaggo-gin-api-basic/routes"
)

func main() {
	database.StartDb()

	var PORT = ":7070"

	routes.StartServer().Run(PORT)
}