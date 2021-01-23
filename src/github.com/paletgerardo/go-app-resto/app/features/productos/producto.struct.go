package productos

import (
	"github.com/paletgerardo/go-app-resto/app/features/categorias"
	"gorm.io/gorm"
	"time"
)

type Producto struct {
	ID          uint                 `gorm:"primaryKey" json:"id"`
	Nombre      string               `json:"nombre"`
	Descripcion string               `json:"descripcion"`
	Activo      bool                 `gorm:"default:true" json:"activo"`
	Precio      float32              `json:"precio"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	DeletedAt   gorm.DeletedAt       `gorm:"index" json:"deleted_at"`
	CategoriaID uint                 `json:"categoria_id"`
	Categoria   categorias.Categoria `gorm:"references:id" json:"categoria"`
}

type ProductoAPI struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Activo      bool      `gorm:"default:true" json:"activo"`
	Precio      float32   `json:"precio"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CategoriaID uint      `json:"categoria_id"`
}
