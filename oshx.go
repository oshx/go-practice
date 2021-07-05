package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	echo.NotFoundHandler = notFound

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", hello)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Echo!")
}

func notFound(c echo.Context) error {
	return c.String(http.StatusNotFound, "Not Found, 404!")
}
