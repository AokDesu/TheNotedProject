package db

import (
	"database/sql"
)

func CreateUserTable(db *sql.DB) error {
	sqlStatement := `
	DO $$
		BEGIN
    	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'userrole') THEN
        	CREATE TYPE UserRole AS ENUM ('admin', 'member');
    	END IF;
	END $$;
	CREATE TABLE IF NOT EXISTS users (
		Id SERIAL PRIMARY KEY,
		Username TEXT NOT NULL,
		Password TEXT NOT NULL,
		Role UserRole DEFAULT 'member',
		CreatedAt timestamp NOT NULL,
		UpdatedAt timestamp
	);
`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}
