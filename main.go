package main

import (
	"gin-api/database"
	"gin-api/routers"
)

func main() {
	database.StartDB()
	
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
