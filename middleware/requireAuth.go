package middleware

import (
<<<<<<< HEAD
	"FinalProject_Rental-Car-Management/database"
	"FinalProject_Rental-Car-Management/models"
	"fmt"
=======
	"fmt"
	"models"
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	// Get the token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Split the Authorization header value to extract the token
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Split gagal", "authHeaderParts": authHeaderParts})
		return
	}

	tokenString := authHeaderParts[1]

	// Validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g., []byte("my_secret_key")
		return []byte(os.Getenv("JWTKEY")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token is not valid"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check expiration
		expirationTime := int64(claims["exp"].(float64))
		currentTime := time.Now().Unix()
		if currentTime > expirationTime {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check user
		var user models.User
<<<<<<< HEAD
		database.DB.Where("username = ?", claims["sub"]).First(&user)
=======
		models.DB.Where("username = ?", claims["sub"]).First(&user)
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
		if user.UserID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach user to the request context
		c.Set("user", user)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "CLAIM failed"})
	}
}
