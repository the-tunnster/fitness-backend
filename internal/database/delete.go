package database

import (
	"time"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteSession(userId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("sessions")

	_, err := collection.DeleteOne(ctx, bson.M{"userId": userId})
	return err
}