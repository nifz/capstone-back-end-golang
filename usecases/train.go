package usecases

import (
	"back-end-golang/models"
	"back-end-golang/repositories"
)

type StationUsecase interface {
	GetAllStations() ([]models.Station, error)
	GetStationByID(id string) (models.Station, error)
	CreateStation(station *models.Station) error
	UpdateStation(station *models.Station) error
	DeleteStation(id string) error
}

type stationUsecase struct {
	stationRepo repositories.TrainRepository
}

func NewStationUsecase(StationRepo repositories.TrainRepository) stationUsecase {
	return &stationUsecase{StationRepo}
}

func (u *stationUsecase) GetAllStations() ([]models.Station, error) {
	return u.stationRepo.GetAllStations()
}

func (u *stationUsecase) GetStationByID(id string) (models.Station, error) {
	return u.stationRepo.GetStationByID(id)
}

func (u *stationUsecase) CreateStation(item *models.Station) error {
	return u.stationRepo.CreateStation(item)
}

func (u *stationUsecase) UpdateStation(item *models.Station) error {
	return u.stationRepo.UpdateStation(item)
}

func (u *stationUsecase) DeleteStation(id string) error {
	return u.stationRepo.DeleteStation(id)
}
