package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/niyioo/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePost handles the creation of a new blog post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		// Handle error
	}

	// Retrieve the authenticated user from the context
	user, ok := GetUserFromContext(r.Context())
	if !ok {
		// Handle error - user not found in context
	}

	// Set the author of the post to the authenticated user
	post.Author = user

	// Set the creation timestamp
	post.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	// Save the post to the database (implement this)

	// Return response with the created post
}

// GetPosts handles retrieving all blog posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	// Retrieve all blog posts from the database (implement this)

	// Return response with the list of posts
}
