package producto

import (
	"errors"
	"fmt"
	"github.com/paletgerardo/go-app-resto/core/db"
	"log"
	"time"
)

func AcaSeGuardaElProducto(p Producto) error {

	p.Activo = true
	p.Updated = time.Now()
	p.Created = time.Now()

	queryString := `INSERT INTO productos (
                       nombre,
                       descripcion,
                       precio,
                       activo,
                       created,
                       updated,
                       categoriaid
                       ) 
                       VALUES ($1, $2, $3,$4, $5, $6, $7)`

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	statment, err := connectionDB.Prepare(queryString)
	if err != nil {
		log.Print(err)
		return err
	}
	defer statment.Close()

	row, err := statment.Exec(p.Nombre, p.Descripcion, p.Precio, p.Activo, p.Created, p.Updated, p.CategoriaId)
	if err != nil {
		log.Print(err)
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
	errorAlParsearDatos := row.Scan(
		&producto.Id,
		&producto.Nombre,
		&producto.Descripcion,
		&producto.Activo,
		&producto.Created,
		&producto.Updated,
		&producto.Precio,
		&producto.CategoriaId,
	)

	if errorAlParsearDatos != nil {
		fmt.Println(errorAlParsearDatos.Error())
		return producto, errors.New("Producto no encontrado, " + errorAlParsearDatos.Error())
	}

	return producto, nil

}

func AcaSeActualizaElProducto(aguardar Producto) error {

	aguardar.Updated = time.Now()

	queryString := `UPDATE productos set 
                     nombre=$1,
                     descripcion=$2,
                     activo=$3,
                     updated=$4,
                     precio=$5,
                     categoriaid=$6                     
					where id = $7`

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	statment, err := connectionDB.Prepare(queryString)
	if err != nil {
		log.Print(err)
		return err
	}
	defer statment.Close()

	row, err := statment.Exec(
		aguardar.Nombre,
		aguardar.Descripcion,
		aguardar.Activo,
		aguardar.Updated,
		aguardar.Precio,
		aguardar.CategoriaId,
		aguardar.Id,
	)
	if err != nil {
		log.Print(err)
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
		log.Print(err)
		return err
	}
	defer statment.Close()

	row, err := statment.Exec(id)
	if err != nil {
		log.Print(err)
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

	rows, err := connectionDB.Query("SELECT  * FROM productos")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		prd := new(Producto)
		if err := rows.Scan(
			&prd.Id,
			&prd.Nombre,
			&prd.Descripcion,
			&prd.Activo,
			&prd.Created,
			&prd.Updated,
			&prd.Precio,
			&prd.CategoriaId,
		); err != nil {
			panic(err)
		}
		productos = append(productos, prd)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return productos, nil
}
