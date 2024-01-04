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
	if c.Request().Method != http.MethodPost {
		return c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	}

	m := "Creating bit"
	if err := c.Bind(&m); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	return c.String(http.StatusOK, m)
}

func main() {
	e := echo.New()

	e.GET("/", home)
	e.GET("/bitbox/view", bitView)
	e.POST("/bitbox/create", bitCreate)

	e.Logger.Fatal(e.Start(":3000"))
}
