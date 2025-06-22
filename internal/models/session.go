package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkoutSession struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	UserID       primitive.ObjectID   `bson:"user_id" json:"user_id"`
	RoutineID    primitive.ObjectID  `bson:"routine_id,omitempty" json:"routine_id,omitempty"`
	Exercises    []WorkoutExercise    `bson:"exercises" json:"exercises"`
	ExerciseIndex int                 `bson:"exercise_index" json:"exercise_index"`
	LastUpdate   time.Time            `bson:"last_update" json:"last_update"`
}
