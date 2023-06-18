package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"FinalProject_Rental-Car-Management/database"
	"FinalProject_Rental-Car-Management/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TaxIndex(c *gin.Context) {
	var tax []models.Taxes
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
	if result := database.DB.Offset(int(offset)).Limit(int(pagination.Limit)).Preload("Car").Find(&tax); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	if result := database.DB.Model(&models.Taxes{}).Count(&count); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Conflict occurred",
		})
		return
	}

	totalPages := count / pagination.Limit

	c.JSON(http.StatusOK, gin.H{
		"Taxes":       tax,
		"Total Pages": totalPages,
	})
}

func TaxShow(c *gin.Context) {
	id := c.Param("id")
	var tax models.Taxes

	if err := database.DB.Preload("Car").First(&tax, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data pajak tidak ditemukan"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Terjadi kesalahan saat mengambil data pajak"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tax": tax})
}

func GetTaxesByCarID(c *gin.Context) {
	id := c.Param("id")
	var taxes []models.Taxes

	pageNumber, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		pageNumber = 1
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10 // default page size
	}

	offset := (pageNumber - 1) * pageSize

	if err := database.DB.Where("car_id = ?", id).Offset(offset).Limit(pageSize).Find(&taxes).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve taxes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"taxes":       taxes,
		"currentPage": pageNumber,
		"pageSize":    pageSize,
	})
}

func TaxCreate(c *gin.Context) {
	var tax models.Taxes

	if err := c.ShouldBindJSON(&tax); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&tax).Error; err != nil {
			return err
		}

		if err := tx.Preload("Car").First(&tax, tax.TaxID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data Maintenance History"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tax": tax})
}

func TaxUpdate(c *gin.Context) {
	var tax models.Taxes
	id := c.Param("id")

	if err := c.ShouldBindJSON(&tax); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&tax).Where("tax_id = ?", id).Updates(&tax).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data pajak"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data pajak berhasil diperbarui"})
}

func TaxDelete(c *gin.Context) {
	var tax models.Taxes

	var input struct {
		ID json.Number `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.ID.Int64()

	if err := database.DB.First(&tax, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data pajak tidak ditemukan"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data pajak"})
		return
	}

	if err := database.DB.Delete(&tax).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data pajak"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
