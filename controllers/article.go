package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleController interface {
	GetAllArticles(c echo.Context) error
	GetArticleByID(c echo.Context) error
	CreateArticle(c echo.Context) error
	UpdateArticle(c echo.Context) error
	DeleteArticle(c echo.Context) error
}

type articleController struct {
	articleUsecase usecases.ArticleUsecase
}

func NewArticleController(articleUsecase usecases.ArticleUsecase) ArticleController {
	return &articleController{articleUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *articleController) GetAllArticles(ctx echo.Context) error {
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

	articles, count, err := c.articleUsecase.GetAllArticles(page, limit)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all article",
			articles,
			page,
			limit,
			count,
		),
	)
}

func (c *articleController) GetArticleByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	article, err := c.articleUsecase.GetArticleByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get article by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get article by id",
			article,
		),
	)

}

func (c *articleController) CreateArticle(ctx echo.Context) error {
	var articleDTO dtos.ArticleInput
	if err := ctx.Bind(&articleDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	article, err := c.articleUsecase.CreateArticle(&articleDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a article",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a article",
			article,
		),
	)
}

func (c *articleController) UpdateArticle(ctx echo.Context) error {

	var articleInput dtos.ArticleInput
	if err := ctx.Bind(&articleInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	article, err := c.articleUsecase.GetArticleByID(uint(id))
	if article.ArticleID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get article by id",
				helpers.GetErrorData(err),
			),
		)
	}

	articleResp, err := c.articleUsecase.UpdateArticle(uint(id), articleInput)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated article",
			articleResp,
		),
	)
}

func (c *articleController) DeleteArticle(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.articleUsecase.DeleteArticle(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted article",
			nil,
		),
	)
}
