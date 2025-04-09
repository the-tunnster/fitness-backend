import os
from typing import Optional, Dict, Any
from datetime import datetime

from util.file_manager import FileManager

from util.data_models import Workout, WorkoutExercise, SetData

BASE_TEMP_WORKOUT_DIR = "data/temp_workouts/"

class CacheManager:
    
    @staticmethod
    def getCachedWorkoutData(user_email: str, routine_name: str) -> tuple[Workout, int]:
        filename = f"{routine_name}.json"
        filepath = FileManager.getFilepath(user_email, BASE_TEMP_WORKOUT_DIR, filename)

        workout = Workout(routine_name=routine_name, exercises=[], date=datetime.now())
        index = -1

        cached_data = {}
        if not os.path.exists(filepath):
            return None, -1
    
        cached_data = FileManager.loadJson(filepath)

        cached_workout = cached_data.get("workout", {})
        if cached_workout == {}:
            return None, -1
        
        for cached_exercise in cached_workout:
            exercise = WorkoutExercise(name=cached_exercise, sets=[])
            
            for cached_set in cached_workout[cached_exercise]:
                set_data = SetData(reps=cached_set["reps"], weight=cached_set["weight"])

                exercise.sets.append(set_data)
        
            workout.exercises.append(exercise)
        
        index = cached_data.get("exercise_index", 0)

        return workout, index

    @staticmethod
    def saveCacheData(user_email: str, routine_name: str, workout_data: Workout, cache_index: int) -> None:
        filename = f"{routine_name}.json"
        filepath = FileManager.getFilepath(user_email, BASE_TEMP_WORKOUT_DIR, filename)

        cache_data = {}
        cache_data["exercise_index"] = cache_index
        cache_data["workout"] = {}

        for exercise in workout_data.exercises:
            cache_data["workout"][exercise.name] = []
            for set_data in exercise.sets:
                cache_data["workout"][exercise.name].append({
                    "reps": set_data.reps,
                    "weight": set_data.weight
                })

        FileManager.saveJson(filepath, cache_data)

    @staticmethod
    def deleteCacheData(user_email: str, routine_name: str) -> None:
        filename = f"{routine_name}.json"
        filepath = FileManager.getFilepath(user_email, BASE_TEMP_WORKOUT_DIR, filename)
        if os.path.exists(filepath):
            os.remove(filepath)
