package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
)

type RecommendationUsecase interface {
	GetAllRecommendations(page, limit int) ([]dtos.RecommendationResponse, int, error)
	GetRecommendationByID(id uint) (dtos.RecommendationResponse, error)
	CreateRecommendation(recommendationInput *dtos.RecommendationInput) (dtos.RecommendationResponse, error)
	UpdateRecommendation(id uint, recommendationInput dtos.RecommendationInput) (dtos.RecommendationResponse, error)
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
// @Tags         Admin - Recommendation
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success      200 {object} dtos.GetAllRecommendationStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/recommendation [get]
// @Security BearerAuth
func (u *recommendationUsecase) GetAllRecommendations(page, limit int) ([]dtos.RecommendationResponse, int, error) {
	recommendations, count, err := u.recommendationRepo.GetAllRecommendations(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var recommendationResponses []dtos.RecommendationResponse
	for _, recommendation := range recommendations {
		recommendationResponse := dtos.RecommendationResponse{
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
// @Tags         Admin - Recommendation
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
func (u *recommendationUsecase) GetRecommendationByID(id uint) (dtos.RecommendationResponse, error) {
	var recommendationResponses dtos.RecommendationResponse
	recommendation, err := u.recommendationRepo.GetRecommendationByID(id)
	if err != nil {
		return recommendationResponses, err
	}
	recommendationResponse := dtos.RecommendationResponse{
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
// @Tags         Admin - Recommendation
// @Accept       json
// @Produce      json
// @Param        request body dtos.RecommendationInput true "Payload Body [RAW]"
// @Success      201 {object} dtos.RecommendationCreeatedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/recommendation [post]
// @Security BearerAuth
func (u *recommendationUsecase) CreateRecommendation(recommendationInput *dtos.RecommendationInput) (dtos.RecommendationResponse, error) {
	var recommendationResponses dtos.RecommendationResponse
	createRecommendation := models.Recommendation{
		Tag: recommendationInput.Tag,
	}

	createdRecommendation, err := u.recommendationRepo.CreateRecommendation(createRecommendation)
	if err != nil {
		return recommendationResponses, err
	}

	recommendationResponse := dtos.RecommendationResponse{
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
// @Tags         Admin - Recommendation
// @Accept       json
// @Produce      json
// @Param id path integer true "ID recommendation"
// @Param        request body dtos.RecommendationInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.RecommendationStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/recommendation/{id} [put]
// @Security BearerAuth
func (u *recommendationUsecase) UpdateRecommendation(id uint, recommendationInput dtos.RecommendationInput) (dtos.RecommendationResponse, error) {
	var recommendation models.Recommendation
	var recommendationResponse dtos.RecommendationResponse

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
// @Tags         Admin - Recommendation
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
		return err
	}

	err = u.recommendationRepo.DeleteRecommendation(recommendation)
	return err
}
