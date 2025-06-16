from balldontlie import BalldontlieAPI
from dotenv import load_dotenv
import os
from pymongo import MongoClient
import certifi
from utils import fetch_players


load_dotenv()

ballDontLieKey = os.getenv('BALLDONTLIE_API_KEY')
mongoURI = os.getenv('MONGO_DB_CONNECTION_STRING')
client = MongoClient(mongoURI, tlsCAFile=certifi.where())
db = client['nba-stats']
playersCollection = db['players']
api = BalldontlieAPI(api_key=ballDontLieKey)
players = fetch_players(api, playersCollection)
