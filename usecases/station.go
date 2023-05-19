package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"fmt"
)

type StationUsecase interface {
	GetAllStations() ([]dtos.StationResponse, error)
	GetStationByID(id uint) (dtos.StationResponse, error)
	CreateStation(station *dtos.StationInput) (dtos.StationResponse, error)
	UpdateStation(id uint, station dtos.StationInput) (dtos.StationResponse, error)
	DeleteStation(id uint) error
}

type stationUsecase struct {
	stationRepo repositories.StationRepository
}

func NewStationUsecase(StationRepo repositories.StationRepository) StationUsecase {
	return &stationUsecase{StationRepo}
}

func (u *stationUsecase) GetAllStations() ([]dtos.StationResponse, error) {
	stations, err := u.stationRepo.GetAllStations()
	if err != nil {
		return nil, err
	}

	var stationResponses []dtos.StationResponse
	for _, station := range stations {
		stationResponse := dtos.StationResponse{
			StationID: station.ID,
			Origin:    station.Origin,
			Name:      station.Name,
			Initial:   station.Initial,
			UpdateAt:  station.UpdatedAt,
		}
		stationResponses = append(stationResponses, stationResponse)
	}

	return stationResponses, nil
}

func (u *stationUsecase) GetStationByID(id uint) (dtos.StationResponse, error) {
	var stationResponses dtos.StationResponse
	station, err := u.stationRepo.GetStationByID(id)
	if err != nil {
		return stationResponses, err
	}
	stationResponse := dtos.StationResponse{
		StationID: station.ID,
		Origin:    station.Origin,
		Name:      station.Name,
		Initial:   station.Initial,
		UpdateAt:  station.UpdatedAt,
	}
	return stationResponse, nil
}

func (u *stationUsecase) CreateStation(station *dtos.StationInput) (dtos.StationResponse, error) {
	var stationResponses dtos.StationResponse
	createStation := models.Station{
		Origin:  station.Origin,
		Name:    station.Name,
		Initial: station.Initial,
	}

	createdStation, err := u.stationRepo.CreateStation(createStation)
	if err != nil {
		return stationResponses, err
	}

	stationResponse := dtos.StationResponse{
		StationID: createdStation.ID,
		Origin:    createdStation.Origin,
		Name:      createdStation.Name,
		Initial:   createdStation.Initial,
		UpdateAt:  createStation.UpdatedAt,
	}
	return stationResponse, nil
}

func (u *stationUsecase) UpdateStation(id uint, stationInput dtos.StationInput) (dtos.StationResponse, error) {
	var station models.Station
	var stationResponse dtos.StationResponse

	station, err := u.stationRepo.GetStationByID(id)
	fmt.Println(station)
	if err != nil {
		return stationResponse, err
	}

	station.Origin = stationInput.Origin
	station.Name = stationInput.Name
	station.Initial = stationInput.Initial

	station, err = u.stationRepo.UpdateStation(station)

	if err != nil {
		return stationResponse, err
	}

	stationResponse.StationID = station.ID
	stationResponse.Origin = station.Origin
	stationResponse.Name = station.Name
	stationResponse.Initial = station.Initial
	stationResponse.UpdateAt = station.UpdatedAt

	return stationResponse, nil

}

func (u *stationUsecase) DeleteStation(id uint) error {
	station, err := u.stationRepo.GetStationByID(id)

	if err != nil {
		return nil
	}
	err = u.stationRepo.DeleteStation(station)
	return err
}
