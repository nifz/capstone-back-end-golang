package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type ReservationRepository interface {
	AdminCreateReservation(reservation models.Reservations) (models.Reservations, error)
	GetAllReservation(page, limit int) ([]models.Reservations, int, error)
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

func (r *reservationRepository) GetAllReservation(page, limit int) ([]models.Reservations, int, error) {
	var (
		reservations []models.Reservations
		count        int64
	)

	err := r.db.Model(&models.Reservations{}).Count(&count).Error
	if err != nil {
		return reservations, int(count), err
	}

	offset := (page - 1) * limit

	err = r.db.Limit(limit).Offset(offset).Find(&reservations).Error

	return reservations, int(count), err
}

// func (r *reservationRepository) GetAllReservation(reservation models.Reservations) (models.Reservations, error) {
// 	var (
// 		reservations []models.Reservations
// 		count        int64
// 	)

// 	err := r.db.Find(&reservations).Count(&count).Error
// 	if err != nil {
// 		return reservations,int(count), err
// 	}
// 	offset := (page - 1) * limit

// 	err = r.db.Limit(limit).Offset(offset).Find(&reservations).Error

// 	return reservations, int(count), err
// }
