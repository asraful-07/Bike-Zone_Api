package dto

type Response struct {
	ID            uint   `json:"id"`
	CustomerID    uint   `json:"customer_id"`
	VehicleID     uint   `json:"vehicle_id"`
	RentStartDate string `json:"rent_start_date"`
	RentEndDate   string `json:"rent_end_date"`
	TotalPrice    int    `json:"total_price"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
}