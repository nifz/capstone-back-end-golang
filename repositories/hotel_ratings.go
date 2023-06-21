package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type HotelRatingsRepository interface {
	// user
	CreateHotelRating(hotelRating models.HotelRating) (models.HotelRating, error)
	GetHotelRatingsByIdOrders(id uint) (models.HotelRating, error)
	CheckExistHotelRating(order_id, user_id uint) (bool, error)
	GetAllHotelRatingsByIdHotels(page, limit int, hotel_id uint) ([]models.HotelRating, int, error)
	GetAllHotelRatingsByIdHotels2(hotel_id uint) ([]models.HotelRating, error)
	// admin
	GetHotelRatingsByHotelID(page, limit int, id uint, filter string) (map[int]int, []models.HotelRating, int, error)
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
func (r *hotelRatingsRepository) GetHotelRatingsByHotelID(page, limit int, id uint, filter string) (map[int]int, []models.HotelRating, int, error) {
	var (
		hotelRatings []models.HotelRating
		count        int64
	)
	offset := (page - 1) * limit
	err := r.db.Where("hotel_id = ?", id).First(&hotelRatings).Count(&count).Error
	if err != nil {
		return nil, hotelRatings, int(count), err
	}
	ratingCounts := make(map[int]int)
	if filter == "latest" {
		err = r.db.Where("hotel_id = ?", id).Order("created_at DESC").Limit(limit).Offset(offset).Find(&hotelRatings).Error
		// for _, rating := range hotelRatings {
		// 	ratingCounts[rating.Rating]++
		// }
	} else if filter == "oldest" {
		err = r.db.Where("hotel_id = ?", id).Order("created_at ASC").Limit(limit).Offset(offset).Find(&hotelRatings).Error
		// for _, rating := range hotelRatings {
		// 	ratingCounts[rating.Rating]++
		// }
	} else {
		err = r.db.Where("hotel_id = ?", id).Limit(limit).Offset(offset).Find(&hotelRatings).Error
		// if err != nil {
		// 	return nil, hotelRatings, int(count), err
		// }
	}

	for _, rating := range hotelRatings {
		ratingCounts[rating.Rating]++
	}

	return ratingCounts, hotelRatings, int(count), err
}
func (r *hotelRatingsRepository) GetHotelRatingsByIdOrders(id uint) (models.HotelRating, error) {
	var hotelRating models.HotelRating
	err := r.db.Where("hotel_order_id = ?", id).First(&hotelRating).Error
	return hotelRating, err
}
func (r *hotelRatingsRepository) CheckExistHotelRating(order_id, user_id uint) (bool, error) {
	hotelRating := models.HotelRating{}
	err := r.db.Where("user_id = ? AND hotel_order_id = ?", user_id, order_id).First(&hotelRating).Error
	if err != nil {
		return false, err
	}
	return true, err
}
func (r *hotelRatingsRepository) GetAllHotelRatingsByIdHotels(page, limit int, hotel_id uint) ([]models.HotelRating, int, error) {
	var (
		hotelRatings []models.HotelRating
		count        int64
	)

	err := r.db.Where("hotel_id = ?", hotel_id).First(&hotelRatings).Count(&count).Error
	if err != nil {
		return hotelRatings, int(count), err
	}

	// err = r.db.Where("hotel_id = ?", hotel_id).Find(&hotelRatings).Count(&count).Error
	// if err != nil {
	// 	return hotelRatings, int(count), err
	// }

	// if len(hotelRatings) == 0 {
	// 	return hotelRatings, int(count), err
	// }

	offset := (page - 1) * limit

	err = r.db.Where("hotel_id = ?", hotel_id).Order("id DESC").Limit(limit).Offset(offset).Find(&hotelRatings).Error

	return hotelRatings, int(count), err
}
func (r *hotelRatingsRepository) GetAllHotelRatingsByIdHotels2(hotel_id uint) ([]models.HotelRating, error) {
	var (
		hotelRatings []models.HotelRating
	)

	err := r.db.Where("hotel_id = ?", hotel_id).Order("id DESC").Limit(10).Find(&hotelRatings).Error

	return hotelRatings, err
}
