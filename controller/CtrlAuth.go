package controller

import (
	"time"
	"strings"
	"net/http"

	"finpro_golang/models"
	"finpro_golang/database"
	"finpro_golang/utils/token"
	"finpro_golang/utils/initializer"
    
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
    var form models.LoginForm

    if err := c.ShouldBindJSON(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Permintaan tidak valid"})
        return
    }

    var user models.User
    if err := database.DB.Where("username = ?", form.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Data pengguna tidak ditemukan"})
        return
    }

    if user.Password != form.Password {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Username atau Kata Sandi tidak valid"})
        return
    }

    existingToken := c.GetHeader("Authorization")
    if existingToken != "" {
        existingToken = strings.Replace(existingToken, "Bearer ", "", 1)
        token.BlacklistToken(existingToken)
    }

    expirationTime := time.Now().Add(time.Duration(initializer.EXP_TOKEN) * time.Minute)
    tokenString, err := token.GenerateTokenString(form, initializer.JWT_KEY, expirationTime)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghasilkan token"})
        return
    }

    expTimeFormatted := expirationTime.Format("2006-01-02 15:04:05")
    c.JSON(http.StatusOK, gin.H{"token": "Bearer "+tokenString, "exp_token": expTimeFormatted})
}

func LogoutHandler(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
        return
    }

    tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

    token.BlacklistToken(tokenString)

    c.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
}