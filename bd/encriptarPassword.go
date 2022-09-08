package bd

import (
	"golang.org/x/crypto/bcrypt"
)

// Metodo para encriptar password
func EncriptarPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	return string(bytes), err
}
