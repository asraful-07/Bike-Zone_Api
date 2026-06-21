package bookings

import (
	"bike_zone_api/internal/domain/bookings/dto"
	"bike_zone_api/internal/domain/vehicles"
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

func bookingErrorResponse(c *echo.Context, err error) error {

	if errors.Is(err, ErrBookingNotFound) {
		return c.JSON(http.StatusNotFound, http_response.Error{
			Code:    http.StatusNotFound,
			Message: "Booking not found",
		})
	}

	if errors.Is(err, vehicles.ErrVehicleNotFound) {
		return c.JSON(http.StatusNotFound, http_response.Error{
			Code:    http.StatusNotFound,
			Message: "Vehicle not found",
		})
	}

	if errors.Is(err, ErrVehicleNotAvailable) {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Vehicle not available",
		})
	}

	if errors.Is(err, ErrInvalidRentDate) {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid rent date",
		})
	}

	if errors.Is(err, ErrInvalidStatus) {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid status",
		})
	}

	return c.JSON(http.StatusInternalServerError, http_response.Error{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}

func (h *handler) CreateBooking(c *echo.Context) error {

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

	response, err := h.service.CreateBooking(req)
	if err != nil {
		return bookingErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *handler) GetAll(c *echo.Context) error {

	response, err := h.service.GetAll()
	if err != nil {
		return bookingErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetMyBookings(c *echo.Context) error {

	customerID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, http_response.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized user",
		})
	}

	response, err := h.service.GetByCustomerID(customerID)
	if err != nil {
		return bookingErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateStatus(c *echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid booking id",
			Details: err.Error(),
		})
	}

	var req dto.UpdateStatusRequest

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

	response, err := h.service.UpdateStatus(uint(id), req)
	if err != nil {
		return bookingErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, response)
}