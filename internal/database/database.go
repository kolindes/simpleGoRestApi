package database

import (
	"database/sql"
	"errors"

	"github.com/kolindes/simpleRestApi/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return nil, err
	}

	if err = createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL UNIQUE,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        );
    `)

	return err
}

func AddUser(username string, email string, password []byte) error {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, password)

	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (*models.User, error) {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, username, email, password FROM users WHERE username = ?", username)
	var id int64
	var storedUsername string
	var storedEmail string
	var storedPassword []byte

	err = row.Scan(&id, &storedUsername, &storedEmail, &storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	user := &models.User{
		ID:             id,
		Username:       storedUsername,
		Email:          storedEmail,
		HashedPassword: storedPassword,
	}

	return user, nil
}
