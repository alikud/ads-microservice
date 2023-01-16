package handler

import (
	"github.com/alikud/ads-microservice/domain"
	_ "github.com/alikud/ads-microservice/pkg/handler/docs"
	"github.com/alikud/ads-microservice/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
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

	h.echo.GET("/", h.HealthCheck)
	h.echo.GET("/offer", h.GetOffers)
	h.echo.GET("/offer/:id", h.GetOfferById)
	h.echo.POST("/offer", h.CreateOffer)
	h.echo.DELETE("/offer", h.DeleteOffer)
	h.echo.PATCH("/offer", h.UpdateOffer)

	h.echo.GET("/swagger/*", echoSwagger.WrapHandler)
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (h *Handler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}

// GetOffers godoc
// @Summary Get all offers by page
// @Description Get offers by page with page query param
// @Tags offers
// @Accept json
// @Produce json
// @Success 201
// @Failure 400
// @Router /offer [get]
func (h *Handler) GetOffers(c echo.Context) error {
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
	offers, err := h.Service.GetAll(limit, offset, orderBy)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
		}{})
	}

	return c.JSON(http.StatusOK, &offers)
}

// GetOfferById godoc
// @Summary Get offer by id
// @Description Get offer by unique id
// @Tags offers
// @Accept json
// @Produce json
// @Success 201
// @Failure 400
// @Router /offer/:id [get]
func (h *Handler) GetOfferById(c echo.Context) error {

	offId := c.Param("id")
	offer, err := h.Service.GetById(offId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"Error": "Wrong id"})
	}
	return c.JSON(http.StatusOK, &offer)
}

// CreateOffer godoc
// @Summary Create offer
// @Description Create offer with json input body
// @Tags offers
// @Accept json
// @Produce json
// @Success 201
// @Failure 400
// @Router /offer [post]
func (h *Handler) CreateOffer(c echo.Context) error {
	var o domain.Offer

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := o.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"Error": "Bad data format"})
	}

	id, err := h.Service.Create(o)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"Error": err.Error()})
	}

	type createOfferResponse struct {
		Id string
	}

	resp := createOfferResponse{Id: id}
	return c.JSON(http.StatusOK, resp)
}

// DeleteOffer godoc
// @Summary Delete offer
// @Description Delete offer with id
// @Tags offers
// @Accept json
// @Produce json
// @Success 201
// @Failure 400
// @Router /offer [delete]
func (h *Handler) DeleteOffer(c echo.Context) error {
	offId := c.QueryParam("id")
	err := h.Service.Delete(offId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"Error": err.Error()})
	}
	return c.JSON(http.StatusOK, offId)
}

// UpdateOffer godoc
// @Summary Update offer
// @Description Update offer with id and json body
// @Tags offers
// @Accept json
// @Produce json
// @Success 201
// @Failure 400
// @Router /offer [delete]
func (h *Handler) UpdateOffer(c echo.Context) error {
	var o domain.Offer

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	offId := c.QueryParam("id")
	if offId == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"Error": "Id is empty"})
	}

	err := h.Service.Update(offId, &o)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusAccepted, offId)
}
