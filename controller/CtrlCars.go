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

func CarsIndex(c *gin.Context) {
	var cars []models.Car
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
	if result := database.DB.Offset(int(offset)).Limit(int(pagination.Limit)).Find(&cars); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
	}

	if result := database.DB.Model(&cars).Count(&count); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
	}

	totalPages := count / pagination.Limit

	c.JSON(http.StatusOK, gin.H{
		"Cars":        cars,
		"Total Pages": totalPages,
	})
=======

	models.DB.Find(&cars)
	c.JSON(http.StatusOK, gin.H{"cars": cars})
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
}

func CarsShow(c *gin.Context) {
	id := c.Param("id")
	var cars models.Car

<<<<<<< HEAD
	if err := database.DB.First(&cars, id).Error; err != nil {
=======
	if err := models.DB.First(&cars, id).Error; err != nil {
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data mobil tidak ditemukan"})
			return
		default:
<<<<<<< HEAD
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data mobil"})
=======
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data mobil tidak ditemukan"})
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
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

<<<<<<< HEAD
	if err := database.DB.Create(&cars).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data mobil"})
		return
	}

=======
	models.DB.Create(&cars)
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
	c.JSON(http.StatusOK, gin.H{"cars": cars})
}

func CarsUpdate(c *gin.Context) {
<<<<<<< HEAD
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
=======
	var cars models.Car
	id := c.Param("id")

	if err := c.ShouldBindJSON(&cars); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&cars).Where("car_id = ?", id).Updates(&cars).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data mobil"})
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data mobil berhasil diperbarui"})
}

<<<<<<< HEAD
=======
func CarsUpdateAvailability(c *gin.Context) {
	var cars models.Car
	id := c.Param("id")

	if err := c.ShouldBindJSON(&cars); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updateFields := make(map[string]interface{})
	updateFields["is_available"] = cars.IsAvailable

	if models.DB.Model(&models.Car{}).Where("car_id = ?", id).Updates(updateFields).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Status ketersediaan tidak dapat diupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data ketersediaan mobil berhasil diperbarui"})
}

>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
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

<<<<<<< HEAD
	if err := database.DB.First(&cars, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data mobil tidak ditemukan"})
		return
	}

	if database.DB.Delete(&cars).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data mobil"})
=======
	if err := models.DB.First(&cars, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "cars tidak ditemukan"})
		return
	}

	if models.DB.Delete(&cars).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus cars"})
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
