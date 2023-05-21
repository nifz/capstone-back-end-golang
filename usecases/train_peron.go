package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
)

type TrainPeronUsecase interface {
	GetAllTrainPerons(page, limit int) ([]dtos.TrainPeronResponse, int, error)
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

// GetAllTrainPerons godoc
// @Summary      Get all train peron
// @Description  Get all train peron
// @Tags         Train Peron
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success      200 {object} dtos.GetAllTrainPeronStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/train-peron [get]
// @Security BearerAuth
func (u *trainPeronUsecase) GetAllTrainPerons(page, limit int) ([]dtos.TrainPeronResponse, int, error) {

	trainPerons, count, err := u.trainPeronRepo.GetAllTrainPerons(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var trainPeronResponses []dtos.TrainPeronResponse
	for _, trainPeron := range trainPerons {

		train, err := u.trainPeronRepo.GetTrainByID2(trainPeron.TrainID)
		if err != nil {
			return trainPeronResponses, 0, err
		}

		stationOrigin, err := u.trainPeronRepo.GetStationByID2(train.StationOriginID)
		if err != nil {
			return trainPeronResponses, 0, err
		}

		stationDestination, err := u.trainPeronRepo.GetStationByID2(train.StationDestinationID)
		if err != nil {
			return trainPeronResponses, 0, err
		}

		trainSeat, err := u.trainPeronRepo.GetTrainSeatsByClass(trainPeron.Class)
		if err != nil {
			return trainPeronResponses, 0, err
		}

		var trainSeatResponses []dtos.TrainSeatResponse
		for _, trainSeat := range trainSeat {
			trainSeatRespon := dtos.TrainSeatResponse{
				Name: trainSeat.Name,
			}
			trainSeatResponses = append(trainSeatResponses, trainSeatRespon)
		}

		trainPeronResponse := dtos.TrainPeronResponse{
			TrainPeronID: trainPeron.ID,
			Train: dtos.TrainResponse{
				TrainID:         trainPeron.TrainID,
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
			},
			Class:     trainPeron.Class,
			Name:      trainPeron.Name,
			Seat:      trainSeatResponses,
			Price:     trainPeron.Price,
			Status:    trainPeron.Status,
			CreatedAt: trainPeron.CreatedAt,
			UpdatedAt: trainPeron.UpdatedAt,
		}
		trainPeronResponses = append(trainPeronResponses, trainPeronResponse)
	}
	return trainPeronResponses, count, nil
}

// GetTrainPeronByID godoc
// @Summary      Get train peron by ID
// @Description  Get train peron by ID
// @Tags         Train Peron
// @Accept       json
// @Produce      json
// @Param id path integer true "ID train peron"
// @Success      200 {object} dtos.TrainPeronStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/train-peron/{id} [get]
// @Security BearerAuth
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

	stationOrigin, err := u.trainPeronRepo.GetStationByID2(train.StationOriginID)
	if err != nil {
		return trainPeronResponses, err
	}

	stationDestination, err := u.trainPeronRepo.GetStationByID2(train.StationDestinationID)
	if err != nil {
		return trainPeronResponses, err
	}

	trainSeat, err := u.trainPeronRepo.GetTrainSeatsByClass(trainPeron.Class)
	if err != nil {
		return trainPeronResponses, err
	}

	var trainSeatResponses []dtos.TrainSeatResponse
	for _, trainSeat := range trainSeat {
		trainSeatRespon := dtos.TrainSeatResponse{
			Name: trainSeat.Name,
		}
		trainSeatResponses = append(trainSeatResponses, trainSeatRespon)
	}

	trainPeronResponse := dtos.TrainPeronResponse{
		TrainPeronID: trainPeron.ID,
		Train: dtos.TrainResponse{
			TrainID:         trainPeron.TrainID,
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
		},
		Class:     trainPeron.Class,
		Name:      trainPeron.Name,
		Seat:      trainSeatResponses,
		Price:     trainPeron.Price,
		Status:    trainPeron.Status,
		CreatedAt: trainPeron.CreatedAt,
		UpdatedAt: trainPeron.UpdatedAt,
	}
	return trainPeronResponse, nil
}

// CreateTrainPeron godoc
// @Summary      Create a new train peron
// @Description  Create a new train peron
// @Tags         Train Peron
// @Accept       json
// @Produce      json
// @Param        request body dtos.TrainPeronInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.TrainPeronStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/train-peron [post]
// @Security BearerAuth
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

	train, err := u.trainPeronRepo.GetTrainByID2(createdTrainPeron.TrainID)
	if err != nil {
		return trainPeronResponsee, err
	}

	stationOrigin, err := u.trainPeronRepo.GetStationByID2(train.StationOriginID)
	if err != nil {
		return trainPeronResponsee, err
	}

	stationDestination, err := u.trainPeronRepo.GetStationByID2(train.StationDestinationID)
	if err != nil {
		return trainPeronResponsee, err
	}

	trainSeat, err := u.trainPeronRepo.GetTrainSeatsByClass(trainPeron.Class)
	if err != nil {
		return trainPeronResponsee, err
	}

	var trainSeatResponses []dtos.TrainSeatResponse
	for _, trainSeat := range trainSeat {
		trainSeatRespon := dtos.TrainSeatResponse{
			Name: trainSeat.Name,
		}
		trainSeatResponses = append(trainSeatResponses, trainSeatRespon)
	}

	trainPeronResponse := dtos.TrainPeronResponse{
		TrainPeronID: createdTrainPeron.ID,
		Train: dtos.TrainResponse{
			TrainID:         createdTrainPeron.TrainID,
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
		},
		Class:     createTrainPeron.Class,
		Name:      createTrainPeron.Name,
		Seat:      trainSeatResponses,
		Price:     createTrainPeron.Price,
		Status:    createdTrainPeron.Status,
		CreatedAt: createdTrainPeron.CreatedAt,
		UpdatedAt: createdTrainPeron.UpdatedAt,
	}
	return trainPeronResponse, nil
}

// UpdateTrainPeron godoc
// @Summary      Update train peron
// @Description  Update train peron
// @Tags         Train Peron
// @Accept       json
// @Produce      json
// @Param id path integer true "ID train peron"
// @Param        request body dtos.TrainPeronInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.TrainPeronStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/train-peron [put]
// @Security BearerAuth
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

	stationOrigin, err := u.trainPeronRepo.GetStationByID2(train.StationOriginID)
	if err != nil {
		return trainPeronResponsee, err
	}

	stationDestination, err := u.trainPeronRepo.GetStationByID2(train.StationDestinationID)
	if err != nil {
		return trainPeronResponsee, err
	}

	trainSeat, err := u.trainPeronRepo.GetTrainSeatsByClass(trainPeron.Class)
	if err != nil {
		return trainPeronResponsee, err
	}

	var trainSeatResponses []dtos.TrainSeatResponse
	for _, trainSeat := range trainSeat {
		trainSeatRespon := dtos.TrainSeatResponse{
			Name: trainSeat.Name,
		}
		trainSeatResponses = append(trainSeatResponses, trainSeatRespon)
	}

	trainPeronResponse = dtos.TrainPeronResponse{
		TrainPeronID: trainPeron.ID,
		Train: dtos.TrainResponse{
			TrainID:         trainPeron.TrainID,
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
		},
		Class:     trainPeron.Class,
		Name:      trainPeron.Name,
		Seat:      trainSeatResponses,
		Price:     trainPeron.Price,
		Status:    trainPeron.Status,
		CreatedAt: trainPeron.CreatedAt,
		UpdatedAt: trainPeron.UpdatedAt,
	}

	return trainPeronResponse, nil

}

// DeleteTrainPeron godoc
// @Summary      Delete a train peron
// @Description  Delete a train peron
// @Tags         Train Peron
// @Accept       json
// @Produce      json
// @Param id path integer true "ID train peron"
// @Success      200 {object} dtos.StatusOKDeletedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/train-peron/{id} [delete]
// @Security BearerAuth
func (u *trainPeronUsecase) DeleteTrainPeron(id uint) error {
	trainPeron, err := u.trainPeronRepo.GetTrainPeronByID(id)

	if err != nil {
		return nil
	}
	err = u.trainPeronRepo.DeleteTrainPeron(trainPeron)
	return err
}
