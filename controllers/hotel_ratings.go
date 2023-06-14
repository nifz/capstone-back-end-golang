package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HotelRatingsController interface {
	//user
	CreateHotelRating(c echo.Context) error
	//admin
	GetRatingsByHotelsId(c echo.Context) error
}

type hotelRatingsController struct {
	hotelRatingUsecase usecases.HotelRatingsUsecase
}

func NewHotelRatingsController(hotelRatingUsecase usecases.HotelRatingsUsecase) HotelRatingsController {
	return &hotelRatingsController{hotelRatingUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *hotelRatingsController) CreateHotelRating(ctx echo.Context) error {
	var hotelRatingsInputDTO dtos.HotelRatingInput

	if err := ctx.Bind(&hotelRatingsInputDTO); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding hotel rating",
				helpers.GetErrorData(err),
			),
		)
	}

	ratings, err := c.hotelRatingUsecase.CreateHotelRating(hotelRatingsInputDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create hotel rating",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Success to create hotel rating",
			ratings,
		),
	)
}

func (c *hotelRatingsController) GetRatingsByHotelsId(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	hotelId := uint(id)
	ratings, err := c.hotelRatingUsecase.GetHotelRatingsByHotelID(hotelId)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get hotel rating",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Success to get hotel rating",
			ratings,
		),
	)
}
