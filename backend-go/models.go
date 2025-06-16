// models.go

package main

type Team struct {
	ID           int    `bson:"id" json:"id"`
	Conference   string `bson:"conference" json:"conference"`
	Division     string `bson:"division" json:"division"`
	City         string `bson:"city" json:"city"`
	Name         string `bson:"name" json:"name"`
	FullName     string `bson:"full_name" json:"full_name"`
	Abbreviation string `bson:"abbreviation" json:"abbreviation"`
}

type SeasonAverage struct {
	PlayerID    int     `bson:"player_id" json:"player_id"`
	Season      int     `bson:"season" json:"season"`
	GamesPlayed int     `bson:"games_played" json:"games_played"`
	Pts         float64 `bson:"pts" json:"pts"`
	Ast         float64 `bson:"ast" json:"ast"`
	Reb         float64 `bson:"reb" json:"reb"`
	Stl         float64 `bson:"stl" json:"stl"`
	Blk         float64 `bson:"blk" json:"blk"`
	Turnover    float64 `bson:"turnover" json:"turnover"`
	Min         string  `bson:"min" json:"min"`
	Fgm         float64 `bson:"fgm" json:"fgm"`
	Fga         float64 `bson:"fga" json:"fga"`
	FgPct       float64 `bson:"fg_pct" json:"fg_pct"`
	Fg3m        float64 `bson:"fg3m" json:"fg3m"`
	Fg3a        float64 `bson:"fg3a" json:"fg3a"`
	Fg3Pct      float64 `bson:"fg3_pct" json:"fg3_pct"`
	Ftm         float64 `bson:"ftm" json:"ftm"`
	Fta         float64 `bson:"fta" json:"fta"`
	FtPct       float64 `bson:"ft_pct" json:"ft_pct"`
	Oreb        float64 `bson:"oreb" json:"oreb"`
	Dreb        float64 `bson:"dreb" json:"dreb"`
}

type Player struct {
	ID             int             `bson:"id" json:"id"`
	College        *string         `bson:"college,omitempty" json:"college,omitempty"`
	Country        *string         `bson:"country,omitempty" json:"country,omitempty"`
	DraftNumber    *int            `bson:"draft_number,omitempty" json:"draft_number,omitempty"`
	DraftRound     *int            `bson:"draft_round,omitempty" json:"draft_round,omitempty"`
	DraftYear      *int            `bson:"draft_year,omitempty" json:"draft_year,omitempty"`
	FirstName      string          `bson:"first_name" json:"first_name"`
	Height         *string         `bson:"height,omitempty" json:"height,omitempty"`
	JerseyNumber   *string         `bson:"jersey_number,omitempty" json:"jersey_number,omitempty"`
	LastName       string          `bson:"last_name" json:"last_name"`
	Position       *string         `bson:"position,omitempty" json:"position,omitempty"`
	Team           Team            `bson:"team" json:"team"`
	TeamID         *int            `bson:"team_id,omitempty" json:"team_id,omitempty"`
	Weight         *string         `bson:"weight,omitempty" json:"weight,omitempty"`
	SeasonAverages []SeasonAverage `bson:"season_averages" json:"season_averages"`
}

type PlayerName struct {
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
}
