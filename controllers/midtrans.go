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

func CheckStatusB2B(ctx echo.Context) error {
	res, err := c.GetStatusB2B(ctx.Param("id"))
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

func ApproveTransaction(ctx echo.Context) error {
	res, err := c.ApproveTransaction(ctx.Param("id"))
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

func DenyTransaction(ctx echo.Context) error {
	res, err := c.DenyTransaction(ctx.Param("id"))
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

func CancelTransaction(ctx echo.Context) error {
	res, err := c.CancelTransaction(ctx.Param("id"))
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

func ExpireTransaction(ctx echo.Context) error {
	res, err := c.ExpireTransaction(ctx.Param("id"))
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

func CaptureTransaction(ctx echo.Context) error {
	refundRequest := &coreapi.CaptureReq{
		TransactionID: "TRANSACTION-ID",
		GrossAmt:      10000,
	}
	res, err := c.CaptureTransaction(refundRequest)
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

func RefundTransaction(ctx echo.Context) error {
	refundRequest := &coreapi.RefundReq{
		Amount: 5000,
		Reason: "Item out of stock",
	}

	res, err := c.RefundTransaction("YOUR_ORDER_ID_or_TRANSACTION_ID", refundRequest)
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

func DirectRefundTransaction(ctx echo.Context) error {
	refundRequest := &coreapi.RefundReq{
		RefundKey: "order1-ref1",
		Amount:    5000,
		Reason:    "Item out of stock",
	}

	// Optional: set payment idempotency key to prevent duplicate request
	c.Options.SetPaymentIdempotencyKey("UNIQUE-ID")

	res, err := c.DirectRefundTransaction("YOUR_ORDER_ID_or_TRANSACTION-ID", refundRequest)
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
