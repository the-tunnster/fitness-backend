package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"fitness-tracker/internal/database"
	"fitness-tracker/internal/models"
	"fitness-tracker/internal/service"
	"fitness-tracker/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Missing user_id")
		return
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	var updates bson.M
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	updates["updatedAt"] = time.Now()

	err = database.UpdateUser(userObjID, updates)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, nil)
}

func UpdateExerciseHandler(w http.ResponseWriter, r *http.Request) {
	exerciseID := r.URL.Query().Get("exercise_id")

	if exerciseID == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Missing user_id or routine_id")
		return
	}

	exerciseObjID, err := primitive.ObjectIDFromHex(exerciseID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var updates bson.M
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	updates["updatedAt"] = time.Now()

	err = database.UpdateExercise(exerciseObjID, updates)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update exercise")
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, nil)
}

func UpdateRoutineHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	routineID := r.URL.Query().Get("routine_id")

	if userID == "" || routineID == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Missing user_id or routine_id")
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	routineObjID, err2 := primitive.ObjectIDFromHex(routineID)
	if err != nil || err2 != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var updated_exercise_data []models.RoutineExerciseDTO
	if err := json.NewDecoder(r.Body).Decode(&updated_exercise_data); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	var updated_exercises []models.RoutineExercise
	for _, exercise := range updated_exercise_data {
		exerciseObjID, err := primitive.ObjectIDFromHex(exercise.ExerciseID)
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, "Invalid exercise_id")
			return
		}

		updated_exercises = append(updated_exercises, models.RoutineExercise{
			ExerciseID: exerciseObjID,
			Name:       exercise.Name,
			TargetSets: exercise.TargetSets,
			TargetReps: exercise.TargetReps,
		})
	}

	updates := bson.M{
		"exercises": updated_exercises,
	}

	err = database.UpdateRoutine(routineObjID, userObjID, updates)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update routine")
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, nil)
}

func UpdateSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	exerciseIndexStr := r.URL.Query().Get("exercise_index")

	if sessionID == "" || exerciseIndexStr == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Missing session_id or exercise_index")
		return
	}

	sessionObjID, err := primitive.ObjectIDFromHex(sessionID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid session_id format")
		return
	}

	exIndex, err := strconv.Atoi(exerciseIndexStr)
	if err != nil || exIndex < 0 {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid exercise_index")
		return
	}

	var updated_exercise_data []models.WorkoutExerciseDTO
	if err := json.NewDecoder(r.Body).Decode(&updated_exercise_data); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	var updated_exercises []models.WorkoutExercise
	for _, exercise := range updated_exercise_data {
		exerciseObjID, err := primitive.ObjectIDFromHex(exercise.ExerciseID)
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, "Invalid exercise_id")
			return
		}
		updated_exercises = append(updated_exercises, models.WorkoutExercise{
			ExerciseID: exerciseObjID,
			Equipment:  exercise.Equipment,
			Variation:  exercise.Variation,
			Sets:       exercise.Sets,
		})
	}

	updates := bson.M{
		"exercises":     updated_exercises,
		"exerciseIndex": exIndex,
		"lastUpdated":   primitive.NewDateTimeFromTime(time.Now()),
	}

	err = database.UpdateSession(sessionObjID, updates)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update session")
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, nil)
}

func UpdateExerciseHistoryHandler(w http.ResponseWriter, r *http.Request) {
	workoutID := r.URL.Query().Get("workout_id")
	userID := r.URL.Query().Get("user_id")

	if workoutID == "" || userID == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Missing workout_id or user_id")
		return
	}

	// Delegate to service layer
	if err := service.UpdateExerciseHistory(userID, workoutID); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update workout history")
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, nil)
}
