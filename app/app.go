package app

import (
	"urlshortner/db"
	"urlshortner/services"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := db.NewFileDb("")
	svc := services.NewURLShortner(db)
	handler := NewURLShortnerHandler(svc)

	// Routes
	e.POST("/shorturl", handler.shorten)
	e.GET("/:surl", handler.getOriginalURL)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
