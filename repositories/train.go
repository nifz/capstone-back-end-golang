package repositories

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"errors"

	"gorm.io/gorm"
)

type TrainRepository interface {
	GetAllTrain(page, limit int) ([]models.Train, int, error)
	GetAllTrains(page, limit int) ([]models.TrainCarriage, int, error)
	GetTrainByID(id uint) (models.Train, error)
	TrainStationByTrainID(id uint) (models.TrainStation, error)
	GetTrainStationByTrainID(id uint) ([]models.TrainStation, error)
	SearchTrainAvailable(trainId, originId, destinationId uint) ([]models.TrainStation, error)
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

func (r *trainRepository) GetAllTrain(page, limit int) ([]models.Train, int, error) {
	var (
		trains []models.Train
		count  int64
	)
	err := r.db.Find(&trains).Count(&count).Error

	// Gunakan slice trainCarriages yang berisi hasil query

	if err != nil {
		return trains, int(count), err
	}

	offset := (page - 1) * limit

	err = r.db.Limit(limit).Offset(offset).Find(&trains).Error

	return trains, int(count), err
}

func (r *trainRepository) GetAllTrains(page, limit int) ([]models.TrainCarriage, int, error) {
	var (
		trains []models.TrainCarriage
		count  int64
	)
	// err := r.db.Joins("JOIN trains ON train_carriages.train_id = trains.id").Preload("Train").Find(&trains).Count(&count).Error
	// var trainCarriages []models.TrainCarriage

	err := r.db.Raw(`
	SELECT tc.*
	FROM train_carriages tc
	JOIN trains t ON tc.train_id = t.id
	WHERE tc.id IN (
		SELECT MIN(id)
		FROM train_carriages
		GROUP BY class, train_id
	);
`).Scan(&trains).Error

	// Gunakan slice trainCarriages yang berisi hasil query

	if err != nil {
		return trains, int(count), err
	}

	// offset := (page - 1) * limit

	// err = r.db.Limit(limit).Offset(offset).Find(&trains).Error

	return trains, int(count), err
}

func (r *trainRepository) GetTrainByID(id uint) (models.Train, error) {
	var train models.Train
	err := r.db.Where("id = ?", id).First(&train).Error
	return train, err
}

func (r *trainRepository) GetTrainStationByTrainID(id uint) ([]models.TrainStation, error) {
	var train []models.TrainStation
	err := r.db.Where("train_id = ?", id).Find(&train).Error
	return train, err
}

func (r *trainRepository) TrainStationByTrainID(id uint) (models.TrainStation, error) {
	var train models.TrainStation
	err := r.db.Where("train_id = ?", id).First(&train).Error
	return train, err
}

func (r *trainRepository) SearchTrainAvailable(trainId, originId, destinationId uint) ([]models.TrainStation, error) {
	var (
		train []models.TrainStation
		count int64
	)
	err := r.db.Where("train_id = ?", trainId).Where("station_id = ? OR station_id = ?", originId, destinationId).Find(&train).Count(&count).Error
	// Cek apakah ada data dengan 'arrive time' yang descending
	for i := 0; i < len(train)-1; i++ {
		if train[i].ArriveTime > train[i+1].ArriveTime {
			err = errors.New("Train not available")
			train = nil // Reset data jika ada 'arrive time' descending
			break
		}
	}

	if err != nil {
		return train, err
	}

	return train, err
}

func (r *trainRepository) GetStationByID(id uint) (dtos.StationInput, error) {
	var station dtos.StationInput
	err := r.db.Where("id = ?", id).Find(&station).Error
	return station, err
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
	if err != nil {
		return err
	}

	trainStation := models.TrainStation{
		TrainID: train.ID,
	}
	err = r.db.Where("train_id = ?", trainStation.TrainID).Delete(&trainStation).Error
	if err != nil {
		return err
	}

	trainCarriage := models.TrainCarriage{
		TrainID: train.ID,
	}
	err = r.db.Where("train_id = ?", trainCarriage.TrainID).Delete(&trainCarriage).Error
	if err != nil {
		return err
	}

	return nil
}
