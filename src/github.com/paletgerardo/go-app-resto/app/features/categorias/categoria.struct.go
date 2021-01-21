package categorias

import (
	"gorm.io/gorm"
	"time"
)

type Categoria struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Nombre      string         `json:"nombre"`
	Descripcion string         `json:"descripcion"`
	Activo      bool           `gorm:"default:true" json:"activo"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
