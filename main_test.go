package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PW486/gost/db"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	db.Open()
	db.Migration()
}

func TestSuccessGetRootRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSuccessCreateRootRoute(t *testing.T) {
	router := setupRouter()

	var jsonStr = []byte(`{ "email": "TestEmail", "name": "TestName" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestValidationErrorCreateRootRoute(t *testing.T) {
	router := setupRouter()

	var jsonStr = []byte(`{ "email": "TestEmail" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
