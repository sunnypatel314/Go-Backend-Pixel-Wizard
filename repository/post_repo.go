package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostRepository struct {
	collection *mongo.Collection
}

func NewPostRepository(client *mongo.Client) *PostRepository {
	dbName := os.Getenv("MONGO_DB_NAME")
	return &PostRepository{
		collection: client.Database(dbName).Collection("posts"),
	}
}

func (r *PostRepository) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post models.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) CreatePost(ctx context.Context, post *models.Post) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, post, &options.InsertOneOptions{})
}

func (r *PostRepository) DeletePost(ctx context.Context, postID string) error {
	objectID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with the given ID")
	}
	return err
}
