package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// install echo middleware pkg
// $ go get github.com/labstack/echo/v4/middleware

func main() {
	e := echo.New()

	// Logger middleware logs the information about each HTTP request.
	e.Use(middleware.Logger())

    // base route '/'
	e.GET("/", func(c echo.Context) error {		
		return c.String(http.StatusOK, "Hello, World!")
	})

	// log error if server failed to run
    e.Logger.Fatal(e.Start(":1323"))
}