package jwt

import (
	"encoding/json"
	"net/http"
)

func Access(w http.ResponseWriter, r *http.Request) {

	// 1- validar datos del usuarioDeAcceso.
	usuarioDeAcceso, errorEnLosDatosDeEntrada := AcaSeValidanLosDatosDeEntrada(r)
	if errorEnLosDatosDeEntrada != nil {
		http.Error(w, "Error en los datos de entrada", 400)
		return
	}

	// 2- buscar si el usuarioDeAcceso esta en la base de datos y devolver usuario.
	usuario, _, errorAlBuscarUsuario := AcaSeBuscaElUsuario(usuarioDeAcceso.Email)
	if errorAlBuscarUsuario != nil {
		http.Error(w, "Error al buscar usuario: "+errorAlBuscarUsuario.Error(), 400)
		return
	}

	// 3- generar JWT con los datos del usuario.
	tokenGenerado, errorAlCrearToken := AcaSeGeneraElToken(usuario)
	if errorAlCrearToken != nil {
		http.Error(w, "Error al crear token", 400)
		return
	}

	// 4- devolver respuesta: setear header con json y el token.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LoginToken{
		Token: tokenGenerado,
	})
}
