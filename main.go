// File: main.go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	return c.String(http.StatusOK, "BITBOX")
}

func bitView(c echo.Context) error {
	return c.String(http.StatusOK, "Displaying bit")
}

func bitCreate(c echo.Context) error {
	return c.String(http.StatusOK, "Creating bit")
}

func main() {
	e := echo.New()

	e.GET("/", home)
	e.GET("/bitbox/view", bitView)
	e.GET("/bitbox/create", bitCreate)

	e.Logger.Fatal(e.Start(":3000"))
}
