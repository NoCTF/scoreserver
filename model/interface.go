package model

import "database/sql"

// Player type
// model
type Player struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Email    sql.NullString `json:"email"`
	TeamID   string         `json:"team_id"`
	Password string         `json:"password"`
}

// Players is list of players
type Players []Player

// Problem type
// model
type Problem struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	IsFlag    bool   `json:"is_flag"`
	Flag      string `json:"flag"`
	Statement string `json:"statement"`
	Score     int    `json:"score"`
}

// Problems is list of problems
type Problems []Problem

// Team type
// model
type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Teams is list of teams
type Teams []Team
