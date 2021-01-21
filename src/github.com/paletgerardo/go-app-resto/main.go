package main

import (
	"github.com/paletgerardo/go-app-resto/app/features/categorias"
	"github.com/paletgerardo/go-app-resto/app/handlers"
	"github.com/paletgerardo/go-app-resto/core/db"
)

func main() {
	gormDB := db.GetPostgresConnection()
	gormDB.AutoMigrate(&categorias.Categoria{})
	handlers.HandlerRequest()
}
