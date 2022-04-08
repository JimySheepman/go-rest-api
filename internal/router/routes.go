package router

import (
	"encoding/json"
	"net/http"

	"github.com/JimySheepman/go-rest-api/internal/handler"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(db *mongo.Database) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"Hello": "World"})
	})

	router.HandleFunc("/api/v1/fetch-data", handler.GetFetchDataHandler(db)).Methods("POST")
	router.HandleFunc("/api/v1/in-memory", handler.PostInMemeoryDataHandler()).Methods("POST")
	router.HandleFunc("/api/v1/in-memory", handler.GetInMemeoryDataHandler()).Methods("GET")
}
