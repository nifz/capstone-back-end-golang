package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type HistorySearchRepository interface {
	HistorySearchGetById(id uint) (models.HistorySearch, error)
	HistorySearchGetByUserId(userId uint) ([]models.HistorySearch, error)
	HistorySearchCreate(historySearch models.HistorySearch) (models.HistorySearch, error)
	HistorySearchUpdate(historySearch models.HistorySearch) (models.HistorySearch, error)
}

type historySearchRepository struct {
	db *gorm.DB
}

func NewHistorySearchRepository(db *gorm.DB) HistorySearchRepository {
	return &historySearchRepository{db}
}

func (r *historySearchRepository) HistorySearchGetById(id uint) (models.HistorySearch, error) {
	var historySearch models.HistorySearch
	err := r.db.Where("id = ?", id).First(&historySearch).Error
	return historySearch, err
}

func (r *historySearchRepository) HistorySearchGetByUserId(userId uint) ([]models.HistorySearch, error) {
	var historySearch []models.HistorySearch
	err := r.db.Where("user_id = ?", userId).Find(&historySearch).Error
	return historySearch, err
}

func (r *historySearchRepository) HistorySearchCreate(historySearch models.HistorySearch) (models.HistorySearch, error) {
	err := r.db.Create(&historySearch).Error
	return historySearch, err
}

func (r *historySearchRepository) HistorySearchUpdate(historySearch models.HistorySearch) (models.HistorySearch, error) {
	err := r.db.Save(&historySearch).Error
	return historySearch, err
}
