package productos

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request) {

	productoAguardar, errorEnDatosDeEntrada := AcaSeValidanLosDatosDeEntrada(r)
	if errorEnDatosDeEntrada != nil {
		http.Error(w, "Error: "+errorEnDatosDeEntrada.Error(), 400)
		return
	}

	productoGuardado, errorAlGuardarProducto := AcaSeGuardaElProducto(productoAguardar)
	if errorAlGuardarProducto != nil {
		http.Error(w, "Error: "+errorAlGuardarProducto.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(productoGuardado)

}

func Read(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID invalido: "+err.Error(), 400)
		return
	}

	producto, errorAlBuscar := BuscarUnProductoPorId(idInt)
	if errorAlBuscar != nil {
		http.Error(w, errorAlBuscar.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(producto)

}

func Update(w http.ResponseWriter, r *http.Request) {

	productoAguardar, errorEnDatosDeEntrada := AcaSeValidanLosDatosDeEntrada(r)
	if errorEnDatosDeEntrada != nil {
		http.Error(w, "Error: "+errorEnDatosDeEntrada.Error(), 400)
		return
	}

	productoActualizado, errorAlGuardarProducto := AcaSeActualizaElProducto(productoAguardar)
	if errorAlGuardarProducto != nil {
		http.Error(w, "Error: "+errorAlGuardarProducto.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(productoActualizado)

}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID invalido: "+err.Error(), 400)
		return
	}

	errorAlBorrar := BorrarUnProductoPorId(idInt)
	if errorAlBorrar != nil {
		http.Error(w, "Error al borrar el producto. ", 400)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func ReadAll(w http.ResponseWriter, r *http.Request) {

	productos, err := BuscarTodosLosProducto()
	if err != nil {
		http.Error(w, "Error al obtener los productos. ", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(productos)

}
