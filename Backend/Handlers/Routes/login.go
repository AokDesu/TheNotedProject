package routes

import (
	authpk "Backend/Handlers/Middleware/Auth"
	models "Backend/Models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func HandlerLogin(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !authpk.ValidateUser(db, user.Username, user.Password) {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		token := authpk.GenerateJWT(db, user.Username)
		if token == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		response := map[string]string{"token": token}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}
