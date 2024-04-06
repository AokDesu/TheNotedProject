package routes

import (
	db "Backend/Databases"
	authpk "Backend/Handlers/Middleware/Auth"
	models "Backend/Models"
	"encoding/json"
	"fmt"
	"log"
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

func testGetAll(userId float64, username, role string, tags []string) (map[string][]interface{}, error) {

	DB, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	defer DB.Close()
	userIdInt := int(userId)
	sqlStatement := `select distinct notes.id, notes.title, notes.detail, notes.createdat, notes.updatedat 
	from note_categories inner join notes on note_categories.note_id = notes.id inner join users on users.id = notes.userid 
	WHERE users.id = $1;
	`
	var id int
	var create, update time.Time
	var title, detail string

	rows, err := DB.Query(sqlStatement, userIdInt)
	data := make(map[string][]interface{})
	data["User"] = append(data["User"], models.User{Username: username, Role: models.UserRole(role), CreatedAt: time.Now(), UpdatedAt: time.Now()})

	for rows.Next() {
		err := rows.Scan(&id, &title, &detail, &create, &update)
		if err != nil {
			fmt.Println("Select all error:", err)
			panic(err)
		}
		//data["Notes"] = append(data["Notes"], models.Note{Id: id, Title: title, Detail: detail, CreatedAt: create, UpdatedAt: update})
		queryCatetories := `select categories.name from note_categories inner join categories on note_categories.category_id = categories.id 
		WHERE note_categories.note_id = $1`
		query, err := DB.Query(queryCatetories, id)
		var categories []string
		var category string
		for query.Next() {
			err = query.Scan(&category)
			if err != nil {
				fmt.Println("Select all categories error:", err)
				panic(err)
			}
			categories = append(categories, category)
		}
		data["Notes"] = append(data["Notes"], Notes{Id: id, Title: title, Detail: detail, CreatedAt: create, UpdatedAt: update, Categories: categories})

	}
	fmt.Println(data)
	return data, nil
}

func testGetCategories(userId float64, username, role string, tags []string) (map[string][]interface{}, error) {
	var numstr string
	var Id int
	userIdInt := int(userId)
	data := make(map[string][]interface{})
	var err error
	for index, name := range tags {
		if index == len(tags)-1 {
			numstr += name
		} else {
			numstr += name + ","
		}
	}

	sqlGetNotebyCategories := `select distinct note_categories.note_id
   	from
	 note_categories
	 inner join categories on note_categories.category_id = categories.id
 	where
	 note_categories.category_id in ` + "(" + numstr + ");"
	DB, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	//fmt.Println(sqlGetNotebyCategories)

	data["User"] = append(data["User"], models.User{Username: username, Role: models.UserRole(role), CreatedAt: time.Now(), UpdatedAt: time.Now()})
	getNoteId, err := DB.Query(sqlGetNotebyCategories)

	for getNoteId.Next() {
		err := getNoteId.Scan(&Id)
		if err != nil {
			return nil, err
		}
		sqlGetNote := `select distinct notes.id, notes.title, notes.detail, notes.createdat, notes.updatedat 
		from note_categories inner join notes on note_categories.note_id = notes.id inner join users on users.id = notes.userid 
		WHERE users.id = $1 AND notes.id = $2;`
		getNote, err := DB.Query(sqlGetNote, userIdInt, Id)
		for getNote.Next() {
			var id int
			var create, update time.Time
			var title, detail string
			err := getNote.Scan(&id, &title, &detail, &create, &update)
			if err != nil {
				return nil, err
			}

			queryCatetories := `select categories.name from note_categories inner join categories on note_categories.category_id = categories.id 
		WHERE note_categories.note_id = $1`
			query, err := DB.Query(queryCatetories, id)
			var categories []string
			var category string
			for query.Next() {
				err = query.Scan(&category)
				if err != nil {
					fmt.Println("Select all categories error:", err)
					panic(err)
				}
				categories = append(categories, category)
			}
			data["Notes"] = append(data["Notes"], Notes{Id: id, Title: title, Detail: detail, CreatedAt: create, UpdatedAt: update, Categories: categories})

		}
	}
	// data["Test"] = append(data["test"], numstr)
	return data, err
}

type Notes struct {
	Id         int       `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Detail     string    `json:"detail,omitempty"`
	UserId     int       `json:"user_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	Categories []string  `json:"categories"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		queryParams := r.URL.Query()
		fmt.Println("QE", queryParams["category_id"])
		claims, ok := authpk.ClaimsToken(r.Header.Get("Authorization"))
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userId, username, role := claims["userId"].(float64), claims["username"].(string), claims["role"].(string)
		var data map[string][]interface{}
		var err error
		if len(queryParams["category_id"]) > 0 {
			data, err = testGetCategories(userId, username, role, queryParams["category_id"])
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println("error: ", err)
				return
			}
		} else {
			data, err = testGetAll(userId, username, role, queryParams["category_id"])
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println("error: ", err)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)

	case http.MethodPost:
		//InsertNote(w, r)
		testInsert(w, r)
	case http.MethodPut:
		testUpdateNote(w, r)
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
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer DB.Close()

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

// func testInsert(w http.ResponseWriter, r *http.Request) {
// 	queryParams := r.URL.Query()
// 	fmt.Println("params", queryParams["categories_name"])

// 	DB, err := db.InitDB()
// 	if err != nil {
// 		fmt.Println("error insert note", err)
// 		panic(err)
// 	}
// 	for _, name := range queryParams {

// 		sqlInsertNote := "INSERT INTO categories (name, createdat, updatedat) VALUES ($1,$2,$3);"
// 		_, err = DB.Exec(sqlInsertNote, name, time.Now(), time.Now())

// 		if err != nil {

// 		}
// 	}
// 	defer DB.Close()
// }

type Post struct {
	Title      string `json:"title"`
	Detail     string `json:"detail"`
	Categories []int  `json:"categories"`
}

func testUpdateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	DB, err := db.InitDB()
	sqlStatement := `update notes
	set title = $1, detail = $2, updatedat = $3
	where id = $4;			
	`
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	_, err = DB.Exec(sqlStatement, note.Title, note.Detail, time.Now(), note.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	fmt.Println("UPDATE route", note.Title)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"response": "Updated Yata desu ne~~"})
}

func testInsert(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	fmt.Println("Insert")
	claims, ok := authpk.ClaimsToken(r.Header.Get("Authorization"))
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		panic(err)
	}

	userId := claims["userId"]

	DB, err := db.InitDB()
	if err != nil {
		fmt.Println("error insert note", err)
		panic(err)
	}
	var insertedNoteID int
	sqlInsertPost := "INSERT INTO notes (title,detail,userid,createdat,updatedat) VALUES ($1, $2, $3, $4, $5) RETURNING id;"

	sqlInsertNoteCategories := "INSERT INTO note_categories (note_id ,category_id) VALUES ($1, $2);"

	err = DB.QueryRow(sqlInsertPost, post.Title, post.Detail, userId, time.Now(), time.Now()).Scan(&insertedNoteID)
	if err != nil {
		fmt.Println("Anyerror")
		panic(err)
	}
	for _, id := range post.Categories {
		_, err = DB.Exec(sqlInsertNoteCategories, insertedNoteID, id)
		if err != nil {
			panic(err)
		}
	}
	defer DB.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"response": "Yata desu ne~~"})

	fmt.Println(userId)
	fmt.Println(post.Title)
	fmt.Println(post.Detail)
	fmt.Println(post.Categories)
}
