package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBeers(c *gin.Context) {
	beers, err := db.FindBeers()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, beers)
}

func AddBeer(c *gin.Context) {
	var beer Beer
	if err := c.ShouldBindJSON(&beer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Bad request"})
		return
	}
	db.SaveBeer(beer)
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
