package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

    // GET /show?team=avenger&member=ironman
    e.GET("/show", show)

    // log error if server failed to run
    e.Logger.Fatal(e.Start(":1323"))
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team: " + team + ", member: " + member)
}