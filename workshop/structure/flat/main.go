package main

import (
	"log"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// Global variable
var db Storage

func setupRouter() *gin.Engine {
	var err error
	db, err = NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	PopulateBeers()

	r := gin.New()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	r.GET("/beers", GetBeers)
	r.POST("/beers", AddBeer)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
