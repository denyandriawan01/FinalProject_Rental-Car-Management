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

	if err := database.DB.Preload("Car").Find(&tax).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data pajak"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tax": tax})
}

func TaxShow(c *gin.Context) {
	id := c.Param("id")
	var tax models.Taxes

	if err := database.DB.Preload("Car").First(&tax, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data pajak tidak ditemukan"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data pajak"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tax": tax})
}

func TaxCreate(c *gin.Context) {
	var tax models.Taxes

	if err := c.ShouldBindJSON(&tax); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&tax).Error; err != nil {
			return err
		}

		if err := tx.Preload("Car").First(&tax, tax.ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data pajak"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tax": tax})
}

func TaxUpdate(c *gin.Context) {
	var tax models.Taxes
	id := c.Param("id")

	if err := c.ShouldBindJSON(&tax); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := database.DB.Model(&tax).Where("tax_id = ?", id).Updates(&tax).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data pajak"})
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