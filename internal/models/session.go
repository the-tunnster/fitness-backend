package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkoutSession struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID        primitive.ObjectID `bson:"userID" json:"user_id"`
	RoutineID     primitive.ObjectID `bson:"routineID,omitempty" json:"routine_id,omitempty"`
	Exercises     []WorkoutExercise  `bson:"exercises" json:"exercises"`
	ExerciseIndex int                `bson:"exerciseIndex" json:"exercise_index"`
	LastUpdate    primitive.DateTime `bson:"lastUpdated" json:"last_update"`
}
