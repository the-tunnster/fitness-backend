package handlers

import (
	"encoding/json"
	"net/http"
	"sort"

	"fitness-tracker/internal/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	emailID := r.URL.Query().Get("email")
	if emailID == "" {
		http.Error(w, "Missing email parameter", http.StatusBadRequest)
		return
	}

	user, err := database.GetUserByEmail(emailID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "No user for this email address", http.StatusNotFound)
		} else {
			http.Error(w, "Couldn't find user", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetOverseerHandler(w http.ResponseWriter, r *http.Request) {
	emailID := r.URL.Query().Get("email")
	if emailID == "" {
		http.Error(w, "Missing email parameter", http.StatusBadRequest)
		return
	}

	overseer, err := database.GetOverseerByEmail(emailID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "No user for this email address", http.StatusNotFound)
		} else {
			http.Error(w, "Couldn't find user", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(overseer)
}

func GetExerciseIDHandler(w http.ResponseWriter, r *http.Request) {
	exerciseNames := r.URL.Query()["exercise_name"]
	if len(exerciseNames) == 0 {
		http.Error(w, "Missing exercise_name parameter(s)", http.StatusBadRequest)
		return
	}

	exerciseIDs := make([]string, len(exerciseNames))

	for i, v := range exerciseNames {
		exercise_id, err := database.GetExerciseID(v)
		if err == nil {
			exerciseIDs[i] = exercise_id
		} else {
			exerciseIDs[i] = primitive.NilObjectID.Hex()
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exerciseIDs)
}

func GetExerciseNameHandler(w http.ResponseWriter, r *http.Request) {
	exerciseIDs := r.URL.Query()["exercise_id"]
	if len(exerciseIDs) == 0 {
		http.Error(w, "Missing exercise_id parameter(s)", http.StatusBadRequest)
		return
	}

	exerciseObjIDs := make([]primitive.ObjectID, len(exerciseIDs))
	exerciseNames := make([]string, len(exerciseIDs))

	for i, v := range exerciseIDs {
		exerciseObjID, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			exerciseObjIDs[i] = primitive.ObjectID{}
		} else {
			exerciseObjIDs[i] = exerciseObjID
		}
	}

	for i, v := range exerciseObjIDs {
		exerciseName, err := database.GetExerciseName(v)
		if err == nil {
			exerciseNames[i] = exerciseName
		} else {
			exerciseNames[i] = "unknown"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exerciseNames)
}

func GetExerciseListHandler(w http.ResponseWriter, r *http.Request) {
	exerciseList := database.GetExerciseList()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exerciseList)
}

func GetCardioListHandler(w http.ResponseWriter, r *http.Request) {
	cardioList := database.GetCardioList()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cardioList)
}

func GetExerciseDataHandler(w http.ResponseWriter, r *http.Request) {
	exerciseID := r.URL.Query().Get("exercise_id")
	if exerciseID == "" {
		http.Error(w, "Missing exercise_id parameter", http.StatusBadRequest)
		return
	}

	exerciseObjID, err := primitive.ObjectIDFromHex(exerciseID)
	if err != nil {
		http.Error(w, "Invalid exercise_id format", http.StatusBadRequest)
		return
	}

	exercise, err := database.GetExerciseData(exerciseObjID)
	if err != nil {
		http.Error(w, "Couldn't find that exercise", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercise)
}

func GetRoutineListHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user_id format", http.StatusBadRequest)
		return
	}

	routineList, err := database.GetUserRoutines(userObjID)
	if err != nil {
		http.Error(w, "Couldn't find any routines", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routineList)
}

func GetRoutineDataHandler(w http.ResponseWriter, r *http.Request) {
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

	routine, err := database.GetRoutineData(userObjID, routineObjID)
	if err != nil {
		http.Error(w, "Couldn't find any routine data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routine)
}

func GetWorkoutListHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	routineID := r.URL.Query().Get("routine_id")

	if userID == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	_, err = primitive.ObjectIDFromHex(routineID)
	if err != nil {
		http.Error(w, "Invalid routine_id", http.StatusBadRequest)
		return
	}

	workoutList, err := database.GetUserWorkouts(userObjID)
	if err != nil {
		http.Error(w, "Couldn't find any workouts", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workoutList)
}

func GetWorkoutDataHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	workoutID := r.URL.Query().Get("workout_id")

	if userID == "" || workoutID == "" {
		http.Error(w, "Missing user_id or workout_id", http.StatusBadRequest)
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	workoutObjID, err := primitive.ObjectIDFromHex(workoutID)
	if err != nil {
		http.Error(w, "Invalid workout_id", http.StatusBadRequest)
		return
	}

	workout, err := database.GetWorkoutData(userObjID, workoutObjID)
	if err != nil {
		http.Error(w, "Couldn't fetch workout data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workout)
}

func GetSessionHandler(w http.ResponseWriter, r *http.Request) {
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

	session, err := database.GetUserSessionData(userObjID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Couldn't find any sessions", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Couldn't fetch session data", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

func CountWorkoutHandler(w http.ResponseWriter, r *http.Request) {
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

	workout_count, _ := database.CountWorkouts(userObjID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workout_count)
}

func GetExerciseHistoryHandler(w http.ResponseWriter, r *http.Request) {
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

	exerciseHistory, err := database.GetExerciseHistoryData(exerciseObjID, userObjID)
	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	type ProcessedHistory struct {
		Date   string  `json:"date"`
		Weight float64 `json:"weight,omitempty"`
		Volume float64 `json:"volume,omitempty"`
	}

	// Group by date and calculate max weight and volume
	dailyMap := make(map[string]struct {
		MaxW   float64
		Volume float64
	})

	for _, day := range exerciseHistory.Sets {
		dateStr := day.Date.Time().Format("2006-01-02")
		for _, s := range day.WorkoutSets {
			if s.Reps > 0 && s.Weight > 0 {
				entry := dailyMap[dateStr]
				if s.Weight > entry.MaxW {
					entry.MaxW = s.Weight
				}
				entry.Volume += s.Weight * float64(s.Reps)
				dailyMap[dateStr] = entry
			}
		}
	}

	// Sort dates and prepare results
	sortedDates := make([]string, 0, len(dailyMap))
	for date := range dailyMap {
		sortedDates = append(sortedDates, date)
	}
	sort.Strings(sortedDates)

	if len(sortedDates) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
		return
	}

	results := make([]ProcessedHistory, 0, len(sortedDates))
	for _, date := range sortedDates {
		entry := dailyMap[date]
		results = append(results, ProcessedHistory{
			Date:   date,
			Weight: entry.MaxW,
			Volume: entry.Volume,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func GetCardioHistoryHandler(w http.ResponseWriter, r *http.Request) {
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

	cardioHistory, err := database.GetCardioHistoryData(cardioObjID, userObjID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "No history found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve cardio history data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cardioHistory)
}

func GetWorkoutComparisonHandler(w http.ResponseWriter, r *http.Request) {
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

	workouts, err := database.GetLastTwoWorkouts(userObjID, routineObjID)
	if err != nil {
		http.Error(w, "Failed to fetch workouts", http.StatusInternalServerError)
		return
	}

	if len(workouts) < 2 {
		http.Error(w, "Not enough workouts to compare", http.StatusNotFound)
		return
	}

	workout1 := workouts[0] // Latest workout
	workout2 := workouts[1] // Second latest workout

	type MetricChange struct {
		ExerciseName string  `json:"exercise_name"`
		Variation    string  `json:"variation"`
		MaxWeight    float64 `json:"max_weight"`
		TotalReps    int     `json:"reps"`
		TotalVolume  float64 `json:"volume"`
		WeightChange float64 `json:"weight_change"`
		RepsChange   int     `json:"reps_change"`
		VolumeChange float64 `json:"volume_change"`
	}

	exerciseMetricMap := make(map[string]MetricChange)

	for _, exercise := range workout1.Exercises {
		exercise_name, _ := database.GetExerciseName(exercise.ExerciseID)
		key := exercise_name + "|" + exercise.Variation

		max_weight, total_reps, total_volume := 0.0, 0, 0.0

		for _, set := range exercise.Sets {
			total_reps += set.Reps
			total_volume += float64(set.Reps) * set.Weight

			if set.Weight > max_weight {
				max_weight = set.Weight
			}
		}

		exerciseMetricMap[key] = MetricChange{
			ExerciseName: exercise_name,
			Variation:    exercise.Variation,
			MaxWeight:    max_weight,
			TotalReps:    total_reps,
			TotalVolume:  total_volume,
			WeightChange: max_weight,
			RepsChange:   total_reps,
			VolumeChange: total_volume,
		}
	}

	for _, exercise := range workout2.Exercises {
		exercise_name, _ := database.GetExerciseName(exercise.ExerciseID)
		key := exercise_name + "|" + exercise.Variation

		max_weight, total_reps, total_volume := 0.0, 0, 0.0

		for _, set := range exercise.Sets {
			total_reps += set.Reps
			total_volume += float64(set.Reps) * set.Weight

			if set.Weight > max_weight {
				max_weight = set.Weight
			}
		}

		last_workout_metrics, ok := exerciseMetricMap[key]
		if !ok {
			continue
		}

		last_workout_metrics.RepsChange = last_workout_metrics.TotalReps - total_reps
		last_workout_metrics.WeightChange = last_workout_metrics.MaxWeight - max_weight
		last_workout_metrics.VolumeChange = last_workout_metrics.TotalVolume - total_volume

		exerciseMetricMap[key] = last_workout_metrics
	}

	changes := make([]MetricChange, 0, len(exerciseMetricMap))
	for _, metrics := range exerciseMetricMap {
		changes = append(changes, metrics)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(changes)
}
