package token

import (
<<<<<<< HEAD
	"FinalProject_Rental-Car-Management/models"
	"errors"
=======
	"errors"
	"models"
>>>>>>> 3789ae5c6753f40b0970d347d395440182ea9a98
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateTokenString(form models.LoginForm, secretKey string) (string, error) {
	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": form.Username,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("Failed to create token")
	}

	return tokenString, nil
}
