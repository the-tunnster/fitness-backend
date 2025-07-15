package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"fitness-tracker/internal/database"
	"fitness-tracker/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		UserID:    userObjID,
		Name:      routine_data.Name,
		Exercises: exercises,
	}

	routineID, err := database.CreateRoutine(newRoutine)
	if err != nil {
		http.Error(w, "Failed to create routine", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(routineID.Hex())
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

	full_routine, err := database.GetRoutineData(userObjID, routineObjID)
	if err != nil {
		http.Error(w, "Couldn't find a matching routine for the provided routine_id", http.StatusBadRequest)
		return
	}

	var last_workout_exercises []models.WorkoutExercise
	routine_exercises := full_routine.Exercises

	var workout_exercises []models.WorkoutExercise

	last_workout, err := database.GetLastWorkoutForRoutine(userObjID, routineObjID)
	if err == nil {
		last_workout_exercises = last_workout.Exercises
	}

	lastWorkoutIndex := 0

	for _, routine_exercise := range routine_exercises {

		var exercise_sets []models.WorkoutSet
		historic_exercise_equipment := "None"
		historic_exercise_variation := "None"
		i := 0

		if lastWorkoutIndex < len(last_workout_exercises) && last_workout_exercises[lastWorkoutIndex].ExerciseID == routine_exercise.ExerciseID {
			
			lastWorkoutExercise := last_workout_exercises[lastWorkoutIndex]
			historic_exercise_equipment = lastWorkoutExercise.Equipment
			historic_exercise_variation = lastWorkoutExercise.Variation

			// Use the sets from the last workout
			for ; i < routine_exercise.TargetSets && i < len(lastWorkoutExercise.Sets); i++ {
				exercise_sets = append(exercise_sets, models.WorkoutSet{
					Reps:   lastWorkoutExercise.Sets[i].Reps,
					Weight: lastWorkoutExercise.Sets[i].Weight,
				})
			}

			lastWorkoutIndex++
		} else { // Fallback to exercise history if no last workout data or no match
			exercise_history, err := database.GetExerciseHistoryData(routine_exercise.ExerciseID, userObjID)

			if err != mongo.ErrNoDocuments && len(exercise_history.Sets) > 0 {
				historic_workout_sets := exercise_history.Sets[len(exercise_history.Sets)-1].WorkoutSets
				historic_exercise_equipment = exercise_history.Sets[len(exercise_history.Sets)-1].Equipment
				historic_exercise_variation = exercise_history.Sets[len(exercise_history.Sets)-1].Variation

				for ; i < routine_exercise.TargetSets && i < len(historic_workout_sets); i++ {
					exercise_sets = append(exercise_sets, models.WorkoutSet{
						Reps:   historic_workout_sets[i].Reps,
						Weight: historic_workout_sets[i].Weight,
					})
				}
			}
		}

		// Fill remaining sets with default values
		for ; i < routine_exercise.TargetSets; i++ {
			exercise_sets = append(exercise_sets, models.WorkoutSet{
				Reps:   routine_exercise.TargetReps,
				Weight: 0.0,
			})
		}

		workout_exercises = append(workout_exercises, models.WorkoutExercise{
			ExerciseID: routine_exercise.ExerciseID,
			Equipment:  historic_exercise_equipment,
			Variation:  historic_exercise_variation,
			Sets:       exercise_sets,
		})

	}

	workout_session := models.WorkoutSession{
		UserID:        userObjID,
		RoutineID:     routineObjID,
		Exercises:     workout_exercises,
		ExerciseIndex: 0,
		LastUpdate:    primitive.NewDateTimeFromTime(time.Now()),
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

	workout := models.FullWorkout{
		ID:          primitive.NilObjectID,
		UserID:      workout_session.UserID,
		RoutineID:   workout_session.RoutineID,
		WorkoutDate: primitive.NewDateTimeFromTime(time.Now()),
		Exercises:   workout_session.Exercises,
	}

	workoutID, err := database.CreateWorkout(workout)
	if err != nil {
		http.Error(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(workoutID.Hex())
}

func CreateExerciseHistoryHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	exerciseID := r.URL.Query().Get("exercise_id")

	if userID == "" || exerciseID == "" {
		http.Error(w, "Missing user_id or exercise_id", http.StatusBadRequest)
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	exerciseObjID, err := primitive.ObjectIDFromHex(exerciseID)
	if err != nil {
		http.Error(w, "Invalid exercise_id", http.StatusBadRequest)
		return
	}

	exerciseHistory := models.ExerciseHistory{
		ExerciseID: exerciseObjID,
		UserID:     userObjID,
		Sets:       []models.ExerciseSets{},
	}

	historyID, err := database.CreateExerciseHistory(exerciseHistory)
	if err != nil {
		http.Error(w, "Failed to create exercise history: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(historyID.Hex())
}

func CreateCardioHistoryHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	cardioID := r.URL.Query().Get("cardio_id")

	if userID == "" || cardioID == "" {
		http.Error(w, "Missing user_id or cardio_id", http.StatusBadRequest)
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	cardioObjID, err := primitive.ObjectIDFromHex(cardioID)
	if err != nil {
		http.Error(w, "Invalid cardio_id", http.StatusBadRequest)
		return
	}

	cardioHistory := models.CardioHistory{
		CardioID: cardioObjID,
		UserID:   userObjID,
		Sessions: []models.CardioSession{},
	}

	historyID, err := database.CreateCardioHistory(cardioHistory)
	if err != nil {
		http.Error(w, "Failed to create cardio history: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(historyID.Hex())
}
