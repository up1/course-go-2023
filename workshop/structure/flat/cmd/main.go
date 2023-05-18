package main

import (
	"demo/beers"
	"log"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// Global variable
// var db Storage

func setupRouter() *gin.Engine {
	db, err := beers.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	database := beers.Database{DB: db}

	beers.PopulateBeers(database)

	r := gin.New()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	r.GET("/beers", beers.GetBeers(database))
	r.POST("/beers", database.AddBeer)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
