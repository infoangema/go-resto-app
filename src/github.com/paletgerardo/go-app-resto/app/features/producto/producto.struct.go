package producto

import (
	"github.com/paletgerardo/go-app-resto/core/structs"
)

type Producto struct {
	Id          int              `json:"id"`
	Nombre      string           `json:"nombre"`
	Descripcion string           `json:"descripcion"`
	Precio      float32          `json:"precio"`
	Activo      bool             `json:"activo"`
	Created     structs.NullTime `json:"created"`
	Updated     structs.NullTime `json:"updated"`
	CategoriaId int              `json:"categoriaId"`
}
