package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"fitness-tracker/internal/database"
	"fitness-tracker/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var dto models.FullRoutineDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(dto.UserID)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	var exercises []models.RoutineExercise
	for _, ex := range dto.Exercises {
		exID, err := primitive.ObjectIDFromHex(ex.ExerciseID)
		if err != nil {
			http.Error(w, "Invalid exercise_id", http.StatusBadRequest)
			return
		}

		exercises = append(exercises, models.RoutineExercise{
			ExerciseID: exID,
			Name:       ex.Name,
			TargetSets: ex.TargetSets,
			TargetReps: ex.TargetReps,
		})
	}

	// Build full routine
	newRoutine := models.FullRoutine{
		UserID:      userObjID,
		Name:        dto.Name,
		Description: dto.Description,
		Exercises:   exercises,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}

	// Insert into DB
	routineID, err := database.CreateRoutine(newRoutine)
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
	userID := r.URL.Query().Get("user_id")
	routineID := r.URL.Query().Get("routine_id")

	if userID == "" || routineID == "" {
		http.Error(w, "Missing user_id or routine_id", http.StatusBadRequest)
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	routineObjID, err := primitive.ObjectIDFromHex(routineID)
	if err != nil {
		http.Error(w, "Invalid routine_id", http.StatusBadRequest)
		return
	}

	var exercisesDTO []models.WorkoutExerciseDTO
	if err := json.NewDecoder(r.Body).Decode(&exercisesDTO); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var exercises []models.WorkoutExercise
	for _, dto := range exercisesDTO {
		exID, err := primitive.ObjectIDFromHex(dto.ExerciseID)
		if err != nil {
			http.Error(w, "Invalid exercise_id in exercises", http.StatusBadRequest)
			return
		}

		var sets []models.WorkoutSet
		for _, s := range dto.Sets {
			sets = append(sets, models.WorkoutSet{
				SetNumber: s.SetNumber,
				Reps:      s.Reps,
				Weight:    s.Weight,
			})
		}

		exercises = append(exercises, models.WorkoutExercise{
			ExerciseID: exID,
			Sets:       sets,
		})
	}

	newSession := models.WorkoutSession{
		UserID:        userObjID,
		RoutineID:     routineObjID,
		Exercises:     exercises,
		ExerciseIndex: 0,
		LastUpdate:    time.Now(),
	}

	sessionID, err := database.CreateSession(newSession)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sessionID.Hex())
}