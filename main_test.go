package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PW486/go-grpc-micro/config"
	"github.com/PW486/go-grpc-micro/database"
	"github.com/PW486/go-grpc-micro/entity"
	"github.com/PW486/go-grpc-micro/router"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	config.Init()

	db := database.Init()
	db.AutoMigrate(&entity.Account{})
}

func TestSuccessGetAccounts(t *testing.T) {
	router := router.Init()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/accounts", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSuccessPostAccount(t *testing.T) {
	router := router.Init()

	var jsonStr = []byte(`{ "email": "TestEmail", "name": "TestName", "password": "486" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestValidationErrorPostAccount(t *testing.T) {
	router := router.Init()

	var jsonStr = []byte(`{ "email": "TestEmail" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSuccessLogIn(t *testing.T) {
	router := router.Init()

	var jsonStr = []byte(`{ "email": "TestEmail", "password": "486" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFailureLogIn(t *testing.T) {
	router := router.Init()

	var jsonStr = []byte(`{ "email": "TestEmail", "password": "123" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestValidationErrorLogIn(t *testing.T) {
	router := router.Init()

	var jsonStr = []byte(`{ "email": "TestEmail" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
