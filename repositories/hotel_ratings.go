package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type HotelRatingsRepository interface {
	CreateHotelRating(hotelRating models.HotelRating) (models.HotelRating, error)
	GetHotelRatingsByHotelID(id uint) (map[int]int, []models.HotelRating, error)
}

type hotelRatingsRepository struct {
	db *gorm.DB
}

func NewHotelRatingsRepository(db *gorm.DB) HotelRatingsRepository {
	return &hotelRatingsRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *hotelRatingsRepository) CreateHotelRating(hotelRating models.HotelRating) (models.HotelRating, error) {
	err := r.db.Create(&hotelRating).Error
	return hotelRating, err
}

func (r *hotelRatingsRepository) GetHotelRatingsByHotelID(id uint) (map[int]int, []models.HotelRating, error) {
	var hotelRatings []models.HotelRating
	err := r.db.Where("hotel_id = ?", id).Find(&hotelRatings).Error
	if err != nil {
		return nil, nil, err
	}

	ratingCounts := make(map[int]int)

	for _, rating := range hotelRatings {
		ratingCounts[rating.Rating]++
	}

	return ratingCounts, hotelRatings, nil
}
