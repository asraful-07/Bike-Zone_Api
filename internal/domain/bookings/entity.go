package bookings

import (
	"bike_zone_api/internal/domain/bookings/dto"
	"bike_zone_api/internal/domain/users"
	"bike_zone_api/internal/domain/vehicles"

	"gorm.io/gorm"
)

const (
	BookingActive    = "active"
	BookingCancelled = "cancelled"
	BookingReturned  = "returned"
)

type Booking struct {
	gorm.Model
	CustomerID    uint `json:"customer_id" gorm:"not null"`
	Customer      users.User `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
	VehicleID     uint `json:"vehicle_id" gorm:"not null"`
	Vehicle       vehicles.Vehicle `json:"vehicle,omitempty" gorm:"foreignKey:VehicleID"`
	RentStartDate string `json:"rent_start_date" gorm:"type:date;not null"`
	RentEndDate   string `json:"rent_end_date" gorm:"type:date;not null"`
	TotalPrice    int `json:"total_price" gorm:"not null"`
	Status        string `json:"status" gorm:"type:varchar(20);default:'active';not null"`
}

func (b *Booking) ToResponse() *dto.Response {
	return &dto.Response{
		ID:             b.ID,
		CustomerID:     b.CustomerID,
		VehicleID:      b.VehicleID,
		RentStartDate:  b.RentStartDate,
		RentEndDate:    b.RentEndDate,
		TotalPrice:     b.TotalPrice,
		Status:         b.Status,
		CreatedAt:      b.CreatedAt.String(),
	}
}