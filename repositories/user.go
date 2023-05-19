package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	UserGetById(id uint) (models.User, error)
	UserGetByEmail(email string) (models.User, error)
	UserCreate(user models.User) (models.User, error)
	UserUpdate(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) UserGetById(id uint) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *userRepository) UserGetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) UserCreate(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) UserUpdate(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}
