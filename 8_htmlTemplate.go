package main

import (
	"net/http"
	"io"
	"github.com/labstack/echo/v4"
	"html/template"
)

type M map[string]interface{}

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}

    e.Renderer = t
	e.GET("/hello", Hello)

	// log error if server failed to run
    e.Logger.Fatal(e.Start(":1323"))
}

func Hello(c echo.Context) error {
	data := M{"name": "Andy", "age": 22}
	return c.Render(http.StatusOK, "hello.html", data)
}
