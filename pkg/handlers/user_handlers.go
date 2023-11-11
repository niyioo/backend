// pkg/handlers/user_handlers.go

package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/niyioo/backend/pkg/database"
	"github.com/niyioo/backend/pkg/models"
	"github.com/niyioo/backend/pkg/utils"
)

// GetUserFromContext retrieves the user from the request context.
func GetUserFromContext(r *http.Request) (*models.User, error) {
	user, ok := r.Context().Value(database.ContextUserKey).(*models.User)
	if !ok {
		return nil, errors.New("user not found in request context")
	}
	return user, nil
}

// Sample code to fix undefined savedUser

// CreateAccount handles the creation of a new user account.
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Hash the password before saving to the database
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error hashing password")
		return
	}

	newUser.Password = hashedPassword

	// Save the user to the database
	savedUser, err := database.SaveUser(newUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, savedUser)
}

// HandleUserRegistration handles user registration.
func HandleUserRegistration(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get user registration details
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	// Hash the user's password before saving to the database
	hashedPassword, err := database.HashPassword(newUser.Password)
	if err != nil {
		http.Error(w, "failed to hash password", http.StatusInternalServerError)
		return
	}

	// Assign the hashed password to the user
	newUser.Password = hashedPassword

	// Save the user to the database
	if err := database.SaveUser(&newUser); err != nil {
		http.Error(w, "failed to save user to database", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}
