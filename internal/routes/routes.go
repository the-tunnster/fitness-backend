package routes

import (
    "net/http"

    "fitness-tracker/internal/handlers"
    "fitness-tracker/internal/middleware"
)

func RegisterRoutes(mux *http.ServeMux) {

    // PUBLIC
    mux.HandleFunc("/login", handlers.HandleLogin)

    // USER
    mux.HandleFunc("/user/create", handlers.CreateUserHandler)
    mux.Handle("/user/update", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.UpdateUserHandler)))
    mux.Handle("/user/id", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetUserHandler)))

    // EXERCISE
    mux.Handle("/exercise/create", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.CreateExerciseHandler)))
    mux.Handle("/exercise/update", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.UpdateExerciseHandler)))

    // ROUTINE
    mux.Handle("/routines/create", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.CreateRoutineHandler)))
    mux.Handle("/routines/list", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetRoutineListHandler)))
    mux.Handle("/routines/data", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetRoutineDataHandler)))
    mux.Handle("/routines/delete", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.DeleteRoutineHandler)))
    mux.Handle("/routines/update", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.UpdateRoutineHandler)))

    // WORKOUT
    mux.Handle("/workouts/create", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.CreateWorkoutHandler)))
    mux.Handle("/workouts/list", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetWorkoutListHandler)))
    mux.Handle("/workouts/data", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetWorkoutDataHandler)))

    // SESSION
    mux.Handle("/session/create", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.CreateSessionHandler)))
    mux.Handle("/session/data", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetSessionHandler)))

    // AUTH
    mux.Handle("/me", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.HandleMe)))
}
