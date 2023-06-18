package controller

import (
	"encoding/json"
	"net/http"

	"FinalProject_Rental-Car-Management/database"
	"FinalProject_Rental-Car-Management/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RentalIndex(c *gin.Context) {
	var rental []models.Rental
	var pagination struct {
		Page  int64 `json:"page"`
		Limit int64 `json:"limit"`
	}
	var count int64

	c.ShouldBindJSON(&pagination)

	if pagination.Page == 0 {
		pagination.Page = 1
	}

	if pagination.Limit == 0 {
		pagination.Limit = 5
	}

	offset := (pagination.Page - 1) * pagination.Limit
	if result := database.DB.Offset(int(offset)).Limit(int(pagination.Limit)).Preload("User").Preload("Car").Find(&rental); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	if result := database.DB.Model(&models.Rental{}).Count(&count); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	totalPages := count / pagination.Limit

	c.JSON(http.StatusOK, gin.H{
		"Rental":      rental,
		"Total Pages": totalPages,
	})
}

func RentalShow(c *gin.Context) {
	id := c.Param("id")
	var rental models.Rental

	if err := database.DB.Preload("User").Preload("Car").First(&rental, id).Error; err != nil {
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

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&rental).Error; err != nil {
			return err
		}

		if err := tx.Preload("User").Preload("Car").First(&rental, rental.RentalID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data rental"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rental": rental})
}

func RentalUpdate(c *gin.Context) {
	var rental models.Rental
	id := c.Param("id")

	if err := database.DB.First(&rental, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"message": "Data rental tidak ditemukan"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data rental"})
			return
		}
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
