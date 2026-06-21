package dto

type CreateRequest struct {
	CustomerID    uint   `json:"customer_id" validate:"required"`
	VehicleID     uint   `json:"vehicle_id" validate:"required"`
	RentStartDate string `json:"rent_start_date" validate:"required"`
	RentEndDate   string `json:"rent_end_date" validate:"required"`
	TotalPrice    int    `json:"total_price" validate:"required,gt=0"`
	Status        string `json:"status" validate:"required,oneof=active cancelled returned"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=active cancelled returned"`
}