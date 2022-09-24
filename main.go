package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", shorten)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func shorten(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from URL shortner!")
}
