package controller

import (
	"encoding/json"
	"net/http"

	"models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PaymentIndex(c *gin.Context) {
	var payment []models.Payment

	models.DB.Find(&payment)
	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

func PaymentShow(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment

	if err := models.DB.First(&payment, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pembayaran tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pembayaran tidak ditemukan"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

func PaymentCreate(c *gin.Context) {
	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&payment)
	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

func PaymentUpdate(c *gin.Context) {
	var payment models.Payment
	id := c.Param("id")

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&payment).Where("payment_id = ?", id).Updates(&payment).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data pembayaran"})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.ID.Int64()

	if err := models.DB.First(&payment, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pembayaran tidak ditemukan"})
		return
	}

	if models.DB.Delete(&payment).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data pembayaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
