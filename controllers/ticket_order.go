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

type TicketOrderController interface {
	CreateTicketOrder(c echo.Context) error
	UpdateTicketOrder(c echo.Context) error
}

type ticketOrderController struct {
	ticketOrderUsecase usecases.TicketOrderUsecase
}

func NewTicketOrderController(ticketOrderUsecase usecases.TicketOrderUsecase) TicketOrderController {
	return &ticketOrderController{ticketOrderUsecase}
}

func (c *ticketOrderController) CreateTicketOrder(ctx echo.Context) error {
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

	var ticketOrderInput dtos.TicketOrderInput
	if err := ctx.Bind(&ticketOrderInput); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding ticket order",
				helpers.GetErrorData(err),
			),
		)
	}

	ticketOrder, err := c.ticketOrderUsecase.CreateTicketOrder(userId, ticketOrderInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a ticket order",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a ticket order",
			ticketOrder,
		),
	)
}

func (c *ticketOrderController) UpdateTicketOrder(ctx echo.Context) error {
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

	var ticketOrderInput dtos.TicketOrderInput
	if err := ctx.Bind(&ticketOrderInput); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding ticket order",
				helpers.GetErrorData(err),
			),
		)
	}

	ticketOrderIDParam := ctx.QueryParam("ticket_order_id")
	ticketOrderID, _ := strconv.Atoi(ticketOrderIDParam)

	statusParam := ctx.QueryParam("status")

	ticketOrder, err := c.ticketOrderUsecase.UpdateTicketOrder(userId, uint(ticketOrderID), statusParam)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to update a ticket order",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to update a ticket order",
			ticketOrder,
		),
	)

}
