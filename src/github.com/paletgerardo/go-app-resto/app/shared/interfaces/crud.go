package interfaces

import "net/http"

type Crud interface {
	Create(w http.ResponseWriter, r *http.Request) http.ResponseWriter
	Read(w http.ResponseWriter, r *http.Request) http.ResponseWriter
	Update(w http.ResponseWriter, r *http.Request) http.ResponseWriter
	Delete(w http.ResponseWriter, r *http.Request) http.ResponseWriter
	ReadAll(w http.ResponseWriter, r *http.Request) http.ResponseWriter
}
