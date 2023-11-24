package routes

import (
	models "Backend/Models"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func InsertUser(db *sql.DB, username, password string) bool {

	sqlStatement := "INSERT INTO users (username, password, createdat, updatedat) VALUES ($1, $2, $3, $4);"

	hashpsw, err := HashPassword(password)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(sqlStatement, username, hashpsw, time.Now(), time.Now())
	if err != nil {
		return err == nil
	}
	return true
}

func HandlersReg(DBcon *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			panic(err)
		}

		if !InsertUser(DBcon, user.Username, user.Password) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		response := map[string]string{"register_successed": "Welcome to Website!"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	}
}

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
