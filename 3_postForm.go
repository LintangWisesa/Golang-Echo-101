package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// POST /save body = application/x-www-form-urlencoded
    e.POST("/save", save)

	// log error if server failed to run
    e.Logger.Fatal(e.Start(":1323"))
}

func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}