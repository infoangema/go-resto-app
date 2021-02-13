package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strings"
)

var Email string
var IdUsuario int

func TokenValidator(tk string) (*Claim, int, error) {

	clave := []byte(os.Getenv("CLAVETOKEN"))
	claims := &Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, -1, errors.New("Formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, errorAlParsearClaims := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return clave, nil
	})

	if errorAlParsearClaims != nil {
		return claims, -1, errors.New("Token invalido: " + (errorAlParsearClaims.Error()))
	}

	if tkn.Valid {
		_, id, errorAlbuscarUsuario := AcaSeBuscaElUsuario(claims.Email)
		if errorAlbuscarUsuario == nil {
			Email = claims.Email
			IdUsuario = id
		}
		return claims, IdUsuario, nil
	}

	return claims, IdUsuario, errorAlParsearClaims
}
