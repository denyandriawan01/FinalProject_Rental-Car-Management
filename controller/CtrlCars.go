package controller

import (
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCars(c *gin.Context) {
	var cars []models.Car
	models.DB.Find(&cars)
	c.JSON(http.StatusOK, gin.H{"data": cars})
}

func GetCar(c *gin.Context) {
	var car models.Car
	if err := models.DB.First(&car, c.Param("car_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": car})
}

func CreateCar(c *gin.Context) {
	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func UpdateCar(c *gin.Context) {
	var car models.Car
	if err := models.DB.First(&car, c.Param("car_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&car).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": car})
}

func DeleteCar(c *gin.Context) {
	var car models.Car
	if err := models.DB.First(&car, c.Param("car_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&car)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
