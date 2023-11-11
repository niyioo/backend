// pkg/database/user.go

package database

import (
	"context"
	"time"
)

// User represents the user model.
type User struct {
	// Define your user model fields here
	Username string `bson:"username"`
	Password string `bson:"password"`
	// Add more fields as needed
}

// SaveUser saves the user information to the database.
func SaveUser(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Assuming you have a collection named "users" in your MongoDB database
	collection := client.Database(DatabaseName).Collection("users")

	// Insert the user into the "users" collection
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
