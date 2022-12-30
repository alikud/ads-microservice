package handler

import (
	"github.com/alikud/ads-microservice/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	h.echo.Use(middleware.Logger())
	h.echo.Use(middleware.Recover())
	h.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	h.echo.GET("/ad", h.getOffer)
	h.echo.GET("/ad/:id", h.getOfferById)
	h.echo.POST("/ad", h.createOffer)

}

func (h *Handler) getOffer(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "Not implemented")
}

func (h *Handler) getOfferById(c echo.Context) error {
	return c.String(http.StatusOK, "Not implemented")
}

func (h *Handler) createOffer(c echo.Context) error {
	return c.String(http.StatusOK, "Not implemented")
}
