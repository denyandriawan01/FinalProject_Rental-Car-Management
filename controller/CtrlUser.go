package controller

import (
	"encoding/json"
	"finpro_golang/models"
	"finpro_golang/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserIndex(c *gin.Context) {
	var user []models.User

	database.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UserShow(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pengguna tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pengguna tidak ditemukan"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UserCreate(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UserUpdate(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&user).Where("user_id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data pengguna"})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.ID.Int64()

	if err := database.DB.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pengguna tidak ditemukan"})
		return
	}

	if database.DB.Delete(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}