package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StationController interface {
	GetAllStations(c echo.Context) error
	GetStationByID(c echo.Context) error
	CreateStation(c echo.Context) error
	UpdateStation(c echo.Context) error
	DeleteStation(c echo.Context) error
}

type stationController struct {
	stationUsecase usecases.StationUsecase
}

func NewStationController(stationUsecase usecases.StationUsecase) StationController {
	return &stationController{stationUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *stationController) GetAllStations(ctx echo.Context) error {
	stations, err := c.stationUsecase.GetAllStations()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	var stationDTOs []dtos.TrainStationDTO

	for _, station := range stations {
		stationDTO := dtos.TrainStationDTO{
			ID:      station.ID,
			Name:    station.Name,
			Initial: station.Initial,
		}
		stationDTOs = append(stationDTOs, stationDTO)
	}

	return ctx.JSON(http.StatusOK, dtos.TrainStationDTOsResponse{
		Message: "Successfully get all station.",
		Data:    stationDTOs,
	})
}

func (c *stationController) GetStationByID(ctx echo.Context) error {
	id := ctx.Param("id")
	station, err := c.stationUsecase.GetStationByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, dtos.TrainStationDTOsResponse{
		Message: "Successfully get station by id.",
		Data: dtos.TrainStationDTO{
			ID:        station.ID,
			Name:      station.Name,
			Initial:   station.Initial,
			CreatedAt: station.CreatedAt,
			UpdatedAt: station.UpdatedAt,
		},
	})
}

func (c *stationController) CreateStation(ctx echo.Context) error {
	var stationDTO dtos.TrainStationDTO
	if err := ctx.Bind(&stationDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	station := models.Station{
		Name:    stationDTO.Name,
		Initial: stationDTO.Initial,
	}

	err := c.stationUsecase.CreateStation(&station)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, dtos.TrainStationDTOsResponse{
		Message: "Successfully created station.",
		Data: dtos.TrainStationDTO{
			ID:        stationDTO.ID,
			Name:      stationDTO.Name,
			Initial:   stationDTO.Initial,
			CreatedAt: stationDTO.CreatedAt,
			UpdatedAt: stationDTO.UpdatedAt,
		},
	})
}

func (c *stationController) UpdateStation(ctx echo.Context) error {
	id := ctx.Param("id")

	var stationDTO dtos.TrainStationDTO
	if err := ctx.Bind(&stationDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	station, err := c.stationUsecase.GetStationByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	station.Name = stationDTO.Name
	station.Initial = stationDTO.Initial

	err = c.stationUsecase.UpdateStation(&station)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, dtos.TrainStationDTOsResponse{
		Message: "Successfully updated station.",
		Data: dtos.TrainStationDTO{
			ID:        stationDTO.ID,
			Name:      stationDTO.Name,
			Initial:   stationDTO.Initial,
			CreatedAt: stationDTO.CreatedAt,
			UpdatedAt: stationDTO.UpdatedAt,
		},
	})
}

func (c *stationController) DeleteStation(ctx echo.Context) error {
	id := ctx.Param("id")

	err := c.stationUsecase.DeleteStation(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, dtos.ErrorDTO{
		Message: "Successfully deleted station.",
	})
}
