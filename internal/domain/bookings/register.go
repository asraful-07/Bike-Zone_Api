package bookings

import (
	"bike_zone_api/internal/config"
	"bike_zone_api/internal/middlewares"
	"time"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"

	"bike_zone_api/internal/auth"
	"bike_zone_api/internal/domain/vehicles"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	// repos
	bookingRepo := NewRepository(db)
	vehicleRepo := vehicles.NewRepository(db)

	// auth
	jwtService := auth.NewJWTService(cfg.JWTSecretKey, 24*time.Hour)

	// service
	bookingService := NewService(bookingRepo, vehicleRepo)

	// handler
	bookingHandler := NewHandler(bookingService)

	api := e.Group("/api/v1/bookings")

	// AUTH PROTECTED ROUTES
	protected := api.Group("")
	protected.Use(middlewares.AuthMiddleware(jwtService))

	// Create Booking (Customer/Admin)
	protected.POST("", bookingHandler.CreateBooking)

	// Get My Bookings (Customer)
	protected.GET("/me", bookingHandler.GetMyBookings)

	// Update Booking Status (Customer/Admin)
	protected.PUT("/:id", bookingHandler.UpdateStatus)

	// ADMIN ONLY ROUTES (optional middleware later)
	protected.GET("", bookingHandler.GetAll)
}