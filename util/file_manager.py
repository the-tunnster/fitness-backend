import os
import json
from typing import Dict, Any

# Base directories and exercise list file
ROUTINE_DIR = "data/routines/"
HISTORY_DIR = "data/history/"
CACHE_DIR = "data/cache/"
EXERCISE_LIST = "data/exercises/list.csv"

class FileManager:
    """Handles file system operations and JSON serialization."""
    
    @staticmethod
    def safeEmail(user_email: str) -> str:
        return user_email.split("@")[0].replace(".", "_")

    @staticmethod
    def getFolder(user_email: str, base_folder: str) -> str:
        safe_email = FileManager.safeEmail(user_email)
        user_folder = os.path.join(base_folder, safe_email)
        os.makedirs(user_folder, exist_ok=True)
        return user_folder

    @staticmethod
    def getFilepath(user_email: str, base_folder: str, filename: str) -> str:
        folder = FileManager.getFolder(user_email, base_folder)
        return os.path.join(folder, filename)

    @staticmethod
    def loadJson(filepath: str) -> Dict[str, Any]:
        if os.path.exists(filepath):
            with open(filepath, "r") as f:
                return json.load(f)
        return {}

    @staticmethod
    def saveJson(filepath: str, data: Dict[str, Any]) -> None:
        with open(filepath, "w") as f:
            json.dump(data, f, indent=4)

    @staticmethod
    def deleteFile(filepath: str) -> None:
        if os.path.exists(filepath):
            os.remove(filepath)
