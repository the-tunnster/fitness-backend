package handlers

import (
	"net/http"

	"fitness-tracker/internal/database"
	"fitness-tracker/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteRoutineHandler(w http.ResponseWriter, r *http.Request) {
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

	err = database.DeleteRoutine(userObjID, routineObjID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Couldn't delete routine")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")

	if sessionID == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Missing session_id")
		return
	}

	sessionObjID, err := primitive.ObjectIDFromHex(sessionID)

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	err = database.DeleteSession(sessionObjID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Couldn't delete session")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
