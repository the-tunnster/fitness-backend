package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `bson:"name" json:"name"`
	Category        string             `bson:"category" json:"category"`
	PrimaryMuscle   string             `bson:"primaryMuscle" json:"primaryMuscle"`
	SecondaryMuscle string             `bson:"secondaryMuscle" json:"secondaryMuscle"`
	TertiaryMuscle  string             `bson:"tertiaryMuscle" json:"tertiaryMuscle"`
	Variations      []string           `bson:"variations" json:"variations"`
	Equipment       []string           `bson:"equipment" json:"equipment"`
	CreatedAt       primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt       primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}
