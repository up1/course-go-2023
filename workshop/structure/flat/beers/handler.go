package beers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Database struct {
	DB Storage
}

func GetBeers(d Database) gin.HandlerFunc {
	println("Inital :: only once")
	return func(c *gin.Context) {
		println("Every request")
		beers, err := d.DB.FindBeers()
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
	d.DB.SaveBeer(beer)
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
