package repository

import (
	"context"
	"os"

	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User repo struct for database communication
type UserRepository struct {
	collection *mongo.Collection
}

// Function to set the struct with a collection; returns pointer to that struct.
func NewUserRepository(client *mongo.Client) *UserRepository {
	dbName := os.Getenv("MONGO_DB_NAME")
	return &UserRepository{
		collection: client.Database(dbName).Collection("users"),
	}
}

// Function to use collection in struct instance to find user by email or username
func (r *UserRepository) FindUserByEmailOrUsername(ctx context.Context, identifier string) (*models.User, error) {
	var user models.User
	filter := bson.M{
		"$or": []bson.M{
			{"email": identifier},
			{"username": identifier},
		},
	}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Function to use collection in struct instance to create a new user.
func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	// Insert the user into the users collection
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
