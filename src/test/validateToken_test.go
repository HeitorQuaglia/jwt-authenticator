package test

import (
	jwtTest "jwt-authenticator/src/jwt"
	"testing"
)

func TestValidateJWTToken(t *testing.T) {
	// Set up input parameters
	username := "testuser"
	secretKey := "testkey"

	// Generate JWT token
	tokenString, err := jwtTest.GenerateJWTToken(username, secretKey)
	if err != nil {
		t.Errorf("Unexpected error generating JWT token: %v", err)
	}

	// Validate JWT token
	valid, err := jwtTest.ValidateJWTToken(tokenString, secretKey)
	if err != nil {
		t.Errorf("Unexpected error validating JWT token: %v", err)
	}

	// Verify validation result
	if !valid {
		t.Error("JWT token is not valid")
	}
}
