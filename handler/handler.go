package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	Yamabiko struct {
		Message string `json:"message"`
	}
	Handler struct{}
)

func (h *Handler) GetWelcome(c echo.Context) error {
	y := new(Yamabiko)
	y.Message = "Welcome to Yamabiko!!"
	return c.JSON(http.StatusOK, y)
}

func (h *Handler) PostEcho(c echo.Context) error {
	y := new(Yamabiko)
	if err := c.Bind(y); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, y)
}
