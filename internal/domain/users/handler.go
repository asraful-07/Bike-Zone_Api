package users

import (
	"bike_zone_api/internal/domain/users/dto"
	"bike_zone_api/internal/http_response"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateUser(c *echo.Context) error {
	var req dto.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.CreateUser(req)
	if err != nil {
		// Handle duplicate email with 409 Conflict, not 500
		if err.Error() == "email already exists" {
			return c.JSON(http.StatusConflict, http_response.Error{
				Code:    http.StatusConflict,
				Message: "Email already registered",
				Details: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, http_response.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *handler) LoginUser(c *echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, http_response.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.LoginUser(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, http_response.Error{
			Code:    http.StatusUnauthorized,
			Message: "Invalid credentials",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetMe(c *echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, http_response.Error{
			Code:    http.StatusUnauthorized,
			Message: "Cannot get user information",
			Details: "missing user id in context",
		})
	}

	// Return all available claims from JWT context
	return c.JSON(http.StatusOK, dto.Response{
		ID:    userID,
		Name:  c.Get("user_name").(string),
		Email: c.Get("user_email").(string),
		Role:  c.Get("user_role").(string),
		Phone: c.Get("user_phone").(string),
		CreatedAt: c.Get("create_At").(string),
	})
}