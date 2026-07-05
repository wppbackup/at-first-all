package repository

import (
	"database/sql"
	"log"
)

func DbTests(db *sql.DB) {

	err := CreateTable(db)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}
	log.Println("Table created successfully")

	user := User{
		Name:  "Cassio",
		Email: "cassio@gmail.com",
	}
	err = CreateUser(db, &user)
	if err != nil {
		log.Fatal("Error creating user in the database")
	}
	log.Println("User created successfully in the database")

}
