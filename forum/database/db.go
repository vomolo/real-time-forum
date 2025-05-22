package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func InitDB(filepath string) (*DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
