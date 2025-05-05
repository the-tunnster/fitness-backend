package handlers

import (
	"log"
	"net/http"
    "encoding/json"

    "fitness-tracker/internal/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    emailID := r.URL.Query().Get("email")
    if emailID == "" {
        http.Error(w, "Missing email parameter", http.StatusBadRequest)
        return
    }

    user, err := database.GetUser(emailID)
    if err != nil {
        log.Println("Couldn't fetch user info")
        log.Println(err)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
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
        log.Println("Couldn't fetch user routines")
        log.Println(err)
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
        log.Println("Couldn't fetch routine data")
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
        log.Println("Couldn't fetch user workouts")
        log.Println(err)
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
        log.Println("Couldn't fetch workout data")
        log.Println(err)
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

    session, err := database.GetSessionData(userObjID)
    if err != nil {
        log.Println("Couldn't fetch workout data")
        log.Println(err)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(session)
}