package main

import (
	"log"
	"net/http"

	"github.com/kolindes/simpleRestApi/internal/config"
	"github.com/kolindes/simpleRestApi/internal/database"
	"github.com/kolindes/simpleRestApi/internal/handlers"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// sqlite for a while.
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", handlers.PathNotFound)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	log.Fatal(http.ListenAndServe(config.Main.Host+":"+config.Main.Port, nil))
}
