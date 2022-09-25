package app

import (
	"fmt"
	"net"
	"net/http"
	"urlshortner/pkg/utils"
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

	_ = c.Bind(&in)

	surl, err := u.svc.CreateShortURL(in.Url)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	serverPort := c.Echo().Listener.Addr().(*net.TCPAddr).Port
	shortURL := fmt.Sprintf("%v:%v/%s", utils.GetOutboundIP(), serverPort, surl)
	return c.String(http.StatusOK, shortURL)
}
