package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserByAdminController interface {
	GetAllUserByAdmin(c echo.Context) error
	GetUserByAdminByID(c echo.Context) error
	CreateUserByAdmin(c echo.Context) error
	UpdateUserByAdmin(c echo.Context) error
	DeleteUserByAdmin(c echo.Context) error
}

type userByAdminController struct {
	userByadminUsecase usecases.UserByAdminUsecase
}

func NewUserByAdminController(userByAdminUsecase usecases.UserByAdminUsecase) UserByAdminController {
	return &userByAdminController{userByAdminUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *userByAdminController) GetAllUserByAdmin(ctx echo.Context) error {
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

	users, count, err := c.userByadminUsecase.GetAllUserByAdmin(page, limit)
	if err != nil {

		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.NewErrorResponse(
				http.StatusInternalServerError,
				"Failed fetching users",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all user",
			users,
			page,
			limit,
			count,
		),
	)
}

func (c *userByAdminController) GetUserByAdminByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := c.userByadminUsecase.GetUserByAdminByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get user by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get user by id",
			user,
		),
	)

}

func (c *userByAdminController) CreateUserByAdmin(ctx echo.Context) error {
	var userDTO dtos.UserRegisterInputByAdmin
	if err := ctx.Bind(&userDTO); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding user",
				helpers.GetErrorData(err),
			),
		)
	}

	user, err := c.userByadminUsecase.CreateUserByAdmin(userDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a user",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a user",
			user,
		),
	)
}

func (c *userByAdminController) UpdateUserByAdmin(ctx echo.Context) error {

	var userInput dtos.UserRegisterInputByAdmin
	if err := ctx.Bind(&userInput); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding user",
				helpers.GetErrorData(err),
			),
		)
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	user, err := c.userByadminUsecase.GetUserByAdminByID(uint(id))
	if user.ID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get user by id",
				helpers.GetErrorData(err),
			),
		)
	}

	userResp, err := c.userByadminUsecase.UpdateUserByAdmin(uint(id), userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding user",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated user",
			userResp,
		),
	)
}

func (c *userByAdminController) DeleteUserByAdmin(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.userByadminUsecase.DeleteUserByAdmin(uint(id))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to delete user",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted user",
			nil,
		),
	)
}
