package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(path string) (*Database, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	initializeSQLTable := ``

	_, err = db.Exec(initializeSQLTable)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil

}
