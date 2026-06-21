package bookings

import (
	"errors"

	"gorm.io/gorm"
)

var ErrBookingNotFound = errors.New("booking not found")

type Repository interface {
	Create(booking *Booking) error
	GetAll() ([]*Booking, error)
	GetByID(id uint) (*Booking, error)
	GetByCustomerID(customerID uint) ([]*Booking, error)
	Update(booking *Booking) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create Booking
func (r *repository) Create(booking *Booking) error {
	return r.db.Create(booking).Error
}

// Get All Bookings (Admin)
func (r *repository) GetAll() ([]*Booking, error) {
	var bookings []*Booking

	err := r.db.
		Preload("Customer").
		Preload("Vehicle").
		Find(&bookings).Error

	if err != nil {
		return nil, err
	}

	return bookings, nil
}

// Get Booking By ID
func (r *repository) GetByID(id uint) (*Booking, error) {
	var booking Booking

	err := r.db.
		Preload("Customer").
		Preload("Vehicle").
		First(&booking, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrBookingNotFound
	}

	if err != nil {
		return nil, err
	}

	return &booking, nil
}

// Get Bookings By Customer ID
func (r *repository) GetByCustomerID(customerID uint) ([]*Booking, error) {
	var bookings []*Booking

	err := r.db.
		Preload("Customer").
		Preload("Vehicle").
		Where("customer_id = ?", customerID).
		Find(&bookings).Error

	if err != nil {
		return nil, err
	}

	return bookings, nil
}

// Update Booking
func (r *repository) Update(booking *Booking) error {
	return r.db.Save(booking).Error
}