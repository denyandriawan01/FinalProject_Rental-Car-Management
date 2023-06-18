package controller

import (
	"encoding/json"
	"net/http"

<<<<<<< HEAD
	"FinalProject_Rental-Car-Management/database"
	"FinalProject_Rental-Car-Management/models"
=======
	"models"
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PaymentIndex(c *gin.Context) {
	var payment []models.Payment
<<<<<<< HEAD
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
	if result := database.DB.Offset(int(offset)).Limit(int(pagination.Limit)).Find(&payment); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	if err := database.DB.Preload("Rental").Preload("Rental.User").Preload("Rental.Car").Find(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data pembayaran"})
		return
	}

	if result := database.DB.Model(&payment).Count(&count); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	totalPages := count / pagination.Limit

	c.JSON(http.StatusOK, gin.H{
		"Payment":     payment,
		"Total Pages": totalPages,
	})
=======

	models.DB.Find(&payment)
	c.JSON(http.StatusOK, gin.H{"payment": payment})
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
}

func PaymentShow(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment

<<<<<<< HEAD
	if err := database.DB.Preload("Rental").Preload("Rental.User").Preload("Rental.Car").First(&payment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data pembayaran tidak ditemukan"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data pembayaran"})
		return
=======
	if err := models.DB.First(&payment, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pembayaran tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pembayaran tidak ditemukan"})
			return
		}
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
	}

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

func PaymentCreate(c *gin.Context) {
	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&payment).Error; err != nil {
			return err
		}

		if err := tx.Preload("Rental").Preload("Rental.User").Preload("Rental.Car").First(&payment, payment.PaymentID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data pembayaran"})
		return
	}

=======
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&payment)
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

func PaymentUpdate(c *gin.Context) {
	var payment models.Payment
	id := c.Param("id")

	if err := c.ShouldBindJSON(&payment); err != nil {
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&payment).Where("payment_id = ?", id).Updates(&payment).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data pembayaran"})
=======
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&payment).Where("payment_id = ?", id).Updates(&payment).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data pembayaran"})
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
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
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
=======
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
		return
	}

	id, _ := input.ID.Int64()

<<<<<<< HEAD
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
=======
	if err := models.DB.First(&payment, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pembayaran tidak ditemukan"})
		return
	}

	if models.DB.Delete(&payment).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data pembayaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
}
