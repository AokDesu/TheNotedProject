package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func ConfigCORS(next http.Handler) http.Handler {

	return handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(next)
}
