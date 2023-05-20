package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"fmt"
)

type TrainUsecase interface {
	GetAllTrains() ([]dtos.TrainResponse, error)
	GetTrainByID(id uint) (dtos.TrainResponse, error)
	CreateTrain(train *dtos.TrainInput) (dtos.TrainResponse, error)
	UpdateTrain(id uint, trainInput dtos.TrainInput) (dtos.TrainResponse, error)
	DeleteTrain(id uint) error
}

type trainUsecase struct {
	trainRepo repositories.TrainRepository
}

func NewTrainUsecase(TrainRepo repositories.TrainRepository) TrainUsecase {
	return &trainUsecase{TrainRepo}
}

func (u *trainUsecase) GetAllTrains() ([]dtos.TrainResponse, error) {

	trains, err := u.trainRepo.GetAllTrains()
	if err != nil {
		return nil, err
	}

	var trainResponses []dtos.TrainResponse
	for _, train := range trains {

		stationOrigin, err := u.trainRepo.GetStationByID2(train.StationOriginID)
		if err != nil {
			return trainResponses, err
		}

		stationDestination, err := u.trainRepo.GetStationByID2(train.StationDestinationID)
		if err != nil {
			return trainResponses, err
		}
		trainResponse := dtos.TrainResponse{
			TrainID:         train.ID,
			StationOriginID: train.StationOriginID,
			StationOrigin: dtos.StationInput{
				Origin:  stationOrigin.Origin,
				Name:    stationOrigin.Name,
				Initial: stationOrigin.Initial,
			},
			StationDestinationID: train.StationDestinationID,
			StationDestination: dtos.StationInput{
				Origin:  stationDestination.Origin,
				Name:    stationDestination.Name,
				Initial: stationDestination.Initial,
			},
			DepartureTime: train.DepartureTime,
			ArriveTime:    train.ArriveTime,
			Name:          train.Name,
			Route:         train.Route,
			Status:        train.Status,
		}
		trainResponses = append(trainResponses, trainResponse)
	}
	fmt.Println(trainResponses)
	return trainResponses, nil
}

func (u *trainUsecase) GetTrainByID(id uint) (dtos.TrainResponse, error) {
	var trainResponses dtos.TrainResponse
	train, err := u.trainRepo.GetTrainByID(id)
	if err != nil {
		return trainResponses, err
	}

	stationOrigin, err := u.trainRepo.GetStationByID2(train.StationOriginID)
	if err != nil {
		return trainResponses, err
	}

	stationDestination, err := u.trainRepo.GetStationByID2(train.StationDestinationID)
	if err != nil {
		return trainResponses, err
	}

	trainResponse := dtos.TrainResponse{
		TrainID:         train.ID,
		StationOriginID: train.StationOriginID,
		StationOrigin: dtos.StationInput{
			Origin:  stationOrigin.Origin,
			Name:    stationOrigin.Name,
			Initial: stationOrigin.Initial,
		},
		StationDestinationID: train.StationDestinationID,
		StationDestination: dtos.StationInput{
			Origin:  stationDestination.Origin,
			Name:    stationDestination.Name,
			Initial: stationDestination.Initial,
		},
		DepartureTime: train.DepartureTime,
		ArriveTime:    train.ArriveTime,
		Name:          train.Name,
		Route:         train.Route,
		Status:        train.Status,
	}
	return trainResponse, nil
}

func (u *trainUsecase) CreateTrain(train *dtos.TrainInput) (dtos.TrainResponse, error) {
	var trainResponsee dtos.TrainResponse

	createTrain := models.Train{
		StationOriginID:      train.StationOriginID,
		StationDestinationID: train.StationDestinationID,
		DepartureTime:        train.DepartureTime,
		ArriveTime:           train.ArriveTime,
		Name:                 train.Name,
		Route:                train.Route,
		Status:               train.Status,
	}

	createdTrain, err := u.trainRepo.CreateTrain(createTrain)
	if err != nil {
		return trainResponsee, err
	}

	stationOrigin, err := u.trainRepo.GetStationByID2(createdTrain.StationOriginID)
	if err != nil {
		return trainResponsee, err
	}

	stationDestination, err := u.trainRepo.GetStationByID2(createdTrain.StationDestinationID)
	if err != nil {
		return trainResponsee, err
	}

	trainResponse := dtos.TrainResponse{
		TrainID:         createdTrain.ID,
		StationOriginID: createdTrain.StationOriginID,
		StationOrigin: dtos.StationInput{
			Origin:  stationOrigin.Origin,
			Name:    stationOrigin.Name,
			Initial: stationOrigin.Initial,
		},
		StationDestinationID: createdTrain.StationDestinationID,
		StationDestination: dtos.StationInput{
			Origin:  stationDestination.Origin,
			Name:    stationDestination.Name,
			Initial: stationDestination.Initial,
		},
		DepartureTime: createdTrain.DepartureTime,
		ArriveTime:    createdTrain.ArriveTime,
		Name:          createdTrain.Name,
		Route:         createdTrain.Route,
		Status:        createdTrain.Status,
	}
	return trainResponse, nil
}

func (u *trainUsecase) UpdateTrain(id uint, trainInput dtos.TrainInput) (dtos.TrainResponse, error) {
	var train models.Train
	var trainResponse dtos.TrainResponse
	var trainResponsee dtos.TrainResponse

	train, err := u.trainRepo.GetTrainByID(id)

	if err != nil {
		return trainResponse, err
	}

	train.StationOriginID = trainInput.StationOriginID
	train.StationDestinationID = trainInput.StationDestinationID
	train.DepartureTime = trainInput.DepartureTime
	train.ArriveTime = trainInput.ArriveTime
	train.Name = trainInput.Name
	train.Route = trainInput.Route
	train.Status = trainInput.Status

	train, err = u.trainRepo.UpdateTrain(train)

	if err != nil {
		return trainResponse, err
	}

	stationOrigin, err := u.trainRepo.GetStationByID2(train.StationOriginID)
	if err != nil {
		return trainResponsee, err
	}

	stationDestination, err := u.trainRepo.GetStationByID2(train.StationDestinationID)
	if err != nil {
		return trainResponsee, err
	}

	trainResponse.TrainID = train.ID
	trainResponse.StationOriginID = train.StationOriginID
	trainResponse.StationOrigin.Origin = stationOrigin.Origin
	trainResponse.StationOrigin.Name = stationOrigin.Name
	trainResponse.StationOrigin.Initial = stationOrigin.Initial
	trainResponse.StationDestinationID = train.StationDestinationID
	trainResponse.StationDestination.Origin = stationDestination.Origin
	trainResponse.StationDestination.Name = stationDestination.Name
	trainResponse.StationDestination.Initial = stationDestination.Initial
	trainResponse.DepartureTime = train.DepartureTime
	trainResponse.ArriveTime = train.ArriveTime
	trainResponse.Name = train.Name
	trainResponse.Route = train.Route
	trainResponse.Status = train.Status
	trainResponse.UpdateAt = train.UpdatedAt

	return trainResponse, nil

}

func (u *trainUsecase) DeleteTrain(id uint) error {
	train, err := u.trainRepo.GetTrainByID(id)

	if err != nil {
		return nil
	}
	err = u.trainRepo.DeleteTrain(train)
	return err
}
