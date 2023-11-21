package main

import (
	middleware "Backend/Handlers/Middleware"
	routes "Backend/Handlers/Routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/", routes.IndexHandler).Methods("GET")

	// Add CORS config
	corsHandler := middleware.ConfigCORS(r)
	
	//Serve HTTP at localhost:8080
	fmt.Println("Serve at port: 8080")
	http.ListenAndServe(":8080", corsHandler)
}
