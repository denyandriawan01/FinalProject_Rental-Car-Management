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

	if err := database.DB.Preload("User").Preload("Car").Find(&rental).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data rental"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rental": rental})
}

func RentalShow(c *gin.Context) {
	id := c.Param("id")
	var rental models.Rental

	if err := database.DB.Preload("User").Preload("Car").First(&rental, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data rental tidak ditemukan"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data rental"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rental": rental})
}

func RentalCreate(c *gin.Context) {
	var rental models.Rental

	if err := c.ShouldBindJSON(&rental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&rental).Error; err != nil {
			return err
		}

		if err := tx.Preload("User").Preload("Car").First(&rental, rental.ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data rental"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rental": rental})
}

func RentalUpdate(c *gin.Context) {
	id := c.Param("id")

	var rental models.Rental
	if err := database.DB.First(&rental, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data rental tidak ditemukan"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data rental"})
		return
	}

	if err := c.ShouldBindJSON(&rental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := database.DB.Save(&rental).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui data rental"})
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
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.ID.Int64()

	if err := database.DB.First(&rental, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data rental tidak ditemukan"})
			return
			}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data rental"})
		return
	}

	if err := database.DB.Delete(&rental).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data rental"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}