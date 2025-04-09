from dataclasses import dataclass
from typing import List
from datetime import datetime

@dataclass
class SetData:
    reps: int
    weight: float

@dataclass
class WorkoutExercise:
    name: str
    sets: List[SetData]

@dataclass
class RoutineExercise:
    name: str
    sets: int
    reps: int

@dataclass
class Workout:
    routine_name: str
    date: datetime
    exercises: List[WorkoutExercise]

@dataclass
class Routine:
    name: str
    exercises: List[RoutineExercise]