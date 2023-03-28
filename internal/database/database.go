package database

import (
	"errors"

	"github.com/kolindes/simpleRestApi/internal/models"
	"github.com/kolindes/simpleRestApi/internal/svcerr"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(dialector gorm.Dialector) error {
	var err error
	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}

func AddUser(username, email string, hashedPassword []byte) error {
	if user, _ := GetUserByUsername(username); user != nil {
		return errors.New(svcerr.UserAlreadyExists)
	}

	user := &models.User{
		Username:       username,
		Email:          email,
		HashedPassword: hashedPassword,
	}

	err := db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := db.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := db.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(id int64) (*models.User, error) {
	user := &models.User{}
	err := db.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// package database

// import (
// 	"database/sql"
// 	"errors"

// 	"github.com/kolindes/simpleRestApi/internal/models"
// 	"github.com/kolindes/simpleRestApi/internal/svcerr"
// 	_ "github.com/mattn/go-sqlite3"
// )

// func InitDB() (*sql.DB, error) {
// 	db, err := sql.Open("sqlite3", "./data.db")
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = createTables(db); err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

// func createTables(db *sql.DB) error {
// 	_, err := db.Exec(`
//         CREATE TABLE IF NOT EXISTS users (
//             id INTEGER PRIMARY KEY AUTOINCREMENT,
//             username TEXT NOT NULL UNIQUE,
//             email TEXT NOT NULL UNIQUE,
//             password TEXT NOT NULL
//         );
//     `)

// 	return err
// }

// func AddUser(username string, email string, password []byte) error {
// 	db, err := sql.Open("sqlite3", "./data.db")
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	if err != nil {
// 		return err
// 	}

// 	if _, err = GetUserByUsername(username); err == nil {
// 		return errors.New(svcerr.UserAlreadyExists)
// 	}

// 	_, err = db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, password)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func GetUserByUsername(username string) (*models.User, error) {
// 	db, err := sql.Open("sqlite3", "./data.db")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer db.Close()

// 	row := db.QueryRow("SELECT id, username, email, password FROM users WHERE username = ?", username)
// 	var id int64
// 	var storedUsername string
// 	var storedEmail string
// 	var storedPassword []byte

// 	err = row.Scan(&id, &storedUsername, &storedEmail, &storedPassword)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, errors.New("user not found")
// 		}
// 		return nil, err
// 	}

// 	user := &models.User{
// 		ID:             id,
// 		Username:       storedUsername,
// 		Email:          storedEmail,
// 		HashedPassword: storedPassword,
// 	}

// 	return user, nil
// }
