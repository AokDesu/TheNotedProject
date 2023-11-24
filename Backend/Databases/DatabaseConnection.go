package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	databasePASSWORD := os.Getenv("DATABASE_PASSWORD")

	db, err := sql.Open("postgres", fmt.Sprintf("user=postgres password=%s host=%s port=5432 dbname=postgres", databasePASSWORD, databaseURL))
	if err != nil {
		return nil, err
	}

	err = CreateUserTable(db)
	if err != nil {
		return nil, err
	}

	err = CreateNoteTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
