package dto

type CustomerResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type VehicleResponse struct {
	ID                 uint   `json:"id"`
	VehicleName        string `json:"vehicle_name"`
	Type               string `json:"type"`
	RegistrationNumber string `json:"registration_number"`
	DailyRentPrice     int    `json:"daily_rent_price"`
	AvailabilityStatus string `json:"availability_status"`
}

type Response struct {
	ID            uint              `json:"id"`
	CustomerID    uint              `json:"customer_id"`
	Customer      *CustomerResponse `json:"customer,omitempty"`
	VehicleID     uint              `json:"vehicle_id"`
	Vehicle       *VehicleResponse  `json:"vehicle,omitempty"`
	RentStartDate string            `json:"rent_start_date"`
	RentEndDate   string            `json:"rent_end_date"`
	TotalPrice    int               `json:"total_price"`
	Status        string            `json:"status"`
	CreatedAt     string            `json:"created_at"`
}