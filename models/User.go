// models/user.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}