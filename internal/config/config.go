package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	MongoDB  string

	Port              string
	StaticExercises   StaticExercises
	ExercisesJSONPath string
}

var AppConfig Config

type StaticExercises struct {
	WarmupID   string
	CooldownID string
}

func LoadConfig() {
	log.Println("Loading .env file...")
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found — using existing env vars")
	}

	uri := getEnvWithDefault("MONGODB_URI", "mongodb://root:example@localhost:27017/?authSource=admin")
	db := getEnvWithDefault("MONGODB_DBNAME", "fitness_app")
	port := getEnvWithDefault("PORT", "8080")

	exercisesJSON := getEnvWithDefault("EXERCISES_JSON_PATH", "exercises.json")

	// Do not use hardcoded defaults; rely on DB resolution at startup.
	// If provided explicitly via env, they will be used initially and overridden by DB lookups.
	warmupID := getEnvWithDefault("WARMUP_ID", "")
	cooldownID := getEnvWithDefault("COOLDOWN_ID", "")

	if uri == "" || db == "" {
		log.Fatal("Missing environment variables: MONGODB_URI and/or MONGODB_DBNAME")
	}

	AppConfig = Config{
		MongoURI: uri,
		MongoDB:  db,
		Port:     port,

		ExercisesJSONPath: exercisesJSON,

		StaticExercises: StaticExercises{
			WarmupID:   warmupID,
			CooldownID: cooldownID,
		},
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
