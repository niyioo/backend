package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Post represents a blog post
type Post struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	Author    string             `json:"author" bson:"author"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at"`
}
