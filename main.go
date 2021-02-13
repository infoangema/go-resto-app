package main

import (
	"go-app-resto/app/features/categorias"
	"go-app-resto/app/features/productos"
	"go-app-resto/app/handlers"
	"go-app-resto/core/db"
	"go-app-resto/core/structs"
)

func main() {
	gormDB := db.GetGormConnection()
	_ = gormDB.AutoMigrate(&productos.Producto{}, &categorias.Categoria{}, &structs.Usuarios{})

	handlers.HandlerRequest()
}
