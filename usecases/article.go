package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"fmt"
)

type ArticleUsecase interface {
	GetAllArticles(page, limit int) ([]dtos.ArticleResponse, int, error)
	GetArticleByID(id uint) (dtos.ArticleResponse, error)
	CreateArticle(articleInput *dtos.ArticleInput) (dtos.ArticleResponse, error)
	UpdateArticle(id uint, articleInput dtos.ArticleInput) (dtos.ArticleResponse, error)
	DeleteArticle(id uint) error
}

type articleUsecase struct {
	articleRepo repositories.ArticleRepository
}

func NewArticleUsecase(ArticleRepo repositories.ArticleRepository) ArticleUsecase {
	return &articleUsecase{ArticleRepo}
}

// GetAllArticles godoc
// @Summary      Get all articles
// @Description  Get all articles
// @Tags         Admin - Article
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success      200 {object} dtos.GetAllArticleStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/article [get]
// @Security BearerAuth
func (u *articleUsecase) GetAllArticles(page, limit int) ([]dtos.ArticleResponse, int, error) {
	articles, count, err := u.articleRepo.GetAllArticles(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var articleResponses []dtos.ArticleResponse
	for _, article := range articles {
		articleResponse := dtos.ArticleResponse{
			ArticleID:   article.ID,
			Title:       article.Title,
			Image:       article.Image,
			Description: article.Description,
			Label:       article.Label,
			Slug:        article.Slug,
			CreatedAt:   article.CreatedAt,
			UpdatedAt:   article.UpdatedAt,
		}
		articleResponses = append(articleResponses, articleResponse)
	}

	return articleResponses, count, nil
}

// GetArticleByID godoc
// @Summary      Get article by ID
// @Description  Get article by ID
// @Tags         Admin - Article
// @Accept       json
// @Produce      json
// @Param id path integer true "ID article"
// @Success      200 {object} dtos.ArticleStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/article/{id} [get]
// @Security BearerAuth
func (u *articleUsecase) GetArticleByID(id uint) (dtos.ArticleResponse, error) {
	var articleResponses dtos.ArticleResponse
	article, err := u.articleRepo.GetArticleByID(id)
	if err != nil {
		return articleResponses, err
	}
	articleResponse := dtos.ArticleResponse{
		ArticleID:   article.ID,
		Title:       article.Title,
		Image:       article.Image,
		Description: article.Description,
		Label:       article.Label,
		Slug:        article.Slug,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}
	return articleResponse, nil
}

// CreateArticle godoc
// @Summary      Create a new article
// @Description  Create a new article
// @Tags         Admin - Article
// @Accept       json
// @Produce      json
// @Param        request body dtos.ArticleInput true "Payload Body [RAW]"
// @Success      201 {object} dtos.ArticleCreeatedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/article [post]
// @Security BearerAuth
func (u *articleUsecase) CreateArticle(articleInput *dtos.ArticleInput) (dtos.ArticleResponse, error) {
	var articleResponses dtos.ArticleResponse
	slug := helpers.CreateSlug(articleInput.Title)
	CreateArticle := models.Article{
		Title:       articleInput.Title,
		Image:       articleInput.Image,
		Description: articleInput.Description,
		Label:       articleInput.Label,
		Slug:        slug,
	}

	createdArticle, err := u.articleRepo.CreateArticle(CreateArticle)
	fmt.Println(CreateArticle)
	if err != nil {
		return articleResponses, err
	}

	articleResponse := dtos.ArticleResponse{
		ArticleID:   createdArticle.ID,
		Title:       createdArticle.Title,
		Image:       createdArticle.Image,
		Description: createdArticle.Description,
		Label:       createdArticle.Label,
		Slug:        createdArticle.Slug,
		CreatedAt:   createdArticle.CreatedAt,
		UpdatedAt:   createdArticle.UpdatedAt,
	}
	return articleResponse, nil
}

// UpdateArticle godoc
// @Summary      Update article
// @Description  Update article
// @Tags         Admin - Article
// @Accept       json
// @Produce      json
// @Param id path integer true "ID article"
// @Param        request body dtos.ArticleInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.ArticleStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/article [put]
// @Security BearerAuth
func (u *articleUsecase) UpdateArticle(id uint, articleInput dtos.ArticleInput) (dtos.ArticleResponse, error) {
	var article models.Article
	var articleResponse dtos.ArticleResponse

	article, err := u.articleRepo.GetArticleByID(id)
	if err != nil {
		return articleResponse, err
	}

	slug := helpers.CreateSlug(articleInput.Title)

	article.Title = articleInput.Title
	article.Image = articleInput.Image
	article.Description = articleInput.Description
	article.Label = articleInput.Label
	article.Slug = slug

	article, err = u.articleRepo.UpdateArticle(article)

	if err != nil {
		return articleResponse, err
	}

	articleResponse.ArticleID = article.ID
	articleResponse.Title = article.Title
	articleResponse.Image = article.Image
	articleResponse.Description = article.Description
	articleResponse.Label = article.Label
	articleResponse.Slug = article.Slug
	articleResponse.CreatedAt = article.CreatedAt
	articleResponse.UpdatedAt = article.UpdatedAt

	return articleResponse, nil

}

// DeleteArticle godoc
// @Summary      Delete a article
// @Description  Delete a article
// @Tags         Admin - Article
// @Accept       json
// @Produce      json
// @Param id path integer true "ID article"
// @Success      200 {object} dtos.StatusOKDeletedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/article/{id} [delete]
// @Security BearerAuth
func (u *articleUsecase) DeleteArticle(id uint) error {
	article, err := u.articleRepo.GetArticleByID(id)

	if err != nil {
		return err
	}

	err = u.articleRepo.DeleteArticle(article)
	return err
}
