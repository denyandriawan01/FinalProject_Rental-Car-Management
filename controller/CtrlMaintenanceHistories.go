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

	database.DB.Find(&MaintenanceHistory)
	c.JSON(http.StatusOK, gin.H{"MaintenanceHistory": MaintenanceHistory})
}

func MaintenanceHistoryShow(c *gin.Context) {
	id := c.Param("id")
	var MaintenanceHistory models.MaintenanceHistory

	if err := database.DB.First(&MaintenanceHistory, id).Error; err != nil {
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
}

func MaintenanceHistoryCreate(c *gin.Context) {
	var MaintenanceHistory models.MaintenanceHistory

	if err := c.ShouldBindJSON(&MaintenanceHistory); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&MaintenanceHistory)
	c.JSON(http.StatusOK, gin.H{"MaintenanceHistory": MaintenanceHistory})
}

func MaintenanceHistoryUpdate(c *gin.Context) {
	var MaintenanceHistory models.MaintenanceHistory
	id := c.Param("id")

	if err := c.ShouldBindJSON(&MaintenanceHistory); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&MaintenanceHistory).Where("maintenance_id = ?", id).Updates(&MaintenanceHistory).RowsAffected == 0 {
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.ID.Int64()

	if err := database.DB.First(&MaintenanceHistory, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data histori perbaikan tidak ditemukan"})
		return
	}

	if database.DB.Delete(&MaintenanceHistory).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data histori perbaikan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data histori perbaikan berhasil dihapus"})
}