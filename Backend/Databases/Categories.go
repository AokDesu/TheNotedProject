package db

import (
	"database/sql"
)

func CreateCategoriesTable(db *sql.DB) error {
	sqlStatement := `
	CREATE TABLE IF NOT EXISTS categories (
		Id SERIAL PRIMARY KEY,
		Name TEXT NOT NULL UNIQUE,
		UserId INT,
		CreatedAt timestamp,
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
