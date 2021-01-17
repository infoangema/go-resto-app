package handlers

import (
	"github.com/gorilla/mux"
	"github.com/paletgerardo/go-app-resto/app/features/categorias"
	"github.com/paletgerardo/go-app-resto/app/features/producto"
	"github.com/paletgerardo/go-app-resto/core/config"
	"github.com/paletgerardo/go-app-resto/core/interceptors"
	"github.com/paletgerardo/go-app-resto/core/jwt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func HandlerRequest() {
	config.LoadEnvironment()
	router := mux.NewRouter()

	// RUTAS
	router.HandleFunc("/productos/create", interceptors.InterceptToken(producto.Create)).Methods("POST")
	router.HandleFunc("/productos/get/{id}", interceptors.InterceptToken(producto.Read)).Methods("GET")
	router.HandleFunc("/productos/update", interceptors.InterceptToken(producto.Update)).Methods("PUT")
	router.HandleFunc("/productos/delete/{id}", interceptors.InterceptToken(producto.Delete)).Methods("DELETE")
	router.HandleFunc("/productos/get", interceptors.InterceptToken(producto.ReadAll)).Methods("GET")

	router.HandleFunc("/categorias/create", interceptors.InterceptToken(categorias.Create)).Methods("POST")
	router.HandleFunc("/categorias/get/{id}", interceptors.InterceptToken(categorias.Read)).Methods("GET")
	router.HandleFunc("/categorias/update", interceptors.InterceptToken(categorias.Update)).Methods("PUT")
	router.HandleFunc("/categorias/delete/{id}", interceptors.InterceptToken(categorias.Delete)).Methods("DELETE")
	router.HandleFunc("/categorias/get", interceptors.InterceptToken(categorias.ReadAll)).Methods("GET")

	router.HandleFunc("/login", jwt.Access).Methods("POST")

	// START APP
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
