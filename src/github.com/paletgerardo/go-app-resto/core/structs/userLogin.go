package structs

import (
	"gorm.io/gorm"
	"time"
)

type Usuarios struct {
	Id              int            `json:"id"`
	Nombre          string         `json:"nombre"`
	Apellido        string         `json:"apellido"`
	Email           string         `json:"email"`
	Password        string         `json:"password"`
	FechaNacimiento time.Time      `json:"fechaNacimiento"`
	CreatedAt       time.Time      `json:"CreatedAt"`
	UpdatedAt       time.Time      `json:"UpdatedAt,omitempty"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (p Usuarios) TableName() string {
	return "usuarios" // default table name
}
