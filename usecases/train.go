package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"fmt"
)

type TrainUsecase interface {
	GetAllTrains(page, limit int) ([]dtos.TrainResponse, int, error)
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

// GetAllTrains godoc
// @Summary      Get all train
// @Description  Get all train
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success      200 {object} dtos.GetAllTrainStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/train [get]
// @Security BearerAuth
func (u *trainUsecase) GetAllTrains(page, limit int) ([]dtos.TrainResponse, int, error) {

	trains, count, err := u.trainRepo.GetAllTrains(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var trainResponses []dtos.TrainResponse
	for _, train := range trains {

		stationOrigin, err := u.trainRepo.GetStationByID2(train.StationOriginID)
		if err != nil {
			return trainResponses, 0, err
		}

		stationDestination, err := u.trainRepo.GetStationByID2(train.StationDestinationID)
		if err != nil {
			return trainResponses, 0, err
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
			CreatedAt:     train.CreatedAt,
			UpdatedAt:     train.UpdatedAt,
		}
		trainResponses = append(trainResponses, trainResponse)
	}
	fmt.Println(trainResponses)
	return trainResponses, count, nil
}

// GetTrainByID godoc
// @Summary      Get train by ID
// @Description  Get train by ID
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param id path integer true "ID train"
// @Success      200 {object} dtos.TrainStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/train/{id} [get]
// @Security BearerAuth
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
		CreatedAt:     train.CreatedAt,
		UpdatedAt:     train.UpdatedAt,
	}
	return trainResponse, nil
}

// CreateTrain godoc
// @Summary      Create a new train
// @Description  Create a new train
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param        request body dtos.TrainInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.TrainStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/train [post]
// @Security BearerAuth
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
		CreatedAt:     createdTrain.CreatedAt,
		UpdatedAt:     createdTrain.UpdatedAt,
	}
	return trainResponse, nil
}

// UpdateTrain godoc
// @Summary      Update train
// @Description  Update train
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param id path integer true "ID train"
// @Param        request body dtos.TrainInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.TrainStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/train [put]
// @Security BearerAuth
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
	trainResponse.CreatedAt = train.CreatedAt
	trainResponse.UpdatedAt = train.UpdatedAt

	return trainResponse, nil

}

// DeleteTrain godoc
// @Summary      Delete a train
// @Description  Delete a train
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param id path integer true "ID train"
// @Success      200 {object} dtos.StatusOKDeletedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/train/{id} [delete]
// @Security BearerAuth
func (u *trainUsecase) DeleteTrain(id uint) error {
	train, err := u.trainRepo.GetTrainByID(id)

	if err != nil {
		return nil
	}
	err = u.trainRepo.DeleteTrain(train)
	return err
}
