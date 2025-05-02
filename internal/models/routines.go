package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Routine struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID      primitive.ObjectID `bson:"userId" json:"user_id"`
    Name        string             `bson:"name" json:"name"`
    Description string             `bson:"description,omitempty" json:"description,omitempty"`
    CreatedAt   primitive.DateTime `bson:"createdAt" json:"createdAt"`
    UpdatedAt   primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}

type RoutineExercise struct {
    ExerciseID  primitive.ObjectID `bson:"exerciseId" json:"exercise_id"`
    TargetSets  int                `bson:"targetSets" json:"target_sets"`
    TargetReps  []int              `bson:"targetReps" json:"target_reps"`
}

type FullRoutine struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID      primitive.ObjectID `bson:"userId" json:"user_id"`
    Name        string             `bson:"name" json:"name"`
    Description string             `bson:"description,omitempty" json:"description,omitempty"`
    Exercises   []RoutineExercise  `bson:"exercises" json:"exercises"`
    CreatedAt   primitive.DateTime `bson:"createdAt" json:"created_at"`
    UpdatedAt   primitive.DateTime `bson:"updatedAt" json:"updated_at"`
}
