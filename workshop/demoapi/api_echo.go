package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", doXXHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func doXXHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
