package db

import (
	"database/sql"
	"time"
)

// Config for Database connection
type Config struct {
	DSN string
}

// Player in database
// db
type Player struct {
	ID        int
	Name      string
	Email     sql.NullString
	TeamID    string
	Password  string
	CreatedOn time.Time
}

// Problem in database
// db
type Problem struct {
	ID        int
	Name      string
	Statement string
	Score     int
	IsFlag    bool
	Flag      sql.NullString
}

// Team in database
// db
type Team struct {
	ID   int
	Name string
}

// Scanner interface
type Scanner interface {
	Scan(...interface{}) error
}
