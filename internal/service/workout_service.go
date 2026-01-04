package service

import (
	"time"

	"fitness-tracker/internal/database"
	"fitness-tracker/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GenerateWorkoutSession builds a workout session based on the user's routine,
// previous workout, and exercise history.
func GenerateWorkoutSession(userID, routineID string) (*models.WorkoutSession, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	routineObjID, err := primitive.ObjectIDFromHex(routineID)
	if err != nil {
		return nil, err
	}

	fullRoutine, err := database.GetRoutineData(userObjID, routineObjID)
	if err != nil {
		return nil, err
	}

	var lastWorkout *models.FullWorkout
	lw, err := database.GetLastWorkoutForRoutine(userObjID, routineObjID)
	if err == nil {
		lastWorkout = &lw
	}

	workoutExercises := buildExercisesFromRoutine(fullRoutine, lastWorkout, userObjID)

	session := &models.WorkoutSession{
		UserID:        userObjID,
		RoutineID:     routineObjID,
		Exercises:     workoutExercises,
		ExerciseIndex: 0,
		LastUpdate:    primitive.NewDateTimeFromTime(time.Now()),
	}

	return session, nil
}

func buildExercisesFromRoutine(routine models.FullRoutine, lastWorkout *models.FullWorkout, userID primitive.ObjectID) []models.WorkoutExercise {
	var workoutExercises []models.WorkoutExercise

	var lastExercises []models.WorkoutExercise
	lastIndex := 0
	if lastWorkout != nil {
		lastExercises = lastWorkout.Exercises
	}

	for _, rEx := range routine.Exercises {
		sets := []models.WorkoutSet{}
		equipment := "None"
		variation := "None"
		i := 0

		if lastWorkout != nil && lastIndex < len(lastExercises) && lastExercises[lastIndex].ExerciseID == rEx.ExerciseID {
			// Prefer last workout's matched exercise
			lwEx := lastExercises[lastIndex]
			equipment = lwEx.Equipment
			variation = lwEx.Variation
			for ; i < rEx.TargetSets && i < len(lwEx.Sets); i++ {
				sets = append(sets, models.WorkoutSet{
					Reps:   lwEx.Sets[i].Reps,
					Weight: lwEx.Sets[i].Weight,
				})
			}
			lastIndex++
		} else {
			// Fallback to last history entry if present
			hist, err := database.GetExerciseHistoryData(rEx.ExerciseID, userID)
			if err != mongo.ErrNoDocuments && len(hist.Sets) > 0 {
				lastSet := hist.Sets[len(hist.Sets)-1]
				equipment = lastSet.Equipment
				variation = lastSet.Variation
				for ; i < rEx.TargetSets && i < len(lastSet.WorkoutSets); i++ {
					sets = append(sets, models.WorkoutSet{
						Reps:   lastSet.WorkoutSets[i].Reps,
						Weight: lastSet.WorkoutSets[i].Weight,
					})
				}
			}
		}

		// Fill remaining sets with defaults
		for ; i < rEx.TargetSets; i++ {
			sets = append(sets, models.WorkoutSet{
				Reps:   rEx.TargetReps,
				Weight: 0.0,
			})
		}

		workoutExercises = append(workoutExercises, models.WorkoutExercise{
			ExerciseID: rEx.ExerciseID,
			Equipment:  equipment,
			Variation:  variation,
			Sets:       sets,
		})
	}

	return workoutExercises
}
