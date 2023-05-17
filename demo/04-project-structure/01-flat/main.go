package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Global variable
var db Storage

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	var err error
	db, err = NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	PopulateBeers()

	r := gin.New()
	r.GET("/beers", GetBeers)
	r.POST("/beers", AddBeer)
	return r
}
