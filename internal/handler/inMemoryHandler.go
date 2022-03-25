package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type MemoryRequestPayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MemoryErrorResponsePayload struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// POST Endpoint (In-memory)
func PostInMemeoryDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var memoryRequestPayload MemoryRequestPayload

		// Reead request body
		req, err := ioutil.ReadAll(r.Body)
		if err != nil {
			json.NewEncoder(w).Encode(MemoryErrorResponsePayload{
				Code:    1,
				Message: "Error: could not complete read from request body",
			})
			log.Println("Could not complete read from request body")
			return
		}

		// Unmarshal request body
		err = json.Unmarshal(req, &memoryRequestPayload)
		if err != nil {
			json.NewEncoder(w).Encode(MemoryErrorResponsePayload{
				Code:    2,
				Message: "Error: could not complete unmarshal body",
			})
			log.Println("Could not complete unmarshal body")
			return
		}

		json.NewEncoder(w).Encode(memoryRequestPayload)
	}
}

//  GET Endpoint (In-memory)
func GetInMemeoryDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		if key == "" {
			json.NewEncoder(w).Encode(MemoryErrorResponsePayload{
				Code:    3,
				Message: "Error: Url Param 'key' is missing",
			})
			log.Println("Url Param 'key' is missing")
			return
		}

		// ? what is value. How to find value? Value is random ?
		json.NewEncoder(w).Encode(MemoryRequestPayload{
			Key:   key,
			Value: "getir",
		})
	}
}
