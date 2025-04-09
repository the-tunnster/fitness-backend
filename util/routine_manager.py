import os
from typing import List, Dict, Any

from util.file_manager import FileManager

from util.data_models import Routine, RoutineExercise

ROUTINE_DIR = "data/routines/"
HISTORY_DIR = "data/history/"
CACHE_DIR = "data/cache/"
EXERCISE_LIST = "data/exercises/list.csv"

class RoutineManager:
    
    @staticmethod
    def getRoutineData(user_email: str, routine_name: str) -> Routine:
        routine = Routine(name=routine_name, exercises=[])

        filepath = FileManager.getFilepath(user_email, ROUTINE_DIR, f"{routine_name}.json")
        jsonData =  FileManager.loadJson(filepath)

        for exercise_data in jsonData["exercises"]:
            exercise = RoutineExercise(name=exercise_data["exercise_name"], sets=exercise_data["sets"], reps=exercise_data["reps"])

            routine.exercises.append(exercise)

        return routine
    
    @staticmethod
    def getRoutineList(user_email: str) -> List[str]:
        folderpath = FileManager.getFolder(user_email, ROUTINE_DIR)
        routine_names = [os.path.splitext(filename)[0] for filename in os.listdir(folderpath)]
        return routine_names
    
    @staticmethod
    def saveRoutine(user_email: str, routine_data: Routine) -> None:
        filename = routine_data.name.replace(" ", "_").lower() + ".json"
        filepath = FileManager.getFilepath(user_email, ROUTINE_DIR, filename)

        routine = {}
        routine["exercises"] = []

        for exercise in routine_data.exercises:
                routine["exercises"].append({
                "exercise_name": exercise.name,
                "sets": exercise.sets,
                "reps": exercise.reps
                })

        FileManager.saveJson(filepath, routine)

    @staticmethod
    def deleteRoutine(user_email: str, routine_name: str) -> None:
        filename = routine_name + ".json"
        filepath = FileManager.getFilepath(user_email, ROUTINE_DIR, filename)
        FileManager.deleteFile(filepath)