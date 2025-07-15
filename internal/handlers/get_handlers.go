package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"sort"
	"time"

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
		Date               string  `json:"date"`
		Weight             float64 `json:"weight,omitempty"`
		InterpolatedWeight float64 `json:"interpolated_weight"`
		Volume             float64 `json:"volume,omitempty"`
	}

	// --- Group by date ---
	type dailyStats struct {
		Date   time.Time
		MaxW   float64
		Volume float64
	}
	dailyMap := map[string]*dailyStats{}

	for _, day := range exerciseHistory.Sets {
		dateTime := day.Date.Time()
		dateStr := dateTime.Format("2006-01-02")
		if _, exists := dailyMap[dateStr]; !exists {
			dailyMap[dateStr] = &dailyStats{Date: dateTime}
		}
		for _, s := range day.WorkoutSets {
			if s.Reps > 0 && s.Weight > 0 {
				if s.Weight > dailyMap[dateStr].MaxW {
					dailyMap[dateStr].MaxW = s.Weight
				}
				dailyMap[dateStr].Volume += s.Weight * float64(s.Reps)
			}
		}
	}

	// --- Sort dates ---
	var sortedDates []string
	for d := range dailyMap {
		sortedDates = append(sortedDates, d)
	}
	sort.Strings(sortedDates)

	if len(sortedDates) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
		return
	}

	// --- Fill in missing days and interpolate ---
	firstDate, _ := time.Parse("2006-01-02", sortedDates[0])
	lastDate, _ := time.Parse("2006-01-02", sortedDates[len(sortedDates)-1])

	results := []ProcessedHistory{}
	var prevDate string
	var prevW float64

	for d := firstDate; !d.After(lastDate); d = d.AddDate(0, 0, 1) {
		ds := d.Format("2006-01-02")
		if entry, ok := dailyMap[ds]; ok {
			results = append(results, ProcessedHistory{
				Date:               ds,
				Weight:             entry.MaxW,
				InterpolatedWeight: entry.MaxW,
				Volume:             entry.Volume,
			})
			prevDate = ds
			prevW = entry.MaxW
			continue
		}

		// Interpolate if possible
		// Find next known
		var nextDate string
		var nextW float64
		for i := range sortedDates {
			if sortedDates[i] > ds {
				nextDate = sortedDates[i]
				nextW = dailyMap[nextDate].MaxW
				break
			}
		}
		if prevDate == "" || nextDate == "" {
			continue
		}

		d0, _ := time.Parse("2006-01-02", prevDate)
		dn, _ := time.Parse("2006-01-02", nextDate)
		totalGap := dn.Sub(d0).Hours() / 24
		curGap := d.Sub(d0).Hours() / 24

		interpolated := prevW + (nextW-prevW)*(1-math.Cos(math.Pi*curGap/totalGap))/2

		results = append(results, ProcessedHistory{
			Date:               ds,
			InterpolatedWeight: interpolated,
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
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cardioHistory)
}
