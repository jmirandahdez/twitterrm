package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jmirandahdez/twitterrm/bd"
	"github.com/jmirandahdez/twitterrm/models"
)

var Email string
var IdUsuario string

func ProcesaToken(tk string) (*models.Claim, bool, string, error) {
	miclave := []byte("MastersdelDesarrollo_gurpodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	// Esta linea es la que valida el token
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miclave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChecaSiExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IdUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IdUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
