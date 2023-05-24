package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"back-end-golang/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type reservationController struct {
	reservationUsecase   usecases.ReservationUsecase
	reservationImageRepo repositories.ReservationImageRepository
}

func NewReservationController(reservationUsecase usecases.ReservationUsecase, reservationImageRepo repositories.ReservationImageRepository) reservationController {
	return reservationController{
		reservationUsecase:   reservationUsecase,
		reservationImageRepo: reservationImageRepo,
	}
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
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to retrieve the image file",
				helpers.GetErrorData(err),
			),
		)
	}
	reservation, err := c.reservationUsecase.AdminCreateReservation(reservationInput, file)
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

	// Check if there is an image file
	if reservationInput.ImageFile != "" {
		// Create the reservation image entity
		reservationImage := models.ReservationImages{
			ReservationId: reservation.ID,
			Image:         reservationInput.ImageFile,
		}

		// Save the reservation image
		err = c.reservationImageRepo.SaveReservationImage(reservationImage)
		if err != nil {
			return ctx.JSON(
				http.StatusBadRequest,
				helpers.NewErrorResponse(
					http.StatusBadRequest,
					"Failed to save reservation image",
					helpers.GetErrorData(err),
				),
			)
		}
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

func (c *reservationController) GetAllReservation(ctx echo.Context) error {
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

	reservations, total, err := c.reservationUsecase.GetAllReservation(page, limit)
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.NewErrorResponse(
				http.StatusInternalServerError,
				"Failed fetching recommendations",
				helpers.GetErrorData(err),
			),
		)
	}
	baseURL := ctx.Scheme() + "://" + ctx.Request().Host + "/api/v1/admin"
	// Menambahkan URL gambar dengan URL dasar aplikasi menggunakan ctx.BaseURL()
	for i := range reservations {
		reservations[i].Image = baseURL + "/images/" + reservations[i].Image
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all reservations",
			reservations,
			page,
			limit,
			total,
		),
	)
}
