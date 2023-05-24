package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"sort"
	"strings"
)

type TrainUsecase interface {
	// admin
	GetAllTrains(page, limit int) ([]dtos.TrainResponse, int, error)
	GetTrainByID(id uint) (dtos.TrainResponse, error)
	CreateTrain(train *dtos.TrainInput) (dtos.TrainResponse, error)
	UpdateTrain(id uint, trainInput dtos.TrainInput) (dtos.TrainResponse, error)
	DeleteTrain(id uint) error

	// user
	SearchTrainAvailable(page, limit, stationOrigin, stationDestination, sortByTrainId int, sortClassName, sortByPrice, sortByArriveTime string) ([]dtos.TrainResponse, int, error)
}

type trainUsecase struct {
	trainRepo repositories.TrainRepository
}

func NewTrainUsecase(TrainRepo repositories.TrainRepository) TrainUsecase {
	return &trainUsecase{TrainRepo}
}

// =============================== ADMIN END ================================== \\

// GetAllTrains godoc
// @Summary      Get all train
// @Description  Get all train
// @Tags         Admin - Train
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

	// Track train IDs and classes
	trainMap := make(map[uint]map[string]bool)

	var trainResponses []dtos.TrainResponse
	for _, train := range trains {
		getTrain, err := u.trainRepo.GetTrainByID(train.TrainID)
		if err != nil {
			return trainResponses, count, err
		}

		getTrainStation, err := u.trainRepo.GetTrainStationByTrainID(getTrain.ID)
		if err != nil {
			return trainResponses, count, err
		}

		// Check if train ID exists in the map
		if classMap, ok := trainMap[getTrain.ID]; ok {
			// Check if class already exists for the train ID
			if classMap[train.Class] {
				continue
			}
		} else {
			// Create a new class map for the train ID
			trainMap[getTrain.ID] = make(map[string]bool)
		}

		// Add class to the train ID map
		trainMap[getTrain.ID][train.Class] = true

		var trainStationResponses []dtos.TrainStationResponse

		for _, train := range getTrainStation {
			getStation, err := u.trainRepo.GetStationByID2(train.StationID)
			if err != nil {
				return trainResponses, count, err
			}

			trainStationResponse := dtos.TrainStationResponse{
				StationID: train.StationID,
				Station: dtos.StationInput{
					Origin:  getStation.Origin,
					Name:    getStation.Name,
					Initial: getStation.Initial,
				},
				ArriveTime: train.ArriveTime,
			}

			trainStationResponses = append(trainStationResponses, trainStationResponse)
		}

		trainResponse := dtos.TrainResponse{
			TrainID:   getTrain.ID,
			CodeTrain: getTrain.CodeTrain,
			Name:      getTrain.Name,
			Class:     train.Class,
			Price:     train.Price,
			Route:     trainStationResponses,
			Status:    getTrain.Status,
			CreatedAt: getTrain.CreatedAt,
			UpdatedAt: getTrain.UpdatedAt,
		}
		trainResponses = append(trainResponses, trainResponse)
	}
	return trainResponses, count, nil
}

// GetTrainByID godoc
// @Summary      Get train by ID
// @Description  Get train by ID
// @Tags         Admin - Train
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

	trainResponse := dtos.TrainResponse{
		TrainID:   train.ID,
		CodeTrain: train.CodeTrain,
		Name:      train.Name,
		Status:    train.Status,
		CreatedAt: train.CreatedAt,
		UpdatedAt: train.UpdatedAt,
	}
	return trainResponse, nil
}

// CreateTrain godoc
// @Summary      Create a new train
// @Description  Create a new train
// @Tags         Admin - Train
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
		CodeTrain: train.CodeTrain,
		Name:      train.Name,
		Status:    train.Status,
	}

	createdTrain, err := u.trainRepo.CreateTrain(createTrain)
	if err != nil {
		return trainResponsee, err
	}

	for _, train := range train.Route {
		station, err := u.trainRepo.GetStationByID2(train.StationID)
		if err != nil {
			return trainResponsee, err
		}
		trainStation := models.TrainStation{
			TrainID:    createdTrain.ID,
			StationID:  station.ID,
			ArriveTime: train.ArriveTime,
		}
		_, err = u.trainRepo.CreateTrainStation(trainStation)
		if err != nil {
			return trainResponsee, err
		}
	}

	getTrain, err := u.trainRepo.GetTrainByID(createdTrain.ID)
	if err != nil {
		return trainResponsee, err
	}

	getTrainStation, err := u.trainRepo.GetTrainStationByTrainID(getTrain.ID)
	if err != nil {
		return trainResponsee, err
	}

	var trainStationResponses []dtos.TrainStationResponse
	for _, train := range getTrainStation {
		getStation, err := u.trainRepo.GetStationByID2(train.StationID)
		if err != nil {
			return trainResponsee, err
		}
		trainStationResponse := dtos.TrainStationResponse{
			StationID: train.StationID,
			Station: dtos.StationInput{
				Origin:  getStation.Origin,
				Name:    getStation.Name,
				Initial: getStation.Initial,
			},
			ArriveTime: train.ArriveTime,
		}
		trainStationResponses = append(trainStationResponses, trainStationResponse)
	}

	trainResponse := dtos.TrainResponse{
		TrainID:   getTrain.ID,
		CodeTrain: getTrain.CodeTrain,
		Name:      getTrain.Name,
		Route:     trainStationResponses,
		Status:    getTrain.Status,
		CreatedAt: getTrain.CreatedAt,
		UpdatedAt: getTrain.UpdatedAt,
	}
	return trainResponse, nil
}

// UpdateTrain godoc
// @Summary      Update train
// @Description  Update train
// @Tags         Admin - Train
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

	train, err := u.trainRepo.GetTrainByID(id)

	if err != nil {
		return trainResponse, err
	}

	train.CodeTrain = trainInput.CodeTrain
	train.Name = trainInput.Name
	train.Status = trainInput.Status

	train, err = u.trainRepo.UpdateTrain(train)

	if err != nil {
		return trainResponse, err
	}

	trainResponse.TrainID = train.ID
	trainResponse.CodeTrain = train.CodeTrain
	trainResponse.Name = train.Name
	trainResponse.Status = train.Status
	trainResponse.CreatedAt = train.CreatedAt
	trainResponse.UpdatedAt = train.UpdatedAt

	return trainResponse, nil
}

// DeleteTrain godoc
// @Summary      Delete a train
// @Description  Delete a train
// @Tags         Admin - Train
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

// =============================== ADMIN END ================================== \\

// =============================== USER ================================== \\

// SearchTrainAvailable godoc
// @Summary      Search Train Available
// @Description  Search Train
// @Tags         User - Train
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param station_origin_id query int true "Station origin id"
// @Param station_destination_id query int true "Station destination id"
// @Param sort_by_train_id query int false "Filter by train id"
// @Param sort_by_class query string false "Filter by class name"
// @Param sort_by_price query string false "Filter by price"
// @Param sort_by_arrive_time query string false "Filter by arrive time"
// @Success      200 {object} dtos.GetAllTrainStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/train/search [get]
// @Security BearerAuth
func (u *trainUsecase) SearchTrainAvailable(page, limit, stationOriginId, stationDestinationId, sortByTrainId int, sortClassName, sortByPrice, sortByArriveTime string) ([]dtos.TrainResponse, int, error) {
	trains, count, err := u.trainRepo.GetAllTrains(page, limit)
	if err != nil {
		return nil, 0, err
	}

	// Track train IDs and classes
	trainMap := make(map[uint]map[string]bool)

	var trainResponses []dtos.TrainResponse

	for _, train := range trains {
		getTrain, err := u.trainRepo.GetTrainByID(train.TrainID)
		if err != nil {
			return trainResponses, count, err
		}

		getTrainStation, err := u.trainRepo.SearchTrainAvailable(getTrain.ID, uint(stationOriginId), uint(stationDestinationId))
		if err != nil {
			return trainResponses, count, err
		}

		// Check if route[0] matches stationOriginId and route[1] matches stationDestinationId
		if getTrainStation[0].StationID != uint(stationOriginId) || getTrainStation[1].StationID != uint(stationDestinationId) {
			continue
		}

		// Check if train ID exists in the map
		if classMap, ok := trainMap[getTrain.ID]; ok {
			// Check if class already exists for the train ID
			if classMap[strings.ToLower(train.Class)] || (sortClassName != "" && strings.ToLower(train.Class) != strings.ToLower(sortClassName)) {
				continue
			}
		} else {
			// Create a new class map for the train ID
			trainMap[getTrain.ID] = make(map[string]bool)
		}

		// Check if class has already been added or doesn't match the filterClass
		if getTrain.ID != uint(sortByTrainId) && sortByTrainId != 0 {
			continue
		}

		// Add class to the train ID map
		trainMap[getTrain.ID][strings.ToLower(train.Class)] = true

		var trainStationResponses []dtos.TrainStationResponse

		for _, trainStation := range getTrainStation {
			getStation, err := u.trainRepo.GetStationByID2(trainStation.StationID)
			if err != nil {
				return trainResponses, count, err
			}

			trainStationResponse := dtos.TrainStationResponse{
				StationID: trainStation.StationID,
				Station: dtos.StationInput{
					Origin:  getStation.Origin,
					Name:    getStation.Name,
					Initial: getStation.Initial,
				},
				ArriveTime: trainStation.ArriveTime,
			}

			trainStationResponses = append(trainStationResponses, trainStationResponse)
		}

		trainResponse := dtos.TrainResponse{
			TrainID:   getTrain.ID,
			CodeTrain: getTrain.CodeTrain,
			Name:      getTrain.Name,
			Class:     train.Class,
			Price:     train.Price,
			Route:     trainStationResponses,
			Status:    getTrain.Status,
			CreatedAt: getTrain.CreatedAt,
			UpdatedAt: getTrain.UpdatedAt,
		}
		trainResponses = append(trainResponses, trainResponse)

		// Sort trainResponses based on price
		if strings.ToLower(sortByPrice) == "asc" {
			sort.SliceStable(trainResponses, func(i, j int) bool {
				return trainResponses[i].Price < trainResponses[j].Price
			})
		} else if strings.ToLower(sortByPrice) == "desc" {
			sort.SliceStable(trainResponses, func(i, j int) bool {
				return trainResponses[i].Price > trainResponses[j].Price
			})
		}

		// Sort trainResponses based on arrive time
		if strings.ToLower(sortByArriveTime) == "asc" {
			sort.SliceStable(trainResponses, func(i, j int) bool {
				return trainResponses[i].Route[0].ArriveTime < trainResponses[j].Route[0].ArriveTime
			})
		} else if strings.ToLower(sortByArriveTime) == "desc" {
			sort.SliceStable(trainResponses, func(i, j int) bool {
				return trainResponses[i].Route[0].ArriveTime > trainResponses[j].Route[0].ArriveTime
			})
		}
	}
	return trainResponses, len(trainResponses), nil
}
