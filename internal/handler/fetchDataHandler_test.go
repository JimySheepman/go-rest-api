package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/JimySheepman/go-rest-api/config/db"
	"github.com/JimySheepman/go-rest-api/config/env"
)

func init() {
	_, err := env.LoadEnvironmentConfigure("../../.env")
	if err != nil {
		log.Fatal("Loading .env file failed")
	}
}

func TestGetFetchDataHandler(t *testing.T) {

	// ! finish GetFetchDataHandler back to finish
	t.Run("mongodb fetch data and filter value", func(t *testing.T) {

		testBody := &RecordsRequestPayload{
			StartDate: "2016-01-26",
			EndDate:   "2018-02-02",
			MinCount:  2700,
			MaxCount:  3000,
		}

		body, _ := json.Marshal(testBody)

		req, err := http.NewRequest(http.MethodPost, "/api/v1/fetch-data", strings.NewReader(string(body)))
		if err != nil {
			t.Errorf("Request creation failed: ERROR: %v", err)
		}

		res := httptest.NewRecorder()
		handler := assertHandler()
		handler.ServeHTTP(res, req)

		expectedResponse := &RecordsRequestPayload{
			StartDate: "2016-01-26",
			EndDate:   "2018-02-02",
			MinCount:  2700,
			MaxCount:  3000,
		}
		marshalExpectedResponse, _ := json.Marshal(expectedResponse)
		expected := string(marshalExpectedResponse)
		if !reflect.DeepEqual(res.Body.String(), marshalExpectedResponse) {
			t.Errorf("Handler returned unexpected body: got\n %v want\n %v", res.Body.String(), string(marshalExpectedResponse))
		}
	})

	t.Run("status method allowed POST", func(t *testing.T) {

		testBody := &RecordsRequestPayload{
			StartDate: "2016-01-26",
			EndDate:   "2018-02-02",
			MinCount:  2700,
			MaxCount:  3000,
		}

		body, _ := json.Marshal(testBody)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/fetch-data", strings.NewReader(string(body)))
		if err != nil {
			t.Errorf("Request creation failed: ERROR: %v", err)
		}

		res := httptest.NewRecorder()
		handler := assertHandler()
		handler.ServeHTTP(res, req)

		if req.Method != "POST" {
			t.Errorf("Request method is not 'POST': got\n %v want\n %v", req.Method, http.MethodPost)
		}

	})

	t.Run("status method not allowed GET", func(t *testing.T) {

		testBody := &RecordsRequestPayload{
			StartDate: "2016-01-26",
			EndDate:   "2018-02-02",
			MinCount:  2700,
			MaxCount:  3000,
		}

		body, _ := json.Marshal(testBody)
		req, err := http.NewRequest(http.MethodGet, "/api/v1/fetch-data", strings.NewReader(string(body)))
		if err != nil {
			t.Errorf("Request creation failed: ERROR: %v", err)
		}

		res := httptest.NewRecorder()
		handler := assertHandler()
		handler.ServeHTTP(res, req)

		if req.Method != "GET" {
			t.Errorf("Request method is not 'POST': got\n %v want\n %v", req.Method, http.MethodPost)
		}

	})

	t.Run("wrong time format", func(t *testing.T) {

		testBody := &RecordsRequestPayload{
			StartDate: "2016-01-26",
			EndDate:   "2018-02-02",
			MinCount:  2700,
		}

		body, _ := json.Marshal(testBody)

		req, err := http.NewRequest(http.MethodPost, "/api/v1/fetch-data", strings.NewReader(string(body)))
		if err != nil {
			t.Errorf("Request creation failed: ERROR: %v", err)
		}

		res := httptest.NewRecorder()
		handler := assertHandler()
		handler.ServeHTTP(res, req)

		expectedResponse := &RecordsResponsePayload{
			Code:    3,
			Message: "Error: wrong time format ",
			Records: []Record{},
		}
		marshalExpectedResponse, _ := json.Marshal(expectedResponse)
		expected := string(marshalExpectedResponse)
		fmt.Println(res.Body)
		if !reflect.DeepEqual(res.Body.String(), expectedResponse) {
			t.Errorf("Handler returned unexpected body: got\n %v want\n %v", res.Body.String(), expected)
		}
	})

}

func assertHandler() http.HandlerFunc {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	handler := http.HandlerFunc(GetFetchDataHandler(database))
	return handler
}
