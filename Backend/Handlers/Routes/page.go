package routes

import (
	db "Backend/Databases"
	authpk "Backend/Handlers/Middleware/Auth"
	models "Backend/Models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetNote(userId float64, username, role string) (map[string][]interface{}, error) {
	DB, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	defer DB.Close()

	sqlStatement := `
		SELECT notes.id, notes.title, notes.detail, notes.createdat, notes.updatedat
		FROM notes 
		WHERE notes.userid = $1
	`

	userIdInt := int(userId)

	rows, err := DB.Query(sqlStatement, userIdInt)
	if err != nil {
		fmt.Println("Err query")
		return nil, err
	}
	var id int
	var create, update time.Time
	var title, detail string
	data := make(map[string][]interface{})

	data["User"] = append(data["User"], models.User{Username: username, Role: models.UserRole(role), CreatedAt: time.Now(), UpdatedAt: time.Now()})
	for rows.Next() {
		err := rows.Scan(&id, &title, &detail, &create, &update)
		if err != nil {

			return nil, err
		}
		data["Notes"] = append(data["Notes"], models.Note{Id: id, Title: title, Detail: detail, CreatedAt: create, UpdatedAt: update})
	}
	fmt.Println(data)
	return data, nil

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		claims, ok := authpk.ClaimsToken(r.Header.Get("Authorization"))
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userId, username, role := claims["userId"].(float64), claims["username"].(string), claims["role"].(string)

		data, err := GetNote(userId, username, role)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("error: ", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)

	case http.MethodPost:
		InsertNote(w, r)
	}

}

func InsertNote(w http.ResponseWriter, r *http.Request) {

	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	claims, ok := authpk.ClaimsToken(r.Header.Get("Authorization"))
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userId := claims["userId"]

	DB, err := db.InitDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sqlStatement := "INSERT INTO notes (title,detail,userid,createdat,updatedat) VALUES ($1, $2, $3, $4, $5);"

	_, err = DB.Exec(sqlStatement, note.Title, note.Detail, userId, time.Now(), time.Now())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"response": "Yata desu ne~~"})

}
