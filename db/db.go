package db

import (
	"database/sql"
)

var DB *sql.DB

func Connect(provider string, dsn string) error {
	database, err := sql.Open(provider, dsn)
	DB = database
	return err
}
