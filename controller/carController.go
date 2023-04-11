package controller

import (
	"fmt"
	"net/http"
	"swaggo-gin-api-basic/database"
	"swaggo-gin-api-basic/models"

	"github.com/gin-gonic/gin"
)

// GetAllCars godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars [get]
func GetAllCars(c *gin.Context){
	var db = database.GetDB()

	var cars []models.Car
	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("error get car")
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}