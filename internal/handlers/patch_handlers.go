package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"fitness-tracker/internal/database"

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

	updates["updatedAt"] = time.Now()

	err = database.UpdateRoutine(routineObjID, userObjID, updates)
	if err != nil {
		http.Error(w, "Failed to update routine", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	if sessionID == "" {
		http.Error(w, "Missing session_id", http.StatusBadRequest)
		return
	}

	sessionObjID, err := primitive.ObjectIDFromHex(sessionID)
	if err != nil {
		http.Error(w, "Invalid session_id format", http.StatusBadRequest)
		return
	}

	var updates bson.M
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	updates["updatedAt"] = time.Now()

	err = database.UpdateSession(sessionObjID, updates)
	if err != nil {
		http.Error(w, "Failed to update session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}