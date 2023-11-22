package main

import (
	db "Backend/Databases"
	middleware "Backend/Handlers/Middleware"
	routes "Backend/Handlers/Routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Connect to Database
	db, err := db.InitDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	// Router
	r := mux.NewRouter()
	r.HandleFunc("/", routes.IndexHandler).Methods("GET")

	// Add CORS config
	corsHandler := middleware.ConfigCORS(r)

	//Serve HTTP at localhost:8080
	fmt.Println("Serve at port: 8080!")
	fmt.Println("Connect to postgresql successed!")
	http.ListenAndServe(":8080", corsHandler)
}
