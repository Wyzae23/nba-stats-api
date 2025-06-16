# NBA Stats Scraper (Python)

This script fetches NBA player data and season averages from the [BallDontLie API](https://www.balldontlie.io) and updates a MongoDB collection (`nba-stats.players`) with that information.

It is intended to be used as a backend data ingestion pipeline, feeding the player data consumed by a Go-based REST API (or any other frontend/backend).

---

## Directory Structure

```
scripts/
├── updateDB.py         # Main entry point for data sync
├── utils.py            # Helpers for formatting and DB updates
├── requirements.txt    # Python dependencies
├── .env                # API and DB credentials (excluded from Git)
```

---

## Prerequisites

- Python 3.9 or newer
- Access to a MongoDB database (Atlas or local)
- A BallDontLie API key (get one from https://www.balldontlie.io)

---

## Setup Instructions

### 1. Clone the repo and enter the `scripts/` directory

```bash
cd scripts
```

### 2. Create a `.env` file

```bash
touch .env
```

Fill it with:

```env
BALLDONTLIE_API_KEY=your_api_key_here
MONGO_DB_CONNECTION_STRING=your_mongo_connection_string
```

> Do **not** commit this file. Make sure `.env` is in `.gitignore`.

### 3. Install dependencies

```bash
pip install -r requirements.txt
```

---

## Running the Scraper

```bash
python updateDB.py
```

This will:
- Fetch all players (using cursor-based pagination)
- For **new players**, fetch season averages (from draft year to current year)
- For **existing players**, update only the latest season's data
- Upsert the records into your MongoDB collection: `nba-stats.players`

---

## How It Works

- Uses `ThreadPoolExecutor` to parallelize player updates
- Normalizes season averages into JSON format
- Efficiently updates only changed player data to reduce DB writes

---

## Customization

You can limit the number of players fetched by editing the call in `fetch_players()`:

```python
players = fetch_players(api, playersCollection, limit=100)
```

You can also change which seasons are fetched, or add logic for playoff stats.

---

## Output Sample

```bash
Starting player data fetch...
Fetching players with cursor: None...
Received 25 players from API.
New player 237 detected; fetching all season data...
Retrieved 13 valid seasons of data.
Inserted new player data for player 237.
```

---

## Dependencies

From `requirements.txt`:

- `balldontlie` — BallDontLie API SDK
- `python-dotenv` — loads your `.env` file into `os.environ`
- `pymongo` — MongoDB client for Python
- `dnspython` — DNS parsing for MongoDB Atlas URIs
- `certifi` — trusted CA bundle (used with TLS)

---

## Best Practices

- Do not commit your `.env` file
- Add `balldontlie/` and `__pycache__/` to `.gitignore` if present
- Run this script as a nightly cron job or scheduled container

---

## License

MIT — feel free to use, fork, and adapt.
