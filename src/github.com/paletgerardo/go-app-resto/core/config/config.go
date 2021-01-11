package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar .env")
	}
}
