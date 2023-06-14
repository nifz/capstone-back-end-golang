package controllers

import (
	"back-end-golang/helpers"
	"back-end-golang/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NotificationController interface {
	GetNotificationByUserID(c echo.Context) error
}

type notificationController struct {
	notificationUsecase usecases.NotificationUsecase
}

func NewNotificationController(notificationUsecase usecases.NotificationUsecase) NotificationController {
	return &notificationController{notificationUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *notificationController) GetNotificationByUserID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	notification, err := c.notificationUsecase.GetNotificationByUserID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get notification by user id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get notification by user id",
			notification,
		),
	)

}
