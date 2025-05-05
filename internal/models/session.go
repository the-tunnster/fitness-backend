package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	UserID          primitive.ObjectID `bson:"userId"`
	ActiveRoutineID primitive.ObjectID `bson:"activeRoutineId,omitempty"`
	WorkoutDraft    WorkoutDraft       `bson:"workoutDraft,omitempty"`
	CurrentStep     WorkoutStep        `bson:"currentStep,omitempty"`
	Stage           string             `bson:"stage"`
	LastUpdated     time.Time          `bson:"lastUpdated"`
}

type WorkoutDraft struct {
	Exercises []ExerciseProgress `bson:"exercises"`
	Notes     string             `bson:"notes,omitempty"`
	StartTime time.Time          `bson:"startTime"`
}

type ExerciseProgress struct {
	ExerciseID primitive.ObjectID `bson:"exerciseId"`
	Sets       []SetProgress      `bson:"sets"`
}

type SetProgress struct {
	SetNumber int  `bson:"setNumber"`
	Done      bool `bson:"done"`
}

type WorkoutStep struct {
	ExerciseIndex int `bson:"exerciseIndex"`
	SetIndex      int `bson:"setIndex"`
}
