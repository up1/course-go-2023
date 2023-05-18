package main

import (
	"log"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// Global variable
// var db Storage

func setupRouter() *gin.Engine {
	db, err := NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	database := Database{db: db}

	PopulateBeers(database)

	r := gin.New()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	r.GET("/beers", GetBeers(database))
	r.POST("/beers", database.AddBeer)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
