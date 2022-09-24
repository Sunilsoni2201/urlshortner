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
	e.GET("/:surl", getOriginalURL)

	fmt.Println(time.Now().String())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

var urls cache.URLCacher = cache.Init()

func getOriginalURL(c echo.Context) error {

	surl := c.Param("surl")
	ourl, err := urls.GetActualURL(surl)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, err.Error())
	}

	c.Response().Header().Set("Location", ourl)

	return c.String(http.StatusMovedPermanently, "")
}

// Handler
func shorten(c echo.Context) error {

	type In struct {
		Url string `json:"url"`
	}
	in := In{}

	c.Bind(&in)

	surl, err := urls.CreateShortURL(in.Url)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, surl)
}
