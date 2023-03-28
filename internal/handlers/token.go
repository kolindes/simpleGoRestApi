package handlers

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kolindes/simpleRestApi/internal/config"
)

func generateToken(userID int64) (string, error) {
	config, err := config.Load()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %v", err)
	}

	expirationTime := time.Now().Add(time.Hour * time.Duration(config.JWT.ExpiresInHours)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": expirationTime,
	})

	tokenString, err := token.SignedString([]byte(config.JWT.SecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}
