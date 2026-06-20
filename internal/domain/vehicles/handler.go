package vehicles

import (
	"bike_zone_api/internal/domain/vehicles/dto"
	"bike_zone_api/internal/http_response"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(s *service) *handler {
	return &handler{service: s}
}

func vehicleErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err, ErrVehicleNotFound) {
		return c.JSON(http.StatusNotFound, http_response.Error{
			Code:    http.StatusNotFound,
			Message: "Vehicle not found",
		})
	}

	return c.JSON(http.StatusInternalServerError, http_response.Error{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}

func (h *handler) CreateVehicle(c *echo.Context) error {
	var req dto.CreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.CreateVehicle(req)
	if err != nil {
		return vehicleErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *handler) GetVehicles(c *echo.Context) error {
	events, err := h.service.GetVehicles()
	if err != nil {
		return vehicleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, events)
}

func (h *handler) GetVehiclesByID(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid event id",
			Details: err.Error(),
		})
	}

	response, err := h.service.GetVehicleByID(uint(id)) 

	if err != nil {
		return vehicleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateVehicle(c *echo.Context) error {
	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid event id",
			Details: err.Error(),
		})
	}

	var req dto.UpdateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.UpdateVehicle(uint(eventId), req)
	if err != nil {
		return vehicleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, response)
}
