package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TrainController interface {
	GetAllTrains(c echo.Context) error
	GetTrainByID(c echo.Context) error
	CreateTrain(c echo.Context) error
	UpdateTrain(c echo.Context) error
	DeleteTrain(c echo.Context) error
}

type trainController struct {
	trainUsecase usecases.TrainUsecase
}

func NewTrainController(trainUsecase usecases.TrainUsecase) TrainController {
	return &trainController{trainUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *trainController) GetAllTrains(ctx echo.Context) error {
	trains, err := c.trainUsecase.GetAllTrains()
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all train",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully get all trains",
			trains,
		),
	)
}

func (c *trainController) GetTrainByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	train, err := c.trainUsecase.GetTrainByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get train by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get train by id",
			train,
		),
	)

}

func (c *trainController) CreateTrain(ctx echo.Context) error {
	var trainDTO dtos.TrainInput
	if err := ctx.Bind(&trainDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	train, err := c.trainUsecase.CreateTrain(&trainDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a train",
				helpers.GetErrorData(err),
			),
		)
	}

	trainResponse := dtos.TrainResponse{
		TrainID:         train.TrainID,
		StationOriginID: train.StationOriginID,
		StationOrigin: dtos.StationInput{
			Origin:  train.StationOrigin.Origin,
			Name:    train.StationOrigin.Name,
			Initial: train.StationOrigin.Initial,
		},
		StationDestinationID: train.StationDestinationID,
		StationDestination: dtos.StationInput{
			Origin:  train.StationDestination.Origin,
			Name:    train.StationDestination.Name,
			Initial: train.StationDestination.Initial,
		},
		DepartureTime: train.DepartureTime,
		ArriveTime:    train.ArriveTime,
		Name:          train.Name,
		Route:         train.Route,
		Status:        train.Status,
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a train",
			trainResponse,
		),
	)
}

func (c *trainController) UpdateTrain(ctx echo.Context) error {

	var trainInput dtos.TrainInput
	if err := ctx.Bind(&trainInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	train, err := c.trainUsecase.GetTrainByID(uint(id))
	if train.TrainID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get train by id",
				helpers.GetErrorData(err),
			),
		)
	}

	trainResp, err := c.trainUsecase.UpdateTrain(uint(id), trainInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to updated a train",
				helpers.GetErrorData(err),
			),
		)
	}

	trainResponse := dtos.TrainResponse{
		TrainID:         trainResp.TrainID,
		StationOriginID: trainResp.StationOriginID,
		StationOrigin: dtos.StationInput{
			Origin:  trainResp.StationOrigin.Origin,
			Name:    trainResp.StationOrigin.Name,
			Initial: trainResp.StationOrigin.Initial,
		},
		StationDestinationID: trainResp.StationDestinationID,
		StationDestination: dtos.StationInput{
			Origin:  trainResp.StationDestination.Origin,
			Name:    trainResp.StationDestination.Name,
			Initial: trainResp.StationDestination.Initial,
		},
		DepartureTime: trainResp.DepartureTime,
		ArriveTime:    trainResp.ArriveTime,
		Name:          trainResp.Name,
		Route:         trainResp.Route,
		Status:        trainResp.Status,
		UpdateAt:      trainResp.UpdateAt,
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated train",
			trainResponse,
		),
	)
}

func (c *trainController) DeleteTrain(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.trainUsecase.DeleteTrain(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted train",
			nil,
		),
	)
}
