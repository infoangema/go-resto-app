package main

import (
	"github.com/paletgerardo/go-app-resto/app/features/categorias"
	"github.com/paletgerardo/go-app-resto/app/features/productos"
	"github.com/paletgerardo/go-app-resto/app/handlers"
	"github.com/paletgerardo/go-app-resto/core/db"
	"github.com/paletgerardo/go-app-resto/core/structs"
)

func main() {
	gormDB := db.GetGormConnection()
	_ = gormDB.AutoMigrate(&productos.Producto{}, &categorias.Categoria{}, &structs.Usuarios{})
	handlers.HandlerRequest()
}
