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

func CreateExercise(exercise models.Exercise) (exerciseID primitive.ObjectID, err error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    collection := GetCollection("exercises")

    result, err := collection.InsertOne(ctx, exercise)
    if err != nil {
        log.Println("Couldn't create exercise")
        return
    }

    exerciseID = result.InsertedID.(primitive.ObjectID)
    return
}

func CreateRoutine(routine models.FullRoutine) (routineID primitive.ObjectID, err error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    collection := GetCollection("routines")

    result, err := collection.InsertOne(ctx, routine)
    if err != nil {
        log.Println("Couldn't create routine")
        return
    }

    routineID = result.InsertedID.(primitive.ObjectID)
    return
}

func CreateWorkout(workout models.FullWorkout) (workoutID primitive.ObjectID, err error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    collection := GetCollection("workouts")

    result, err := collection.InsertOne(ctx, workout)
    if err != nil {
        log.Println("Couldn't create workout")
        return
    }

    workoutID = result.InsertedID.(primitive.ObjectID)
    return
}

func CreateSession(session models.WorkoutSession) (sessionID primitive.ObjectID, err error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    collection := GetCollection("sessions")

    result, err := collection.InsertOne(ctx, session)
    if err != nil {
        log.Println("Couldn't create workout")
        return
    }

    sessionID = result.InsertedID.(primitive.ObjectID)
    return
}

func CreateHistory(history models.ExerciseHistory) (historyID primitive.ObjectID, err error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    collection := GetCollection("history")

    result, err := collection.InsertOne(ctx, history)
    if err != nil {
        log.Println("Couldn't create historic data")
        return
    }

    historyID = result.InsertedID.(primitive.ObjectID)
    return
}