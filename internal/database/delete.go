package database

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteSession(userID primitive.ObjectID) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("sessions")

	_, err = collection.DeleteOne(ctx, bson.M{"userID": userID})
	return
}

func DeleteRoutine(userID, routiuneID primitive.ObjectID) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("routines")

	result, err := collection.DeleteOne(ctx, bson.M{
		"_id":    routiuneID,
		"userID": userID,
	})

	if err != nil {
		log.Println("Failed to delete routine")
		return err
	}

	if result.DeletedCount == 0 {
		log.Println("Coudn't find specified routine")
		return errors.New("no document with matching routineID and userID")
	}

	return
}