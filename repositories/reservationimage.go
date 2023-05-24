package repositories

import (
	"back-end-golang/models"
	"fmt"

	"gorm.io/gorm"
)

type ReservationImageRepository interface {
	SaveReservationImage(image models.ReservationImages) error
	GetReservationImagesByReservationID(reservationID int) (models.ReservationImages, error)
}

type reservationImageRepository struct {
	db *gorm.DB
}

func NewReservationImageRepository(db *gorm.DB) *reservationImageRepository {
	return &reservationImageRepository{db}
}

func (r *reservationImageRepository) SaveReservationImage(image models.ReservationImages) error {
	fmt.Println("image", image)
	err := r.db.Create(&image).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *reservationImageRepository) GetReservationImagesByReservationID(reservationID int) (models.ReservationImages, error) {
	var image models.ReservationImages
	err := r.db.Where("reservation_id = ?", reservationID).First(&image).Error
	if err != nil {
		return models.ReservationImages{}, err
	}
	return image, nil
}
