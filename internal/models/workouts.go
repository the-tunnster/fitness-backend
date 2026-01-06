package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Workout struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"userID" json:"user_id"`
	RoutineID   primitive.ObjectID `bson:"routineID" json:"routine_id"`
	WorkoutDate primitive.DateTime `bson:"workoutDate" json:"workout_date"`
}

type WorkoutSet struct {
	Reps   int     `bson:"reps" json:"reps"`
	Weight float64 `bson:"weight" json:"weight"`
}

type WorkoutExercise struct {
	ExerciseID primitive.ObjectID `bson:"exerciseID" json:"exercise_id"`
	Equipment  string             `bson:"equipment" json:"equipment"`
	Variation  string             `bson:"variation" json:"variation"`
	Sets       []WorkoutSet       `bson:"sets" json:"sets"`
	Name       string             `bson:"name" json:"name"`
}

type WorkoutExerciseDTO struct {
	ExerciseID string       `json:"exercise_id"`
	Equipment  string       `json:"equipment"`
	Variation  string       `json:"variation"`
	Sets       []WorkoutSet `json:"sets"`
	Name       string       `json:"name"`
}

type FullWorkout struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"userID" json:"user_id"`
	RoutineID   primitive.ObjectID `bson:"routineID" json:"routine_id"`
	WorkoutDate primitive.DateTime `bson:"workoutDate" json:"workout_date"`
	Exercises   []WorkoutExercise  `bson:"exercises" json:"exercises"`
}
