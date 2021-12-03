package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()

	// POST /users body = json / xml / form / query
	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u) // return as JSON
		// or
		// return c.XML(http.StatusCreated, u) // return as XML
	})

	// log error if server failed to run
	e.Logger.Fatal(e.Start(":1323"))
}
