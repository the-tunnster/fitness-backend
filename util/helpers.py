import os
import streamlit

def folderSetup():
	os.makedirs('data/', exist_ok=True)

	os.makedirs('data/history/', exist_ok=True)
	os.makedirs('data/routines/', exist_ok=True)

	os.makedirs('data/temp_workouts/', exist_ok=True)

def initSessionState():
	if "user_data" not in streamlit.session_state:
		streamlit.session_state["user_data"] = None
	if "routine_exercises" not in streamlit.session_state:
		streamlit.session_state["routine_exercises"] = []
	if "current_exercise_index" not in streamlit.session_state:
		streamlit.session_state.current_exercise_index = 0
	if "workout_data" not in streamlit.session_state:
		streamlit.session_state.workout_data = {}

def isUserLoggedIn():
	return streamlit.experimental_user.is_logged_in

def loginUser(auth_provider:str):
	streamlit.login(auth_provider)

def logoutUser():
	streamlit.logout()

def hideSidebar():
	streamlit.markdown("""
		<style>
			[data-testid="stSidebarNav"] {
				display: none !important;
			}
			[data-testid="stSidebarContent"] {
				padding-top: 0px;
			}
			[data-testid="stColumn"] {
    				width: calc(33.3333% - 1rem) !important;
    				flex: 1 1 calc(33.3333% - 1rem) !important;
    				min-width: calc(33% - 1rem) !important;
			}	
		</style>
	""", unsafe_allow_html=True)

def actualSidebar():
	streamlit.sidebar.title("Navigation Shiz")
	streamlit.sidebar.page_link(page="./FitnessTracker.py", label="Home")
	streamlit.sidebar.page_link(page="./pages/routines.py", label="User Routines")
	streamlit.sidebar.page_link(page="./pages/workout.py", label="Workout")
	streamlit.sidebar.page_link(page="./pages/analytics.py", label="Analytics")
	streamlit.sidebar.markdown("---")

	if streamlit.session_state["user_data"] is None:
		google_button = streamlit.sidebar.button("Log in with Google")
		if google_button:
			loginUser("google")
	else:
		streamlit.sidebar.title(f"Hello, {streamlit.session_state["user_data"]["name"]}!")
		logout_button = streamlit.sidebar.button("Log Out")	
		if logout_button:
			logoutUser()

def getUserInfo():
	streamlit.session_state["user_data"] = {
		"name": streamlit.experimental_user.name,
		"email": streamlit.experimental_user.email,
	}
	streamlit.rerun()

def clearSessionState():
	if "routine_exercises" in streamlit.session_state:
		del streamlit.session_state.routine_exercises

def quickSetUp():
	folderSetup()
	initSessionState()
	hideSidebar()
	actualSidebar()