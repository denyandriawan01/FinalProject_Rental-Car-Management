package controller

import (
	"encoding/json"
	"net/http"

	"FinalProject_Rental-Car-Management/database"
	"FinalProject_Rental-Car-Management/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CarsIndex(c *gin.Context) {
	var cars []models.Car
	var count int64

	if result := database.DB.Find(&cars); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	if result := database.DB.Model(&cars).Count(&count); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Cars":       cars,
		"TotalCount": count,
	})

}

func CarsShow(c *gin.Context) {
	id := c.Param("id")
	var cars models.Car

	if err := database.DB.First(&cars, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data mobil tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data mobil"})
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

	if err := database.DB.Create(&cars).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data mobil"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cars": cars})
}

func CarsUpdate(c *gin.Context) {
	id := c.Param("id")

	var car models.Car
	if err := database.DB.First(&car, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"message": "Data mobil tidak ditemukan"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data mobil"})
			return
		}
	}

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := database.DB.Save(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui data mobil"})
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

	if err := database.DB.First(&cars, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data mobil tidak ditemukan"})
		return
	}

	if database.DB.Delete(&cars).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data mobil"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
