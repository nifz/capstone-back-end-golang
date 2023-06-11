package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type TrainSeatRepository interface {
	GetTrainSeatByID(id uint) (models.TrainSeat, error)
	GetTrainSeatByClass(class string) ([]models.TrainSeat, error)
}

type trainSeatRepository struct {
	db *gorm.DB
}

func NewTrainSeatRepository(db *gorm.DB) TrainSeatRepository {
	return &trainSeatRepository{db}
}

func (r *trainSeatRepository) GetTrainSeatByID(id uint) (models.TrainSeat, error) {
	var trainSeat models.TrainSeat
	err := r.db.Where("id = ?", id).First(&trainSeat).Error
	return trainSeat, err
}

func (r *trainSeatRepository) GetTrainSeatByClass(class string) ([]models.TrainSeat, error) {
	var trainSeat []models.TrainSeat
	err := r.db.Where("class = ?", class).Find(&trainSeat).Error
	return trainSeat, err
}
