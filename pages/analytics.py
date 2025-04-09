import streamlit as st
import pandas as pd

from util.helpers import *
from util.routine_manager import RoutineManager
from util.history_manager import HistoryManager

# Configure the Streamlit page.
st.set_page_config(
    page_title="Workout Analytics",
    page_icon="🏋️",
    layout="wide",
    initial_sidebar_state="collapsed"
)

# Perform initial setup.
quickSetUp()

# Redirect to login if the user is not logged in.
if not isUserLoggedIn():
    st.switch_page("FitnessTracker.py")

# Retrieve user information if not already set.
if st.session_state["user_data"] is None:
    getUserInfo()

user_email = st.session_state["user_data"]["email"]
routine_list = RoutineManager.getRoutineList(user_email)

if not routine_list:
    st.warning("No routines available. Please create one first.")
    st.stop()

# Let the user select a routine.
routine_name = st.selectbox("Select a Routine", routine_list, index=0)
st.divider()

# Load routine details and get the list of exercises.
routine_data = RoutineManager.getRoutineData(user_email, routine_name)
historical_data = HistoryManager.loadConsolidatedHistory(user_email)

# Create an exercises dictionary where each key is an exercise name and each value 
# is a list of sessions. Each session contains a list of sets with keys 'reps' and 'weight'.
exercises = {
    exercise.name: historical_data[exercise.name]
    for exercise in routine_data.exercises if exercise.name in historical_data
}

if not exercises:
    st.warning("No exercises found for this routine.")
    st.stop()

# Calculate the max weight and volume metrics for each exercise.
max_weight_data = {}
volume_data = {}

for exercise, sessions in exercises.items():
    max_weight_series = []
    volume_series = []
    for session in sessions:
        if session:  # Check if the session is not empty.
            # Maximum weight for the session.
            session_max = max(set_data["weight"] for set_data in session)
            # Volume is the sum of (reps * weight) for each set.
            session_volume = sum(set_data["reps"] * set_data["weight"] for set_data in session)
        else:
            session_max = None
            session_volume = None
        max_weight_series.append(session_max)
        volume_series.append(session_volume)
    max_weight_data[exercise] = max_weight_series
    volume_data[exercise] = volume_series

# Create two tabs using Streamlit's tab functionality.
tab1, tab2 = st.tabs(["Max Weight", "Volume"])

with tab1:
    st.subheader("Maximum Weight per Session")
    # Convert the max weight data to a DataFrame.
    # Each column is an exercise, and the rows are session numbers.
    df_max = pd.DataFrame({exercise: pd.Series(weights) for exercise, weights in max_weight_data.items()})
    # Adjust the index to start at 1 (i.e., Session 1, Session 2, ...).
    df_max.index = df_max.index + 1
    st.line_chart(df_max)

with tab2:
    st.subheader("Volume per Session")
    # Convert the volume data to a DataFrame.
    df_volume = pd.DataFrame({exercise: pd.Series(volumes) for exercise, volumes in volume_data.items()})
    df_volume.index = df_volume.index + 1
    st.line_chart(df_volume)
