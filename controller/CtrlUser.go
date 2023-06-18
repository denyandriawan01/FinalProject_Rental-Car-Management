package controller

import (
	"FinalProject_Rental-Car-Management/database"
	"FinalProject_Rental-Car-Management/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserIndex(c *gin.Context) {
	var user []models.User
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
	if result := database.DB.Offset(int(offset)).Limit(int(pagination.Limit)).Find(&user); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	if result := database.DB.Model(&models.User{}).Count(&count); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	totalPages := count / pagination.Limit

	c.JSON(http.StatusOK, gin.H{
		"User":        user,
		"Total Pages": totalPages,
	})
}

func UserShow(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"message": "Data pengguna tidak ditemukan"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data pengguna"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UserCreate(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UserUpdate(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := database.DB.Model(&user).Where("user_id = ?", id).Updates(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui data pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data pengguna berhasil diperbarui"})
}

func UserDelete(c *gin.Context) {
	var user models.User

	var input struct {
		ID json.Number `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.ID.Int64()

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data pengguna tidak ditemukan"})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
