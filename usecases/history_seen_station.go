package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"errors"
)

type HistorySeenStationUsecase interface {
	GetAllHistorySeenStations(page, limit int, userId uint) ([]dtos.HistorySeenStationResponse, int, error)
	CreateHistorySeenStation(userId uint, historySeenStationInput dtos.HistorySeenStationInput) (dtos.HistorySeenStationResponse, error)
}

type historySeenStationUsecase struct {
	historySeenStationRepo repositories.HistorySeenStationRepository
	stationRepo            repositories.StationRepository
}

func NewHistorySeenStationUsecase(historySeenStationRepo repositories.HistorySeenStationRepository, stationRepo repositories.StationRepository) HistorySeenStationUsecase {
	return &historySeenStationUsecase{historySeenStationRepo, stationRepo}
}

// GetAllHistorySeenStations godoc
// @Summary      Get all history seen station by user id
// @Description  Get all history seen station by user id
// @Tags         User - History Seen
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success      200 {object} dtos.GetAllHistorySeenStationStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/history-seen-station [get]
// @Security BearerAuth
func (u *historySeenStationUsecase) GetAllHistorySeenStations(page, limit int, userId uint) ([]dtos.HistorySeenStationResponse, int, error) {
	historySeenStations, count, err := u.historySeenStationRepo.GetAllHistorySeenStation(page, limit, userId)
	if err != nil {
		return nil, 0, err
	}

	var historySeenStationResponses []dtos.HistorySeenStationResponse
	for _, historySeenStation := range historySeenStations {
		getStationOrigin, err := u.stationRepo.GetStationByID(historySeenStation.StationOriginID)
		if err != nil {
			continue
		}
		getStationDestination, err := u.stationRepo.GetStationByID(historySeenStation.StationDestinationID)
		if err != nil {
			continue
		}
		historySeenStationResponse := dtos.HistorySeenStationResponse{
			ID: historySeenStation.ID,
			StationOrigin: dtos.StationResponseSimply2{
				StationID: getStationOrigin.ID,
				Origin:    getStationOrigin.Origin,
				Name:      getStationOrigin.Name,
				Initial:   getStationOrigin.Initial,
			},
			StationDestination: dtos.StationResponseSimply2{
				StationID: getStationDestination.ID,
				Origin:    getStationDestination.Origin,
				Name:      getStationDestination.Name,
				Initial:   getStationDestination.Initial,
			},
			CreatedAt: historySeenStation.CreatedAt,
			UpdatedAt: historySeenStation.UpdatedAt,
		}
		historySeenStationResponses = append(historySeenStationResponses, historySeenStationResponse)
	}

	return historySeenStationResponses, count, nil
}

func (u *historySeenStationUsecase) CreateHistorySeenStation(userId uint, historySeenStationInput dtos.HistorySeenStationInput) (dtos.HistorySeenStationResponse, error) {
	var historySeenStationResponses dtos.HistorySeenStationResponse

	if historySeenStationInput.StationDestinationID < 1 || historySeenStationInput.StationOriginID < 1 {
		return historySeenStationResponses, errors.New("Failed to create history seen station")
	}

	createHistorySeenStation := models.HistorySeenStation{
		UserID:               userId,
		StationOriginID:      historySeenStationInput.StationOriginID,
		StationDestinationID: historySeenStationInput.StationDestinationID,
	}

	getHistorySeenStation, _ := u.historySeenStationRepo.GetHistorySeenStationByID(historySeenStationInput.StationOriginID, historySeenStationInput.StationDestinationID, userId)
	if getHistorySeenStation.ID > 0 {
		createHistorySeenStation, _ = u.historySeenStationRepo.UpdateHistorySeenStation(getHistorySeenStation)
	} else {
		createHistorySeenStation, _ = u.historySeenStationRepo.CreateHistorySeenStation(createHistorySeenStation)
	}

	getStationOrigin, _ := u.stationRepo.GetStationByID(createHistorySeenStation.StationOriginID)
	getStationDestination, _ := u.stationRepo.GetStationByID(createHistorySeenStation.StationDestinationID)

	historySeenStationResponse := dtos.HistorySeenStationResponse{
		ID: createHistorySeenStation.ID,
		StationOrigin: dtos.StationResponseSimply2{
			StationID: getStationOrigin.ID,
			Origin:    getStationOrigin.Origin,
			Name:      getStationOrigin.Name,
			Initial:   getStationOrigin.Initial,
		},
		StationDestination: dtos.StationResponseSimply2{
			StationID: getStationDestination.ID,
			Origin:    getStationDestination.Origin,
			Name:      getStationDestination.Name,
			Initial:   getStationDestination.Initial,
		},
		CreatedAt: createHistorySeenStation.CreatedAt,
		UpdatedAt: createHistorySeenStation.UpdatedAt,
	}
	return historySeenStationResponse, nil
}
