package repositories

import (
	"back-end-golang/models"
	"strconv"

	"gorm.io/gorm"
)

type TrainRepository interface {
	GetAllStations() ([]models.Station, error)
	GetStationByID(id string) (models.Station, error)
	CreateStation(station *models.Station) error
	UpdateStation(station *models.Station) error
	DeleteStation(id string) error
}

type trainRepository struct {
	db *gorm.DB
}

func NewTrainRepository(db *gorm.DB) trainRepository {
	return &trainRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *trainRepository) GetAllStations() ([]models.Station, error) {
	var stations []models.Station
	err := r.db.Find(&stations).Error
	return stations, err
}

func (r *trainRepository) GetStationByID(id string) (models.Station, error) {
	var station models.Station
	err := r.db.First(&station).Error
	return station, err
}

func (r *trainRepository) CreateStation(station *models.Station) error {
	var _ = r.db.Create(station).Error
	return r.db.First(&station).Error
}

func (r *trainRepository) UpdateStation(station *models.Station) error {
	return r.db.Save(station).Error
}

func (r *trainRepository) DeleteStation(id string) error {
	Id, _ := strconv.Atoi(id)
	return r.db.Delete(&models.Station{Model: gorm.Model{ID: uint(Id)}}).Error
}
