package handlers

import (
	"encoding/json"
	"net/http"

	"fitness-tracker/internal/models"
	"fitness-tracker/internal/database"
	"fitness-tracker/internal/middleware"
	"fitness-tracker/internal/login"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := database.GetUserByUsername(req.Username)
	if err != nil || !database.ValidatePassword(user.HashedPassword, req.Password) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := login.GenerateJWT(user.ID.Hex())
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func HandleMe(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
        http.Error(w, "Invalid user_id", http.StatusBadRequest)
        return
    }

	user, err := database.GetUserByID(userObjID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}