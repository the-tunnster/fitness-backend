import streamlit
from util.helpers import *

streamlit.set_page_config(
    page_title="Workout Tracker",
    page_icon="🏋️",
    layout="wide",
    initial_sidebar_state="expanded"
)

quickSetUp()

if not isUserLoggedIn():
    streamlit.warning("Please log in to workout/access your workout data.")
    streamlit.stop()


if streamlit.session_state["user_data"] is None:
    streamlit.session_state["user_data"] = {
        "name": streamlit.experimental_user.name,
        "email": streamlit.experimental_user.email,
    }
    streamlit.rerun()

streamlit.markdown("""

# Welcome to the Workout Tracker </br>
                   
Congratulations on making it this far. </br>
I was sure something would've broken by now. </br>
But here we are. </br>
                   
In any case, let's get started. </br>
The sidebar navigation should help you get around. </br>
If you come across any issues, please let me know, and I'll get onto fixing it ASAP. </br>
For now though, it is in super early development, so please try and break i?. </br>
                   
Good luck, and have a G lift! </br>


""", unsafe_allow_html=True)
