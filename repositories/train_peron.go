package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type TrainPeronRepository interface {
	GetAllTrainPerons(page, limit int) ([]models.TrainPeron, int, error)
	GetTrainPeronByID(id uint) (models.TrainPeron, error)
	GetTrainByID2(id uint) (models.Train, error)
	GetStationByID2(id uint) (models.Station, error)
	GetTrainSeatsByClass(class string) ([]models.TrainSeat, error)
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

func (r *trainPeronRepository) GetAllTrainPerons(page, limit int) ([]models.TrainPeron, int, error) {
	var (
		trainPerons []models.TrainPeron
		count       int64
	)
	err := r.db.Find(&trainPerons).Count(&count).Error
	if err != nil {
		return trainPerons, int(count), err
	}

	offset := (page - 1) * limit

	err = r.db.Limit(limit).Offset(offset).Find(&trainPerons).Error

	return trainPerons, int(count), err
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

func (r *trainPeronRepository) GetStationByID2(id uint) (models.Station, error) {
	var station models.Station
	err := r.db.Where("id = ?", id).First(&station).Error
	return station, err
}

func (r *trainPeronRepository) GetTrainSeatsByClass(class string) ([]models.TrainSeat, error) {
	var trainSeats []models.TrainSeat
	err := r.db.Where("class = ?", class).Find(&trainSeats).Error
	return trainSeats, err
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
