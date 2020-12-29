package producto

import (
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	p, err := ValidationDataAndParseToJson(w, r)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}
	err = CreateProducto(p)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	DeleteProducto(w, r)
}

func FinfById(w http.ResponseWriter, r *http.Request) {
	FindByIdProducto(w, r)
}
