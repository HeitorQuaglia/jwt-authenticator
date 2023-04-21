package jwt

import (
	"errors"
	jwt "github.com/golang-jwt/jwt/v4"
)

func ValidateJWTToken(tokenString string, secretKey string) (bool, error) {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing algorithm")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, err
	}

	// Check if token is valid
	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return false, errors.New("invalid token")
	}

	return true, nil
}
