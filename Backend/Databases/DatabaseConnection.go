package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=bkokdoa0weqw31-i9!#@sdag1215 host=db.whcdsgtfrjzmcxcrrfca.supabase.co port=5432 dbname=postgres")

	if err != nil {
		return nil, err
	}
	err = db.Ping()
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
