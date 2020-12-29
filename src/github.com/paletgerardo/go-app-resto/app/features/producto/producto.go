package producto

import (
	"errors"
	"github.com/paletgerardo/go-app-resto/core/db"
	"log"
	"net/http"
	"time"
)

type Producto struct {
	Id          int
	Nombre      string
	Descripcion string
	Activo      bool
	Created     time.Time
	Updated     time.Time
}

func CreateProducto(p Producto) error {
	queryString := `INSERT INTO productos (nombre, descripcion) VALUES ($1, $2 )`

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	statment, err := connectionDB.Prepare(queryString)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer statment.Close()

	row, err := statment.Exec(p.Nombre, p.Descripcion)
	if err != nil {
		log.Fatal(err)
		return err
	}

	i, err := row.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba una sola fila afectada")
	}

	return nil
}

func DeleteProducto(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteProducto"))
}

func FindByIdProducto(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FindByIdProducto"))
}
