package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"fitness-tracker/internal/database"
	"fitness-tracker/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	var updated_exercise_data []models.RoutineExerciseDTO
	if err := json.NewDecoder(r.Body).Decode(&updated_exercise_data); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	var updated_exercises []models.RoutineExercise
	for _, exercise := range updated_exercise_data {
		exerciseObjID, err := primitive.ObjectIDFromHex(exercise.ExerciseID)
		if err != nil {
			http.Error(w, "Invalid exercise_id", http.StatusBadRequest)
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
		http.Error(w, "Failed to update routine", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	exerciseIndexStr := r.URL.Query().Get("exercise_index")

	if sessionID == "" || exerciseIndexStr == "" {
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

	var updated_exercise_data []models.WorkoutExerciseDTO
	if err := json.NewDecoder(r.Body).Decode(&updated_exercise_data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var updated_exercises []models.WorkoutExercise
	for _, exercise := range updated_exercise_data {
		exerciseObjID, err := primitive.ObjectIDFromHex(exercise.ExerciseID)
		if err != nil {
			http.Error(w, "Invalid exercise_id", http.StatusBadRequest)
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
		"exercises":      updated_exercises,
		"exercise_index": exIndex,
		"last_update":    primitive.NewDateTimeFromTime(time.Now()),
	}

	err = database.UpdateSession(sessionObjID, updates)
	if err != nil {
		http.Error(w, "Failed to update session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateHistoryHandler(w http.ResponseWriter, r *http.Request) {
	workoutID := r.URL.Query().Get("workout_id")
	userID := r.URL.Query().Get("user_id")

	if workoutID == "" || userID == "" {
		http.Error(w, "Missing workout_id or user_id", http.StatusBadRequest)
		return
	}

	workoutObjID, err := primitive.ObjectIDFromHex(workoutID)
	if err != nil {
		http.Error(w, "Invalid workout_id format", http.StatusBadRequest)
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user_id format", http.StatusBadRequest)
		return
	}

	workoutData, err := database.GetWorkoutData(userObjID, workoutObjID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "No workout data found for given workout id", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Failed to fetch workout data", http.StatusInternalServerError)
			return
		}
	}

	for _, exercise := range(workoutData.Exercises){
		exerciseSets := models.ExerciseSets{
			Date: workoutData.WorkoutDate,
			Equipment: exercise.Equipment,
			Variation: exercise.Variation,
			WorkoutSets: exercise.Sets,
		}
	
		updates := bson.M{
			"$push": bson.M{
				"exercise_sets": exerciseSets,
			},
		}

		err = database.UpdateHistory(exercise.ExerciseID, userObjID, updates)
		if err != nil {
			http.Error(w, "Failed to update workout history", http.StatusInternalServerError)
			return
		}
	}
	
	w.WriteHeader(http.StatusNoContent)
}