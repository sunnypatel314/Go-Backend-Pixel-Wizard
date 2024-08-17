// models/post.go
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CreatorID   primitive.ObjectID `bson:"creatorId"`
	CreatorName string             `bson:"creator"`
	Prompt      string             `bson:"prompt"`
	PhotoURL    string             `bson:"photo_url"`
	CreatedDate time.Time          `bson:"created_date"`
}
