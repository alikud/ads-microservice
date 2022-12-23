package handler

import (
	"github.com/alikud/ads-microservice/pkg/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Service *service.Service
	echo    *echo.Echo
}

func NewHandler(service *service.Service, echo *echo.Echo) *Handler {
	return &Handler{Service: service, echo: echo}
}

func (h *Handler) InitRoutes() {

	h.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	h.echo.GET("/ad", h.getAd)
	h.echo.GET("/ad/:id", h.getAdById)
	h.echo.POST("/ad", h.createAd)

}

func (h *Handler) getAd(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "Not implemented")
}

func (h *Handler) getAdById(c echo.Context) error {
	return c.String(http.StatusOK, "Not implemented")
}

func (h *Handler) createAd(c echo.Context) error {
	return c.String(http.StatusOK, "Not implemented")
}
