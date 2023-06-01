package controller

import (
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRentals(c *gin.Context) {
	var rentals []models.Rental
	models.DB.Find(&rentals)
	c.JSON(http.StatusOK, gin.H{"data": rentals})
}

func GetRental(c *gin.Context) {
	var rental models.Rental
	if err := models.DB.First(&rental, c.Param("rental_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rental})
}

func CreateRental(c *gin.Context) {
	var input models.Rental
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func UpdateRental(c *gin.Context) {
	var rental models.Rental
	if err := models.DB.First(&rental, c.Param("rental_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input models.Rental
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&rental).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": rental})
}

func DeleteRental(c *gin.Context) {
	var rental models.Rental
	if err := models.DB.First(&rental, c.Param("rental_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&rental)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
