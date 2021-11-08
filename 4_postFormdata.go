package main

import (
	"net/http"
	"io"
	"os"
	"github.com/labstack/echo/v4"
	"path/filepath"
)

func main() {
	e := echo.New()

	// POST /save body = multipart/formdata
    e.POST("/save", save)

	// log error if server failed to run
    e.Logger.Fatal(e.Start(":1323"))
}

func save(c echo.Context) error {
	// Get name (text)
	name := c.FormValue("name")
	// Get avatar (file)
  	avatar, err := c.FormFile("avatar")
  	if err != nil {
 		return err
 	}
 
 	// Source
 	src, err := avatar.Open()
 	if err != nil {
 		return err
 	}
 	defer src.Close()
 
 	// Destination: file will be saved on root folder
 	// dst, err := os.Create(avatar.Filename)

	// Destination: file will be saved on "upload" folder, 
	// need to create upload folder & import pkg "path/filepath"
	dst, err := os.Create(filepath.Join("upload", filepath.Base(avatar.Filename)))
 	if err != nil {
 		return err
 	}
 	defer dst.Close()
 
 	// Copy
 	if _, err = io.Copy(dst, src); err != nil {
  		return err
  	}

	return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
}