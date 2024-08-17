package repository

import (
	"context"
	"os"

	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	dbName := os.Getenv("MONGO_DB_NAME")
	return &UserRepository{
		collection: client.Database(dbName).Collection("users"),
	}
}

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

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	// Insert the user into the users collection
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
