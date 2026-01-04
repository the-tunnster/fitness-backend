package database

import (
	"context"
	"log"
	"time"

	"fitness-tracker/internal/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDatabase *mongo.Database

func InitMongo() {
	cfg := config.AppConfig

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}

	// Verify the database is reachable
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	MongoClient = client
	MongoDatabase = client.Database(cfg.MongoDB)

	if err := InitIndexes(ctx, MongoDatabase); err != nil {
		log.Fatalf("Failed to initialize MongoDB indexes: %v", err)
	}

	// Resolve Warm-Up and Cool-Down first; seed or upsert if missing
	warmupHex, warmErr := GetExerciseID("Warm-Up")
	cooldownHex, coolErr := GetExerciseID("Cool-Down")

	if warmErr == nil && warmupHex != "" && coolErr == nil && cooldownHex != "" {
		// Both exist; set config and return
		config.AppConfig.StaticExercises.WarmupID = warmupHex
		config.AppConfig.StaticExercises.CooldownID = cooldownHex
		return
	}

	exercises := MongoDatabase.Collection("exercises")
	count, cntErr := exercises.CountDocuments(ctx, bson.M{})
	if cntErr != nil {
		log.Printf("Warning: failed to count exercises: %v", cntErr)
	}

	if count == 0 {
		// Fresh deployment: seed from JSON and ensure placeholders
		if err := EnsureExercisesSeeded(ctx, MongoDatabase); err != nil {
			log.Printf("Warning: seeding exercises from JSON failed: %v", err)
		}
	} else {
		// Existing DB missing one or both: ensure placeholders only
		if err := ensureWarmupCooldownDefaults(ctx, exercises); err != nil {
			log.Printf("Warning: ensuring Warm-Up/Cool-Down defaults failed: %v", err)
		}
	}

	// Re-resolve and cache IDs
	if warmupHex, err := GetExerciseID("Warm-Up"); err == nil && warmupHex != "" {
		config.AppConfig.StaticExercises.WarmupID = warmupHex
	}
	if cooldownHex, err := GetExerciseID("Cool-Down"); err == nil && cooldownHex != "" {
		config.AppConfig.StaticExercises.CooldownID = cooldownHex
	}
}

func GetCollection(name string) *mongo.Collection {
	return MongoDatabase.Collection(name)
}
