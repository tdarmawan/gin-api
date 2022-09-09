package controllers

import (
	"fmt"
	"gin-api/database"
	"gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCars(c * gin.Context){
	var db = database.GetDB()

	var cars []models.Car
	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting car datas :", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

func GetOneCars(c * gin.Context){
	var db = database.GetDB()

	var carOne models.Car
	// err := db.Table("Car").Where("Id = ?", c.Param("id")).First(&car).Error
	err := db.First(&carOne, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": carOne})
}

func CreateCars(c *gin.Context) {
	var db = database.GetDB()
	// Validate input
	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	carinput := models.Car{Merk: input.Merk, Harga: input.Harga, Typecars: input.Typecars}
	db.Create(&carinput)

	c.JSON(http.StatusOK, gin.H{"data": carinput})
}

func UpdateCars(c *gin.Context){
	var db = database.GetDB()
	
	var car models.Car
	err := db.First(&car, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&car).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": car})
}

func DeleteCars(c *gin.Context){
	var db = database.GetDB()
	// Get model if exist
	var carDelete models.Car
	err := db.First(&carDelete, "Id = ?", c.Param("id")).Error; 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&carDelete)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
