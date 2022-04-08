package main

import (
	"log"
	"net/http"
	"time"

	"github.com/JimySheepman/go-rest-api/config/db"
	"github.com/JimySheepman/go-rest-api/config/env"
	"github.com/JimySheepman/go-rest-api/internal/router"
)

func init() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	_, err := env.LoadEnvironmentConfigure(".env")
	if err != nil {
		log.Fatal("Loading .env file failed")
	}

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	mux := router.Router(database)

	srv := &http.Server{
		Handler:      mux,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
