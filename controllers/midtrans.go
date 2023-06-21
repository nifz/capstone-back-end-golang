package controllers

import (
	"back-end-golang/configs"
	"back-end-golang/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var c coreapi.Client

func initiateCoreApiClient() {
	c.New(configs.EnvMidtransServerKey(), midtrans.Sandbox)
}

// CheckTransaction godoc
// @Summary      Get Transaction order by midtrans
// @Description  Get Transaction order by midtrans
// @Tags         User - Order
// @Accept       json
// @Produce      json
// @Param order_id query string true "Order id"
// @Success      200 {object} dtos.StatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /public/transaction [get]
func CheckTransaction(ctx echo.Context) error {
	initiateCoreApiClient()
	res, err := c.CheckTransaction(ctx.QueryParam("order_id"))
	if err != nil {
		// do something on error handle
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully get transaction",
			res,
		),
	)
}
