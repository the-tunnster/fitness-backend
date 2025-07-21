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
    mux.HandleFunc("/exercise/data", handlers.GetExerciseDataHandler)
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
    mux.HandleFunc("/workouts/count", handlers.CountWorkoutHandler)
    mux.HandleFunc("/workouts/create", handlers.CreateWorkoutHandler)
    mux.HandleFunc("/workouts/comparison", handlers.GetWorkoutComparisonHandler)

    // SESSION
    mux.HandleFunc("/session/data", handlers.GetSessionHandler)
    mux.HandleFunc("/session/create", handlers.CreateSessionHandler)
    mux.HandleFunc("/session/update", handlers.UpdateSessionHandler)
    mux.HandleFunc("/session/delete", handlers.DeleteSessionHandler)

    // HISTORY
    mux.HandleFunc("/history/create", handlers.CreateExerciseHistoryHandler)
    mux.HandleFunc("/history/data", handlers.GetExerciseHistoryHandler)
    mux.HandleFunc("/history/update", handlers.UpdateExerciseHistoryHandler)

    // CARDIO
    mux.HandleFunc("/cardio/create", handlers.CreateCardioHistoryHandler)
    mux.HandleFunc("/cardio/update", handlers.UpdateCardioHistoryHandler)
    mux.HandleFunc("/cardio/data", handlers.GetCardioHistoryHandler)
    mux.HandleFunc("/cardio/list", handlers.GetCardioListHandler)
    
    // AUTH
    mux.HandleFunc("/me", handlers.GetUserHandler)
}
