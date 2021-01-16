package jwt

import (
	"encoding/json"
	"errors"
	"net/http"
)

func AcaSeValidanLosDatosDeEntrada(r *http.Request) (UsuarioDeAcceso, error) {

	var usuarioLogin UsuarioDeAcceso

	errorAlParsearLosDatos := json.NewDecoder(r.Body).Decode(&usuarioLogin)
	if errorAlParsearLosDatos != nil {
		return usuarioLogin, errors.New("Error al parsear los datos.")
	}

	if len(usuarioLogin.Email) < 6 || len(usuarioLogin.Password) < 6 {
		return usuarioLogin, errors.New("Error en el email o password.")
	}

	return usuarioLogin, nil
}
