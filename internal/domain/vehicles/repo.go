package vehicles

import (
	"errors"

	"gorm.io/gorm"
)

var ErrVehicleNotFound = errors.New("vehicle not found")

type Repository interface {
	Create(vehicle *Vehicle) error
	GetAll() ([]*Vehicle, error)
	GetByID(vehicleID uint) (*Vehicle, error)
	GetByRegistrationNumber(registrationNumber string) (*Vehicle, error)
	Update(vehicle *Vehicle) error
	Delete(vehicle *Vehicle) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(vehicle *Vehicle) error {
	return r.db.Create(vehicle).Error
}

func (r *repository) GetAll() ([]*Vehicle, error) {
	var vehicles []*Vehicle

	if err := r.db.Find(&vehicles).Error; err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (r *repository) GetByID(vehicleID uint) (*Vehicle, error) {
	var vehicle Vehicle

	err := r.db.First(&vehicle, vehicleID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrVehicleNotFound
		}
		return nil, err
	}

	return &vehicle, nil
}

func (r *repository) GetByRegistrationNumber(reg string) (*Vehicle, error) {
	var vehicle Vehicle

	err := r.db.Where("registration_number = ?", reg).First(&vehicle).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrVehicleNotFound
		}
		return nil, err
	}

	return &vehicle, nil
}

func (r *repository) Update(vehicle *Vehicle) error {
	return r.db.Save(vehicle).Error
}

func (r *repository) Delete(vehicle *Vehicle) error {
	return r.db.Delete(vehicle).Error
}