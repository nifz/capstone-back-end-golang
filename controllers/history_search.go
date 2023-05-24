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

type HistorySearchController struct {
	historyUsecase usecases.HistorySearchUseCase
}

func NewHistorySearchController(historyUsecase usecases.HistorySearchUseCase) HistorySearchController {
	return HistorySearchController{historyUsecase}
}

func (h *HistorySearchController) HistorySearchGetAll(ctx echo.Context) error {

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
	histories, err := h.historyUsecase.HistorySearchGetAll(userId)
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.NewErrorResponse(
				http.StatusInternalServerError,
				"Failed fetching history search",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully get all history search",
			histories,
		),
	)
}

func (h *HistorySearchController) HistorySearchCreate(ctx echo.Context) error {
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

	var historyInput dtos.HistorySearchInput
	err = ctx.Bind(&historyInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create history search data",
				helpers.GetErrorData(err),
			),
		)
	}

	history, err := h.historyUsecase.HistorySearchCreate(userId, historyInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create history search data",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully create history",
			history,
		),
	)
}

func (h *HistorySearchController) HistorySearchDelete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

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

	_, err = h.historyUsecase.HistorySearchDelete(userId, uint(id))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to delete history search data",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully delete history",
			nil,
		),
	)
}
