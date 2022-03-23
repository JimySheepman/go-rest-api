package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestGetDataHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/getData", nil)
	if err != nil {
		t.Errorf("Request creation failed: ERROR: %v", err)
	}

	rr := httptest.NewRecorder()

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	handler := http.HandlerFunc(GetDataHandler(database))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"startDate":"2015-01-01","endDate":"2017-01-28","minCount":4,"maxCount":6304}`
	if !reflect.DeepEqual(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got\n %v want\n %v", rr.Body.String(), expected)
	}
}

/*

t.Run("read env file", func(t *testing.T) {
	loaded, _ := LoadEnvironmentConfigure("../../.env")
	expected := true

	if loaded != expected {
		t.Errorf("expected %v but got %v", expected, loaded)
	}
})

t.Run("read env file error", func(t *testing.T) {
	loaded, _ := LoadEnvironmentConfigure("../.env")
	expected := false

	if loaded != expected {
		t.Errorf("expected %v but got %v", expected, loaded)
	}
}) */
