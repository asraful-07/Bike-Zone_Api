package vehicles

import (
	"errors"

	"bike_zone_api/internal/domain/vehicles/dto"
)

var ErrVehicleAlreadyExists = errors.New("vehicle with this registration number already exists")

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}
 
// Create Vehicle
func (s *service) CreateVehicle(req dto.CreateRequest) (*dto.Response, error) {

	// Check registration number already exists
	existing, err := s.repo.GetByRegistrationNumber(req.RegistrationNumber)
	if err == nil && existing != nil {
		return nil, ErrVehicleAlreadyExists
	}

	if err != nil && !errors.Is(err, ErrVehicleNotFound) {
		return nil, err
	}

	vehicle := Vehicle{
		VehicleName:        req.VehicleName,
		Type:               req.Type,
		RegistrationNumber: req.RegistrationNumber,
		DailyRentPrice:     req.DailyRentPrice,
		AvailabilityStatus: VehicleAvailable,
		UserID:             req.UserID,
	}

	if err := s.repo.Create(&vehicle); err != nil {
		return nil, err
	}

	return vehicle.ToResponse(), nil
}

// Get All Vehicles
func (s *service) GetVehicles() ([]dto.Response, error) {
	vehicles, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.Response

	for _, vehicle := range vehicles {
		responses = append(responses, *vehicle.ToResponse())
	}

	return responses, nil
}

// Get Vehicle By ID
func (s *service) GetVehicleByID(vehicleID uint) (*dto.Response, error) {
	vehicle, err := s.repo.GetByID(vehicleID)
	if err != nil {
		return nil, err
	}

	return vehicle.ToResponse(), nil
}

// Update Vehicle
func (s *service) UpdateVehicle(vehicleID uint, req dto.UpdateRequest) (*dto.Response, error) {

	vehicle, err := s.repo.GetByID(vehicleID)
	if err != nil {
		return nil, err
	}

	if req.VehicleName != "" {
		vehicle.VehicleName = req.VehicleName
	}

	if req.Type != "" {
		vehicle.Type = req.Type
	}

	if req.RegistrationNumber != "" &&
		req.RegistrationNumber != vehicle.RegistrationNumber {

		existing, err := s.repo.GetByRegistrationNumber(req.RegistrationNumber)
		if err == nil && existing != nil {
			return nil, ErrVehicleAlreadyExists
		}

		if err != nil && !errors.Is(err, ErrVehicleNotFound) {
			return nil, err
		}

		vehicle.RegistrationNumber = req.RegistrationNumber
	}

	if req.DailyRentPrice != 0 {
		vehicle.DailyRentPrice = req.DailyRentPrice
	}

	if req.AvailabilityStatus != "" {
		vehicle.AvailabilityStatus = req.AvailabilityStatus
	}

	if err := s.repo.Update(vehicle); err != nil {
		return nil, err
	}

	return vehicle.ToResponse(), nil
}

// Delete Vehicle
func (s *service) DeleteVehicle(vehicleID uint) error {
	vehicle, err := s.repo.GetByID(vehicleID)
	if err != nil {
		return err
	}

	return s.repo.Delete(vehicle)
}