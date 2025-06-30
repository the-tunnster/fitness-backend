package handlers

import (
	"encoding/json"
	"log"
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
	userID := r.URL.Query().Get("user_id")
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	var routine_data models.FullRoutineDTO
	if err := json.NewDecoder(r.Body).Decode(&routine_data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var exercises []models.RoutineExercise
	for _, exercise := range routine_data.Exercises {
		exerciseObjID, err := primitive.ObjectIDFromHex(exercise.ExerciseID)
		if err != nil {
			http.Error(w, "Invalid exercise_id", http.StatusBadRequest)
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
		UserID:      userObjID,
		Name:        routine_data.Name,
		Exercises:   exercises,
	}

	routineID, err := database.CreateRoutine(newRoutine)
	if err != nil {
		http.Error(w, "Failed to create routine", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(routineID.Hex())
}

func CreateWorkoutSessionHandler(w http.ResponseWriter, r *http.Request) {
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

	full_routine, err := database.GetRoutineData(userObjID, routineObjID)
	if err != nil {
		http.Error(w, "Couldn't find a matching routine for the provided routine_id", http.StatusBadRequest)
		return
	}

	var workout_exercises []models.WorkoutExercise
	for _, routine_exercise := range(full_routine.Exercises) {

		var exercise_sets []models.WorkoutSet
		for i := 0 ; i < routine_exercise.TargetSets ; i++ {
			exercise_sets = append(exercise_sets, models.WorkoutSet{
				Reps: routine_exercise.TargetReps,
				Weight: 0.0,
			})
		}
		
		workout_exercises = append(workout_exercises, models.WorkoutExercise{
			ExerciseID: routine_exercise.ExerciseID,
			Equipment: "none",
			Variation: "none",
			Sets: exercise_sets,
		})

	}

	workout_session := models.WorkoutSession{
		UserID: userObjID,
		RoutineID: routineObjID,
		Exercises: workout_exercises,
		ExerciseIndex: 0,
		LastUpdate: primitive.NewDateTimeFromTime(time.Now()),
	}

	sessionID, err := database.CreateSession(workout_session)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	log.Print("succesfully created workout session", sessionID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sessionID.Hex())
}

func CreateWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	sessionObjID, err := primitive.ObjectIDFromHex(sessionID)
	if err != nil {
		http.Error(w, "Invalid session_id", http.StatusBadRequest)
		return
	}

	workout_session, err := database.GetSessionData(sessionObjID)
	if err != nil {
		http.Error(w, "Failed to find session", http.StatusInternalServerError)
		return
	}

	var workout_exercises []models.WorkoutExercise
	for _, exercise := range(workout_session.Exercises) {

		var exercise_sets []models.WorkoutSet
		for i := 0 ; i < len(exercise.Sets) ; i++ {
			exercise_sets = append(exercise_sets, models.WorkoutSet{
				Reps: exercise.Sets[i].Reps,
				Weight: exercise.Sets[i].Weight,
			})
		}

		workout_exercises = append(workout_exercises, models.WorkoutExercise{
			ExerciseID: exercise.ExerciseID,
			Equipment: exercise.Equipment,
			Variation: exercise.Variation,
			Sets: exercise_sets,
		})
	}

	workout := models.FullWorkout{
		ID: primitive.NilObjectID,
		UserID: workout_session.UserID,
		RoutineID: workout_session.RoutineID,
		WorkoutDate: primitive.NewDateTimeFromTime(time.Now()),
		Exercises: workout_exercises,
	}

	workoutID, err := database.CreateWorkout(workout)
	if err != nil {
		http.Error(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(workoutID.Hex())
}