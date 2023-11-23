package main

import (
	dbp "Backend/Databases"
	middleware "Backend/Handlers/Middleware"
	authpk "Backend/Handlers/Middleware/Auth"
	routes "Backend/Handlers/Routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to Database
	db, err := dbp.InitDB()
	if err != nil {
		panic(err)
	}
	// Check database connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	// Router
	r := mux.NewRouter()
	r.HandleFunc("/", authpk.ValidateToken(routes.IndexHandler)).Methods("GET")
	r.HandleFunc("/api/v1/register", routes.HandlersReg(db)).Methods("POST")
	r.HandleFunc("/api/v1/login", routes.HandlerLogin(db)).Methods("POST")

	// Add CORS config
	corsHandler := middleware.ConfigCORS(r)
	// Serve HTTP at localhost:8080
	fmt.Println("Serve at port: 8080!")
	fmt.Println("Connect to postgresql successed!")
	http.ListenAndServe(":8080", corsHandler)
}
