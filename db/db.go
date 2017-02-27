package db

import (
	"database/sql"
	"errors"
	"log"
	"regexp"

	"github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"
)

var (
	// ErrDatabaseNotInitialized is raised on begin transaction but db has not been initialized
	ErrDatabaseNotInitialized = errors.New("database has not benn initialized")
)

var db *sql.DB // global database connection

// Tx is a wrapper of sql.Tx
type Tx struct {
	Commited bool
	*sql.Tx
}

// Init initializes database connection
func Init(dsn string) (err error) {
	if dsn == "" {
		dsn, err = configureDSN()
		if err != nil {
			return err
		}
	}
	log.Printf("Connect to DB (%s) ...", dsn)
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		return err
	}
	db = conn
	return nil
}

func configureDSN() (string, error) {
	c := mysql.Config{
		User:      "root",
		DBName:    "noctf",
		Addr:      "localhost",
		Net:       "tcp",
		ParseTime: true,
	}
	portSuffix, err := regexp.Compile(`:\d+$`)
	if err != nil {
		return "", err
	}
	if !portSuffix.MatchString(c.Addr) {
		c.Addr = c.Addr + ":3306"
	}
	return c.FormatDSN(), nil
}

// Begin returns database transaction
func Begin() (*Tx, error) {
	if db == nil {
		return nil, ErrDatabaseNotInitialized
	}
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{Tx:tx, Commited: false}, nil
}

// Commit commits the transaction
func (tx *Tx) Commit() error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	tx.Commited = true
	return nil
}

// Rollback rollbacks the transaction if the transaction has not been commited
func (tx *Tx) Rollback() error {
	err := tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}
