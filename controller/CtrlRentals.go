package controller

import (
	"encoding/json"
	"net/http"

	"finpro_golang/models"
	"finpro_golang/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RentalIndex(c *gin.Context) {
	var rental []models.Rental

	database.DB.Find(&rental)
	c.JSON(http.StatusOK, gin.H{"rental": rental})
}

func RentalShow(c *gin.Context) {
	id := c.Param("id")
	var rental models.Rental

	if err := database.DB.First(&rental, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data rental tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data rental tidak ditemukan"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"rental": rental})
}

func RentalCreate(c *gin.Context) {
	var rental models.Rental

	if err := c.ShouldBindJSON(&rental); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&rental)
	c.JSON(http.StatusOK, gin.H{"rental": rental})
}

func RentalUpdate(c *gin.Context) {
	var rental models.Rental
	id := c.Param("id")

	if err := c.ShouldBindJSON(&rental); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&rental).Where("rental_id = ?", id).Updates(&rental).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data rental"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data rental berhasil diperbarui"})
}

func RentalDelete(c *gin.Context) {
	var rental models.Rental

	var input struct {
		ID json.Number `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.ID.Int64()

	if err := database.DB.First(&rental, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data rental tidak ditemukan"})
		return
	}

	if database.DB.Delete(&rental).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data rental"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}