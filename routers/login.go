package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jmirandahdez/twitterrm/bd"
	"github.com/jmirandahdez/twitterrm/jwt"
	"github.com/jmirandahdez/twitterrm/models"
)

// Logeo a la app
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o contrasena invalido"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentaLogin(t.Email, t.Password)

	if !existe {
		http.Error(w, "Usuario y/o contrasena invalido", 400)
		return
	}

	jwtkey, err := jwt.GeneraJwt(documento)
	if err != nil {
		http.Error(w, "Ocurrio un errror al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtkey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expiratioTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expiratioTime,
	})
}
