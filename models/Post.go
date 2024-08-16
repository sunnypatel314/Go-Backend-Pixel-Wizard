// models/post.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Creator     string             `bson:"creator"`
	Prompt      string             `bson:"prompt"`
	PhotoURL    string             `bson:"photo_url"`
	CreatedDate string             `bson:"created_date"`
}
