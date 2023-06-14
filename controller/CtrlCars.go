package controller

import (
	"encoding/json"
	"net/http"

	"models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CarsIndex(c *gin.Context) {
	var cars []models.Car
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
	if result := models.DB.Offset(int(offset)).Limit(int(pagination.Limit)).Find(&cars); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
	}

	if result := models.DB.Model(&cars).Count(&count); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
	}

	totalPages := count / pagination.Limit

	c.JSON(http.StatusOK, gin.H{
		"cars":       cars,
		"totalPages": totalPages,
	})
}

func CarsShow(c *gin.Context) {
	id := c.Param("id")
	var cars models.Car

	if err := models.DB.First(&cars, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data mobil tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data mobil tidak ditemukan"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"cars": cars})
}

func CarsCreate(c *gin.Context) {
	var cars models.Car

	if err := c.ShouldBindJSON(&cars); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&cars)
	c.JSON(http.StatusOK, gin.H{"cars": cars})
}

func CarsUpdate(c *gin.Context) {
	var cars models.Car
	id := c.Param("id")

	if err := c.ShouldBindJSON(&cars); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&cars).Where("car_id = ?", id).Updates(&cars).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data mobil"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data mobil berhasil diperbarui"})
}

func CarsDelete(c *gin.Context) {
	var cars models.Car

	var input struct {
		ID json.Number `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.ID.Int64()

	if err := models.DB.First(&cars, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "cars tidak ditemukan"})
		return
	}

	if models.DB.Delete(&cars).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus cars"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
