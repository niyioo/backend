// pkg/database/database.go

package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseURL is the MongoDB URL.
const DatabaseURL = "mongodb://localhost:27017"

// DatabaseName is the name of the MongoDB database.
const DatabaseName = "blog"

var client *mongo.Client

// Connect initializes the MongoDB connection.
func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(DatabaseURL)
	var err error // Change the variable name to avoid shadowing
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
}
