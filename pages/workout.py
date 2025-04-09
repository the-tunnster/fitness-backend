import streamlit as st

from util.helpers import *

from util.cache_manager import CacheManager
from util.routine_manager import RoutineManager
from util.history_manager import HistoryManager

from util.data_models import Routine, Workout, RoutineExercise, WorkoutExercise, SetData

st.set_page_config(
    page_title="Workout Tracker",
    page_icon="🏋️",
    layout="wide",
    initial_sidebar_state="collapsed"
)

quickSetUp()
if not isUserLoggedIn():
    st.switch_page("FitnessTracker.py")

if st.session_state.get("user_data") is None:
    getUserInfo()

user_email = st.session_state["user_data"]["email"]
routine_list = [None] + RoutineManager.getRoutineList(user_email)

# Load routine information.
routine_name = st.selectbox("Select a Routine", routine_list, index=0)
if len(routine_list) == 1:
    st.warning("No routines available. Please create one first.")
    st.stop()
if routine_name is None:
    st.stop()

# Load routine data and build exercise list.
routine_data = RoutineManager.getRoutineData(user_email, routine_name)
exercise_list = [exercise.name for exercise in routine_data.exercises]

# Load historical workout data from the consolidated history.
historical_workout_data = HistoryManager.getLastWorkout(user_email, routine_data)
cached_workout_data, cached_workout_index = CacheManager.getCachedWorkoutData(user_email, routine_name)

current_workout = None

if cached_workout_data is not None:
    st.info("Using cached workout data.")
    
    current_workout = cached_workout_data
    st.session_state.exercise_index = cached_workout_index
else:
    st.info("Using historical workout data.")

    current_workout = historical_workout_data
    st.session_state.exercise_index = 0

    CacheManager.saveCacheData(user_email, routine_name, current_workout, st.session_state.exercise_index)

st.divider()

# --- Exercise Rendering Fragment ---
@st.fragment
def workout_fragment(workout_obj: Workout):
    current_exercise_name = exercise_list[st.session_state.exercise_index]
    exercise_data = workout_obj.exercises[st.session_state.exercise_index]
    st.header(f"{current_exercise_name}")

    col1, col2 = st.columns(2, gap="small")
    col1.write("Reps")
    col2.write("Weight in kg")

    new_sets = []
    # Iterate over each set.
    for i, set_item in enumerate(exercise_data.sets):
        reps_input = col1.number_input(
            label=f"Set {i+1} reps",
            label_visibility="hidden",
            value=set_item.reps,
            min_value=0,
            key=f"{current_exercise_name}_reps_{i}"
        )
        weight_input = col2.number_input(
            label=f"Set {i+1} weight",
            label_visibility="hidden",
            value=float(set_item.weight),
            min_value=-50.0,
            step=0.01,
            format="%.2f",
            key=f"{current_exercise_name}_weight_{i}"
        )
        new_sets.append(SetData(reps=reps_input, weight=weight_input))
    
    # Update the exercise data.
    workout_obj.exercises[st.session_state.exercise_index].sets = new_sets

    CacheManager.saveCacheData(user_email, routine_name, workout_obj, st.session_state.exercise_index)

    # Navigation buttons.
    col_prev, col_next = st.columns(2)
    if col_prev.button("Previous") and st.session_state.exercise_index > 0:
        st.session_state.exercise_index -= 1
        CacheManager.saveCacheData(user_email, routine_name, workout_obj, st.session_state.exercise_index)
        st.rerun(scope="fragment")
    if col_next.button("Next") and st.session_state.exercise_index < len(exercise_list) - 1:
        st.session_state.exercise_index += 1
        CacheManager.saveCacheData(user_email, routine_name, workout_obj, st.session_state.exercise_index)
        st.rerun(scope="fragment")

    return workout_obj

# Render the fragment and update workout_obj.
workout_obj = workout_fragment(current_workout)

st.divider()

# --- Save Workout ---
if st.button("Save Workout"):
    HistoryManager.buildConsolidatedHistory(user_email)
    HistoryManager.appendWorkout(user_email, current_workout)
    HistoryManager.saveWorkout(user_email, current_workout)
    CacheManager.deleteCacheData(user_email, routine_name)
    st.success("Workout recorded!")
