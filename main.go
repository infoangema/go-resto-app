package main

import (
	"./src/github.com/paletgerardo/go-app-resto/app/features/categorias"
	"./src/github.com/paletgerardo/go-app-resto/app/features/productos"
	"./src/github.com/paletgerardo/go-app-resto/app/handlers"
	"./src/github.com/paletgerardo/go-app-resto/core/db"
	"./src/github.com/paletgerardo/go-app-resto/core/structs"
)

func main() {
	gormDB := db.GetGormConnection()
	_ = gormDB.AutoMigrate(&productos.Producto{}, &categorias.Categoria{}, &structs.Usuarios{})

	handlers.HandlerRequest()
}
