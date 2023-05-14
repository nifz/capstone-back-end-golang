package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/middlewares"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"errors"
	"time"
)

type UserUsecase interface {
	Login(input dtos.UserLoginInput) (string, error)
	Register(input dtos.UserRegisterInput) (dtos.UserRegisterResponse, error)
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) Login(input dtos.UserLoginInput) (string, error) {
	var accessToken string

	user, _ := u.userRepo.GetUserByEmail(input.Email)
	if user.ID == 0 {
		return accessToken, errors.New("Email or password is wrong")
	}

	valid := helpers.ComparePassword(input.Password, user.Password)
	if !valid {
		return accessToken, errors.New("Email or password is wrong")
	}

	accessToken, err := middlewares.CreateToken(user.ID)
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}

func (u *userUsecase) Register(input dtos.UserRegisterInput) (dtos.UserRegisterResponse, error) {
	var (
		user         models.User
		userResponse dtos.UserRegisterResponse
	)

	user, _ = u.userRepo.GetUserByEmail(input.Email)
	if user.ID > 0 {
		return userResponse, errors.New("Email already used")
	}

	password, err := helpers.HashPassword(input.Password)
	if err != nil {
		return userResponse, err
	}

	if input.Email == "" || input.FullName == "" || input.Password == "" {
		return userResponse, errors.New("Failed to create user")
	}

	dateNow := "2006-01-02"
	birthDateParse, _ := time.Parse(dateNow, input.BirthDate)

	user.FullName = input.FullName
	user.Email = input.Email
	user.Password = password
	user.PhoneNumber = input.PhoneNumber
	user.Gender = input.Gender
	user.BirthDate = birthDateParse
	user.ProfilePicture = input.ProfilePicture
	user.Role = input.Role

	user, err = u.userRepo.CreateUser(user)
	if err != nil {
		return userResponse, err
	}

	userResponse.FullName = user.FullName
	userResponse.Email = user.Email
	userResponse.PhoneNumber = user.PhoneNumber
	userResponse.Gender = user.Gender
	userResponse.BirthDate = user.BirthDate
	userResponse.ProfilePicture = user.ProfilePicture
	userResponse.Role = user.Role
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return userResponse, err
}
