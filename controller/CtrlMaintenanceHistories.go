package controller

import (
	"encoding/json"
	"net/http"

	"finpro_golang/models"
	"finpro_golang/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MaintenanceHistoryIndex(c *gin.Context) {
	var MaintenanceHistory []models.MaintenanceHistory

	if err := database.DB.Preload("Car").Find(&MaintenanceHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data histori perbaikan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"MaintenanceHistory": MaintenanceHistory})
}

func MaintenanceHistoryShow(c *gin.Context) {
	id := c.Param("id")
	var MaintenanceHistory models.MaintenanceHistory

	if err := database.DB.Preload("Car").First(&MaintenanceHistory, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data histori perbaikan tidak ditemukan"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data histori perbaikan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"MaintenanceHistory": MaintenanceHistory})
}

func MaintenanceHistoryCreate(c *gin.Context) {
	var MaintenanceHistory models.MaintenanceHistory

	if err := c.ShouldBindJSON(&MaintenanceHistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&MaintenanceHistory).Error; err != nil {
			return err
		}

		if err := tx.Preload("Car").First(&MaintenanceHistory, MaintenanceHistory.ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data histori perbaikan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"MaintenanceHistory": MaintenanceHistory})
}

func MaintenanceHistoryUpdate(c *gin.Context) {
	var MaintenanceHistory models.MaintenanceHistory
	id := c.Param("id")

	if err := c.ShouldBindJSON(&MaintenanceHistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&MaintenanceHistory).Where("maintenance_id = ?", id).Updates(&MaintenanceHistory).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data histori perbaikan"})
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
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.ID.Int64()

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
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data histori perbaikan berhasil dihapus"})
}