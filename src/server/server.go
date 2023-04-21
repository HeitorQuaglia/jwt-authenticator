package main

import (
	"fmt"
	"jwt-authenticator/src/jwt"
	"net/http"
	"time"
)

func main() {
	// Secret key for JWT
	secretKey := "my-secret-key"

	// HTTP route for generating JWT tokens
	http.HandleFunc("/generate-token", func(w http.ResponseWriter, r *http.Request) {
		// Get username from query parameter
		username := r.PostFormValue("username")

		// Generate JWT token
		token, err := jwt.GenerateJWTToken(username, secretKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write token to response
		fmt.Fprintf(w, "%s", token)
	})

	// HTTP route for validating JWT tokens
	http.HandleFunc("/validate-token", func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		tokenString := r.Header.Get("Authorization")[7:] // remove "Bearer " prefix

		// Validate JWT token
		valid, err := jwt.ValidateJWTToken(tokenString, secretKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write validation status to response
		fmt.Fprint(w, map[bool]string{valid: "Valid token", !valid: "Invalid token"}[true])
	})

	// Start HTTP server
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 30,
	}
	fmt.Println("Starting server on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
