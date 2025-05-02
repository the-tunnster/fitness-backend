package routes

import (
    "net/http"

    "fitness-tracker/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {

    mux.HandleFunc("/user/id", handlers.GetUserHandler)
    mux.HandleFunc("/routines/list", handlers.GetRoutineListHandler)
    mux.HandleFunc("/routines/data", handlers.GetRoutineDataHandler)
    mux.HandleFunc("/workouts/list", handlers.GetWorkoutListHandler)
    mux.HandleFunc("/workouts/data", handlers.GetWorkoutDataHandler)

}
