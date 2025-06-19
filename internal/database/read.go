package database

import (
	"context"
	"log"
	"time"

	"fitness-tracker/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserByEmail(emailID string) (user models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("users")

	err = collection.FindOne(ctx, bson.M{"email": emailID}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return
}

func GetUserByID(userID primitive.ObjectID) (user models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("users")

	err = collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return
}

func GetUserRoutines(userID primitive.ObjectID) (routineList []models.Routine, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("routines")

	cursor, err := collection.Find(ctx, bson.M{"userID": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var routine models.Routine
		if err := cursor.Decode(&routine); err != nil {
			continue
		}
		routineList = append(routineList, routine)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error decoding some documents")
	}

	return
}

func GetRoutineData(userID, routineID primitive.ObjectID) (routine models.FullRoutine, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("routines")

	err = collection.FindOne(ctx, bson.M{
		"_id":    routineID,
		"userID": userID,
	}).Decode(&routine)

	return
}

func GetUserWorkouts(userID primitive.ObjectID) (workoutList []models.Workout, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("workouts")

	cursor, err := collection.Find(ctx, bson.M{
		"userID": userID,
	})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var workout models.Workout
		if err := cursor.Decode(&workout); err != nil {
			continue
		}
		workoutList = append(workoutList, workout)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error decoding some documents")
		log.Println(err)
	}

	return
}

func GetWorkoutData(userID, workoutID primitive.ObjectID) (workout models.FullWorkout, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("workouts")

	err = collection.FindOne(ctx, bson.M{
		"_id":    workoutID,
		"userID": userID,
	}).Decode(&workout)

	return
}

func GetSessionData(userID primitive.ObjectID) (session models.Session, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("sessions")

	err = collection.FindOne(ctx, bson.M{"userID": userID}).Decode(&session)

	return
}

func GetExerciseList() (exercises []models.Exercise) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("exercises")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil
	}

	for cursor.Next(ctx) {
		var exercise models.Exercise
		if err := cursor.Decode(&exercise); err != nil {
			continue
		}
		exercises = append(exercises, exercise)
	}

	return
}

func GetExerciseID(exerciseName string) (exerciseID string, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("exercises")

	var exercise models.Exercise

	err = collection.FindOne(ctx, bson.M{
		"name":    exerciseName,
	}).Decode(&exercise)

	exerciseID = exercise.ID.Hex()

	return
}

func GetExerciseName(exerciseID primitive.ObjectID) (exerciseName string, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("exercises")

	var exercise models.Exercise

	err = collection.FindOne(ctx, bson.M{
		"_id":    exerciseID,
	}).Decode(&exercise)

	exerciseName = exercise.Name

	return
}