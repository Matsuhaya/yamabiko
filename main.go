package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Yamabiko struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	logger := middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logFormat(),
		Output: os.Stdout,
	})
	e.Use(logger)
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(bodyDumpHandler))

	e.GET("/", getWelcome)
	e.POST("/v1/echo", postEcho)
	e.Logger.Fatal(e.Start(":1323"))
}

func getWelcome(c echo.Context) error {
	y := new(Yamabiko)
	y.Message = "Welcome to Yamabiko!!"
	return c.JSON(http.StatusOK, y)
}

func postEcho(c echo.Context) error {
	y := new(Yamabiko)
	if err := c.Bind(y); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, y)
}

func logFormat() string {
	var format string
	format += "time:${time_rfc3339}\t"
	format += "host:${remote_ip}\t"
	format += "forwardedfor:${header:x-forwarded-for}\t"
	format += "req:-\t"
	format += "status:${status}\t"
	format += "method:${method}\t"
	format += "uri:${uri}\t"
	format += "size:${bytes_out}\t"
	format += "referer:${referer}\t"
	format += "ua:${user_agent}\t"
	format += "reqtime_ns:${latency}\t"
	format += "cache:-\t"
	format += "runtime:-\t"
	format += "apptime:-\t"
	format += "vhost:${host}\t"
	format += "reqtime_human:${latency_human}\t"
	format += "x-request-id:${id}\t"
	format += "host:${host}\n"

	return format
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}
