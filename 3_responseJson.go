package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message   string `json:"message"`
	Language  string `json:"language"`
	Framework string `json:"framework"`
}

func main() {
	e := echo.New()

	// base route '/'
	e.GET("/", func(c echo.Context) error {
		resp := &Response{
			Message:   "Hello World!",
			Language:  "Go",
			Framework: "Echo",
		}
		return c.JSON(http.StatusOK, resp)
	})

	// log error if server failed to run
	e.Logger.Fatal(e.Start(":1000"))
}
