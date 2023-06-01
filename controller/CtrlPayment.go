package controller

import (
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPayments(c *gin.Context) {
	var payments []models.Payment
	models.DB.Find(&payments)
	c.JSON(http.StatusOK, gin.H{"data": payments})
}

func GetPayment(c *gin.Context) {
	var payment models.Payment
	if err := models.DB.First(&payment, c.Param("payment_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

func CreatePayment(c *gin.Context) {
	var input models.Payment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func UpdatePayment(c *gin.Context) {
	var payment models.Payment
	if err := models.DB.First(&payment, c.Param("payment_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input models.Payment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&payment).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

func DeletePayment(c *gin.Context) {
	var payment models.Payment
	if err := models.DB.First(&payment, c.Param("payment_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&payment)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
