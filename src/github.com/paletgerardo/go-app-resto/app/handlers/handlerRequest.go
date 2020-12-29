package handlers

import (
	"github.com/gorilla/mux"
	"github.com/paletgerardo/go-app-resto/app/features/producto"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func HandlerRequest() {
	router := mux.NewRouter()

	// RUTAS
	router.HandleFunc("/productos/create", producto.Create).Methods("POST")
	router.HandleFunc("/productos/read", producto.Create).Methods("GET")
	router.HandleFunc("/productos/delete", producto.Delete).Methods("DELETE")
	router.HandleFunc("/productos/get", producto.FinfById).Methods("GET")
	router.HandleFunc("/productos/get/{id}", producto.FinfById).Methods("GET")

	// START APP
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
