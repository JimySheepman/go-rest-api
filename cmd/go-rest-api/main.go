package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/JimySheepman/go-rest-api/config/db"
	"github.com/JimySheepman/go-rest-api/config/env"
	"github.com/JimySheepman/go-rest-api/internal/handler"
	"github.com/gorilla/mux"
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

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"Hello": "World"})
	})

	router.HandleFunc("/api/v1/fetch-data", handler.GetFetchDataHandler(database)).Methods("POST")
	router.HandleFunc("/api/v1/in-memory", handler.PostInMemeoryDataHandler(database)).Methods("POST")
	router.HandleFunc("/api/v1/in-memory", handler.GetInMemeoryDataHandler(database)).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
