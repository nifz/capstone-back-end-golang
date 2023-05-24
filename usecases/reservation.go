package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"errors"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
)

type ReservationUsecase interface {
	AdminCreateReservation(input dtos.ReservationCreateInput, image *multipart.FileHeader) (dtos.ReservationCreateResponse, error)
	GetAllReservation(page, limit int) ([]dtos.ReservationResponse, int, error)
}

type reservationUsecase struct {
	reservationRepo      repositories.ReservationRepository
	reservationImageRepo repositories.ReservationImageRepository
}

type Image struct {
	File     multipart.File
	Filename string
}

func NewReservationUsecase(reservationRepo repositories.ReservationRepository, reservationImageRepo repositories.ReservationImageRepository) ReservationUsecase {
	return &reservationUsecase{
		reservationRepo:      reservationRepo,
		reservationImageRepo: reservationImageRepo,
	}
}

func (u *reservationUsecase) AdminCreateReservation(input dtos.ReservationCreateInput, image *multipart.FileHeader) (dtos.ReservationCreateResponse, error) {
	// Admin create reservation
	var (
		reservation         models.Reservations
		reservationResponse dtos.ReservationCreateResponse
	)

	// Pengecekan semua inputan
	if input.Name == "" || input.Province_name == "" || input.Regency_name == "" || input.District_name == "" || input.Village_name == "" || input.Postal_code == "" || input.Full_address == "" || input.Type == "" || input.Price == 0 || input.Description == "" || input.Tags == "" || input.Status == "" {
		return reservationResponse, errors.New("all fields must be filled")
	}

	// Convert input.Type string to typeReservation
	var reservationType models.TypeReservation
	switch input.Type {
	case "hotels":
		reservationType = models.Hotels
	case "villa":
		reservationType = models.Villa
	case "guest_house":
		reservationType = models.GustHouse
	default:
		return reservationResponse, errors.New("invalid reservation type")
	}

	// Convert input.Status string to status
	var reservationStatus models.Status
	switch input.Status {
	case "available":
		reservationStatus = models.Avaliable
	case "unavailable":
		reservationStatus = models.Unavaliable
	default:
		return reservationResponse, errors.New("invalid reservation status")
	}

	// Map input to reservation entity
	reservation.Name = input.Name
	reservation.Province_name = input.Province_name
	reservation.Regency_name = input.Regency_name
	reservation.District_name = input.District_name
	reservation.Village_name = input.Village_name
	reservation.Postal_code = input.Postal_code
	reservation.Full_address = input.Full_address
	reservation.Type = reservationType
	reservation.Price = input.Price
	reservation.Thumbnail = input.Thumbnail
	reservation.Description = input.Description
	reservation.Tags = input.Tags
	reservation.Status = reservationStatus

	// Create the reservation
	createdReservation, err := u.reservationRepo.AdminCreateReservation(reservation)
	if err != nil {
		return reservationResponse, err
	}

	// Handle image upload
	if image != nil {
		imageName := GenerateFileName(image.Filename)
		imagePath := "./images/" + imageName
		err = saveImage(image, imagePath)
		if err != nil {
			return reservationResponse, err
		}
		// Create the reservation image entity
		reservationImage := models.ReservationImages{
			ReservationId: createdReservation.ID,
			Image:         imageName,
		}
		// Save the reservation image
		err = u.reservationImageRepo.SaveReservationImage(reservationImage)
		if err != nil {
			return reservationResponse, err
		}
		// Assign image path to the response
		reservationResponse.Image = reservationImage.Image
	}

	reservationResponse.ID = createdReservation.ID
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

func (u *reservationUsecase) GetAllReservation(page, limit int) ([]dtos.ReservationResponse, int, error) {
	// Panggil metode GetAllReservation dari repository untuk mendapatkan data reservasi
	reservations, total, err := u.reservationRepo.GetAllReservation(page, limit)
	if err != nil {
		return nil, 0, err
	}

	// Membuat slice untuk menyimpan responsenya
	reservationResponses := make([]dtos.ReservationResponse, len(reservations))

	// Mengonversi data reservasi menjadi responsenya
	for i, reservation := range reservations {
		// Mendapatkan data ReservationImages terkait
		reservationImages, err := u.reservationImageRepo.GetReservationImagesByReservationID(int(reservation.ID))
		if err != nil {
			return nil, 0, err
		}

		// Menyimpan data gambar pertama (jika ada)
		// var image string
		// if len(reservationImages) > 0 {
		// 	image = reservationImages[0].Image
		// }
		reservationResponses[i] = dtos.ReservationResponse{
			ReservationID: reservation.ID,
			Name:          reservation.Name,
			Province_name: reservation.Province_name,
			Regency_name:  reservation.Regency_name,
			District_name: reservation.District_name,
			Village_name:  reservation.Village_name,
			Postal_code:   reservation.Postal_code,
			Full_address:  reservation.Full_address,
			Type:          string(reservation.Type),
			Price:         reservation.Price,
			Thumbnail:     reservation.Thumbnail,
			Description:   reservation.Description,
			Tags:          reservation.Tags,
			Status:        string(reservation.Status),
			Image:         reservationImages.Image,
		}
	}

	return reservationResponses, total, nil
}
func saveImage(image *multipart.FileHeader, path string) error {
	src, err := image.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func RandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateFileName(fileName string) string {
	return "Image-" + RandomString(6) + filepath.Ext(fileName)
}
