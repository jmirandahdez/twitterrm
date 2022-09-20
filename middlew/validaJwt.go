package middlew

import (
	"net/http"

	"github.com/jmirandahdez/twitterrm/routers"
)

func ValidaJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesaToken(r.Header.Get("Autorizathion"))

		if err != nil {
			http.Error(w, "Error en el tocket"+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
