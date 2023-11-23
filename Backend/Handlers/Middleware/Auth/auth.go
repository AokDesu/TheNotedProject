package authpk

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = "daskdosakdoakodakovcs1321"

func GenerateJWT(DB *sql.DB, username string) string {
	sqlStatement := "select id, role from users where username = $1"

	var id int
	var role string

	rows, err := DB.Query(sqlStatement, username)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &role)
		if err != nil {
			panic(err)
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	claims["userId"] = id
	claims["role"] = role

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	return tokenString
}

func ValidateUser(DB *sql.DB, username, password string) bool {
	sqlStatement := "select (password) from users where username = $1"

	rows, err := DB.Query(sqlStatement, username)
	if err != nil {
		return false
	}

	var hashPassword string
	for rows.Next() {
		if err := rows.Scan(&hashPassword); err != nil {
			return false
		}
	}

	return CheckPasswordHash(hashPassword, password)
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
