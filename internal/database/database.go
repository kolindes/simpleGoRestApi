package database

import (
	"errors"

	"github.com/kolindes/simpleRestApi/internal/config"
	"github.com/kolindes/simpleRestApi/internal/models"
	"github.com/kolindes/simpleRestApi/internal/svcerr"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(config config.DBConfig) error {
	dialect := sqlite.Open("./data.db")

	var err error

	db, err = gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	// defer sqlDB.Close()
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)

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
