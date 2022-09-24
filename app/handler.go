package app

import (
	"fmt"
	"net/http"
	"urlshortner/services"

	"github.com/labstack/echo"
)

type urlShortnerHandler struct {
	svc services.URLShortner
}

func NewURLShortnerHandler(svc services.URLShortner) *urlShortnerHandler {
	return &urlShortnerHandler{
		svc: svc,
	}
}

func (u *urlShortnerHandler) getOriginalURL(c echo.Context) error {

	surl := c.Param("surl")
	ourl, err := u.svc.GetActualURL(surl)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, err.Error())
	}

	c.Response().Header().Set("Location", ourl)

	return c.String(http.StatusMovedPermanently, "")
}

// Handler
func (u *urlShortnerHandler) shorten(c echo.Context) error {

	type In struct {
		Url string `json:"url"`
	}
	in := In{}

	c.Bind(&in)

	surl, err := u.svc.CreateShortURL(in.Url)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "justurl.com:8080/"+surl)
}
