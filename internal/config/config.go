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

	OIDCClientID     string
	OIDCClientSecret string
	OIDCRedirectURL  string
}

var AppConfig Config

func LoadConfig() {
	log.Println("Loading .env file...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found — using existing env vars")
	}

	uri := getEnvWithDefault("MONGODB_URI", "mongodb://root:example@localhost:27017/?authSource=admin")
	db := getEnvWithDefault("MONGODB_DBNAME", "fitness_app")
	port := getEnvWithDefault("PORT", "8080")

	oidcClientID := getEnvWithDefault("OIDC_CLIENT_ID", "")
	oidcClientSecret := getEnvWithDefault("OIDC_CLIENT_SECRET", "")
	oidcRedirectURL := getEnvWithDefault("OIDC_REDIRECT_URL", "")

	if uri == "" || db == "" {
		log.Fatal("Missing environment variables: MONGODB_URI and/or MONGODB_DBNAME")
	}

	AppConfig = Config{
		MongoURI: uri,
		MongoDB:  db,
		Port:     port,
		OIDCClientID: oidcClientID,
		OIDCClientSecret: oidcClientSecret,
		OIDCRedirectURL: oidcRedirectURL,
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
