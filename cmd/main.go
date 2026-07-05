package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/CassioAA/drive/internal/repository"
)

func main() {

	connString := "host=localhost port=5432 dbname=postgres user=postgres password=senha123 sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error opening the database: ", err)
	}
	defer db.Close()
	log.Println("Connection string checked successfully")

	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	log.Println("Connection to the database checked successfully")

	repository.DbTests(db)
}
