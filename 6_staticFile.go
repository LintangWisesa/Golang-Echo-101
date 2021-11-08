package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Serve any file from "file" directory for path /static/* 
	// GET http://localhost:1323/static/tes.png
    e.Static("/static", "file")

	// log error if server failed to run
    e.Logger.Fatal(e.Start(":1323"))
}