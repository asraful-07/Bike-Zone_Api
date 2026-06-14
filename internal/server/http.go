package server

import (
	"bike_zone_api/internal/config"
	"net/http"
	"os/user"

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
	if err := db.AutoMigrate(&user.User{}); err != nil {
		panic("failed to migrate database")
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.RequestLogger())

	// All Routes
	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Bike zone api running")
	})


	if err := e.Start(":" + cfg.PORT); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}