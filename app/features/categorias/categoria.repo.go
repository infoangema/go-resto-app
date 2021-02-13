package categorias

import (
	"errors"
	"go-app-resto/core/db"
)

func AcaSeGuardaLaCategoria(categoria Categoria) (Categoria, error) {
	gormdb := db.GetGormConnection()
	result := gormdb.Create(&categoria)
	if result.Error != nil {
		return categoria, errors.New("Error alguardar la categoria: ")
	}
	return categoria, nil
}

func BuscarUnaCategoriaPorId(id int) (Categoria, error) {

	var categoria Categoria
	gormdb := db.GetGormConnection()
	result := gormdb.First(&categoria, id)
	if result.Error != nil {
		return categoria, errors.New("Error al buscar la categoria: ")
	}
	return categoria, nil

}

func AcaSeActualizaLaCategoria(categoria Categoria) (Categoria, error) {
	gormdb := db.GetGormConnection()
	result := gormdb.Save(&categoria)
	if result.Error != nil {
		return categoria, errors.New("Error al actualizar la categoria: ")
	}
	return categoria, nil
}

func BorrarUnaCategoriaPorId(id int) error {

	var categoria Categoria
	gormdb := db.GetGormConnection()
	result := gormdb.Delete(&categoria, id)
	if result.Error != nil {
		return errors.New("Error al buscar la categoria")
	}
	return nil
}

func BuscarTodasLasCategorias() ([]*Categoria, error) {

	categorias := make([]*Categoria, 0)
	gormdb := db.GetGormConnection()
	result := gormdb.Find(&categorias)
	if result.Error != nil {
		return categorias, errors.New("Error al buscar las categoria")
	}
	return categorias, nil
}
