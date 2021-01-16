package categorias

import (
	"time"
)

type Categoria struct {
	Id          int       `json:"id"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Activo      bool      `json:"activo"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}
