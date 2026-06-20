package dto

type CreateRequest struct {
	VehicleName        string `json:"vehicle_name" validate:"required,min=2,max=100"`
	Type               string `json:"type" validate:"required,oneof=car bike van SUV"`
	RegistrationNumber string `json:"registration_number" validate:"required"`
	DailyRentPrice     int    `json:"daily_rent_price" validate:"required,gt=0"`
	AvailabilityStatus string `json:"availability_status" validate:"required,oneof=available booked"`
	UserID             uint   `json:"user_id" validate:"required"`
}

type UpdateRequest struct {
	VehicleName        string `json:"vehicle_name" validate:"omitempty,min=2,max=100"`
	Type               string `json:"type" validate:"omitempty,oneof=car bike van SUV"`
	RegistrationNumber string `json:"registration_number" validate:"omitempty"`
	DailyRentPrice     int    `json:"daily_rent_price" validate:"omitempty,gt=0"`
	AvailabilityStatus string `json:"availability_status" validate:"omitempty,oneof=available booked"`
}
