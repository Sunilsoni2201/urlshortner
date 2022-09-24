package main

import (
	"fmt"
	"net/http"
	"time"
	"urlshortner/cache"

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
	e.POST("/shorturl", shorten)

	fmt.Println(time.Now().String())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func shorten(c echo.Context) error {

	type In struct {
		Url string `json:"url"`
	}
	in := In{}

	c.Bind(&in)

	urls := cache.Init()
	surl, err := urls.CreateShortURL(in.Url)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, surl)
}
