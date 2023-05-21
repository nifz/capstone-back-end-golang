package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"errors"
)

type ReservationUsecase interface {
	AdminCreateReservation(input dtos.ReservationCreateInput) (dtos.ReservationCreateResponse, error)
}

type reservationUsecase struct {
	reservationRepo repositories.ReservationRepository
}

func NewReservationUsecase(reservationRepo repositories.ReservationRepository) ReservationUsecase {
	return &reservationUsecase{reservationRepo}
}

func (u *reservationUsecase) AdminCreateReservation(input dtos.ReservationCreateInput) (dtos.ReservationCreateResponse, error) {
	// admin create reservation
	var (
		reservation         models.Reservations
		reservationResponse dtos.ReservationCreateResponse
	)

	// reservation, _ = u.reservationRepo.AdminCreateReservation(reservation)
	// if reservation.ID == 0 {
	// 	return reservationResponse, errors.New("failed to create reservation")
	// }

	// pengecekan semua inputan
	if input.Name == "" || input.Province_name == "" || input.Regency_name == "" || input.District_name == "" || input.Village_name == "" || input.Postal_code == "" || input.Full_address == "" || input.Type == "" || input.Price == 0 || input.Thumbnail == "" || input.Description == "" || input.Tags == "" || input.Status == "" {
		return reservationResponse, errors.New("all fields must be filled")
	}

	reservationResponse.Name = reservation.Name
	reservationResponse.Province_name = reservation.Province_name
	reservationResponse.Regency_name = reservation.Regency_name
	reservationResponse.District_name = reservation.District_name
	reservationResponse.Village_name = reservation.Village_name
	reservationResponse.Postal_code = reservation.Postal_code
	reservationResponse.Full_address = reservation.Full_address
	reservationResponse.Type = string(reservation.Type)
	reservationResponse.Price = reservation.Price
	reservationResponse.Thumbnail = reservation.Thumbnail
	reservationResponse.Description = reservation.Description
	reservationResponse.Tags = reservation.Tags
	reservationResponse.Status = string(reservation.Status)

	return reservationResponse, nil

}
