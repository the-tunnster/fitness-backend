# Fitness Tracker 🏋️

A comprehensive fitness tracking application built with Python and Streamlit. This project allows users to manage workout routines, track their progress, and analyze workout data.

## Features

- **Routine Management**: Create, edit, and delete workout routines.
- **Workout Tracking**: Log workouts with detailed exercise data, including sets, reps, and weights.
- **Historical Data**: Access past workout data and use it to pre-fill new workouts.
- **Analytics**: Visualize workout progress with metrics like maximum weight and volume per session.
- **User Authentication**: Secure login and session management using Streamlit's experimental user features.
- **Data Persistence**: Save and load data using JSON files for routines, workouts, and history.

## Project Structure
```
fitness-is-my-passion/
├── data/ # Stores user data (routines, history, temp workouts) 
├── pages/ # Streamlit pages for different app functionalities 
│ ├── analytics.py # Workout analytics page 
│ ├── routines.py # Routine management page 
│ └── workout.py # Workout tracking page 
├── util/ # Utility modules for core functionality 
│ ├── cache_manager.py # Handles temporary workout data caching 
│ ├── data_models.py # Defines data models for routines and workouts 
│ ├── exercise_manager.py # Manages exercise list data 
│ ├── file_manager.py # Handles file operations (JSON read/write) 
│ ├── helpers.py # Helper functions for setup and session management 
│ ├── history_manager.py # Manages workout history 
│ └── routine_manager.py # Manages workout routines 
├── certificates/ # SSL certificates (ignored by Git) 
├── FitnessTracker.py # Main entry point for the app 
└── README.md # Project documentation
```


## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/fitness-is-my-passion.git
   cd fitness-is-my-passion
   ```
2. Set up a virtual environment:
	```bash
	python3 -m venv .venv
	source .venv/bin/activate
	```
3. Install dependencies:
	```bash
	pip install -r requirements.txt
	```
4. Run the app:
	```bash
	streamlit run FitnessTracker.py
	```

## Usage

- **Login**: Log in using the sidebar to access your data.
- **Manage Routines**: Navigate to the "User Routines" page to create or edit workout routines.
- **Track Workouts**: Use the "Workout" page to log your exercises and save progress.
- **Analyze Progress**: View workout analytics on the "Analytics" page.

## File Management

- **Routine Data**: Stored in data/routines/ as JSON files.
- **Workout History**: Stored in data/history/ as JSON files.
- **Temporary Workouts**: Cached in data/temp_workouts/.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## Acknowledgments

- Built with Streamlit.
- Inspired by a passion for fitness and data-driven progress tracking.