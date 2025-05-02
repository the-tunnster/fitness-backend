package main

import (
    "log"
    "net/http"
    
    "fitness-tracker/internal/config"
    "fitness-tracker/internal/database"
    "fitness-tracker/internal/routes"
)

func main() {
    log.Println("Loding app config...")
    cfg := config.LoadConfig()

    log.Println("Initialising database connection...")
    database.InitMongo(cfg)

    log.Println("Registering routes and multiplexer...")
    mux := http.NewServeMux()
    routes.RegisterRoutes(mux)

    log.Printf("🚀 Server running on :%s", cfg.Port)
    log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
