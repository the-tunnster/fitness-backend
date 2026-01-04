package database

import (
	"context"
	"log"
	"time"

	"fitness-tracker/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertSession(session models.WorkoutSession) (sessionID primitive.ObjectID, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("sessions")

	// Single active session per user
	filter := bson.M{"userID": session.UserID}
	// Ensure TTL is current
	session.LastUpdate = primitive.NewDateTimeFromTime(time.Now())
	update := bson.M{"$set": session}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var result models.WorkoutSession
	err = collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Shouldn't happen with ReturnDocumentAfter, but handle defensively
			log.Println("UpsertSession: no document returned after upsert")
		} else {
			log.Println("Error upserting session", err)
		}
		return primitive.NilObjectID, err
	}
	return result.ID, nil
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
		log.Println("Error updating user information")
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
			"_id": exerciseID,
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

func UpdateSession(sessionID primitive.ObjectID, updates bson.M) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("sessions")

	result, err := collection.UpdateOne(ctx,
		bson.M{
			"_id": sessionID,
		},
		bson.M{
			"$set": updates,
		},
	)

	if err != nil {
		log.Println("Error updating session")
		return
	}

	if result.MatchedCount == 0 {
		log.Println("No mathcing session found")
		return
	}

	return
}

func UpdateExerciseHistory(exerciseID, userID primitive.ObjectID, updates bson.M) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("exerciseHistory")

	opts := options.Update().SetUpsert(true)
	_, err = collection.UpdateOne(ctx,
		bson.M{
			"exerciseID": exerciseID,
			"userID":     userID,
		},
		updates,
		opts,
	)

	return
}
