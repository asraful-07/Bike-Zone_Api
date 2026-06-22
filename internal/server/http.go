package server

import (
	"bike_zone_api/internal/config"
	"bike_zone_api/internal/domain/bookings"
	"bike_zone_api/internal/domain/users"
	"bike_zone_api/internal/domain/vehicles"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"gorm.io/gorm"
)

type CustomValidator struct {
  validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
  return cv.validator.Struct(i)
}

func StartServer( db *gorm.DB, cfg *config.Config) {
	// Ensure fresh schema for development: drop and recreate tables
	// NOTE: This will delete existing data in these tables.
	if err := db.AutoMigrate(&users.User{}, &vehicles.Vehicle{}, &bookings.Booking{}); err != nil {
		panic("failed to migrate database")
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.RequestLogger())

	
	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Bike zone api running")
	})

    // All Routes
	users.RegisterRoutes(e, db, cfg)
	vehicles.RegisterRoutes(e, db)
	bookings.RegisterRoutes(e, db, cfg)

	if err := e.Start(":" + cfg.PORT); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}