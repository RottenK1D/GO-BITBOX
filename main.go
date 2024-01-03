// File: main.go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Home Page")
}

func about(c echo.Context) error {
	return c.String(http.StatusOK, "About Page")
}

func main() {
	e := echo.New()

	e.GET("/", home)
	e.GET("/about", about)

	e.Logger.Fatal(e.Start(":3000"))
}
