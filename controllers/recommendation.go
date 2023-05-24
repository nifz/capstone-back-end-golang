package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RecommendationController interface {
	GetAllRecommendations(c echo.Context) error
	GetRecommendationByID(c echo.Context) error
	CreateRecommendation(c echo.Context) error
	UpdateRecommendation(c echo.Context) error
	DeleteRecommendation(c echo.Context) error
}

type recommendationController struct {
	recommendationUsecase usecases.RecommendationUsecase
}

func NewRecommendationController(recommendationUsecase usecases.RecommendationUsecase) RecommendationController {
	return &recommendationController{recommendationUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *recommendationController) GetAllRecommendations(ctx echo.Context) error {
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

	recommendations, count, err := c.recommendationUsecase.GetAllRecommendations(page, limit)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed fetching recommendations",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all recommendation",
			recommendations,
			page,
			limit,
			count,
		),
	)
}

func (c *recommendationController) GetRecommendationByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	recommendation, err := c.recommendationUsecase.GetRecommendationByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get recommendation by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get recommendation by id",
			recommendation,
		),
	)

}

func (c *recommendationController) CreateRecommendation(ctx echo.Context) error {
	var recommendationDTO dtos.RecommendationInput
	if err := ctx.Bind(&recommendationDTO); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding article",
				helpers.GetErrorData(err),
			),
		)
	}

	recommendation, err := c.recommendationUsecase.CreateRecommendation(&recommendationDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a recommendation",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a recommendation",
			recommendation,
		),
	)
}

func (c *recommendationController) UpdateRecommendation(ctx echo.Context) error {

	var recommendationInput dtos.RecommendationInput
	if err := ctx.Bind(&recommendationInput); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding article",
				helpers.GetErrorData(err),
			),
		)
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	recommendation, err := c.recommendationUsecase.GetRecommendationByID(uint(id))
	if recommendation.RecommendationID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get recommendation by id",
				helpers.GetErrorData(err),
			),
		)
	}

	recommendationResp, err := c.recommendationUsecase.UpdateRecommendation(uint(id), recommendationInput)
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.NewErrorResponse(
				http.StatusInternalServerError,
				"Failed update recommendation",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated recommendation",
			recommendationResp,
		),
	)
}

func (c *recommendationController) DeleteRecommendation(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.recommendationUsecase.DeleteRecommendation(uint(id))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed deleting recommendation",
				helpers.GetErrorData(err),
			),
		)
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted recommendation",
			nil,
		),
	)
}
