package main

import (
	"log"
	"net/http"

	"github.com/kolindes/simpleRestApi/internal/config"
	"github.com/kolindes/simpleRestApi/internal/database"
	"github.com/kolindes/simpleRestApi/internal/handlers"
	"github.com/kolindes/simpleRestApi/internal/logger"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log := logger.New(config.LoggingConfig)

	err = database.InitDB(config.DB)
	if err != nil {
		log.Fatal(err.Error())
	}

	http.HandleFunc("/", handlers.PathNotFound)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	log.Info("Listening on " + config.Main.Host + ":" + config.Main.Port)
	log.Fatal(http.ListenAndServe(config.Main.Host+":"+config.Main.Port, nil).Error())
}
