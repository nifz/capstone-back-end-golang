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
	pageParam := ctx.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 1
	}

	limitParam := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 10
	}
	trainPerons, count, err := c.trainPeronUsecase.GetAllTrainPerons(page, limit)
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
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all train peron",
			trainPerons,
			page,
			limit,
			count,
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

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a train peron",
			trainPeron,
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

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated train peron",
			trainPeronResp,
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
