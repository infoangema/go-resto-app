package producto

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ValidationDataAndParseToJson(w http.ResponseWriter, r *http.Request) (Producto, error) {
	var p Producto

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		return p, errors.New("Tipo de datos incorrectos.")
	}

	if len(p.Nombre) < 3 {
		return p, errors.New("El nombre del producto debe tener 3 o mas caracteres")
	}
	return p, nil
}
