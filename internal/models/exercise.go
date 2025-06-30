package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `bson:"name" json:"name"`
	Category        string             `bson:"category" json:"category"`
	PrimaryMuscle   string             `bson:"primaryMuscle" json:"primary_muscle"`
	SecondaryMuscle string             `bson:"secondaryMuscle" json:"secondary_muscle"`
	TertiaryMuscle  string             `bson:"tertiaryMuscle" json:"tertiary_muscle"`
	Variations      []string           `bson:"variations" json:"variations"`
	Equipment       []string           `bson:"equipment" json:"equipment"`
}
