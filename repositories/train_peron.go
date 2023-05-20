package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type TrainPeronRepository interface {
	GetAllTrainPerons() ([]models.TrainPeron, error)
	GetTrainPeronByID(id uint) (models.TrainPeron, error)
	GetTrainByID2(id uint) (models.Train, error)
	CreateTrainPeron(trainPeron models.TrainPeron) (models.TrainPeron, error)
	UpdateTrainPeron(trainPeron models.TrainPeron) (models.TrainPeron, error)
	DeleteTrainPeron(trainPeron models.TrainPeron) error
}

type trainPeronRepository struct {
	db *gorm.DB
}

func NewTrainPeronRepository(db *gorm.DB) TrainPeronRepository {
	return &trainPeronRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *trainPeronRepository) GetAllTrainPerons() ([]models.TrainPeron, error) {
	var trainPerons []models.TrainPeron
	err := r.db.Find(&trainPerons).Error
	return trainPerons, err
}

func (r *trainPeronRepository) GetTrainPeronByID(id uint) (models.TrainPeron, error) {
	var trainPeron models.TrainPeron
	err := r.db.Where("id = ?", id).First(&trainPeron).Error
	return trainPeron, err
}

func (r *trainPeronRepository) GetTrainByID2(id uint) (models.Train, error) {
	var train models.Train
	err := r.db.Where("id = ?", id).First(&train).Error
	return train, err
}

func (r *trainPeronRepository) CreateTrainPeron(trainPeron models.TrainPeron) (models.TrainPeron, error) {
	err := r.db.Create(&trainPeron).Error
	return trainPeron, err
}

func (r *trainPeronRepository) UpdateTrainPeron(trainPeron models.TrainPeron) (models.TrainPeron, error) {
	err := r.db.Save(trainPeron).Error
	return trainPeron, err
}

func (r *trainPeronRepository) DeleteTrainPeron(trainPeron models.TrainPeron) error {
	err := r.db.Delete(&trainPeron).Error
	return err
}
