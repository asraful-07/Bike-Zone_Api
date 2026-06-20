package dto

type Response struct {
	ID                 uint   `json:"id"`
	UserID             uint   `json:"user_id"`
	VehicleName        string `json:"vehicle_name"`
	Type               string `json:"type"`
	RegistrationNumber string `json:"registration_number"`
	DailyRentPrice     int    `json:"daily_rent_price"`
	AvailabilityStatus string `json:"availability_status"`
	CreatedAt          string `json:"created_at"`
}