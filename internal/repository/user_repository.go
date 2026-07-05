package repository

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateUser(db *sql.DB, user *User) error {

	query := `
	INSERT INTO users (name, email)
	VALUES ($1, $2)
	RETURNING id, created_at
	`

	return db.QueryRow(query, user.Name, user.Email).Scan(&user.ID, &user.CreatedAt)
}

func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(query)

	return err
}
