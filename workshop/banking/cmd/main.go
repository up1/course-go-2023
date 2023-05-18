package main

import (
	"banking"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create DB connection
	my := banking.MySQL{}

	// initial banking
	b := banking.Banking{}

	e := echo.New()
	e.GET("/balance", banking.GetBalance(b, my.CreateConnection()))
	e.Logger.Fatal(e.Start(":1323"))
}
