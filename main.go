package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})
	e.GET("/fatur", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello fatur")
	})
	e.GET("/home", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ini halaman home")
	})
	e.Start(":8000")
}
