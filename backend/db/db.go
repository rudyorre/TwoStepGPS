package db

import (
	"backend/models"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

func NewDB() (*DB, error) {
	var connStr string
	if os.Getenv("ENV") == "development" {
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASS"),
			os.Getenv("POSTGRES_NAME"),
		)
	} else {
		// connStr = fmt.Sprintf(
		// 	"postgres://%s:%s@%s:%s/%s",
		// 	os.Getenv("POSTGRES_USER"),
		// 	os.Getenv("POSTGRES_PASS"),
		// 	os.Getenv("POSTGRES_HOST"),
		// 	os.Getenv("POSTGRES_PORT"),
		// 	os.Getenv("POSTGRES_NAME"),
		// )
		connStr = os.Getenv("POSTGRES_CONN_STR")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &DB{Conn: db}, nil
}

func (d *DB) GetAllUsers() (*sql.Rows, error) {
	return d.Conn.Query("SELECT * FROM users;")
}

func (d *DB) CreateUser(user models.User) error {
	_, err := d.Conn.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2);", user.Username, user.PasswordHash)
	return err
}

func (d *DB) UserExists(username string) (bool, error) {
	var exists bool
	err := d.Conn.QueryRow("SELECT exists (SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return exists, nil
}

func (d *DB) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := d.Conn.QueryRow("SELECT * FROM users WHERE username=$1;", username).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *DB) GetUserByID(id int) (*models.User, error) {
	return nil, nil
}
