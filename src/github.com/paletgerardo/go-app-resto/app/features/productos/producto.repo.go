package productos

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/paletgerardo/go-app-resto/core/db"
	"log"
)

func AcaSeGuardaElProducto(productoQueGuardo Producto) (Producto, error) {

	gormdb := db.GetGormConnection()
	result := gormdb.Save(&productoQueGuardo)
	if result.Error != nil {
		return productoQueGuardo, errors.New("Error al actualizar el producto: ")
	}
	return productoQueGuardo, nil
}

func BuscarUnProductoPorId(id int) (ProductoAPI, error) {

	var producto ProductoAPI
	gormdb := db.GetGormConnection()
	result := gormdb.Model(&Producto{}).Where("id = ?", id).Find(&producto)
	// gerr: error() no puede imprimir Detail.
	if result.Error != nil {
		return producto, errors.New("Error al buscar el producto: ")
	}
	return producto, nil

}

func AcaSeActualizaElProducto(producto Producto) (Producto, error) {

	gormdb := db.GetGormConnection()
	result := gormdb.Save(producto)
	if result.Error != nil {
		//todo => Crear funcion de salida para logs.
		productoJson, err := json.Marshal(producto)
		if err != nil {
			fmt.Println(err)
			return producto, errors.New("No se puedo actualizar el producto.")
		}
		log.Println("producto: ", string(productoJson))
		return producto, errors.New("No se puedo actualizar el producto.")
	}
	return producto, nil
}

func BorrarUnProductoPorId(id int) error {

	var producto Producto
	gormdb := db.GetGormConnection()
	result := gormdb.Delete(&producto, id)
	if result.Error != nil {
		return errors.New("Error al intentar borrar el producto.")
	}
	return nil

}

func BuscarTodosLosProducto() ([]ProductoAPI, error) {

	var productos []ProductoAPI
	gormdb := db.GetGormConnection()
	// Model recibe un puntero de tipo struct []Producto, pero devuelve []ProductoAPI.
	result := gormdb.Model(&[]Producto{}).Find(&productos)
	if result.Error != nil {
		return productos, errors.New("Error al buscar el producto: ")
	}
	return productos, nil
}

type SyntaxError struct {
	Detail string
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("%d: syntax error", e.Detail)
}
