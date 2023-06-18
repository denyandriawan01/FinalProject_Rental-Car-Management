package controller

import (
<<<<<<< HEAD
	"FinalProject_Rental-Car-Management/database"
	"FinalProject_Rental-Car-Management/models"
	"FinalProject_Rental-Car-Management/utils/token"
	"net/http"
	"os"
=======
	"models"
	"net/http"
	"os"
	"token"
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleLogin(c *gin.Context) {
	var form models.LoginForm

	if err := c.BindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User

<<<<<<< HEAD
	if err := database.DB.First(&user, "username = ?", form.Username).Error; err != nil {
=======
	if err := models.DB.First(&user, "username = ?", form.Username).Error; err != nil {
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Data pengguna tidak ditemukan"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
			return
		}
	}

	// Compare passwords
	if user.Password != form.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	// Generate JWT token
	tokenString, err := token.GenerateTokenString(form, os.Getenv("JWTKEY"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	// Send it back in the response
	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func HandleLogout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}
