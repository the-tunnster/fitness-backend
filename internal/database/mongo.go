package database

import (
    "log"
    "time"
    "context"

    "fitness-tracker/internal/config"

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

    MongoClient = client
    MongoDatabase = client.Database(cfg.MongoDB)

    //if err := InitIndexes(ctx, MongoDatabase); err != nil {
    //   log.Fatalf("Failed to initialize MongoDB indexes: %v", err)
    //}
}

func GetCollection(name string) *mongo.Collection {
    return MongoDatabase.Collection(name)
}