package database

import (
	"context"
	"log"
	"time"

	"fitness-tracker/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertSession(session models.Session) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("sessions")

	filter := bson.M{"userID": session.UserID}
	update := bson.M{"$set": session}
	opts := options.Update().SetUpsert(true)

	_, err = collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println("Error updating session")
	}

	return
}

func UpdateUser(userID primitive.ObjectID, updates bson.M) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("users")
	
	_, err = collection.UpdateOne(ctx,
		bson.M{
			"_id": userID,
		},
		bson.M{
			"$set": updates,
		},
	)
	if err != nil {
		log.Println("Error updating used information")
		log.Println(err)
	}

	return
}

func UpdateExercise(exerciseID primitive.ObjectID, updates bson.M) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("exercises")

	result, err := collection.UpdateOne(ctx,
		bson.M{
			"_id":    exerciseID,
		},
		bson.M{
			"$set": updates,
		},
	)

	if err != nil {
		log.Println("Error updating exercise")
		return
	}

	if result.MatchedCount == 0 {
		log.Println("No mathcing exercise id found")
		return
	}
	
	return
}

func UpdateRoutine(routineID, userID primitive.ObjectID, updates bson.M) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("routines")

	result, err := collection.UpdateOne(ctx,
		bson.M{
			"_id":    routineID,
			"userID": userID,
		},
		bson.M{
			"$set": updates,
		},
	)

	if err != nil {
		log.Println("Error updating routine")
		return
	}

	if result.MatchedCount == 0 {
		log.Println("No mathcing user id found")
		return
	}
	
	return
}