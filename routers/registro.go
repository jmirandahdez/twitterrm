package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jmirandahdez/twitterrm/bd"
	"github.com/jmirandahdez/twitterrm/models"
)

// Funcion para crear en BD el registro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en parametros recibidos: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password requiere al menos 6 caracteres", 400)
		return
	}

	// esta funcion devuelve 3 valores y solo ocupo en esta parte el segundo valor
	_, encontrado, _ := bd.ChecaSiExisteUsuario(t.Email)

	if encontrado {
		http.Error(w, "El email ya existe", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Error al guardar el usuario: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se guardo el usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
