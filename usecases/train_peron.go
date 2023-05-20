package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
)

type TrainPeronUsecase interface {
	GetAllTrainPerons() ([]dtos.TrainPeronResponse, error)
	GetTrainPeronByID(id uint) (dtos.TrainPeronResponse, error)
	CreateTrainPeron(trainPeron *dtos.TrainPeronInput) (dtos.TrainPeronResponse, error)
	UpdateTrainPeron(id uint, trainPeronInput dtos.TrainPeronInput) (dtos.TrainPeronResponse, error)
	DeleteTrainPeron(id uint) error
}

type trainPeronUsecase struct {
	trainPeronRepo repositories.TrainPeronRepository
}

func NewTrainPeronUsecase(TrainPeronRepo repositories.TrainPeronRepository) TrainPeronUsecase {
	return &trainPeronUsecase{TrainPeronRepo}
}

func (u *trainPeronUsecase) GetAllTrainPerons() ([]dtos.TrainPeronResponse, error) {

	trainPerons, err := u.trainPeronRepo.GetAllTrainPerons()
	if err != nil {
		return nil, err
	}

	var trainPeronResponses []dtos.TrainPeronResponse
	for _, trainPeron := range trainPerons {

		train, err := u.trainPeronRepo.GetTrainByID2(trainPeron.TrainID)
		if err != nil {
			return trainPeronResponses, err
		}

		trainPeronResponse := dtos.TrainPeronResponse{
			TrainPeronID: trainPeron.ID,
			TrainID:      trainPeron.TrainID,
			Train: dtos.TrainInput{
				StationOriginID:      train.StationOriginID,
				StationDestinationID: train.StationDestinationID,
				DepartureTime:        train.DepartureTime,
				ArriveTime:           train.ArriveTime,
				Name:                 train.Name,
				Route:                train.Route,
				Status:               train.Status,
			},
			Class:  trainPeron.Class,
			Name:   trainPeron.Name,
			Price:  trainPeron.Price,
			Status: trainPeron.Status,
		}
		trainPeronResponses = append(trainPeronResponses, trainPeronResponse)
	}
	return trainPeronResponses, nil
}

func (u *trainPeronUsecase) GetTrainPeronByID(id uint) (dtos.TrainPeronResponse, error) {
	var trainPeronResponses dtos.TrainPeronResponse
	trainPeron, err := u.trainPeronRepo.GetTrainPeronByID(id)
	if err != nil {
		return trainPeronResponses, err
	}

	train, err := u.trainPeronRepo.GetTrainByID2(trainPeron.TrainID)
	if err != nil {
		return trainPeronResponses, err
	}

	trainPeronResponse := dtos.TrainPeronResponse{
		TrainPeronID: trainPeron.ID,
		TrainID:      trainPeron.ID,
		Train: dtos.TrainInput{
			StationOriginID:      train.StationOriginID,
			StationDestinationID: train.StationDestinationID,
			DepartureTime:        train.DepartureTime,
			ArriveTime:           train.ArriveTime,
			Name:                 train.Name,
			Route:                train.Route,
			Status:               train.Status,
		},
		Class:  trainPeron.Class,
		Name:   trainPeron.Name,
		Price:  trainPeron.Price,
		Status: trainPeron.Status,
	}
	return trainPeronResponse, nil
}

func (u *trainPeronUsecase) CreateTrainPeron(trainPeron *dtos.TrainPeronInput) (dtos.TrainPeronResponse, error) {
	var trainPeronResponsee dtos.TrainPeronResponse

	createTrainPeron := models.TrainPeron{
		TrainID: trainPeron.TrainID,
		Class:   trainPeron.Class,
		Name:    trainPeron.Name,
		Price:   trainPeron.Price,
		Status:  trainPeron.Status,
	}

	createdTrainPeron, err := u.trainPeronRepo.CreateTrainPeron(createTrainPeron)
	if err != nil {
		return trainPeronResponsee, err
	}

	train, err := u.trainPeronRepo.GetTrainByID2(trainPeron.TrainID)
	if err != nil {
		return trainPeronResponsee, err
	}

	trainPeronResponse := dtos.TrainPeronResponse{
		TrainPeronID: createdTrainPeron.ID,
		TrainID:      createdTrainPeron.TrainID,
		Train: dtos.TrainInput{
			StationOriginID:      train.StationOriginID,
			StationDestinationID: train.StationDestinationID,
			DepartureTime:        train.DepartureTime,
			ArriveTime:           train.ArriveTime,
			Name:                 train.Name,
			Route:                train.Route,
			Status:               train.Status,
		},
		Class:  createTrainPeron.Class,
		Name:   createTrainPeron.Name,
		Price:  createTrainPeron.Price,
		Status: createdTrainPeron.Status,
	}
	return trainPeronResponse, nil
}

func (u *trainPeronUsecase) UpdateTrainPeron(id uint, trainPeronInput dtos.TrainPeronInput) (dtos.TrainPeronResponse, error) {
	var trainPeron models.TrainPeron
	var trainPeronResponse dtos.TrainPeronResponse
	var trainPeronResponsee dtos.TrainPeronResponse

	trainPeron, err := u.trainPeronRepo.GetTrainPeronByID(id)

	if err != nil {
		return trainPeronResponse, err
	}

	trainPeron.TrainID = trainPeronInput.TrainID
	trainPeron.Class = trainPeronInput.Class
	trainPeron.Name = trainPeronInput.Name
	trainPeron.Price = trainPeronInput.Price
	trainPeron.Status = trainPeronInput.Status

	trainPeron, err = u.trainPeronRepo.UpdateTrainPeron(trainPeron)

	if err != nil {
		return trainPeronResponsee, err
	}

	train, err := u.trainPeronRepo.GetTrainByID2(trainPeron.TrainID)
	if err != nil {
		return trainPeronResponsee, err
	}

	trainPeronResponse.TrainPeronID = trainPeron.ID
	trainPeronResponse.TrainID = trainPeron.TrainID
	trainPeronResponse.Train.StationOriginID = train.StationOriginID
	trainPeronResponse.Train.StationDestinationID = train.StationDestinationID
	trainPeronResponse.Train.DepartureTime = train.DepartureTime
	trainPeronResponse.Train.ArriveTime = train.ArriveTime
	trainPeronResponse.Train.Name = train.Name
	trainPeronResponse.Train.Route = train.Route
	trainPeronResponse.Train.Status = train.Status
	trainPeronResponse.Class = trainPeron.Class
	trainPeronResponse.Name = trainPeron.Name
	trainPeronResponse.Price = trainPeron.Price
	trainPeronResponse.Status = trainPeron.Status
	trainPeronResponse.UpdateAt = trainPeron.UpdatedAt

	return trainPeronResponse, nil

}

func (u *trainPeronUsecase) DeleteTrainPeron(id uint) error {
	trainPeron, err := u.trainPeronRepo.GetTrainPeronByID(id)

	if err != nil {
		return nil
	}
	err = u.trainPeronRepo.DeleteTrainPeron(trainPeron)
	return err
}
