package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type TrainRepository interface {
	GetAllTrains() ([]models.Train, error)
	GetTrainByID(id uint) (models.Train, error)
	GetStationByID2(id uint) (models.Station, error)
	CreateTrain(train models.Train) (models.Train, error)
	UpdateTrain(train models.Train) (models.Train, error)
	DeleteTrain(train models.Train) error
}

type trainRepository struct {
	db *gorm.DB
}

func NewTrainRepository(db *gorm.DB) TrainRepository {
	return &trainRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *trainRepository) GetAllTrains() ([]models.Train, error) {
	var trains []models.Train
	err := r.db.Find(&trains).Error
	return trains, err
}

func (r *trainRepository) GetTrainByID(id uint) (models.Train, error) {
	var train models.Train
	err := r.db.Where("id = ?", id).First(&train).Error
	return train, err
}

func (r *trainRepository) GetStationByID2(id uint) (models.Station, error) {
	var station models.Station
	err := r.db.Where("id = ?", id).First(&station).Error
	return station, err
}

func (r *trainRepository) CreateTrain(train models.Train) (models.Train, error) {
	err := r.db.Create(&train).Error
	return train, err
}

func (r *trainRepository) UpdateTrain(train models.Train) (models.Train, error) {
	err := r.db.Save(&train).Error
	return train, err
}

func (r *trainRepository) DeleteTrain(train models.Train) error {
	err := r.db.Delete(&train).Error
	return err
}
