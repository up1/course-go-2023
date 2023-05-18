package banking

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BalanceResponse struct {
	Amount int `json:"amount"`
}

func GetBalance(b Banking, db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := BalanceResponse{Amount: b.Balance()}
		return c.JSON(http.StatusOK, r)
	}
}
