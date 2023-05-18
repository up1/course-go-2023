package main

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockStorageMemory struct {
	Err error
}

func (s *MockStorageMemory) FindBeers() ([]Beer, error) {
	beers := []Beer{}
	return beers, s.Err
}

func (s *MockStorageMemory) SaveBeer(beers ...Beer) error {
	return nil
}

func TestGetAllBeersRouteWithMock(t *testing.T) {
	expected := []Beer{}
	db = new(MockStorageMemory)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, router := gin.CreateTestContext(w)

	// Create router for test
	router.GET("/beers", GetBeers)

	c.Request = httptest.NewRequest("GET", "/beers", nil)
	router.ServeHTTP(w, c.Request)

	assert.Equal(t, 200, w.Code)

	var actual []Beer
	json.Unmarshal([]byte(w.Body.String()), &actual)
	assert.Equal(t, expected, actual)
}

func TestErrorGetAllBeersRouteWithMock(t *testing.T) {
	db = &MockStorageMemory{Err: fmt.Errorf("Demo")}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, router := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/beers", nil)

	router.GET("/beers", GetBeers)
	router.ServeHTTP(w, c.Request)

	assert.Equal(t, 400, w.Code)
}
