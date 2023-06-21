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

type HistorySeenHotelController interface {
	GetAllHistorySeenHotels(c echo.Context) error
	CreateHistorySeenHotel(c echo.Context) error
}

type historySeenHotelController struct {
	historySeenHotelUsecase usecases.HistorySeenHotelUsecase
}

func NewHistorySeenHotelController(historySeenHotelUsecase usecases.HistorySeenHotelUsecase) HistorySeenHotelController {
	return &historySeenHotelController{historySeenHotelUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *historySeenHotelController) GetAllHistorySeenHotels(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				"Unauthorized",
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

	historySeenHotels, count, err := c.historySeenHotelUsecase.GetAllHistorySeenHotels(page, limit, userId)
	if err != nil {

		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.NewErrorResponse(
				http.StatusInternalServerError,
				"Failed fetching historySeenHotels",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all historySeenHotel",
			historySeenHotels,
			page,
			limit,
			count,
		),
	)
}

func (c *historySeenHotelController) CreateHistorySeenHotel(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				"Unauthorized",
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
	var historySeenHotelInput dtos.HistorySeenHotelInput
	if err := ctx.Bind(&historySeenHotelInput); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding historySeenHotel",
				helpers.GetErrorData(err),
			),
		)
	}

	historySeenHotel, err := c.historySeenHotelUsecase.CreateHistorySeenHotel(userId, historySeenHotelInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a historySeenHotel",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a historySeenHotel",
			historySeenHotel,
		),
	)
}
