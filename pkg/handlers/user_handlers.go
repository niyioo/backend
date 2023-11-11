// pkg/handlers/user_handlers.go

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/niyioo/backend/pkg/database"
	"github.com/niyioo/backend/pkg/models"
	"github.com/niyioo/backend/pkg/utils"
)

// GetUserFromContext retrieves the user from the request context.
func GetUserFromContext(r *http.Request) *models.User {
	userContext := r.Context().Value(ContextUserKey)
	if userContext != nil {
		if user, ok := userContext.(*models.User); ok {
			return user
		}
	}
	return nil
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
