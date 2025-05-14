package main

import (
	"log"
	"net/http"

	"fitness-tracker/internal/config"
	"fitness-tracker/internal/database"
	"fitness-tracker/internal/login"
	"fitness-tracker/internal/routes"
)

func main() {
	log.Println("Loding app config...")
	config.LoadConfig()

	log.Println("Initialising OIDC...")
	login.InitOIDC()

	log.Println("Initialising database connection...")
	database.InitMongo()

	log.Println("Registering routes and multiplexer...")
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	log.Printf("Server running on :%s", config.AppConfig.Port)
	log.Fatal(http.ListenAndServe(":"+config.AppConfig.Port, mux))
}
