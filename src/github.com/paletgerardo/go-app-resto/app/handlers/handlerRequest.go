package handlers

import (
	"github.com/gorilla/mux"
	"github.com/paletgerardo/go-app-resto/app/features/categorias"
	"github.com/paletgerardo/go-app-resto/app/features/productos"
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
	router.HandleFunc("/productos/create", interceptors.InterceptToken(productos.Create)).Methods("POST")
	router.HandleFunc("/productos/get/{id}", interceptors.InterceptToken(productos.Read)).Methods("GET")
	router.HandleFunc("/productos/update", interceptors.InterceptToken(productos.Update)).Methods("PUT")
	router.HandleFunc("/productos/delete/{id}", interceptors.InterceptToken(productos.Delete)).Methods("DELETE")
	router.HandleFunc("/productos/get", interceptors.InterceptToken(productos.ReadAll)).Methods("GET")

	router.HandleFunc("/categorias/create", interceptors.InterceptToken(categorias.Create)).Methods("POST")
	router.HandleFunc("/categorias/get/{id}", interceptors.InterceptToken(categorias.Read)).Methods("GET")
	router.HandleFunc("/categorias/update", interceptors.InterceptToken(categorias.Update)).Methods("PUT")
	router.HandleFunc("/categorias/delete/{id}", interceptors.InterceptToken(categorias.Delete)).Methods("DELETE")
	router.HandleFunc("/categorias/get", interceptors.InterceptToken(categorias.ReadAll)).Methods("GET")

	router.HandleFunc("/login", jwt.Access).Methods("POST")

	//router.HandleFunc("/peticion/{id}", peticion).Methods("GET")

	// START APP
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func peticion(w http.ResponseWriter, r *http.Request) {
	reg := mux.Vars(r)["id"]
	/*miHttpClient := &http.Client{}
	url := "https://httpbin.org/get"
	peticion, erroEnPeticion := http.NewRequest("GET", url, nil)
	if erroEnPeticion != nil {
		http.Error(w, "error en la peticion", 400)
	}

	respuesta, errorEnRespuesta := miHttpClient.Do(peticion)
	if errorEnRespuesta != nil {
		http.Error(w, "error al intentar obtener la respuesta de la peticion", 400)

	}
	defer respuesta.Body.Close()

	respuestaByte, errorAlLeeresp := ioutil.ReadAll(respuesta.Body)
	if errorAlLeeresp != nil {
		http.Error(w, "error al intentar obtener la respuesta de la peticion", 400)
	}
	respuestaString := string(respuestaByte) + " reg: " + reg*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte(respuestaString))
	w.Write([]byte(reg))

}
