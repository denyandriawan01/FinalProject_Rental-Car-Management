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

func MaintenanceHistoryIndex(c *gin.Context) {
<<<<<<< HEAD
	var maintenancehistory []models.MaintenanceHistory
	var pagination struct {
		Page  int64 `json:"page"`
		Limit int64 `json:"limit"`
	}
	var count int64

	if err := c.ShouldBindJSON(&pagination); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if pagination.Page == 0 {
		pagination.Page = 1
	}

	if pagination.Limit == 0 {
		pagination.Limit = 5
	}

	offset := (pagination.Page - 1) * pagination.Limit
	if result := database.DB.Offset(int(offset)).Limit(int(pagination.Limit)).Preload("Car").Find(&maintenancehistory); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	if result := database.DB.Model(&models.MaintenanceHistory{}).Count(&count); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	totalPages := count / pagination.Limit

	c.JSON(http.StatusOK, gin.H{
		"Maintenance History": maintenancehistory,
		"Total Pages":         totalPages,
	})
=======
	var MaintenanceHistory []models.MaintenanceHistory

	models.DB.Find(&MaintenanceHistory)
	c.JSON(http.StatusOK, gin.H{"MaintenanceHistory": MaintenanceHistory})
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
}

func MaintenanceHistoryShow(c *gin.Context) {
	id := c.Param("id")
<<<<<<< HEAD
	var maintenanceHistory models.MaintenanceHistory

	if err := database.DB.Preload("Car").First(&maintenanceHistory, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data histori perbaikan tidak ditemukan"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Terjadi kesalahan saat mengambil data histori perbaikan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Maintenance History": maintenanceHistory})
=======
	var MaintenanceHistory models.MaintenanceHistory

	if err := models.DB.First(&MaintenanceHistory, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data histori perbaikan tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data histori perbaikan tidak ditemukan"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"MaintenanceHistory": MaintenanceHistory})
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
}

func MaintenanceHistoryCreate(c *gin.Context) {
	var MaintenanceHistory models.MaintenanceHistory

	if err := c.ShouldBindJSON(&MaintenanceHistory); err != nil {
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&MaintenanceHistory).Error; err != nil {
			return err
		}

		if err := tx.Preload("Car").First(&MaintenanceHistory, MaintenanceHistory.MaintenanceID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data histori perbaikan"})
		return
	}

=======
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&MaintenanceHistory)
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
	c.JSON(http.StatusOK, gin.H{"MaintenanceHistory": MaintenanceHistory})
}

func MaintenanceHistoryUpdate(c *gin.Context) {
	var MaintenanceHistory models.MaintenanceHistory
	id := c.Param("id")

	if err := c.ShouldBindJSON(&MaintenanceHistory); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

<<<<<<< HEAD
	if database.DB.Model(&MaintenanceHistory).Where("maintenance_id = ?", id).Updates(&MaintenanceHistory).RowsAffected == 0 {
=======
	if models.DB.Model(&MaintenanceHistory).Where("maintenance_id = ?", id).Updates(&MaintenanceHistory).RowsAffected == 0 {
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data histori perbaikan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data histori perbaikan berhasil diperbarui"})
}

func MaintenanceHistoryDelete(c *gin.Context) {
	var MaintenanceHistory models.MaintenanceHistory

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
	if err := database.DB.First(&MaintenanceHistory, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data histori perbaikan tidak ditemukan"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data histori perbaikan"})
		return
	}

	if database.DB.Delete(&MaintenanceHistory).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data histori perbaikan"})
=======
	if err := models.DB.First(&MaintenanceHistory, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data histori perbaikan tidak ditemukan"})
		return
	}

	if models.DB.Delete(&MaintenanceHistory).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data histori perbaikan"})
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data histori perbaikan berhasil dihapus"})
}
