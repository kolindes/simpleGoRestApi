package main

import (
	"log"
	"net/http"

	"github.com/kolindes/simpleRestApi/internal/database"
	"github.com/kolindes/simpleRestApi/internal/handlers"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", handlers.PathNotFound)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
