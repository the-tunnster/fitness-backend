package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"fitness-tracker/internal/database"
	"fitness-tracker/internal/models"
	"fitness-tracker/internal/service"
	"fitness-tracker/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	userID, err := database.CreateUser(user)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.JSONResponse(w, http.StatusCreated, userID.Hex())
}

func CreateExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var exercise models.Exercise

	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	exerciseID, err := database.CreateExercise(exercise)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create exercise")
		return
	}

	utils.JSONResponse(w, http.StatusCreated, exerciseID.Hex())
}

func CreateRoutineHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	var routine_data models.FullRoutineDTO
	if err := json.NewDecoder(r.Body).Decode(&routine_data); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	var exercises []models.RoutineExercise
	for _, exercise := range routine_data.Exercises {
		exerciseObjID, err := primitive.ObjectIDFromHex(exercise.ExerciseID)
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, "Invalid exercise_id")
			return
		}

		exercises = append(exercises, models.RoutineExercise{
			ExerciseID: exerciseObjID,
			Name:       exercise.Name,
			TargetSets: exercise.TargetSets,
			TargetReps: exercise.TargetReps,
		})
	}

	newRoutine := models.FullRoutine{
		UserID:    userObjID,
		Name:      routine_data.Name,
		Exercises: exercises,
	}

	routineID, err := database.CreateRoutine(newRoutine)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create routine")
		return
	}

	utils.JSONResponse(w, http.StatusCreated, routineID.Hex())
}

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	routineID := r.URL.Query().Get("routine_id")

	if userID == "" || routineID == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Missing user_id or routine_id")
		return
	}

	// Delegate session generation to the service layer

	session, err := service.GenerateWorkoutSession(userID, routineID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Failed to generate session")
		return
	}

	sessionID, err := database.UpsertSession(*session)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create or update session")
		return
	}

	log.Print("successfully upserted workout session", sessionID)

	utils.JSONResponse(w, http.StatusCreated, sessionID.Hex())
}

func CreateWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	sessionObjID, err := primitive.ObjectIDFromHex(sessionID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid session_id")
		return
	}

	workout_session, err := database.GetSessionData(sessionObjID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to find session")
		return
	}

	workout := models.FullWorkout{
		ID:          primitive.NilObjectID,
		UserID:      workout_session.UserID,
		RoutineID:   workout_session.RoutineID,
		WorkoutDate: primitive.NewDateTimeFromTime(time.Now()),
		Exercises:   workout_session.Exercises,
	}

	workoutID, err := database.CreateWorkout(workout)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create workout")
		return
	}

	utils.JSONResponse(w, http.StatusCreated, workoutID.Hex())
}

func CreateExerciseHistoryHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	exerciseID := r.URL.Query().Get("exercise_id")

	if userID == "" || exerciseID == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Missing user_id or exercise_id")
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	exerciseObjID, err := primitive.ObjectIDFromHex(exerciseID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid exercise_id")
		return
	}

	exerciseHistory := models.ExerciseHistory{
		ExerciseID: exerciseObjID,
		UserID:     userObjID,
		Sets:       []models.ExerciseSets{},
	}

	historyID, err := database.CreateExerciseHistory(exerciseHistory)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create exercise history: "+err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, historyID.Hex())
}
