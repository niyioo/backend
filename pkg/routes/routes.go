package routes

import (
	"net/http"

	"github.com/niyioo/backend/pkg/handlers"
	"github.com/niyioo/backend/pkg/middleware"
)

// SetPostRoutes sets the routes for blog posts
func SetPostRoutes() {
	http.HandleFunc("/api/posts", middleware.AuthenticateMiddleware(handlers.CreatePost)).Methods("POST")
	http.HandleFunc("/api/posts", handlers.GetPosts).Methods("GET")
}
