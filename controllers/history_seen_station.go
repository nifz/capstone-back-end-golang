package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/middlewares"
	"back-end-golang/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HistorySeenStationController interface {
	GetAllHistorySeenStations(c echo.Context) error
	CreateHistorySeenStation(c echo.Context) error
}

type historySeenStationController struct {
	historySeenStationUsecase usecases.HistorySeenStationUsecase
}

func NewHistorySeenStationController(historySeenStationUsecase usecases.HistorySeenStationUsecase) HistorySeenStationController {
	return &historySeenStationController{historySeenStationUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *historySeenStationController) GetAllHistorySeenStations(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(nil),
			),
		)
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	if err != nil {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(err),
			),
		)
	}

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

	historySeenStations, count, err := c.historySeenStationUsecase.GetAllHistorySeenStations(page, limit, userId)
	if err != nil {

		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.NewErrorResponse(
				http.StatusInternalServerError,
				"Failed fetching historySeenStations",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all historySeenStation",
			historySeenStations,
			page,
			limit,
			count,
		),
	)
}

func (c *historySeenStationController) CreateHistorySeenStation(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(nil),
			),
		)
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	if err != nil {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(err),
			),
		)
	}
	var historySeenStationInput dtos.HistorySeenStationInput
	if err := ctx.Bind(&historySeenStationInput); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding historySeenStation",
				helpers.GetErrorData(err),
			),
		)
	}

	historySeenStation, err := c.historySeenStationUsecase.CreateHistorySeenStation(userId, historySeenStationInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a historySeenStation",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a historySeenStation",
			historySeenStation,
		),
	)
}
