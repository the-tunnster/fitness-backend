package handlers

import (
	"net/http"

    "fitness-tracker/internal/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteRoutineHandler(w http.ResponseWriter, r *http.Request) {
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
	
	err = database.DeleteRoutine(userObjID, routineObjID)
	if err != nil {
		http.Error(w, "Couldn't delete routine", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")

	if sessionID == "" {
		http.Error(w, "Missing session_id", http.StatusBadRequest)
		return
	}

	sessionObjID, err := primitive.ObjectIDFromHex(sessionID)

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	
	err = database.DeleteSession(sessionObjID)
	if err != nil {
		http.Error(w, "Couldn't delete session", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}