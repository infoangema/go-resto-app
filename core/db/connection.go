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
	URL := "postgres://omttzarnqhjlda:01ead292afd329351220056870602ba07c46eceab148d3fd611ff02b755c187d@ec2-35-174-118-71.compute-1.amazonaws.com:5432/d2ch7fcu9q8og7"
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
