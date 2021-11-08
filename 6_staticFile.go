package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Serve any file from "file" directory for path /static/* 
	// GET http://localhost:1323/static/tes.png => image
	// GET http://localhost:1323/static/tes.html => html
    e.Static("/static", "file")

	// log error if server failed to run
    e.Logger.Fatal(e.Start(":1323"))
}