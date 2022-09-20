package bd

import (
	"github.com/jmirandahdez/twitterrm/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentaLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChecaSiExisteUsuario(email)

	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}

	return usu, true
}
