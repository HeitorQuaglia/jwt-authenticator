package test

import (
	jwtTest "jwt-authenticator/src/jwt"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

func TestGenerateJWTToken(t *testing.T) {
	// Set up input parameters
	username := "testuser"
	secretKey := "testkey"

	// Generate JWT token
	tokenString, err := jwtTest.GenerateJWTToken(username, secretKey)
	if err != nil {
		t.Errorf("Unexpected error generating JWT token: %v", err)
	}

	// Parse JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		t.Errorf("Unexpected error parsing JWT token: %v", err)
	}

	// Verify JWT token
	if !token.Valid {
		t.Error("JWT token is not valid")
	}

	// Verify claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Error("Unexpected type for JWT claims")
	}

	if usernameClaim, ok := claims["username"].(string); !ok || usernameClaim != username {
		t.Errorf("Unexpected username claim value: %v", usernameClaim)
	}
}
