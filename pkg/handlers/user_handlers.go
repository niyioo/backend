package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/niyioo/backend/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// Register handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Handle error
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		// Handle error
	}

	user.Password = string(hashedPassword)

	// Save the user to the database (implement this)

	// Generate JWT
	token, err := GenerateToken(user)
	if err != nil {
		// Handle error
	}

	user.Token = token

	// Return response with user and token
	// (you may want to send just the token in production)
}

// Login handles user login
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Handle error
	}

	// Fetch user from the database by username (implement this)

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(savedUser.Password))
	if err != nil {
		// Handle error - invalid password
	}

	// Generate JWT
	token, err := GenerateToken(savedUser)
	if err != nil {
		// Handle error
	}

	savedUser.Token = token

	// Return response with user and token
	// (you may want to send just the token in production)
}
