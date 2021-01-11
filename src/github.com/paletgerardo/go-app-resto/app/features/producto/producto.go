package producto

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/paletgerardo/go-app-resto/core/db"
	"log"
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
}

func AcaSeGuardaElProducto(p Producto) error {
	queryString := `INSERT INTO productos (nombre, descripcion, precio) VALUES ($1, $2, $3)`

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	statment, err := connectionDB.Prepare(queryString)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer statment.Close()

	row, err := statment.Exec(p.Nombre, p.Descripcion, p.Precio)
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

func BuscarUnProductoPorId(id int) (Producto, error) {

	var producto Producto
	queryString := `select * from productos where id = $1`

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	row := connectionDB.QueryRow(queryString, id)
	errorAlParsearDatos := row.Scan(&producto.Id, &producto.Nombre, &producto.Descripcion, &producto.Precio)

	if errorAlParsearDatos != nil && errorAlParsearDatos != sql.ErrNoRows {
		fmt.Println(errorAlParsearDatos.Error())
		return producto, errorAlParsearDatos
	}

	return producto, nil

}

func AcaSeActualizaElProducto(aguardar Producto) error {
	queryString := `UPDATE productos set 
                     nombre=$1,
                     descripcion=$2,
                     precio=$3
					where id = $4`

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	statment, err := connectionDB.Prepare(queryString)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer statment.Close()

	row, err := statment.Exec(aguardar.Nombre, aguardar.Descripcion, aguardar.Precio, aguardar.Id)
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

func BorrarUnProductoPorId(id int) error {
	queryString := `DELETE from productos where id = $1`

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	statment, err := connectionDB.Prepare(queryString)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer statment.Close()

	row, err := statment.Exec(id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	i, err := row.RowsAffected()
	if i != 1 {
		return errors.New("Error al eliminar el producto: ")
	}

	return nil
}

func BuscarTodosLosProducto() ([]*Producto, error) {

	productos := make([]*Producto, 0)

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	rows, err := connectionDB.Query("SELECT  id, nombre, descripcion, precio FROM productos limit 20")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		prd := new(Producto)
		if err := rows.Scan(&prd.Id, &prd.Nombre, &prd.Descripcion, &prd.Precio); err != nil {
			panic(err)
		}
		productos = append(productos, prd)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return productos, nil
}
