package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Yamabiko is response
type Yamabiko struct {
	Message string `json:"message"`
}

func GetWelcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		y := new(Yamabiko)
		y.Message = "Welcome to Yamabiko!!"
		return c.JSON(http.StatusOK, y)
	}
}

func PostEcho() echo.HandlerFunc {
	return func(c echo.Context) error {
		y := new(Yamabiko)
		if err := c.Bind(y); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, y)
	}
}
