package handler

import (
	"fmt"
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
	h.echo.GET("/ad/:id", h.getOfferById)
	h.echo.POST("/ad", h.createOffer)

}

type Links struct {
}

func (h *Handler) getOffers(c echo.Context) error {
	type getOffersResponse struct {
		Self   string `json:"self"`
		Next   string `json:"next"`
		Result []domain.Offer
	}

	name := c.QueryParam("page")
	var page int
	page, err := strconv.Atoi(name)
	if err != nil {
		page = 1
	}

	//want get column:desc string
	orderBy := c.QueryParam("orderBy")

	//https://docs.oracle.com/en/cloud/saas/service/18c/cxsvc/c_osvc_sorting.html
	limit := 2
	offset := page * limit
	offers, _ := h.Service.GetAll(limit, offset, orderBy)

	nextPage := page + 1
	resp := getOffersResponse{
		Self:   fmt.Sprintf("/offers?=page%d", page),
		Next:   fmt.Sprintf("/offers?page=%d", nextPage),
		Result: offers,
	}

	if resp.Result == nil {
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) getOfferById(c echo.Context) error {
	return c.String(http.StatusOK, "Not implemented")
}

func (h *Handler) createOffer(c echo.Context) error {
	return c.String(http.StatusOK, "Not implemented")
}
