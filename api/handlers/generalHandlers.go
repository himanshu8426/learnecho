package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func HomePage(c echo.Context) error {
	return c.String(http.StatusOK, "You hit home page")
}
