package productos

import (
	"errors"
	"github.com/paletgerardo/go-app-resto/core/db"
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
	if result.Error != nil {
		return producto, errors.New("Error al buscar el producto: ")
	}
	return producto, nil

}

func AcaSeActualizaElProducto(aguardar Producto) (Producto, error) {

	gormdb := db.GetGormConnection()
	result := gormdb.Save(&aguardar)
	if result.Error != nil {
		return aguardar, errors.New("Error al actualizar la categoria: ")
	}
	return aguardar, nil
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

func BuscarTodosLosProducto() ([]*Producto, error) {

	productos := make([]*Producto, 0)

	gormdb := db.GetGormConnection()
	result := gormdb.Find(&productos)
	if result.Error != nil {
		return productos, errors.New("Error al buscar las categoria")
	}
	return productos, nil
}
