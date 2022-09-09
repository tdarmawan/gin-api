package routers

import (
	"gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine{
	router := gin.Default()

	router.GET("/cars/:id", controllers.GetOneCars)
	router.GET("/cars", controllers.GetAllCars)
	router.POST("/cars", controllers.CreateCars)
	router.PATCH("/cars/:id", controllers.UpdateCars)
	router.DELETE("/cars/:id", controllers.DeleteCars)

	return router
}