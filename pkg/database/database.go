package database

import (
	"context"
	"log"

	"github.com/niyioo/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// SavePost saves a blog post to the database
func SavePost(post models.Post) error {
	collection := client.Database("your_database_name").Collection("posts")

	_, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// GetPosts retrieves all blog posts from the database
func GetPosts() ([]models.Post, error) {
	var posts []models.Post
	collection := client.Database("your_database_name").Collection("posts")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var post models.Post
		if err := cursor.Decode(&post); err != nil {
			log.Fatal(err)
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
