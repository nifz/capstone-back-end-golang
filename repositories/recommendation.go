package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type RecommendationRepository interface {
	GetAllRecommendations(page, limit int) ([]models.Recomendation, int, error)
	GetRecommendationByID(id uint) (models.Recomendation, error)
	CreateRecommendation(recommendation models.Recomendation) (models.Recomendation, error)
	UpdateRecommendation(recommendation models.Recomendation) (models.Recomendation, error)
	DeleteRecommendation(recommendation models.Recomendation) error
}

type recommendationRepository struct {
	db *gorm.DB
}

func NewRecommendationRepository(db *gorm.DB) RecommendationRepository {
	return &recommendationRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *recommendationRepository) GetAllRecommendations(page, limit int) ([]models.Recomendation, int, error) {
	var (
		recommendations []models.Recomendation
		count           int64
	)
	err := r.db.Find(&recommendations).Count(&count).Error
	if err != nil {
		return recommendations, int(count), err
	}

	offset := (page - 1) * limit

	err = r.db.Limit(limit).Offset(offset).Find(&recommendations).Error

	return recommendations, int(count), err
}

func (r *recommendationRepository) GetRecommendationByID(id uint) (models.Recomendation, error) {
	var recommendation models.Recomendation
	err := r.db.Where("id = ?", id).First(&recommendation).Error
	return recommendation, err
}

func (r *recommendationRepository) CreateRecommendation(recommendation models.Recomendation) (models.Recomendation, error) {
	err := r.db.Create(&recommendation).Error
	return recommendation, err
}

func (r *recommendationRepository) UpdateRecommendation(recommendation models.Recomendation) (models.Recomendation, error) {
	err := r.db.Save(&recommendation).Error
	return recommendation, err
}

func (r *recommendationRepository) DeleteRecommendation(recommendation models.Recomendation) error {
	err := r.db.Delete(&recommendation).Error
	return err
}
