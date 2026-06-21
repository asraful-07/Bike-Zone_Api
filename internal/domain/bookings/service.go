package bookings

import (
	"errors"
	"time"

	"bike_zone_api/internal/domain/bookings/dto"
	"bike_zone_api/internal/domain/vehicles"
)

var (
	ErrVehicleNotAvailable = errors.New("vehicle is not available")
	ErrInvalidRentDate     = errors.New("rent end date must be after rent start date")
	ErrInvalidStatus   = errors.New("invalid status")
)

type service struct {
	repo        Repository
	vehicleRepo vehicles.Repository
}

func NewService(repo Repository, vehicleRepo vehicles.Repository) *service {
	return &service{
		repo:        repo,
		vehicleRepo: vehicleRepo,
	}
}

// Create Booking
func (s *service) CreateBooking(req dto.CreateRequest) (*dto.Response, error) {

	// Get Vehicle
	vehicle, err := s.vehicleRepo.GetByID(req.VehicleID)
	if err != nil {
		return nil, err
	}

	// Check Availability
	if vehicle.AvailabilityStatus != vehicles.VehicleAvailable {
		return nil, ErrVehicleNotAvailable
	}

	// Parse dates
	startDate, err := time.Parse("2006-01-02", req.RentStartDate)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", req.RentEndDate)
	if err != nil {
		return nil, err
	}

	// Validate Date
	if !endDate.After(startDate) {
		return nil, ErrInvalidRentDate
	}

	// Calculate Number Of Days
	days := int(endDate.Sub(startDate).Hours() / 24)
	if days <= 0 {
		days = 1
	}

	// Calculate Total Price
	totalPrice := vehicle.DailyRentPrice * days

	booking := Booking{
		CustomerID:    req.CustomerID,
		VehicleID:     req.VehicleID,
		RentStartDate: req.RentStartDate,
		RentEndDate:   req.RentEndDate,
		TotalPrice:    totalPrice,
		Status:        BookingActive,
	}

	// Save Booking
	if err := s.repo.Create(&booking); err != nil {
		return nil, err
	}

	// Update Vehicle Status
	vehicle.AvailabilityStatus = vehicles.VehicleBooked

	if err := s.vehicleRepo.Update(vehicle); err != nil {
		return nil, err
	}

	return booking.ToResponse(), nil
}

// Get All Bookings (Admin)
func (s *service) GetAll() ([]*dto.Response, error) {

	bookings, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.Response, 0, len(bookings))

	for _, booking := range bookings {
		responses = append(responses, booking.ToResponse())
	}

	return responses, nil
}

// Get Bookings By Customer (Customer only)
func (s *service) GetByCustomerID(customerID uint) ([]*dto.Response, error) {

	bookings, err := s.repo.GetByCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.Response, 0, len(bookings))

	for _, booking := range bookings {
		responses = append(responses, booking.ToResponse())
	}

	return responses, nil
}

func (s *service) UpdateStatus(bookingID uint, req dto.UpdateStatusRequest) (*dto.Response, error) {

	// Get booking
	booking, err := s.repo.GetByID(bookingID)
	if err != nil {
		return nil, err
	}

	// Update status
	booking.Status = req.Status

	// If returned → vehicle available
	if req.Status == BookingReturned {
		vehicle, err := s.vehicleRepo.GetByID(booking.VehicleID)
		if err != nil {
			return nil, err
		}

		vehicle.AvailabilityStatus = vehicles.VehicleAvailable

		if err := s.vehicleRepo.Update(vehicle); err != nil {
			return nil, err
		}
	}

	// If cancelled → vehicle available
	if req.Status == BookingCancelled {
		vehicle, err := s.vehicleRepo.GetByID(booking.VehicleID)
		if err != nil {
			return nil, err
		}

		vehicle.AvailabilityStatus = vehicles.VehicleAvailable

		if err := s.vehicleRepo.Update(vehicle); err != nil {
			return nil, err
		}
	}

	// Save booking
	if err := s.repo.Update(booking); err != nil {
		return nil, err
	}

	return booking.ToResponse(), nil
}