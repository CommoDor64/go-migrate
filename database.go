package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// GetDB returns the db according to connection uri
func GetDB(connURI string) *sql.DB {
	db, err := sql.Open("postgres", connURI)
	if err != nil {
		panic(err)
	}

	return db
}
