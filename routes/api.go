package routes

import (
	"github.com/gin-gonic/gin"

	"swaggo-gin-api-basic/controller"

	_ "swaggo-gin-api-basic/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

// @title Car API
// @version 1.0
// @description This is Car Service REST API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html 
// @host localhost:7070
// @BasePath /
func StartServer() *gin.Engine{
	router := gin.Default()

	//READ

	router.GET("/cars", controller.GetAllCars)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}