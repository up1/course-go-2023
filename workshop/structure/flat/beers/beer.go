package beers

import "time"

type Beer struct {
	ID        int
	Name      string  `json:"name" binding:"required"`
	Brewery   string  `json:"brewery" binding:"required"`
	Abv       float32 `json:"abv" binding:"required"`
	ShortDesc string  `json:"short_description" binding:"required"`
	Created   time.Time
}
