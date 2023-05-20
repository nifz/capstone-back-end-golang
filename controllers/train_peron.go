package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TrainPeronController interface {
	GetAllTrainPerons(c echo.Context) error
	GetTrainPeronByID(c echo.Context) error
	CreateTrainPeron(c echo.Context) error
	UpdateTrainPeron(c echo.Context) error
	DeleteTrainPeron(c echo.Context) error
}

type trainPeronController struct {
	trainPeronUsecase usecases.TrainPeronUsecase
}

func NewTrainPeronController(trainPeronUsecase usecases.TrainPeronUsecase) TrainPeronController {
	return &trainPeronController{trainPeronUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *trainPeronController) GetAllTrainPerons(ctx echo.Context) error {
	trainPerons, err := c.trainPeronUsecase.GetAllTrainPerons()
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all train peron",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully get all train peron",
			trainPerons,
		),
	)
}

func (c *trainPeronController) GetTrainPeronByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	trainPeron, err := c.trainPeronUsecase.GetTrainPeronByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get train peron by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get train peron by id",
			trainPeron,
		),
	)

}

func (c *trainPeronController) CreateTrainPeron(ctx echo.Context) error {
	var trainPeronDTO dtos.TrainPeronInput
	if err := ctx.Bind(&trainPeronDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	trainPeron, err := c.trainPeronUsecase.CreateTrainPeron(&trainPeronDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a train peron",
				helpers.GetErrorData(err),
			),
		)
	}

	trainPeronResponse := dtos.TrainPeronResponse{
		TrainPeronID: trainPeron.TrainPeronID,
		TrainID:      trainPeron.TrainID,
		Train: dtos.TrainInput{
			StationOriginID:      trainPeron.Train.StationOriginID,
			StationDestinationID: trainPeron.Train.StationDestinationID,
			DepartureTime:        trainPeron.Train.DepartureTime,
			ArriveTime:           trainPeron.Train.ArriveTime,
			Name:                 trainPeron.Train.Name,
			Route:                trainPeron.Train.Route,
			Status:               trainPeron.Train.Status,
		},
		Class:    trainPeron.Class,
		Name:     trainPeron.Name,
		Price:    trainPeron.Price,
		Status:   trainPeron.Status,
		UpdateAt: trainPeron.UpdateAt,
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a train peron",
			trainPeronResponse,
		),
	)
}

func (c *trainPeronController) UpdateTrainPeron(ctx echo.Context) error {

	var trainPeronInput dtos.TrainPeronInput
	if err := ctx.Bind(&trainPeronInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	trainPeron, err := c.trainPeronUsecase.GetTrainPeronByID(uint(id))
	if trainPeron.TrainPeronID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get train peron by id",
				helpers.GetErrorData(err),
			),
		)
	}

	trainPeronResp, err := c.trainPeronUsecase.UpdateTrainPeron(uint(id), trainPeronInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to updated a train peron",
				helpers.GetErrorData(err),
			),
		)
	}

	trainPeronResponse := dtos.TrainPeronResponse{
		TrainPeronID: trainPeronResp.TrainPeronID,
		TrainID:      trainPeronResp.TrainID,
		Train: dtos.TrainInput{
			StationOriginID:      trainPeronResp.Train.StationOriginID,
			StationDestinationID: trainPeronResp.Train.StationDestinationID,
			DepartureTime:        trainPeronResp.Train.DepartureTime,
			ArriveTime:           trainPeronResp.Train.ArriveTime,
			Name:                 trainPeronResp.Train.Name,
			Route:                trainPeronResp.Train.Route,
			Status:               trainPeronResp.Train.Status,
		},
		Class:    trainPeronResp.Class,
		Name:     trainPeronResp.Name,
		Price:    trainPeronResp.Price,
		Status:   trainPeronResp.Status,
		UpdateAt: trainPeronResp.UpdateAt,
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated train peron",
			trainPeronResponse,
		),
	)
}

func (c *trainPeronController) DeleteTrainPeron(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.trainPeronUsecase.DeleteTrainPeron(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted train peron",
			nil,
		),
	)
}
