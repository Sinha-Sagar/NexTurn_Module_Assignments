import os

class Config:
    SECRET_KEY = os.environ.get('SECRET_KEY') or 'default_secret_key'
    SQLALCHEMY_DATABASE_URI = r'sqlite:///D:\Nexturn\SagarSinha-NexTurn-Program\M3_PythonAPIExercise_Assignments\mydb.db'
    SQLALCHEMY_TRACK_MODIFICATIONS = False