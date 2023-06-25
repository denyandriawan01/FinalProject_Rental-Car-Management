package controller

import (
	"encoding/json"
	"net/http"

	"FinalProject_Rental-Car-Management/database"
	"FinalProject_Rental-Car-Management/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TaxIndex(c *gin.Context) {
	var tax []models.Taxes
	var count int64

	if result := database.DB.Preload("Car").Find(&tax); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	if result := database.DB.Model(&models.Taxes{}).Count(&count); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Taxes": tax,
		"Count": count,
	})
}

func TaxShow(c *gin.Context) {
	id := c.Param("id")
	var tax models.Taxes

	if err := database.DB.Preload("Car").First(&tax, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pajak tidak ditemukan"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Terjadi kesalahan saat mengambil data pajak"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tax": tax})
}

func GetTaxesByCarID(c *gin.Context) {
	id := c.Param("id")
	var taxes []models.Taxes

	if err := database.DB.Where("car_id = ?", id).Preload("Car").Find(&taxes).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve taxes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"taxes": taxes,
	})
}

func TaxCreate(c *gin.Context) {
	var tax models.Taxes

	if err := c.ShouldBindJSON(&tax); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&tax).Error; err != nil {
			return err
		}

		if err := tx.Preload("Car").First(&tax, tax.TaxID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data Maintenance History"})
		return
	}

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
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.ID.Int64()

	if err := database.DB.First(&tax, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data pajak tidak ditemukan"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data pajak"})
		return
	}

	if err := database.DB.Delete(&tax).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data pajak"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
