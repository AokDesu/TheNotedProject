package authpk

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey string

func GenerateJWT(DB *sql.DB, username string) string {

	secretKey = os.Getenv("SECRET_KEY")
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
	claims["username"] = username
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

func ValidateToken(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		secretKey = os.Getenv("SECRET_KEY")

		getToken := r.Header.Get("Authorization")
		if len(getToken) < 7 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := getToken[7:]
		parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !parseToken.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		_, ok := parseToken.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	}

}

func ClaimsToken(getToken string) (jwt.MapClaims, bool) {
	secretKey = os.Getenv("SECRET_KEY")
	if len(getToken) < 7 {

		return nil, false
	}
	token := getToken[7:]
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, false
	}
	if !parseToken.Valid {
		return nil, false
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}
	return claims, true
}
