package structs

import (
	"errors"
	"time"
)

/* Struct Nulltime: se utiliza para poder recibir datos nulos de tipo fecha desde
la base de datos */
type NullTime struct {
	Time  time.Time `json:"time"`
	Valid bool      `json:"valid"`
}

func (nt *NullTime) Scan(fecha interface{}) error {
	if fecha == nil {
		nt.Time, nt.Valid = time.Time{}, false
		return nil
	}

	switch v := fecha.(type) {
	case time.Time:
		nt.Time, nt.Valid = v, true
		return nil
	}
	return errors.New("No se pudo convertir %T a time.Time")
}
