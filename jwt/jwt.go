package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jmirandahdez/twitterrm/models"
)

func GeneraJwt(t models.Usuario) (string, error) {
	miclave := []byte("MastersdelDesarrollo_gurpodeFacebook")

	paylod := jwt.MapClaims{
		"email":          t.Email,
		"nombre":         t.Nombre,
		"apellidos":      t.Apellidos,
		"fechaacimiento": t.FechaNacimiento,
		"biografia":      t.Biografia,
		"ubicacion":      t.Ubicacion,
		"sitioWeb":       t.SitioWeb,
		"_id":            t.ID.Hex(),
		"exp":            time.Now().Add(time.Hour * 24).Unix(),
	}

	// Se pasan los valores a cifrar y algoritmo de cifrado
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, paylod)

	// Se indica la clave privada con la que se va a cifrar
	tokenStr, err := token.SignedString(miclave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
