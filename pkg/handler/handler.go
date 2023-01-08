package handler

import (
	"github.com/alikud/ads-microservice/domain"
	"github.com/alikud/ads-microservice/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
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

	h.echo.GET("/offers", h.getOffers)
	h.echo.GET("/offers/:id", h.getOfferById)
	h.echo.POST("/ad", h.createOffer)

}

func (h *Handler) getOffers(c echo.Context) error {
	type getOffersResponse struct {
		Result []domain.Offer
	}

	var page int
	p := c.QueryParam("page")
	if p == "" {
		page = 0
	}
	page, _ = strconv.Atoi(p)

	//want get column:desc string
	//https://docs.oracle.com/en/cloud/saas/service/18c/cxsvc/c_osvc_sorting.html
	orderBy := c.QueryParam("orderBy")

	limit := 2
	offset := page * limit
	offers, _ := h.Service.GetAll(limit, offset, orderBy)

	resp := getOffersResponse{
		Result: offers,
	}

	if resp.Result == nil {
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) getOfferById(c echo.Context) error {

	type getOffersResponse struct {
		Result domain.Offer
	}
	offId := c.Param("id")
	offer, _ := h.Service.GetById(offId)
	return c.JSON(http.StatusOK, offer)
}

func (h *Handler) createOffer(c echo.Context) error {
	return c.String(http.StatusOK, "Not implemented")
}
