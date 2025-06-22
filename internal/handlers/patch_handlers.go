package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"fitness-tracker/internal/database"
	"fitness-tracker/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	var updates bson.M
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	updates["updatedAt"] = time.Now()

	err = database.UpdateUser(userObjID, updates)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateExerciseHandler(w http.ResponseWriter, r *http.Request) {
	exerciseID := r.URL.Query().Get("exercise_id")

	if exerciseID == "" {
		http.Error(w, "Missing user_id or routine_id", http.StatusBadRequest)
		return
	}

	exerciseObjID, err := primitive.ObjectIDFromHex(exerciseID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updates bson.M
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	updates["updatedAt"] = time.Now()

	err = database.UpdateExercise(exerciseObjID, updates)
	if err != nil {
		http.Error(w, "Failed to update exercise", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateRoutineHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	routineID := r.URL.Query().Get("routine_id")

	if userID == "" || routineID == "" {
		http.Error(w, "Missing user_id or routine_id", http.StatusBadRequest)
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	routineObjID, err2 := primitive.ObjectIDFromHex(routineID)
	if err != nil || err2 != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updates bson.M
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Convert exercises (if present)
	if rawExercises, ok := updates["exercises"]; ok {
		exList, ok := rawExercises.([]interface{})
		if !ok {
			http.Error(w, "Invalid format for exercises", http.StatusBadRequest)
			return
		}

		var parsedExercises []models.RoutineExercise

		for _, ex := range exList {
			exMap, ok := ex.(map[string]interface{})
			if !ok {
				http.Error(w, "Invalid exercise object", http.StatusBadRequest)
				return
			}

			exIDStr, ok := exMap["exercise_id"].(string)
			if !ok {
				http.Error(w, "exercise_id must be a string", http.StatusBadRequest)
				return
			}

			exID, err := primitive.ObjectIDFromHex(exIDStr)
			if err != nil {
				http.Error(w, "Invalid exercise_id", http.StatusBadRequest)
				return
			}

			name, _ := exMap["name"].(string)
			targetSets, _ := exMap["target_sets"].(float64)
			rawReps, _ := exMap["target_reps"].([]interface{})

			var targetReps []int
			for _, r := range rawReps {
				if repFloat, ok := r.(float64); ok {
					targetReps = append(targetReps, int(repFloat))
				}
			}

			parsedExercises = append(parsedExercises, models.RoutineExercise{
				ExerciseID: exID,
				Name:       name,
				TargetSets: int(targetSets),
				TargetReps: targetReps,
			})
		}

		updates["exercises"] = parsedExercises
	}

	updates["updatedAt"] = primitive.NewDateTimeFromTime(time.Now())

	err = database.UpdateRoutine(routineObjID, userObjID, updates)
	if err != nil {
		http.Error(w, "Failed to update routine", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	exerciseIndexStr := r.URL.Query().Get("exercise_index")

	if sessionID == "" || "exercise_index" == "" {
		http.Error(w, "Missing session_id or exercise_index", http.StatusBadRequest)
		return
	}

	sessionObjID, err := primitive.ObjectIDFromHex(sessionID)
	if err != nil {
		http.Error(w, "Invalid session_id format", http.StatusBadRequest)
		return
	}

	exIndex, err := strconv.Atoi(exerciseIndexStr)
	if err != nil || exIndex < 0 {
		http.Error(w, "Invalid exercise_index", http.StatusBadRequest)
		return
	}

	var dto models.WorkoutExerciseDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	exerciseObjID, err := primitive.ObjectIDFromHex(dto.ExerciseID)
	if err != nil {
		http.Error(w, "Invalid exercise_id", http.StatusBadRequest)
		return
	}

	var sets []models.WorkoutSet
	for _, s := range dto.Sets {
		sets = append(sets, models.WorkoutSet(s))
	}

	updatedExercise := models.WorkoutExercise{
	ExerciseID: exerciseObjID,
	Equipment:  dto.Equipment,
	Variation:  dto.Variation,
	Sets:       sets,
}

	updates := bson.M{
		fmt.Sprintf("exercises.%d", exIndex): updatedExercise,
		"last_update":                         time.Now(),
	}

	err = database.UpdateSession(sessionObjID, updates)
	if err != nil {
		http.Error(w, "Failed to update session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}