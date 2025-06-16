// handlers.go

package main

import (
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FetchPlayer(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	if firstName == "" || lastName == "" {
		jsonResponse(w, http.StatusBadRequest, ErrorResponse{
			Error: "firstName and/or lastName field is empty",
		})
		return
	}

	filter := bson.M{"first_name": firstName, "last_name": lastName}
	var player Player
	err := playersColl.FindOne(r.Context(), filter).Decode(&player)

	if err != nil {
		jsonResponse(w, http.StatusNotFound, ErrorResponse{
			Error: "could not find player in database",
		})
		return
	}

	jsonResponse(w, http.StatusAccepted, player)
}

func FetchPlayerNames(w http.ResponseWriter, r *http.Request) {
	filter := bson.M{}
	projection := bson.M{"first_name": 1, "last_name": 1, "_id": 0}
	opts := options.Find().SetProjection(projection)

	cursor, err := playersColl.Find(r.Context(), filter, opts)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorResponse{
			Error: "Error fetching player names",
		})
		return
	}
	defer cursor.Close(r.Context())

	var names []PlayerName
	for cursor.Next(r.Context()) {
		var p PlayerName
		if err := cursor.Decode(&p); err != nil {
			jsonResponse(w, http.StatusInternalServerError, ErrorResponse{
				Error: "Error decoding player name",
			})
			return
		}
		names = append(names, p)
	}

	jsonResponse(w, http.StatusOK, names)
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := bson.M{}

	if team := query.Get("team"); team != "" {
		filter["team.full_name"] = team
	}
	if position := query.Get("position"); position != "" {
		filter["position"] = position
	}
	if country := query.Get("country"); country != "" {
		filter["country"] = country
	}
	if draftYear := query.Get("draft_year"); draftYear != "" {
		filter["draft_year"] = draftYear
	}

	cursor, err := playersColl.Find(r.Context(), filter)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorResponse{"Error fetching players"})
		return
	}
	defer cursor.Close(r.Context())

	var players []Player
	if err := cursor.All(r.Context(), &players); err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorResponse{"Error decoding players"})
		return
	}

	jsonResponse(w, http.StatusOK, players)
}

func GetPlayerByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, ErrorResponse{"Invalid player ID"})
		return
	}

	var player Player
	filter := bson.M{"id": id}
	err = playersColl.FindOne(r.Context(), filter).Decode(&player)
	if err != nil {
		jsonResponse(w, http.StatusNotFound, ErrorResponse{"Player not found"})
		return
	}

	jsonResponse(w, http.StatusOK, player)
}

func GetPlayersByPosition(w http.ResponseWriter, r *http.Request) {
	pos := r.URL.Query().Get("position")
	if pos == "" {
		jsonResponse(w, http.StatusBadRequest, ErrorResponse{"Missing position query parameter"})
		return
	}

	filter := bson.M{"position": pos}
	cursor, err := playersColl.Find(r.Context(), filter)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorResponse{"Error fetching players"})
		return
	}
	defer cursor.Close(r.Context())

	var players []Player
	if err := cursor.All(r.Context(), &players); err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorResponse{"Error decoding players"})
		return
	}

	jsonResponse(w, http.StatusOK, players)
}

func GetPlayersByTeam(w http.ResponseWriter, r *http.Request) {
	abbr := r.URL.Query().Get("abbreviation")
	if abbr == "" {
		jsonResponse(w, http.StatusBadRequest, ErrorResponse{"Missing abbreviation query parameter"})
		return
	}

	filter := bson.M{"team.abbreviation": abbr}
	cursor, err := playersColl.Find(r.Context(), filter)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorResponse{"Error fetching players"})
		return
	}
	defer cursor.Close(r.Context())

	var players []Player
	if err := cursor.All(r.Context(), &players); err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorResponse{"Error decoding players"})
		return
	}

	jsonResponse(w, http.StatusOK, players)
}

func GetPlayersByDraftYear(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, ErrorResponse{"Invalid draft year"})
		return
	}

	filter := bson.M{"draft_year": year}
	cursor, err := playersColl.Find(r.Context(), filter)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorResponse{"Error fetching players"})
		return
	}
	defer cursor.Close(r.Context())

	var players []Player
	if err := cursor.All(r.Context(), &players); err != nil {
		jsonResponse(w, http.StatusInternalServerError, ErrorResponse{"Error decoding players"})
		return
	}

	jsonResponse(w, http.StatusOK, players)
}

func GetSeasonAverages(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("player_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, ErrorResponse{"Invalid player ID"})
		return
	}

	var player Player
	filter := bson.M{"id": id}
	err = playersColl.FindOne(r.Context(), filter).Decode(&player)
	if err != nil {
		jsonResponse(w, http.StatusNotFound, ErrorResponse{"Player not found"})
		return
	}

	jsonResponse(w, http.StatusOK, player.SeasonAverages)
}
