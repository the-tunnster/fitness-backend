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

func GetOverseerByEmail(emailID string) (overseer models.Overseer, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("overseers")

	err = collection.FindOne(ctx, bson.M{"email": emailID}).Decode(&overseer)
	if err != nil {
		return models.Overseer{}, err
	}

	return
}

func GetOverseerByID(overseerID primitive.ObjectID) (overseer models.Overseer, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("overseers")

	err = collection.FindOne(ctx, bson.M{"_id": overseerID}).Decode(&overseer)
	if err != nil {
		return models.Overseer{}, err
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

func GetLastWorkoutForRoutine(userID, routineID primitive.ObjectID) (last_workout models.FullWorkout, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("workouts")

	opts := options.FindOne().SetSort(bson.D{{Key: "workoutDate", Value: -1}})
	err = collection.FindOne(ctx, bson.M{
		"userID":    userID,
		"routineID": routineID,
	}, opts).Decode(&last_workout)

	if err != nil {
		return last_workout, err
	}

	return
}

func CountWorkouts(userID primitive.ObjectID) (workout_count int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("workouts")

	workout_count, err = collection.CountDocuments(ctx, bson.M{
		"userID": userID,
	})

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

func GetUserSessionData(userID primitive.ObjectID) (session models.WorkoutSession, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("sessions")

	err = collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&session)

	return
}

func GetSessionData(sessionID primitive.ObjectID) (session models.WorkoutSession, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("sessions")

	err = collection.FindOne(ctx, bson.M{"_id": sessionID}).Decode(&session)

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

func GetCardioList() (cardios []models.Cardio) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("cardio")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil
	}

	for cursor.Next(ctx) {
		var cardio models.Cardio
		if err := cursor.Decode(&cardio); err != nil {
			continue
		}
		cardios = append(cardios, cardio)
	}

	return
}

func GetExerciseID(exerciseName string) (exerciseID string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("exercises")

	var exercise models.Exercise

	err = collection.FindOne(ctx, bson.M{
		"name": exerciseName,
	}).Decode(&exercise)

	exerciseID = exercise.ID.Hex()

	return
}

func GetExerciseName(exerciseID primitive.ObjectID) (exerciseName string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("exercises")

	var exercise models.Exercise

	err = collection.FindOne(ctx, bson.M{
		"_id": exerciseID,
	}).Decode(&exercise)

	exerciseName = exercise.Name

	return
}

func GetExerciseData(exerciseID primitive.ObjectID) (exercise models.Exercise, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("exercises")

	err = collection.FindOne(ctx, bson.M{
		"_id": exerciseID,
	}).Decode(&exercise)

	return
}

func GetExerciseHistoryData(exerciseID primitive.ObjectID, userID primitive.ObjectID) (history models.ExerciseHistory, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("exerciseHistory")

	err = collection.FindOne(ctx, bson.M{
		"exerciseID": exerciseID,
		"userID":     userID,
	}).Decode(&history)

	return
}

func GetCardioHistoryData(cardioID primitive.ObjectID, userID primitive.ObjectID) (history models.CardioHistory, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("cardioHistory")

	err = collection.FindOne(ctx, bson.M{
		"cardioID": cardioID,
		"userID":   userID,
	}).Decode(&history)

	return
}

func GetLastTwoWorkouts(userID, routineID primitive.ObjectID) (workouts []models.FullWorkout, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection("workouts")
	opts := options.Find().SetSort(bson.D{{Key: "workoutDate", Value: -1}}).SetLimit(2)

	cursor, err := collection.Find(ctx, bson.M{
		"userID":    userID,
		"routineID": routineID,
	}, opts)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var workout models.FullWorkout
		err = cursor.Decode(&workout)
		if err != nil {
			return
		}
		workouts = append(workouts, workout)
	}

	return
}
