package categorias

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request) {

	categoriaParaGuardar, errorEnDatosDeEntrada := AcaSeValidanLosDatosDeEntrada(r)
	if errorEnDatosDeEntrada != nil {
		http.Error(w, "Error al validar los datos de entrada: "+errorEnDatosDeEntrada.Error(), 400)
		return
	}

	errorAlGuardar := AcaSeGuardaLaCategoria(categoriaParaGuardar)
	if errorAlGuardar != nil {
		http.Error(w, "Error al guardar la categoria: "+errorAlGuardar.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func Read(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Id invalido: "+err.Error(), 400)
		return
	}

	categoria, errorAlBuscar := BuscarUnaCategoriaPorId(idInt)
	if errorAlBuscar != nil {
		http.Error(w, "Categoria no encontrada: "+errorAlBuscar.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categoria)

}

func Update(w http.ResponseWriter, r *http.Request) {

	categoriaAguardar, errorEnDatosDeEntrada := AcaSeValidanLosDatosDeEntrada(r)
	if errorEnDatosDeEntrada != nil {
		http.Error(w, "Error: "+errorEnDatosDeEntrada.Error(), 400)
		return
	}

	errorAlGuardarCategoria := AcaSeActualizaLaCategoria(categoriaAguardar)
	if errorAlGuardarCategoria != nil {
		http.Error(w, "Error: "+errorAlGuardarCategoria.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Id invalido: "+err.Error(), 400)
		return
	}

	errorAlBorrar := BorrarUnaCategoriaPorId(idInt)
	if errorAlBorrar != nil {
		http.Error(w, "Error al borrar la categoria. ", 400)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func ReadAll(w http.ResponseWriter, r *http.Request) {

	categorias, err := BuscarTodasLasCategorias()
	if err != nil {
		http.Error(w, "Error al obtener los categorias. ", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categorias)

}
