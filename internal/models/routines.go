package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoutineExercise struct {
	ExerciseID primitive.ObjectID `bson:"exercise_id" json:"exercise_id"`
	Name       string             `bson:"name" json:"name"`
	TargetSets int                `bson:"target_sets" json:"target_sets"`
	TargetReps int                `bson:"target_reps" json:"target_reps"`
}

type RoutineExerciseDTO struct {
	ExerciseID string `json:"exercise_id"`
	Name       string `json:"name"`
	TargetSets int    `json:"target_sets"`
	TargetReps int    `json:"target_reps"`
}

type FullRoutine struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"userID" json:"user_id"`
	Name      string             `bson:"name" json:"name"`
	Exercises []RoutineExercise  `bson:"exercises" json:"exercises"`
}

type FullRoutineDTO struct {
	Name      string               `json:"name"`
	Exercises []RoutineExerciseDTO `json:"exercises"`
}

type Routine struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID primitive.ObjectID `bson:"userID" json:"user_id"`
	Name   string             `bson:"name" json:"name"`
}
