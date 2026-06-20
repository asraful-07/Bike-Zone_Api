package vehicles

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	api := e.Group("/api/v1/vehicle")

	api.POST("", handler.CreateVehicle)
	api.GET("", handler.GetVehicles)
	api.GET("/:id", handler.GetVehiclesByID)
	api.PATCH("/:id", handler.UpdateVehicle)
}
