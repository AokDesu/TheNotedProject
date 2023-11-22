package routes

import (
	models "Backend/Models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func InsertUser(db *sql.DB, username, password string) {

	sqlStatement := "INSERT INTO users (username, password, createdat) VALUES ($1, $2, $3);"

	hashpsw, err := HashPassword(password)
	if err != nil {
		panic(err)
	}
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(hashpsw)
	result, err := db.Exec(sqlStatement, username, hashpsw, time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println("Inseted user: ", result)
}

func HandlersReg(DBcon *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			panic(err)
		}
		InsertUser(DBcon, user.Username, user.Password)
	}
}


func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
