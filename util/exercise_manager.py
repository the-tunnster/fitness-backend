import pandas
from typing import List

EXERCISE_LIST = "data/exercises/list.csv"

class ExerciseManager:
    
    @staticmethod
    def getExerciseList() -> List[str]:
        dataFrame = pandas.read_csv(EXERCISE_LIST)
        return dataFrame["exercise-name"].tolist()