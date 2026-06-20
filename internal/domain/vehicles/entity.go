package vehicles

import (
	"bike_zone_api/internal/domain/users"
	"bike_zone_api/internal/domain/vehicles/dto"

	"gorm.io/gorm"
)

const (
	VehicleTypeCar  = "car"
	VehicleTypeBike = "bike"
	VehicleTypeVan  = "van"
	VehicleTypeSUV  = "SUV"
)

const (
	VehicleAvailable = "available"
	VehicleBooked    = "booked"
)

type Vehicle struct {
	gorm.Model
	VehicleName        string     `json:"vehicle_name" gorm:"type:varchar(100);not null"`
	Type               string     `json:"type" gorm:"type:varchar(20);not null"`
	RegistrationNumber string     `json:"registration_number" gorm:"type:varchar(100);unique;not null"`
	DailyRentPrice     int        `json:"daily_rent_price" gorm:"not null"`
	AvailabilityStatus string     `json:"availability_status" gorm:"type:varchar(20);default:'available';not null"`
	UserID             uint       `json:"user_id" gorm:"not null"`
	User               users.User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

func (v *Vehicle) ToResponse() *dto.Response {
	return &dto.Response{
		ID:                 v.ID,
		VehicleName:        v.VehicleName,
		Type:               v.Type,
		RegistrationNumber: v.RegistrationNumber,
		DailyRentPrice:     v.DailyRentPrice,
		AvailabilityStatus: v.AvailabilityStatus,
		CreatedAt:          v.CreatedAt.String(),
	}
}