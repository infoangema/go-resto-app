package categorias

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/paletgerardo/go-app-resto/core/db"
	"log"
	"time"
)

type Categoria struct {
	Id          int       `json:"id"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Activo      bool      `json:"activo"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

func AcaSeGuardaLaCategoria(c Categoria) error {

	queryString := "insert into categorias (nombre, descripcion) values ($1, $2)"

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	statment, err := connectionDB.Prepare(queryString)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer statment.Close()

	row, err := statment.Exec(c.Nombre, c.Descripcion)
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
	errorAlParsearDatos := row.Scan(&categoria.Id, &categoria.Nombre, &categoria.Descripcion)

	if errorAlParsearDatos != nil && errorAlParsearDatos != sql.ErrNoRows {
		fmt.Println(errorAlParsearDatos.Error())
		return categoria, errorAlParsearDatos
	}

	return categoria, nil

}

func AcaSeActualizaLaCategoria(aguardar Categoria) error {
	queryString := `UPDATE categorias set 
                     nombre=$1,
                     descripcion=$2
					where id = $3`

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	statment, err := connectionDB.Prepare(queryString)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer statment.Close()

	row, err := statment.Exec(aguardar.Nombre, aguardar.Descripcion, aguardar.Id)
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

	rows, err := connectionDB.Query("SELECT  id, nombre, descripcion FROM categorias limit 20")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		cat := new(Categoria)
		if err := rows.Scan(&cat.Id, &cat.Nombre, &cat.Descripcion); err != nil {
			panic(err)
		}
		categorias = append(categorias, cat)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return categorias, nil
}
