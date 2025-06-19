package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Routine struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"userID" json:"user_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	CreatedAt   primitive.DateTime `bson:"createdAt" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt" json:"updated_at"`
}

type RoutineExercise struct {
	ExerciseID primitive.ObjectID `bson:"exercise_id" json:"exercise_id"`
	Name       string             `bson:"name" json:"name"`
	TargetSets int                `bson:"target_sets" json:"target_sets"`
	TargetReps []int              `bson:"target_reps" json:"target_reps"`
}

type FullRoutine struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"userID" json:"user_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Exercises   []RoutineExercise  `bson:"exercises" json:"exercises"`
	CreatedAt   primitive.DateTime `bson:"createdAt" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt" json:"updated_at"`
}
