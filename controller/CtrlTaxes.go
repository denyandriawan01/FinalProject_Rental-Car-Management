package controller

import (
	"models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTaxes(c *gin.Context) {
	var taxes []models.Taxes
	models.DB.Find(&taxes)
	c.JSON(http.StatusOK, gin.H{"data": taxes})
}

func GetTax(c *gin.Context) {
	var tax models.Taxes
	if err := models.DB.First(&tax, c.Param("tax_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tax})
}

func CreateTax(c *gin.Context) {
	var input models.Taxes
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func UpdateTax(c *gin.Context) {
	var tax models.Taxes
	if err := models.DB.First(&tax, c.Param("tax_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input models.Taxes
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&tax).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": tax})
}

func DeleteTax(c *gin.Context) {
	var tax models.Taxes
	if err := models.DB.First(&tax, c.Param("tax_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&tax)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
