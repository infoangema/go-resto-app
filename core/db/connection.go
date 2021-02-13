package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var USER = os.Getenv("USERDB")
var PASS = os.Getenv("PASSDB")
var HOST = os.Getenv("HOSTDB")
var PORT = os.Getenv("PORTDB")
var NAME = os.Getenv("DBNAME")

func GetConnectionToDB() *sql.DB {
	URL := "postgres://gerrdev:9003@127.0.0.1:5432/restodb?sslmode=disable"
	//URL := "postgres://"+USER+":"+PASS+"@"+HOST+":"+PORT+"/"+ NAME +"?sslmode=disable"

	connectionToDB, err := sql.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
	}
	return connectionToDB
}

func GetGormConnection() *gorm.DB {
	sqlDB := GetConnectionToDB()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return gormDB
}
