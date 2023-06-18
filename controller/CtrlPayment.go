package controller

import (
	"encoding/json"
	"net/http"

	"finpro_golang/models"
	"finpro_golang/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PaymentIndex(c *gin.Context) {
	var payment []models.Payment

	if err := database.DB.Preload("Rental").Preload("Rental.User").Preload("Rental.Car").Find(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data pembayaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

func PaymentShow(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment

	if err := database.DB.Preload("Rental").Preload("Rental.User").Preload("Rental.Car").First(&payment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data pembayaran tidak ditemukan"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data pembayaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

func PaymentCreate(c *gin.Context) {
	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&payment).Error; err != nil {
			return err
		}

		if err := tx.Preload("Rental").Preload("Rental.User").Preload("Rental.Car").First(&payment, payment.ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data pembayaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

func PaymentUpdate(c *gin.Context) {
	var payment models.Payment
	id := c.Param("id")

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&payment).Where("payment_id = ?", id).Updates(&payment).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data pembayaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data pembayaran berhasil diperbarui"})
}

func PaymentDelete(c *gin.Context) {
	var payment models.Payment

	var input struct {
		ID json.Number `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.ID.Int64()

	if err := database.DB.First(&payment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data pembayaran tidak ditemukan"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data pembayaran"})
		return
	}

	if database.DB.Delete(&payment).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data pembayaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data pembayaran berhasil dihapus"})
}