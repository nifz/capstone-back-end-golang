package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type ReservationRepository interface {
	AdminCreateReservation(reservation models.Reservations) (models.Reservations, error)
	// GetAllReservation() ([]models.Reservations, error)
	// GetReservationById(id int) (models.Reservations, error)
}

type reservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) ReservationRepository {
	return &reservationRepository{db}
}

func (r *reservationRepository) AdminCreateReservation(reservation models.Reservations) (models.Reservations, error) {
	err := r.db.Create(&reservation).Error
	return reservation, err
}
