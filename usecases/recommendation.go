package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
)

type RecommendationUsecase interface {
	GetAllRecommendations(page, limit int) ([]dtos.RecomendationResponse, int, error)
	GetRecommendationByID(id uint) (dtos.RecomendationResponse, error)
	CreateRecommendation(recommendationInput *dtos.RecomendationInput) (dtos.RecomendationResponse, error)
	UpdateRecommendation(id uint, recommendationInput dtos.RecomendationInput) (dtos.RecomendationResponse, error)
	DeleteRecommendation(id uint) error
}

type recommendationUsecase struct {
	recommendationRepo repositories.RecommendationRepository
}

func NewRecommendationUsecase(RecommendationRepo repositories.RecommendationRepository) RecommendationUsecase {
	return &recommendationUsecase{RecommendationRepo}
}

// GetAllRecommendations godoc
// @Summary      Get all recommendation
// @Description  Get all recommendation
// @Tags         Recommendation
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success      200 {object} dtos.GetRecommendationStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/recommendation [get]
// @Security BearerAuth
func (u *recommendationUsecase) GetAllRecommendations(page, limit int) ([]dtos.RecomendationResponse, int, error) {
	recommendations, count, err := u.recommendationRepo.GetAllRecommendations(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var recommendationResponses []dtos.RecomendationResponse
	for _, recommendation := range recommendations {
		recommendationResponse := dtos.RecomendationResponse{
			RecommendationID: recommendation.ID,
			Tag:              recommendation.Tag,
			CreatedAt:        recommendation.CreatedAt,
			UpdatedAt:        recommendation.UpdatedAt,
		}
		recommendationResponses = append(recommendationResponses, recommendationResponse)
	}

	return recommendationResponses, count, nil
}

// GetRecommendationByID godoc
// @Summary      Get recommendation by ID
// @Description  Get recommendation by ID
// @Tags         Recommendation
// @Accept       json
// @Produce      json
// @Param id path integer true "ID recommendation"
// @Success      200 {object} dtos.RecommendationStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/recommendation/{id} [get]
// @Security BearerAuth
func (u *recommendationUsecase) GetRecommendationByID(id uint) (dtos.RecomendationResponse, error) {
	var recommendationResponses dtos.RecomendationResponse
	recommendation, err := u.recommendationRepo.GetRecommendationByID(id)
	if err != nil {
		return recommendationResponses, err
	}
	recommendationResponse := dtos.RecomendationResponse{
		RecommendationID: recommendation.ID,
		Tag:              recommendation.Tag,
		CreatedAt:        recommendation.CreatedAt,
		UpdatedAt:        recommendation.UpdatedAt,
	}
	return recommendationResponse, nil
}

// CreateRecommendation godoc
// @Summary      Create a new recommendation
// @Description  Create a new recommendation
// @Tags         Recommendation
// @Accept       json
// @Produce      json
// @Param        request body dtos.RecomendationInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.StationStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/recommendation [post]
// @Security BearerAuth
func (u *recommendationUsecase) CreateRecommendation(recommendationInput *dtos.RecomendationInput) (dtos.RecomendationResponse, error) {
	var recommendationResponses dtos.RecomendationResponse
	createRecommendation := models.Recomendation{
		Tag: recommendationInput.Tag,
	}

	createdRecommendation, err := u.recommendationRepo.CreateRecommendation(createRecommendation)
	if err != nil {
		return recommendationResponses, err
	}

	recommendationResponse := dtos.RecomendationResponse{
		RecommendationID: createdRecommendation.ID,
		Tag:              createdRecommendation.Tag,
		CreatedAt:        createdRecommendation.CreatedAt,
		UpdatedAt:        createdRecommendation.UpdatedAt,
	}
	return recommendationResponse, nil
}

// UpdateRecommendation godoc
// @Summary      Update recommendation
// @Description  Update recommendation
// @Tags         Recommendation
// @Accept       json
// @Produce      json
// @Param id path integer true "ID recommendation"
// @Param        request body dtos.RecomendationInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.RecommendationStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/recommendation [put]
// @Security BearerAuth
func (u *recommendationUsecase) UpdateRecommendation(id uint, recommendationInput dtos.RecomendationInput) (dtos.RecomendationResponse, error) {
	var recommendation models.Recomendation
	var recommendationResponse dtos.RecomendationResponse

	recommendation, err := u.recommendationRepo.GetRecommendationByID(id)
	if err != nil {
		return recommendationResponse, err
	}

	recommendation.Tag = recommendationInput.Tag

	recommendation, err = u.recommendationRepo.UpdateRecommendation(recommendation)

	if err != nil {
		return recommendationResponse, err
	}

	recommendationResponse.RecommendationID = recommendation.ID
	recommendationResponse.Tag = recommendation.Tag
	recommendationResponse.CreatedAt = recommendation.CreatedAt
	recommendationResponse.UpdatedAt = recommendation.UpdatedAt

	return recommendationResponse, nil

}

// DeleteRecommendation godoc
// @Summary      Delete a recommendation
// @Description  Delete a recommendation
// @Tags         Recommendation
// @Accept       json
// @Produce      json
// @Param id path integer true "ID recommendation"
// @Success      200 {object} dtos.StatusOKDeletedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/recommendation/{id} [delete]
// @Security BearerAuth
func (u *recommendationUsecase) DeleteRecommendation(id uint) error {
	recommendation, err := u.recommendationRepo.GetRecommendationByID(id)

	if err != nil {
		return nil
	}

	err = u.recommendationRepo.DeleteRecommendation(recommendation)
	return err
}
