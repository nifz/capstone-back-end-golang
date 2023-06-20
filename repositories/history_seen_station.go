package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type HistorySeenStationRepository interface {
	GetAllHistorySeenStation(page, limit int, userId uint) ([]models.HistorySeenStation, int, error)
	GetHistorySeenStationByID(station_origin_id, station_destination_id, userId uint) (models.HistorySeenStation, error)
	CreateHistorySeenStation(HistorySeenStation models.HistorySeenStation) (models.HistorySeenStation, error)
	UpdateHistorySeenStation(HistorySeenStation models.HistorySeenStation) (models.HistorySeenStation, error)
	DeleteHistorySeenStation(HistorySeenStation models.HistorySeenStation) (models.HistorySeenStation, error)
}

type historySeenStationRepository struct {
	db *gorm.DB
}

func NewHistorySeenStationRepository(db *gorm.DB) HistorySeenStationRepository {
	return &historySeenStationRepository{db}
}

func (r *historySeenStationRepository) GetAllHistorySeenStation(page, limit int, userId uint) ([]models.HistorySeenStation, int, error) {
	var (
		histories []models.HistorySeenStation
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

func (r *historySeenStationRepository) GetHistorySeenStationByID(station_origin_id, station_destination_id, userId uint) (models.HistorySeenStation, error) {
	var historySeenStation models.HistorySeenStation
	err := r.db.Where("station_origin_id = ? AND station_destination_id = ? AND user_id = ?", station_origin_id, station_destination_id, userId).First(&historySeenStation).Error
	return historySeenStation, err
}

func (r *historySeenStationRepository) CreateHistorySeenStation(historySeenStation models.HistorySeenStation) (models.HistorySeenStation, error) {
	err := r.db.Create(&historySeenStation).Error
	return historySeenStation, err
}

func (r *historySeenStationRepository) UpdateHistorySeenStation(historySeenStation models.HistorySeenStation) (models.HistorySeenStation, error) {
	err := r.db.Save(&historySeenStation).Error
	return historySeenStation, err
}

func (r *historySeenStationRepository) DeleteHistorySeenStation(historySeenStation models.HistorySeenStation) (models.HistorySeenStation, error) {
	err := r.db.Unscoped().Delete(&historySeenStation).Error
	return historySeenStation, err
}
