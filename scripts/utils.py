import json
import logging
from time import sleep
from datetime import datetime
from concurrent.futures import ThreadPoolExecutor

logging.basicConfig(level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s")

# recursively converts any object (dict, list, class) into a JSON-serializable structure
def toJsonSerializable(obj):
    if isinstance(obj, dict):
        return {key: toJsonSerializable(value) for key, value in obj.items()}
    
    elif isinstance(obj, list):
        return [toJsonSerializable(item) for item in obj]
    
    elif hasattr(obj, "__dict__"):
        return {key: toJsonSerializable(value) for key, value in vars(obj).items()}
    
    else:
        return obj

# converts a list of season average objects into a list of dictionarie
def season_averages_to_json(season_averages):
    json_data = []

    for season in season_averages:
        season_dict = {
            "player_id": season.player_id,
            "season": season.season,
            "games_played": season.games_played,
            "pts": season.pts,
            "ast": season.ast,
            "reb": season.reb,
            "stl": season.stl,
            "blk": season.blk,
            "turnover": season.turnover,
            "min": season.min,
            "fgm": season.fgm,
            "fga": season.fga,
            "fg_pct": season.fg_pct,
            "fg3m": season.fg3m,
            "fg3a": season.fg3a,
            "fg3_pct": season.fg3_pct,
            "ftm": season.ftm,
            "fta": season.fta,
            "ft_pct": season.ft_pct,
            "oreb": season.oreb,
            "dreb": season.dreb
        }
        json_data.append(season_dict)

    return json_data

# fetches all available season averages for a player from a start year to the current year
def fetch_all_season_averages(api, player_id, start_year, season_type):
    try:
        current_season = datetime.now().year
        season_averages = []

        print(f"Fetching season averages for player ID {player_id} from {start_year} to {current_season}...")

        for season in range(start_year, current_season + 1):
            try:
                response = api.nba.season_averages.get(season=season , player_id=player_id)

                if response.data[0]:
                    season_averages.append(response.data[0])
                    print(f"Season {season}: Data found.")
                else:
                    print(f"Season {season}: No data available.")

            except Exception as e:
                print(f"Error fetching data for season {season}: {e}")

        print(f"Retrieved {len(season_averages)} valid seasons of data.")
        return season_averages

    except Exception as e:
        print(f"Error in fetching season averages: {e}")
        return []
    
# processes a single player:
# - if already exists in DB, update their current season
# - otherwise, fetch and insert full historical data
def process_player(api, playersCollection, player):
    player_id = player["id"]
    current_season = 2024

    existing_player = playersCollection.find_one({"id": player_id})
    
    if existing_player and existing_player.get("season_averages"):
        print(f"Player {player_id} exists; updating current season data only.")
        try:
            response = api.nba.season_averages.get(season=current_season, player_id=player_id)
            if response.data and len(response.data) > 0:
                new_current_data = vars(response.data[0])
                
                season_averages = existing_player["season_averages"]
                updated = False
                for idx, season_data in enumerate(season_averages):
                    if season_data["season"] == current_season:
                        if season_data != new_current_data:
                            season_averages[idx] = new_current_data
                            updated = True
                        break
                else:
                    season_averages.append(new_current_data)
                    updated = True
                
                if updated:
                    playersCollection.update_one(
                        {"id": player_id},
                        {"$set": {"season_averages": season_averages}}
                    )
                    print(f"Updated current season data for player {player_id}.")
                else:
                    print(f"No changes in current season data for player {player_id}.")
            else:
                print(f"No current season data available for player {player_id}.")
        except Exception as e:
            print(f"Error fetching current season data for player {player_id}: {e}")
    else:
        print(f"New player {player_id} detected; fetching all season data...")
        if player.get("draft_year"):
            start_year = int(player.get("draft_year"))
        else:
            start_year = 1940

        all_seasons_data = fetch_all_season_averages(api, player_id, start_year, "regular")
        season_averages = [vars(season) for season in all_seasons_data] if all_seasons_data else []
        
        update_fields = {
            "first_name": player.get("first_name"),
            "last_name": player.get("last_name"),
            "position": player.get("position"),
            "height": player.get("height"),
            "weight": player.get("weight"),
            "jersey_number": player.get("jersey_number"),
            "college": player.get("college"),
            "country": player.get("country"),
            "draft_year": player.get("draft_year"),
            "draft_round": player.get("draft_round"),
            "draft_number": player.get("draft_number"),
            "team": {
                "id": player["team"].get("id") if player.get("team") else None,
                "conference": player["team"].get("conference") if player.get("team") else None,
                "division": player["team"].get("division") if player.get("team") else None,
                "city": player["team"].get("city") if player.get("team") else None,
                "name": player["team"].get("name") if player.get("team") else None,
                "full_name": player["team"].get("full_name") if player.get("team") else None,
                "abbreviation": player["team"].get("abbreviation") if player.get("team") else None
            },
            "team_id": player.get("team_id"),
            "season_averages": season_averages
        }
        
        playersCollection.update_one(
            {"id": player_id},
            {"$set": update_fields},
            upsert=True
        )
        print(f"Inserted new player data for player {player_id}.")

# fetches all players from the BallDontLie API and processes them
# supports pagination and optional limit
def fetch_players(api, playersCollection, limit=None):
    logging.info("Starting player data fetch...")
    all_players = []
    cursor = None
    total_fetched = 0

    while True:
        try:
            logging.info(f"Fetching players with cursor: {cursor}...")
            response = api.nba.players.list(per_page=100, cursor=cursor) if cursor else api.nba.players.list(per_page=25)
            players = response.data
            playersJSON = toJsonSerializable(players)
            logging.info(f"Received {len(playersJSON)} players from API.")


            with ThreadPoolExecutor(max_workers=3) as executor:
                futures = [executor.submit(process_player, api, playersCollection, player) for player in playersJSON]

                for future in futures:
                    future.result()

            all_players.extend(playersJSON)
            total_fetched += len(playersJSON)
            cursor = response.meta.next_cursor
            logging.info(f"Fetched {len(playersJSON)} players. Total so far: {total_fetched}. Next cursor: {cursor}")

            if not cursor:
                logging.info("Reached the last page. Fetching complete.")
                break
            if limit and total_fetched >= limit:
                logging.info(f"Reached player limit of {limit}. Stopping fetch.")
                break

            sleep(1)

        except Exception as e:
            logging.error(f"Error fetching data: {e}", exc_info=True)
            break

    logging.info("Player data fetch completed.")
    return all_players