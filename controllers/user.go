package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase usecases.UserUsecase
}

func NewUserController(userUsecase usecases.UserUsecase) UserController {
	return UserController{userUsecase}
}

func (c *UserController) Login(ctx echo.Context) error {
	var userInput dtos.UserLoginInput

	err := ctx.Bind(&userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to login",
				helpers.GetErrorData(err),
			),
		)
	}

	tokenString, err := c.userUsecase.Login(userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed getting token",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully logged in",
			echo.Map{
				"token": tokenString,
			},
		),
	)
}

func (c *UserController) Register(ctx echo.Context) error {
	var userInput dtos.UserRegisterInput
	err := ctx.Bind(&userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to register",
				helpers.GetErrorData(err),
			),
		)
	}

	user, err := c.userUsecase.Register(userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to register",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully registered",
			user,
		),
	)
}
