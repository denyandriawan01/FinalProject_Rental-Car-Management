package controller

import (
	"encoding/json"
	"net/http"

	"finpro_golang/models"
	"finpro_golang/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TaxIndex(c *gin.Context) {
	var tax []models.Taxes

	database.DB.Find(&tax)
	c.JSON(http.StatusOK, gin.H{"tax": tax})
}

func TaxShow(c *gin.Context) {
	id := c.Param("id")
	var tax models.Taxes

	if err := database.DB.First(&tax, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pajak tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pajak tidak ditemukan"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"tax": tax})
}

func TaxCreate(c *gin.Context) {
	var tax models.Taxes

	if err := c.ShouldBindJSON(&tax); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&tax)
	c.JSON(http.StatusOK, gin.H{"tax": tax})
}

func TaxUpdate(c *gin.Context) {
	var tax models.Taxes
	id := c.Param("id")

	if err := c.ShouldBindJSON(&tax); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&tax).Where("tax_id = ?", id).Updates(&tax).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data pajak"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data pajak berhasil diperbarui"})
}

func TaxDelete(c *gin.Context) {
	var tax models.Taxes

	var input struct {
		ID json.Number `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.ID.Int64()

	if err := database.DB.First(&tax, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pajak tidak ditemukan"})
		return
	}

	if database.DB.Delete(&tax).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data pajak"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}