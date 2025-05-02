package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username       string             `bson:"username" json:"username"`
	Email          string             `bson:"email" json:"email"`
	Gender         string             `bson:"gender" json:"gender"`
	DateOfBirth    string             `bson:"dateOfBirth" json:"dateOfBirth"`
	Height         float64            `bson:"height" json:"height"`
	Weight         float64            `bson:"weight" json:"weight"`
	UnitPreference string             `bson:"unitPreference" json:"unitPreference"`
}