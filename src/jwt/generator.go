package jwt

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateJWTToken(username string, secretKey string) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Create token with claims and secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
