package categorias

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/paletgerardo/go-app-resto/core/db"
	"log"
	"time"
)

func AcaSeGuardaLaCategoria(c Categoria) error {

	c.Activo = true
	c.Updated = time.Now()
	c.Created = time.Now()

	queryString := "insert into categorias (nombre, descripcion, activo, created, updated) values ($1, $2, $3, $4, $5)"

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	statment, err := connectionDB.Prepare(queryString)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer statment.Close()

	row, err := statment.Exec(c.Nombre, c.Descripcion, c.Activo, c.Created, c.Updated)
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

func BuscarUnaCategoriaPorId(id int) (Categoria, error) {

	var categoria Categoria
	queryString := `select * from categorias where id = $1`

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	row := connectionDB.QueryRow(queryString, id)
	errorAlParsearDatos := row.Scan(&categoria.Id, &categoria.Nombre, &categoria.Descripcion, &categoria.Activo, &categoria.Created, &categoria.Updated)

	if errorAlParsearDatos != nil && errorAlParsearDatos != sql.ErrNoRows {
		fmt.Println(errorAlParsearDatos.Error())
		return categoria, errorAlParsearDatos
	}

	return categoria, nil

}

func AcaSeActualizaLaCategoria(aguardar Categoria) error {
	aguardar.Updated = time.Now()
	queryString := `UPDATE categorias set 
                     	nombre=$1,
                     	descripcion=$2,
                      	activo=$3,
						updated=$4
					where id = $5`

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	statment, err := connectionDB.Prepare(queryString)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer statment.Close()

	row, err := statment.Exec(aguardar.Nombre, aguardar.Descripcion, aguardar.Activo, aguardar.Updated, aguardar.Id)
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

func BorrarUnaCategoriaPorId(id int) error {
	queryString := `DELETE from categorias where id = $1`

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

func BuscarTodasLasCategorias() ([]*Categoria, error) {

	categorias := make([]*Categoria, 0)

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	rows, err := connectionDB.Query("SELECT  * FROM categorias")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		cat := new(Categoria)
		if err := rows.Scan(&cat.Id, &cat.Nombre, &cat.Descripcion, &cat.Activo, &cat.Created, &cat.Updated); err != nil {
			panic(err)
		}
		categorias = append(categorias, cat)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return categorias, nil
}
