package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Database struct {
	db Storage
}

func GetBeers(d Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		beers, err := d.db.FindBeers()
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, beers)
	}
}

func (d *Database) AddBeer(c *gin.Context) {
	var beer Beer
	if err := c.ShouldBindJSON(&beer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Bad request"})
		return
	}
	d.db.SaveBeer(beer)
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
