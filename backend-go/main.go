// main.go

package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	InitMongo(ctx)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})
	http.HandleFunc("/player", FetchPlayer)
	http.HandleFunc("/players", GetPlayers)
	http.HandleFunc("/player-names", FetchPlayerNames)
	http.HandleFunc("/player/id", GetPlayerByID)
	http.HandleFunc("/players/position", GetPlayersByPosition)
	http.HandleFunc("/players/team", GetPlayersByTeam)
	http.HandleFunc("/players/drafted", GetPlayersByDraftYear)
	http.HandleFunc("/averages", GetSeasonAverages)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server Failed: ", err)
	}
}
