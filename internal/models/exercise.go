package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Category   string             `bson:"category" json:"category"`
	Variations []string           `bson:"variations" json:"variations"`
	Equipment  []string           `bson:"equipment" json:"equipment"`
}

type ExerciseSets struct {
	Date        primitive.DateTime `bson:"date" json:"date"`
	Equipment   string             `bson:"equipment" json:"equipment"`
	Variation   string             `bson:"variation" json:"variation"`
	WorkoutSets []WorkoutSet       `bson:"sets" json:"sets"`
}

type ExerciseHistory struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"userID" json:"user_id"`
	ExerciseID primitive.ObjectID `bson:"exerciseID" json:"exercise_id"`
	Sets       []ExerciseSets     `bson:"exerciseSets" json:"exercise_sets"`
}
