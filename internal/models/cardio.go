package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cardio struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Variations []string           `bson:"variations" json:"variations"`
	Equipment  []string           `bson:"equipment" json:"equipment"`
}

type CardioMetric struct {
	TotalDistance  float32   `bson:"totalDistance" json:"total_distance"`
	TotalTime      float32   `bson:"totalTime" json:"total_time"`
	DistanceSplits []float32 `bson:"distanceSplits" json:"distance_splits"`
	TimeSplits     []float32 `bson:"timeSplits" json:"time_splits"`
	AveragePace    float32   `bson:"averagePace" json:"average_pace"`
	CaloriesBurned float32   `bson:"caloriesBurned" json:"calories_burned"`
	HeartRate      int16     `bson:"heartRate" json:"heart_rate"`
}

type CardioSession struct {
	Date          primitive.DateTime `bson:"date" json:"date"`
	Equipment     string             `bson:"equipment" json:"equipment"`
	Variation     string             `bson:"variation" json:"variation"`
	CardioMetrics CardioMetric       `bson:"metrics" json:"metrics"`
}

type CardioSessionDTO struct {
	Equipment     string       `json:"equipment"`
	Variation     string       `json:"variation"`
	CardioMetrics CardioMetric `json:"metrics"`
}

type CardioHistory struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID   primitive.ObjectID `bson:"userID" json:"user_id"`
	CardioID primitive.ObjectID `bson:"cardioID" json:"cardio_id"`
	Sessions []CardioSession    `bson:"sessions" json:"sessions"`
}
