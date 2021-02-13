package categorias

import (
	"encoding/json"
	"errors"
	"net/http"
)

func AcaSeValidanLosDatosDeEntrada(r *http.Request) (Categoria, error) {

	var categoria Categoria
	errorAlPardear := json.NewDecoder(r.Body).Decode(&categoria)
	if errorAlPardear != nil {
		return categoria, errors.New("Error al verificar los datos de entrada.")
	}

	if len(categoria.Nombre) < 3 {
		return categoria, errors.New("El nombre de la categoria no puede ser menor a 3 caracteres")
	}

	return categoria, nil
}
