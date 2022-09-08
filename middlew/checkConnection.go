package middlew

import (
	"net/http"

	"github.com/jmirandahdez/twitterrm/bd"
)

// Middlew para POST
func ChekDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
