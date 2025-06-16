# NBA Stats API (Go + MongoDB)

A lightweight REST API written in Go for querying NBA player information and statistics. It uses the standard library (`net/http`) and connects to a **cloud-hosted MongoDB** instance (e.g., MongoDB Atlas).

This service is designed as a focused microservice within a future **all-in-one sports data platform**, intended to support multiple leagues and sports with modular, language-agnostic components. The NBA module serves as the foundation for scaling to other sports like NFL, MLB, NHL, and more.

---

## Features

- Fetch player data by name, ID, team, draft year, or position
- Get player season averages
- Use Docker and Docker Compose for local development
- Zero dependencies beyond Go and MongoDB

---

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/)
- A MongoDB URI (e.g., MongoDB Atlas)

---

## Project Structure

```
backend-go/
├── main.go
├── handlers.go
├── mongo.go
├── models.go
├── types.go
├── go.mod
├── go.sum
├── .env
├── Dockerfile
└── docker-compose.yml
```

---

## ⚙️ Environment Configuration

Create a `.env` file at the root of `backend-go/`:

```env
MONGO_DB_URI=REPLACE_WITH_YOUR_MONGODB_URI
```

---

## Docker Setup

### Dockerfile

```dockerfile
FROM golang:1.24.1 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app .

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates tzdata && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/app .
COPY .env .env

EXPOSE 8080
CMD ["./app"]
```

---

### docker-compose.yml

```yaml
version: '3.8'

services:
  backend:
    build:
      context: .
    container_name: nba-backend
    ports:
      - "8080:8080"
    env_file:
      - .env
```

---

### Start the API

```bash
docker-compose up --build
```

Your API will be running at:

```
http://localhost:8080
```

---

## API Endpoints

All endpoints return JSON. Replace any placeholder values as needed.

### Root

```http
GET /
```
Returns a plain "Hello" string.

---

### Get a Player by Name

```http
GET /player?first_name=LeBron&last_name=James
```
Returns full player data if found.

---

### Get Player by ID

```http
GET /player/id?id=23
```
Returns full player data by numeric ID.


---

### Get All Players

```http
GET /players
```
Returns an array of all players, optionally filtered by query parameters.


#### Optional query parameters:
- `team` — e.g. `Los Angeles Lakers`
- `position` — e.g. `G`, `F`
- `country` — e.g. `USA`
- `draft_year` — e.g. `2003`

**Example:**

```bash
curl "http://localhost:8080/players?team=Los Angeles Lakers&position=F"
```

---

### Get All Player Names

```http
GET /player-names
```
Returns a list of all player first and last names only.

---

### Get Players by Position

```http
GET /players/position?position=G
```
Returns players filtered by position (e.g. G, F, etc.).

---

### Get Players by Team Abbreviation

```http
GET /players/team?abbreviation=LAL
```
Returns players on a given team using the team abbreviation.

---

### Get Players by Draft Year

```http
GET /players/drafted?year=2003
```
Returns players drafted in the specified year.

---

### Get Season Averages by Player ID

```http
GET /averages?player_id=23
```
Returns an array of per-season average stats for the given player ID.

---

## Testing Locally with Curl

```bash
curl "http://localhost:8080/"
curl "http://localhost:8080/player?first_name=LeBron&last_name=James"
curl "http://localhost:8080/player/id?id=23"
curl "http://localhost:8080/players?position=G"
curl "http://localhost:8080/player-names"
curl "http://localhost:8080/players/drafted?year=2003"
curl "http://localhost:8080/averages?player_id=23"
```

---

## Stop & Cleanup

To stop the API:

```bash
docker-compose down
```

---

## Notes

- All logic is implemented using Go’s `net/http` (no external routers).
- Data is stored in a MongoDB collection named `players` inside the `nba-stats` database.
- Ensure your MongoDB URI has access and correct network rules.

---

## License

MIT — free to use, modify, and share.
