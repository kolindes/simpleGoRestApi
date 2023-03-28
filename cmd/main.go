package main

import (
	"log"
	"net/http"

	"github.com/kolindes/simpleRestApi/internal/config"
	"github.com/kolindes/simpleRestApi/internal/database"
	"github.com/kolindes/simpleRestApi/internal/handlers"
	"github.com/kolindes/simpleRestApi/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	var db *gorm.DB
	dialect := sqlite.Open("./data.db")

	db, err = gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	database.InitDB(dialect)

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	http.HandleFunc("/", handlers.PathNotFound)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	log.Println("Listening on " + config.Main.Host + ":" + config.Main.Port)
	log.Fatal(http.ListenAndServe(config.Main.Host+":"+config.Main.Port, nil))
}
