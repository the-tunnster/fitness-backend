package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Overseer struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Username string               `bson:"username" json:"username"`
	Email    string               `bson:"email" json:"email"`
	Clients  []primitive.ObjectID `bson:"clients" json:"clients"`
}

type OverseerDTO struct {
	ID       string
	Username string
	Email    string
	Clients  []string
}
