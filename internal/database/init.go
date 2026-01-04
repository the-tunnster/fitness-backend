package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitIndexes(ctx context.Context, db *mongo.Database) error {
	if err := initUserIndexes(ctx, db); err != nil {
		return err
	}
	if err := initRoutineIndexes(ctx, db); err != nil {
		return err
	}
	if err := initExerciseIndexes(ctx, db); err != nil {
		return err
	}
	if err := initSessionIndexes(ctx, db); err != nil {
		return err
	}
	if err := initWorkoutIndexes(ctx, db); err != nil {
		return err
	}
	if err := initExerciseHistoryIndexes(ctx, db); err != nil {
		return err
	}
	return nil
}

func initUserIndexes(ctx context.Context, db *mongo.Database) error {
	users := db.Collection("users")
	_, err := users.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

func initRoutineIndexes(ctx context.Context, db *mongo.Database) error {
	routines := db.Collection("routines")
	_, err := routines.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "userID", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "userID", Value: 1}, {Key: "name", Value: 1}},
		},
	})
	return err
}

func initExerciseIndexes(ctx context.Context, db *mongo.Database) error {
	exercises := db.Collection("exercises")

	// Create unique index on name (idempotent if it already exists uniquely)
	_, err := exercises.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

// Cardio feature removed

func initSessionIndexes(ctx context.Context, db *mongo.Database) error {
	sessions := db.Collection("sessions")
	_, err := sessions.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "userID", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "lastUpdated", Value: 1}},
			Options: options.Index().SetExpireAfterSeconds(60 * 60 * 48), // TTL: 48 hours
		},
	})
	return err
}

func initWorkoutIndexes(ctx context.Context, db *mongo.Database) error {
	workouts := db.Collection("workouts")
	_, err := workouts.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "userID", Value: 1}, {Key: "workoutDate", Value: -1}},
		},
		{
			Keys: bson.D{{Key: "routineID", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "userID", Value: 1}, {Key: "routineID", Value: 1}, {Key: "workoutDate", Value: -1}},
		},
	})
	return err
}

func initExerciseHistoryIndexes(ctx context.Context, db *mongo.Database) error {
	exerciseHistory := db.Collection("exerciseHistory")
	_, err := exerciseHistory.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "userID", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "exerciseID", Value: 1}},
		},
	})
	return err
}

// Cardio feature removed
