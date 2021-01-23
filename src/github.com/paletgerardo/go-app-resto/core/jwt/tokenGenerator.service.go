package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/paletgerardo/go-app-resto/core/structs"
	"os"
	"time"
)

func AcaSeGeneraElToken(usuario structs.UserLogin) (string, error) {

	firma := []byte(os.Getenv("CLAVETOKEN"))
	payload := jwt.MapClaims{
		"id":              usuario.Id,
		"nombre":          usuario.Nombre,
		"apellido":        usuario.Apellido,
		"fechaNacimiento": usuario.FechaNacimiento,
		"email":           usuario.Email,
		"exp":             time.Now().Add(48 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenstr, err := token.SignedString(firma)
	if err != nil {
		return tokenstr, err
	}

	return tokenstr, nil
}
