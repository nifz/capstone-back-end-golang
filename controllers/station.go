package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/usecases"
	"net/http"
	"strconv"

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

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully get all stations",
			stations,
		),
	)
}

func (c *stationController) GetStationByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	station, err := c.stationUsecase.GetStationByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get station by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get station by id",
			station,
		),
	)

}

func (c *stationController) CreateStation(ctx echo.Context) error {
	var stationDTO dtos.StationInput
	if err := ctx.Bind(&stationDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	station, err := c.stationUsecase.CreateStation(&stationDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a station",
				helpers.GetErrorData(err),
			),
		)
	}

	stationResponse := dtos.StationResponse{
		StationID: station.StationID,
		Origin:    station.Origin,
		Name:      station.Name,
		Initial:   station.Initial,
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a station",
			stationResponse,
		),
	)
}

func (c *stationController) UpdateStation(ctx echo.Context) error {

	var stationInput dtos.StationInput
	if err := ctx.Bind(&stationInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	station, err := c.stationUsecase.GetStationByID(uint(id))
	if station.StationID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get station by id",
				helpers.GetErrorData(err),
			),
		)
	}

	stationResp, err := c.stationUsecase.UpdateStation(uint(id), stationInput)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	stationResponse := dtos.StationResponse{
		StationID: stationResp.StationID,
		Origin:    stationResp.Origin,
		Name:      stationResp.Name,
		Initial:   stationResp.Initial,
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated station",
			stationResponse,
		),
	)
}

func (c *stationController) DeleteStation(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.stationUsecase.DeleteStation(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted station",
			nil,
		),
	)
}
