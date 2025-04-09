import os
from typing import Dict, Any
from datetime import datetime

from util.file_manager import FileManager

from util.data_models import Routine, Workout, WorkoutExercise, SetData

HISTORY_DIR = "data/history/"
CACHE_DIR = "data/cache/"
EXERCISE_LIST = "data/exercises/list.csv"
class HistoryManager:

    @staticmethod
    def getHistoryFilePath(user_email: str) -> str:
        filename = "_consolidated_history.json"
        return FileManager.getFilepath(user_email, HISTORY_DIR, filename)

    @staticmethod
    def loadConsolidatedHistory(user_email: str) -> Dict[str, Any]:
        history_file = HistoryManager.getHistoryFilePath(user_email)
        return FileManager.loadJson(history_file) if os.path.exists(history_file) else {}

    @staticmethod
    def getLastWorkout(user_email: str, routine_data: Routine) -> Workout:
        history = HistoryManager.loadConsolidatedHistory(user_email)

        last_workout = Workout(routine_name=routine_data.name, exercises=[], date=datetime.now())

        for exercise in routine_data.exercises:
            workout_exercise = WorkoutExercise(name=exercise.name, sets=[])
            
            num_sets = exercise.sets

            exercise_data = history.get(workout_exercise.name, [])
            historical_sets = exercise_data[-1] if isinstance(exercise_data, list) and exercise_data else []

            for historical_set in historical_sets:
                set_data = SetData(reps=0, weight=0.0)

                set_data.reps = historical_set.get("reps", 0)
                set_data.weight = historical_set.get("weight", 0.0)

                workout_exercise.sets.append(set_data)

            if num_sets >= len(historical_sets):
                missing = num_sets - len(historical_sets)
                for _ in range(missing):
                    workout_exercise.sets.append(SetData(reps=0, weight=0.0))

            last_workout.exercises.append(workout_exercise)

        return last_workout

    @staticmethod
    def appendWorkout(user_email: str, workout_data: Workout) -> None:
        history = HistoryManager.loadConsolidatedHistory(user_email)

        for workout_ex in workout_data.exercises:
            # sets_data = ({"reps": s.reps, "weight": s.weight} for s in workout_ex.sets)
            sets_data = [{"reps": s.reps, "weight": s.weight} for s in workout_ex.sets]

            if workout_ex.name not in history or not isinstance(history[workout_ex.name], list):
                history[workout_ex.name] = []

            history[workout_ex.name].extend([sets_data])

        history_file = HistoryManager.getHistoryFilePath(user_email)
        try:
            FileManager.saveJson(history_file, history)
        except Exception as e:
            print(f"Error saving history to {history_file}: {e}")

    @staticmethod
    def saveWorkout(user_email:str, workout_data: Workout) -> None:
        filename = f"{workout_data.routine_name}_{datetime.now().strftime('%Y%m%d')}.json"
        filepath = FileManager.getFilepath(user_email, HISTORY_DIR, filename)
        workout = {}

        for exercise in workout_data.exercises:
            workout[exercise.name] = []
            for set_data in exercise.sets:
                workout[exercise.name].append({
                    "reps": set_data.reps,
                    "weight": set_data.weight
                })
        
        FileManager.saveJson(filepath, workout)

    @staticmethod
    def buildConsolidatedHistory(user_email: str) -> None:
        filepath = FileManager.getFilepath(user_email, HISTORY_DIR, "_consolidated_history.json")
        if os.path.exists(filepath):
            return

        history = {}  # consolidated history data

        # Iterate over all JSON files in the history directory except the consolidated one
        history_folder = FileManager.getFolder(user_email, HISTORY_DIR)
        for filename in os.listdir(history_folder):
            if filename.endswith(".json") and filename != "_consolidated_history.json":
                file_path = os.path.join(history_folder, filename)
                workout = FileManager.loadJson(file_path)

                for exercise_name in workout:
                    set_data = workout[exercise_name]

                    if exercise_name not in history:
                        history[exercise_name] = []

                    history[exercise_name].extend([set_data])

        FileManager.saveJson(filepath, history)