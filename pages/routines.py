import streamlit as st
import pandas as pd
from util.helpers import *
from util.exercise_manager import ExerciseManager
from util.cache_manager import CacheManager
from util.routine_manager import RoutineManager

from util.data_models import RoutineExercise, Routine

st.set_page_config(
    page_title="Routine Manager",
    page_icon="🏋️",
    layout="wide",
    initial_sidebar_state="collapsed"
)

quickSetUp()
if not isUserLoggedIn():
    st.switch_page("FitnessTracker.py")

if st.session_state.get("user_data") is None:
    getUserInfo()

############################
# Select Routine to Edit
############################

user_email = st.session_state["user_data"]["email"]
selected_routine = st.selectbox("Select a Routine", ["Create New"] + RoutineManager.getRoutineList(user_email), index=0)

if selected_routine == "Create New":
    st.session_state.routine_exercises = [RoutineExercise(name="", sets=0, reps=0)]
    routine_name = st.text_input("Enter Routine Name")
else:
    routine_data = RoutineManager.getRoutineData(user_email, selected_routine)
    st.session_state.routine_exercises = routine_data.exercises
    routine_name = routine_data.name

st.divider()

############################
# Routine Editor (Using DataFrame)
############################

exercise_list = ExerciseManager.getExerciseList()

df = st.data_editor(
    pd.DataFrame(st.session_state.routine_exercises),
    use_container_width=True,
    column_config={
        "name": st.column_config.SelectboxColumn("Exercise Name", options=exercise_list, required=True),
        "sets": st.column_config.NumberColumn("Sets", min_value=1, max_value=10, step=1, default=0),
        "reps": st.column_config.NumberColumn("Reps", min_value=1, step=1, default=0),
    },
    num_rows="dynamic",  # Allows adding/removing rows
)

# Save the edited dataframe back to session state
# Convert the DataFrame to a list of RoutineExercise objects
st.session_state.routine_exercises = [
    RoutineExercise(name=row["name"], sets=int(row["sets"]), reps=int(row["reps"]))
    for _, row in df.iterrows()
]

# Actions: Save, Delete
col1, col2 = st.columns([1, 1])
if col1.button("Save Routine"):
    if routine_name:
        routine = Routine(name=routine_name, exercises=st.session_state.routine_exercises)
        RoutineManager.saveRoutine(user_email, routine)
        CacheManager.deleteCacheData(user_email, routine_name)
        clearSessionState()
        st.rerun(scope="app")

if selected_routine != "Create New" and col2.button("Delete Routine"):
    RoutineManager.deleteRoutine(user_email, selected_routine)
    st.rerun(scope="app")
