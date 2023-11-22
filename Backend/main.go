package main

import (
	dbp "Backend/Databases"
	middleware "Backend/Handlers/Middleware"
	routes "Backend/Handlers/Routes"
	models "Backend/Models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Connect to Database
	db, err := dbp.InitDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	// Router
	r := mux.NewRouter()
	r.HandleFunc("/", routes.IndexHandler).Methods("GET")
	r.HandleFunc("/api/v1/register", routes.HandlersReg(db)).Methods("POST")

	// Add CORS config
	corsHandler := middleware.ConfigCORS(r)

	// sql := "SELECT password FROM users"
	// var user models.User
	// rows, err := db.Query(sql)
	// if err != nil {
	// 	panic(err)
	// }
	// for rows.Next() {
	// 	if err := rows.Scan(&user.Password); err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println(user.Password)
	// if routes.CheckPasswordHash(user.Password, "Secret-ultrasafepassword123") {
	// 	fmt.Println("Match.")
	// } else {
	// 	fmt.Println("Doesn't match.")
	// }


	//Serve HTTP at localhost:8080
	fmt.Println("Serve at port: 8080!")
	fmt.Println("Connect to postgresql successed!")
	http.ListenAndServe(":8080", corsHandler)
}
