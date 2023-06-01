package controller

import (
	"github.com/gin-gonic/gin"
	"models"
	"net/http"
)

func GetMaintenanceHistories(c *gin.Context) {
	var maintenanceHistories []models.MaintenanceHistory
	models.DB.Find(&maintenanceHistories)
	c.JSON(http.StatusOK, gin.H{"data": maintenanceHistories})
}

func GetMaintenanceHistory(c *gin.Context) {
	var maintenanceHistory models.MaintenanceHistory
	if err := models.DB.First(&maintenanceHistory, c.Param("maintenance_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": maintenanceHistory})
}

func CreateMaintenanceHistory(c *gin.Context) {
	var input models.MaintenanceHistory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func UpdateMaintenanceHistory(c *gin.Context) {
	var maintenanceHistory models.MaintenanceHistory
	if err := models.DB.First(&maintenanceHistory, c.Param("maintenance_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input models.MaintenanceHistory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&maintenanceHistory).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": maintenanceHistory})
}

func DeleteMaintenanceHistory(c *gin.Context) {
	var maintenanceHistory models.MaintenanceHistory
	if err := models.DB.First(&maintenanceHistory, c.Param("maintenance_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&maintenanceHistory)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
