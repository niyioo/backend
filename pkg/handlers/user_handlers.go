package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/niyioo/backend/pkg/models"
)

// Register handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// Parse the request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Handle error
	}

	// Validate and register the user (implement this)

	// Return response
}

// Login handles user login
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// Parse the request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Handle error
	}

	// Validate and perform login (implement this)

	// Return response
}
