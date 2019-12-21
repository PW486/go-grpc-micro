package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PW486/gost/database"
	"github.com/PW486/gost/entity"
	"github.com/PW486/gost/router"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	db := database.Init()
	db.AutoMigrate(&entity.Account{})
}

func TestSuccessGetRootRoute(t *testing.T) {
	router := router.Init()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/accounts", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSuccessCreateRootRoute(t *testing.T) {
	router := router.Init()

	var jsonStr = []byte(`{ "email": "TestEmail", "name": "TestName", "password": "abc" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestValidationErrorCreateRootRoute(t *testing.T) {
	router := router.Init()

	var jsonStr = []byte(`{ "email": "TestEmail" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogInSucceed(t *testing.T) {
	router := router.Init()

	var jsonStr = []byte(`{ "email": "TestEmail", "password": "abc" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLogInFailed(t *testing.T) {
	router := router.Init()

	var jsonStr = []byte(`{ "email": "TestEmail", "password": "aaa" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogInValidationFailed(t *testing.T) {
	router := router.Init()

	var jsonStr = []byte(`{ "email": "TestEmail" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
