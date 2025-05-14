package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Workout struct {
    ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID     primitive.ObjectID `bson:"userID" json:"user_id"`
    RoutineID  primitive.ObjectID `bson:"routineID" json:"routine_id"`
    WorkoutDate primitive.DateTime `bson:"workoutDate" json:"workout_date"`
}

type WorkoutSet struct {
    SetNumber int     `bson:"setNumber" json:"set_number"`
    Reps      int     `bson:"reps" json:"reps"`
    Weight    float64 `bson:"weight" json:"weight"`
}

type WorkoutExercise struct {
    ExerciseID primitive.ObjectID `bson:"exerciseID" json:"exercise_id"`
    Sets       []WorkoutSet       `bson:"sets" json:"sets"`
}

type FullWorkout struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID      primitive.ObjectID `bson:"userID" json:"user_id"`
    RoutineID   primitive.ObjectID `bson:"routineID" json:"routine_id"`
    WorkoutDate primitive.DateTime `bson:"workoutDate" json:"workout_date"`
    Exercises   []WorkoutExercise  `bson:"exercises" json:"exercises"`
}
