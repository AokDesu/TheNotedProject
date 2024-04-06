package db

import (
	"database/sql"
)

func CreateNoteCategories(db *sql.DB) error {
	sqlStatement := `
	CREATE TABLE IF NOT EXISTS note_categories (
		note_id INTEGER REFERENCES notes(Id),
		category_id INTEGER REFERENCES categories(Id),
		PRIMARY KEY (note_id, category_id)
	);
	`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}
