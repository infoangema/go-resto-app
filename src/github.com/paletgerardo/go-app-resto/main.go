package main

import (
	"github.com/paletgerardo/go-app-resto/app/handlers"
)

func main() {

	//if bd.CheckConection() == 0 {
	//	log.Fatal("Error al conectar con bd")
	//	return
	//}

	handlers.HandlerRequest()

}
