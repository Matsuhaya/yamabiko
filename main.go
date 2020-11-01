package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Yamabiko struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.POST("/v1/echo", postEcho)
	e.Logger.Fatal(e.Start(":1323"))
}

func postEcho(c echo.Context) error {
	y := new(Yamabiko)
	if err := c.Bind(y); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, y)
}
