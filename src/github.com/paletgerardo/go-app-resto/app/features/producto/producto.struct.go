package producto

import (
	"time"
)

type Producto struct {
	Id          int       `json:"id"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Precio      float32   `json:"precio"`
	Activo      bool      `json:"activo"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	CategoriaId int       `json:"categoriaId"`
}
