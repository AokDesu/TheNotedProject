package db

import (
	"database/sql"
)

func CreateNoteTable(db *sql.DB) error {
	sqlStatement := `
	CREATE TABLE IF NOT EXISTS notes (
		Id SERIAL PRIMARY KEY,
		Title TEXT NOT NULL,
		Detail TEXT,
		UserId INT,
		CreatedAt timestamp NOT NULL,
		UpdatedAt timestamp,
		FOREIGN KEY (UserId) REFERENCES users(Id)
	);
`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}
