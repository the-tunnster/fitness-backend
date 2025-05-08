package routes

import (
    "net/http"

    "fitness-tracker/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {

    // USER Handlers
    mux.HandleFunc("/user/id", handlers.GetUserHandler)
    mux.HandleFunc("/user/create", handlers.CreateUserHandler)
    mux.HandleFunc("/user/update", handlers.UpdateUserHandler)

    // ROUTINE Handlers
    mux.HandleFunc("/routines/list", handlers.GetRoutineListHandler)
    mux.HandleFunc("/routines/data", handlers.GetRoutineDataHandler)
    mux.HandleFunc("/routines/delete", handlers.DeleteRoutineHandler)
    mux.HandleFunc("/routines/update", handlers.UpdateRoutineHandler)

    // WORKOUT Handlers
    mux.HandleFunc("/workouts/list", handlers.GetWorkoutListHandler)
    mux.HandleFunc("/workouts/data", handlers.GetWorkoutDataHandler)

    // SESSION Handlers
    mux.HandleFunc("/session/data", handlers.GetSessionHandler)

    // LOGIN Handlers
    mux.HandleFunc("/login", handlers.HandleLogin)
    mux.HandleFunc("/callback", handlers.HandleCallback)
    
}