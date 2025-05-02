package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	MongoDB  string

	Port string
}

func LoadConfig() *Config {
	log.Println("Loading .env file...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found — using existing env vars")
	}

	uri := getEnvWithDefault("MONGODB_URI", "mongodb://root:example@localhost:27017/?authSource=admin")
	db := getEnvWithDefault("MONGODB_DBNAME", "fitness_app")
	port := getEnvWithDefault("PORT", "8080")

	if uri == "" || db == "" {
		log.Fatal("Missing environment variables: MONGODB_URI and/or MONGODB_DBNAME")
	}

	return &Config{
		MongoURI: uri,
		MongoDB:  db,
		Port:     port,
	}
}

func getEnvWithDefault(key, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		log.Println("Using default value for ", key)
		return defaultValue
	}
	return
}
