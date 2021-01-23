package main

import (
	"github.com/paletgerardo/go-app-resto/app/features/categorias"
	"github.com/paletgerardo/go-app-resto/app/features/productos"
	"github.com/paletgerardo/go-app-resto/app/handlers"
	"github.com/paletgerardo/go-app-resto/core/db"
)

func main() {
	gormDB := db.GetPostgresConnection()
	_ = gormDB.AutoMigrate(&productos.Producto{}, &categorias.Categoria{})
	handlers.HandlerRequest()
}
