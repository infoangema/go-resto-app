package producto

import (
	"encoding/json"
	"errors"
	"net/http"
)

/* Aca se hacen todas las validaciones del producto, y devuelve el producto model */
func AcaSeValidanLosDatosDeEntrada(r *http.Request) (Producto, error) {
	var producto Producto

	errorAlParsear := json.NewDecoder(r.Body).Decode(&producto)
	if errorAlParsear != nil {
		return producto, errors.New("Tipo de datos incorrectos.")
	}

	if len(producto.Nombre) < 3 {
		return producto, errors.New("El nombre del producto debe tener 3 o mas caracteres")
	}

	return producto, nil
}
