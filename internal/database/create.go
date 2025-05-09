package database

import (
	"context"
	"log"
	"time"

	"fitness-tracker/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) (userID primitive.ObjectID, err error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    collection := GetCollection("users")

    result, err := collection.InsertOne(ctx, user)
    if err != nil {
        log.Println("Couldn't create user")
        return
    }

    userID = result.InsertedID.(primitive.ObjectID)
    return
}