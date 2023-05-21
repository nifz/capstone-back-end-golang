package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type reservationController struct {
	reservationUsecase usecases.ReservationUsecase
}

func NewReservationController(reservationUsecase usecases.ReservationUsecase) reservationController {
	return reservationController{reservationUsecase}
}

func (c *reservationController) AdminCreateReservation(ctx echo.Context) error {
	var reservationInput dtos.ReservationCreateInput

	err := ctx.Bind(&reservationInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create reservation bad request",
				helpers.GetErrorData(err),
			),
		)
	}

	reservation, err := c.reservationUsecase.AdminCreateReservation(reservationInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create reservation",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully created reservation",
			reservation,
		),
	)
}
