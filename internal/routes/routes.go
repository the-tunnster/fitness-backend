package routes

import (
	"net/http"

	"fitness-tracker/internal/handlers"
	"fitness-tracker/internal/middleware"
)

func RegisterRoutes(mux *http.ServeMux) {

	// USER
	mux.Handle("/user/id", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetUserByIDHandler)))
	mux.Handle("/user/all", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetAllUsersHandler)))
	mux.Handle("/user/create", middleware.AllowMethods([]string{"POST"}, http.HandlerFunc(handlers.CreateUserHandler)))
	mux.Handle("/user/update", middleware.RequireUser(middleware.AllowMethods([]string{"PATCH"}, http.HandlerFunc(handlers.UpdateUserHandler))))

	// OVERSEER removed

	// EXERCISE
	mux.Handle("/exercise/id", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetExerciseIDHandler)))
	mux.Handle("/exercise/name", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetExerciseNameHandler)))
	mux.Handle("/exercise/list", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetExerciseListHandler)))
	mux.Handle("/exercise/data", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetExerciseDataHandler)))
	mux.Handle("/exercise/create", middleware.AllowMethods([]string{"POST"}, http.HandlerFunc(handlers.CreateExerciseHandler)))
	mux.Handle("/exercise/update", middleware.AllowMethods([]string{"PATCH"}, http.HandlerFunc(handlers.UpdateExerciseHandler)))

	// ROUTINE
	mux.Handle("/routines/list", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetRoutineListHandler)))
	mux.Handle("/routines/data", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetRoutineDataHandler)))
	mux.Handle("/routines/create", middleware.AllowMethods([]string{"POST"}, http.HandlerFunc(handlers.CreateRoutineHandler)))
	mux.Handle("/routines/update", middleware.AllowMethods([]string{"PATCH"}, http.HandlerFunc(handlers.UpdateRoutineHandler)))
	mux.Handle("/routines/delete", middleware.AllowMethods([]string{"DELETE"}, http.HandlerFunc(handlers.DeleteRoutineHandler)))

	// WORKOUT
	mux.Handle("/workouts/list", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetWorkoutListHandler)))
	mux.Handle("/workouts/data", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetWorkoutDataHandler)))
	mux.Handle("/workouts/count", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.CountWorkoutHandler)))
	mux.Handle("/workouts/create", middleware.RequireUser(middleware.AllowMethods([]string{"POST"}, http.HandlerFunc(handlers.CreateWorkoutHandler))))
	mux.Handle("/workouts/comparison", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetWorkoutComparisonHandler)))

	// SESSION
	mux.Handle("/session/data", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetSessionHandler)))
	mux.Handle("/session/create", middleware.RequireUser(middleware.AllowMethods([]string{"POST"}, http.HandlerFunc(handlers.CreateSessionHandler))))
	mux.Handle("/session/update", middleware.AllowMethods([]string{"PATCH"}, http.HandlerFunc(handlers.UpdateSessionHandler)))
	mux.Handle("/session/delete", middleware.AllowMethods([]string{"DELETE"}, http.HandlerFunc(handlers.DeleteSessionHandler)))

	// HISTORY
	mux.Handle("/history/create", middleware.AllowMethods([]string{"POST"}, http.HandlerFunc(handlers.CreateExerciseHistoryHandler)))
	mux.Handle("/history/data", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetExerciseHistoryHandler)))
	mux.Handle("/history/update", middleware.AllowMethods([]string{"PATCH"}, http.HandlerFunc(handlers.UpdateExerciseHistoryHandler)))

	// CARDIO removed

	// AUTH
	mux.Handle("/me", middleware.AllowMethods([]string{"GET"}, http.HandlerFunc(handlers.GetUserHandler)))
	// overseer auth route removed
}
