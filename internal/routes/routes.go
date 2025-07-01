package routes

import (
    "net/http"

    "fitness-tracker/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {

    // USER
    mux.HandleFunc("/user/create", handlers.CreateUserHandler)
    mux.HandleFunc("/user/update", handlers.UpdateUserHandler)

    // EXERCISE
    mux.HandleFunc("/exercise/id", handlers.GetExerciseIDHandler)
    mux.HandleFunc("/exercise/name", handlers.GetExerciseNameHandler)
    mux.HandleFunc("/exercise/list", handlers.GetExerciseListHandler)
    mux.HandleFunc("/exercise/create", handlers.CreateExerciseHandler)
    mux.HandleFunc("/exercise/update", handlers.UpdateExerciseHandler)

    // ROUTINE
    mux.HandleFunc("/routines/list", handlers.GetRoutineListHandler)
    mux.HandleFunc("/routines/data", handlers.GetRoutineDataHandler)
    mux.HandleFunc("/routines/create", handlers.CreateRoutineHandler)
    mux.HandleFunc("/routines/update", handlers.UpdateRoutineHandler)
    mux.HandleFunc("/routines/delete", handlers.DeleteRoutineHandler)

    // WORKOUT
    mux.HandleFunc("/workouts/list", handlers.GetWorkoutListHandler)
    mux.HandleFunc("/workouts/data", handlers.GetWorkoutDataHandler)
    mux.HandleFunc("/workouts/create", handlers.CreateWorkoutHandler)

    // SESSION
    mux.HandleFunc("/session/data", handlers.GetSessionHandler)
    mux.HandleFunc("/session/create", handlers.CreateWorkoutSessionHandler)
    mux.HandleFunc("/session/update", handlers.UpdateSessionHandler)
    mux.HandleFunc("/session/delete", handlers.DeleteSessionHandler)


    // AUTH
    mux.HandleFunc("/me", handlers.GetUserHandler)
}
