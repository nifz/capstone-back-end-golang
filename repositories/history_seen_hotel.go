package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type HistorySeenHotelRepository interface {
	GetAllHistorySeenHotel(page, limit int, userId uint) ([]models.HistorySeenHotel, int, error)
	GetHistorySeenHotelByID(hotelId, userId uint) (models.HistorySeenHotel, error)
	CreateHistorySeenHotel(HistorySeenHotel models.HistorySeenHotel) (models.HistorySeenHotel, error)
	UpdateHistorySeenHotel(HistorySeenHotel models.HistorySeenHotel) (models.HistorySeenHotel, error)
	DeleteHistorySeenHotel(HistorySeenHotel models.HistorySeenHotel) (models.HistorySeenHotel, error)
}

type historySeenHotelRepository struct {
	db *gorm.DB
}

func NewHistorySeenHotelRepository(db *gorm.DB) HistorySeenHotelRepository {
	return &historySeenHotelRepository{db}
}

func (r *historySeenHotelRepository) GetAllHistorySeenHotel(page, limit int, userId uint) ([]models.HistorySeenHotel, int, error) {
	var (
		histories []models.HistorySeenHotel
		count     int64
	)
	err := r.db.Order("id DESC").Find(&histories).Count(&count).Error
	if err != nil {
		return histories, int(count), err
	}

	offset := (page - 1) * limit

	err = r.db.Order("id DESC").Limit(limit).Offset(offset).Find(&histories).Error

	return histories, int(count), err
}

func (r *historySeenHotelRepository) GetHistorySeenHotelByID(hotelId, userId uint) (models.HistorySeenHotel, error) {
	var historySeenHotel models.HistorySeenHotel
	err := r.db.Where("hotel_id = ? AND user_id = ?", hotelId, userId).First(&historySeenHotel).Error
	return historySeenHotel, err
}

func (r *historySeenHotelRepository) CreateHistorySeenHotel(historySeenHotel models.HistorySeenHotel) (models.HistorySeenHotel, error) {
	err := r.db.Create(&historySeenHotel).Error
	return historySeenHotel, err
}

func (r *historySeenHotelRepository) UpdateHistorySeenHotel(historySeenHotel models.HistorySeenHotel) (models.HistorySeenHotel, error) {
	err := r.db.Save(&historySeenHotel).Error
	return historySeenHotel, err
}

func (r *historySeenHotelRepository) DeleteHistorySeenHotel(historySeenHotel models.HistorySeenHotel) (models.HistorySeenHotel, error) {
	err := r.db.Unscoped().Delete(&historySeenHotel).Error
	return historySeenHotel, err
}
