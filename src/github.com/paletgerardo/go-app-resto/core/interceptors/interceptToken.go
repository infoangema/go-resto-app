package interceptors

import (
	"github.com/paletgerardo/go-app-resto/core/jwt"
	"net/http"
)

func InterceptToken(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_, _, errorAlValidarToken := jwt.TokenValidator(r.Header.Get("Authorization"))
		if errorAlValidarToken != nil {
			http.Error(w, errorAlValidarToken.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}

}
