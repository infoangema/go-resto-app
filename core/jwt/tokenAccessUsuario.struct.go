package jwt

type UsuarioDeAcceso struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
