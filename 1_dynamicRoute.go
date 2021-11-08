package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

    // GET users/:id route
    e.GET("/users/:id", getUser)

    // log error if server failed to run
    e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
    // User ID from path `users/:id`
    id := c.Param("id")
    return c.String(http.StatusOK, id)
}