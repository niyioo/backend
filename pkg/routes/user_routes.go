package routes

import (
	"github.com/gorilla/mux"
	"github.com/niyioo/backend/pkg/handlers"
)

// SetUserRoutes sets up routes for user-related operations
func SetUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/register", handlers.Register).Methods("POST")
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")
}
