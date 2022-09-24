package app

import (
	"fmt"
	"time"
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

	db := db.NewMemoryDb("~/memorydb.json")
	svc := services.NewURLShortner(db)
	handler := NewURLShortnerHandler(svc)

	// Routes
	e.POST("/shorturl", handler.shorten)
	e.GET("/:surl", handler.getOriginalURL)

	fmt.Println(time.Now().String())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
