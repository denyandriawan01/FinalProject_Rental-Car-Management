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
	var count int64

	if result := database.DB.Preload("User").Preload("Car").Find(&rental); result.Error != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"Rental": rental,
		"Count":  count,
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

	rental.IsCompleted = false

	var car models.Car
	if err := database.DB.First(&car, rental.CarID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Mobil tidak ditemukan"})
		return
	}

	if !car.IsAvailable {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Mobil sedang tidak tersedia"})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&rental).Error; err != nil {
			return err
		}

		if err := tx.Preload("User").Preload("Car").First(&rental, rental.RentalID).Error; err != nil {
			return err
		}

		if rental.IsCompleted {
			if err := tx.Model(&car).Update("is_available", true).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Model(&car).Update("is_available", false).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&rental).Preload("Car").First(&rental).Error; err != nil {
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

	if err := database.DB.Preload("Car").First(&rental, id).Error; err != nil {
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

	if err := database.DB.Model(&rental).Update("is_completed", rental.IsCompleted).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui data rental"})
		return
	}

	if rental.IsCompleted {
		if err := database.DB.Model(&rental.Car).Update("is_available", true).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui status ketersediaan mobil"})
			return
		}
	} else {
		if err := database.DB.Model(&rental.Car).Update("is_available", false).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui status ketersediaan mobil"})
			return
		}
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
