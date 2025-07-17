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
		http.Error(w, "Couldn't find user", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
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
