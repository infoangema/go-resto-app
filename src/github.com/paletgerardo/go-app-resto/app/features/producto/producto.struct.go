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
	Created_at  time.Time `json:"CreatedAt"`
	Updated_at  time.Time `json:"UpdatedAt"`
	CategoriaId int       `json:"categoriaId"`
}
