package main

import (
	"fmt"
	"os"

	"github.com/Matsuhaya/yamabiko/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// middleware
	logger := middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logFormat(),
		Output: os.Stdout,
	})
	e.Use(logger)
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(bodyDumpHandler))

	// routing
	e.GET("/", handler.GetWelcome())
	e.POST("/v1/echo", handler.PostEcho())

	// start server
	e.Logger.Fatal(e.Start(":1323"))
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
