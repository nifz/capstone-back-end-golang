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

type HotelOrderController interface {
	GetHotelOrders(c echo.Context) error
	GetHotelOrdersByAdmin(c echo.Context) error
	GetHotelOrderDetailByAdmin(c echo.Context) error
	GetHotelOrderByID(c echo.Context) error
	CreateHotelOrder(c echo.Context) error
	UpdateHotelOrder(c echo.Context) error
}

type hotelOrderController struct {
	hotelOrderUsecase usecases.HotelOrderUsecase
}

func NewHotelOrderController(hotelOrderUsecase usecases.HotelOrderUsecase) HotelOrderController {
	return &hotelOrderController{hotelOrderUsecase}
}

func (c *hotelOrderController) GetHotelOrders(ctx echo.Context) error {
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
		limit = 1000
	}

	statusParam := ctx.QueryParam("status")

	hotelOrder, count, err := c.hotelOrderUsecase.GetHotelOrders(page, limit, userId, statusParam)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get a hotel order",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully to get order hotels",
			hotelOrder,
			page,
			limit,
			count,
		),
	)
}

func (c *hotelOrderController) GetHotelOrdersByAdmin(ctx echo.Context) error {
	pageParam := ctx.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 1
	}

	limitParam := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 1000
	}

	searchParam := ctx.QueryParam("search")
	dateStartParam := ctx.QueryParam("date_start")
	dateEndParam := ctx.QueryParam("date_end")
	orderByParam := ctx.QueryParam("order_by")
	filterParam := ctx.QueryParam("filter")
	ratingClass, _ := strconv.Atoi(ctx.QueryParam("rating_class"))

	hotelOrder, count, err := c.hotelOrderUsecase.GetHotelOrdersByAdmin(page, limit, ratingClass, searchParam, dateStartParam, dateEndParam, orderByParam, filterParam)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get a hotel order",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully to get order hotels",
			hotelOrder,
			page,
			limit,
			count,
		),
	)
}

func (c *hotelOrderController) GetHotelOrderDetailByAdmin(ctx echo.Context) error {
	hotelOrderIdParam := ctx.QueryParam("hotel_order_id")
	hotelOrderId, _ := strconv.Atoi(hotelOrderIdParam)

	hotelOrder, err := c.hotelOrderUsecase.GetHotelOrdersDetailByAdmin(uint(hotelOrderId))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get a hotel order",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get order hotels",
			hotelOrder,
		),
	)
}

func (c *hotelOrderController) GetHotelOrderByID(ctx echo.Context) error {
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

	hotelOrderIdParam := ctx.QueryParam("hotel_order_id")
	hotelOrderId, _ := strconv.Atoi(hotelOrderIdParam)

	hotelOrder, err := c.hotelOrderUsecase.GetHotelOrderByID(userId, uint(hotelOrderId))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get a hotel order",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get order hotels",
			hotelOrder,
		),
	)
}

func (c *hotelOrderController) CreateHotelOrder(ctx echo.Context) error {
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

	var hotelOrderInput dtos.HotelOrderInput
	if err := ctx.Bind(&hotelOrderInput); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding hotel order",
				helpers.GetErrorData(err),
			),
		)
	}

	hotelOrder, err := c.hotelOrderUsecase.CreateHotelOrder(userId, hotelOrderInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a hotel order",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a hotel order",
			hotelOrder,
		),
	)
}

func (c *hotelOrderController) UpdateHotelOrder(ctx echo.Context) error {
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

	var hotelOrderInput dtos.HotelOrderInput
	if err := ctx.Bind(&hotelOrderInput); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding hotel order",
				helpers.GetErrorData(err),
			),
		)
	}

	hotelOrderIDParam := ctx.QueryParam("hotel_order_id")
	hotelOrderID, _ := strconv.Atoi(hotelOrderIDParam)

	statusParam := ctx.QueryParam("status")

	hotelOrder, err := c.hotelOrderUsecase.UpdateHotelOrder(userId, uint(hotelOrderID), statusParam)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to update a hotel order",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to update a hotel order",
			hotelOrder,
		),
	)

}
