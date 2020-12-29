package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetConnectionToDB() *sql.DB {
	stringConnectionToDB := "postgres://gerrdev:9003@127.0.0.1:5432/restodb?sslmode=disable"
	connectionToDB, err := sql.Open("postgres", stringConnectionToDB)
	if err != nil {
		log.Fatal(err)
	}
	return connectionToDB
}
