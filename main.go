package main

import (
	"log"
	"net/http"
	"time"

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

	srv := &http.Server{
		Handler:      http.HandlerFunc(router.Serve),
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
