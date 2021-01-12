package handlers

import (
	"github.com/gorilla/mux"
	"github.com/paletgerardo/go-app-resto/app/features/categorias"
	"github.com/paletgerardo/go-app-resto/app/features/producto"
	"github.com/paletgerardo/go-app-resto/core/config"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func HandlerRequest() {
	config.LoadEnvironment()
	router := mux.NewRouter()

	// RUTAS
	router.HandleFunc("/productos/create", producto.Create).Methods("POST")
	router.HandleFunc("/productos/get/{id}", producto.Read).Methods("GET")
	router.HandleFunc("/productos/update", producto.Update).Methods("PUT")
	router.HandleFunc("/productos/delete/{id}", producto.Delete).Methods("DELETE")
	router.HandleFunc("/productos/get", producto.ReadAll).Methods("GET")

	router.HandleFunc("/categorias/create", categorias.Create).Methods("POST")
	router.HandleFunc("/categorias/get/{id}", categorias.Read).Methods("GET")
	router.HandleFunc("/categorias/update", categorias.Update).Methods("PUT")
	router.HandleFunc("/categorias/delete/{id}", categorias.Delete).Methods("DELETE")
	router.HandleFunc("/categorias/get", categorias.ReadAll).Methods("GET")

	// START APP
	port := os.Getenv("PORT")
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
