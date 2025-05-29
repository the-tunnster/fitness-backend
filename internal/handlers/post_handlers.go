package handlers

import (
	"encoding/json"
	"net/http"

	"fitness-tracker/internal/database"
	"fitness-tracker/internal/models"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userID, err := database.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userID.Hex())
}

func CreateExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var exercise models.Exercise

	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	exerciseID, err := database.CreateExercise(exercise)
	if err != nil {
		http.Error(w, "Failed to create exercise", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exerciseID.Hex())
}

func CreateRoutineHandler(w http.ResponseWriter, r *http.Request) {
	var routine models.FullRoutine

	if err := json.NewDecoder(r.Body).Decode(&routine); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	routineID, err := database.CreateRoutine(routine)
	if err != nil {
		http.Error(w, "Failed to create routine", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(routineID.Hex())
}

func CreateWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	var workout models.FullWorkout

	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	workoutID, err := database.CreateWorkout(workout)
	if err != nil {
		http.Error(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(workoutID.Hex())
}

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	var session models.Session

	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	sessionID, err := database.CreateSession(session)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sessionID.Hex())
}