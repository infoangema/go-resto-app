package structs

import "time"

type UserLogin struct {
	Id              int       `json:"id"`
	Nombre          string    `json:"nombre"`
	Apellido        string    `json:"apellido"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	FechaNacimiento time.Time `json:"fechaNacimiento"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated,omitempty"`
}
