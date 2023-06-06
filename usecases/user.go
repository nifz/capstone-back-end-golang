package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/middlewares"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"errors"
	"strings"
	"time"
)

type UserUsecase interface {
	UserLogin(input dtos.UserLoginInput) (dtos.UserInformationResponse, error)
	UserRegister(input dtos.UserRegisterInput) (dtos.UserInformationResponse, error)
	UserUpdatePassword(userId uint, input dtos.UserUpdatePasswordInput) (dtos.UserInformationResponse, error)
	UserUpdateProfile(userId uint, input dtos.UserUpdateProfileInput) (dtos.UserInformationResponse, error)
	UserCredential(userId uint) (dtos.UserInformationResponse, error)
	UserUpdatePhotoProfile(userId uint, input dtos.UserUpdatePhotoProfileInput) (dtos.UserInformationResponse, error)
	UserDeletePhotoProfile(userId uint) (dtos.UserInformationResponse, error)
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

// UserLogin godoc
// @Summary      Login
// @Description  Login an account
// @Tags         User - Account
// @Accept       json
// @Produce      json
// @Param        request body dtos.UserLoginInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.UserStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /login [post]
func (u *userUsecase) UserLogin(input dtos.UserLoginInput) (dtos.UserInformationResponse, error) {
	var (
		userResponse dtos.UserInformationResponse
		accessToken  string
	)

	input.Email = strings.ToLower(input.Email)

	user, err := u.userRepo.UserGetByEmail(input.Email)
	if err != nil {
		return userResponse, errors.New("Email or password is wrong")
	}

	valid := helpers.ComparePassword(input.Password, user.Password)
	if !valid {
		return userResponse, errors.New("Email or password is wrong")
	}

	accessToken, err = middlewares.CreateToken(user.ID, user.Role)
	if err != nil {
		return userResponse, err
	}

	userResponse.ID = user.ID
	userResponse.FullName = user.FullName
	userResponse.Email = user.Email
	userResponse.PhoneNumber = user.PhoneNumber
	userResponse.BirthDate = helpers.FormatDateToYMD(user.BirthDate)
	userResponse.ProfilePicture = user.ProfilePicture
	userResponse.Citizen = user.Citizen
	userResponse.Role = user.Role
	userResponse.Token = &accessToken
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return userResponse, nil
}

// UserRegister godoc
// @Summary      Register
// @Description  Register an account
// @Tags         User - Account
// @Accept       json
// @Produce      json
// @Param        request body dtos.UserRegisterInput true "Payload Body [RAW]"
// @Success      201 {object} dtos.UserCreeatedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /register [post]
func (u *userUsecase) UserRegister(input dtos.UserRegisterInput) (dtos.UserInformationResponse, error) {
	var (
		user         models.User
		userResponse dtos.UserInformationResponse
	)

	input.Email = strings.ToLower(input.Email)

	user, err := u.userRepo.UserGetByEmail(input.Email)
	if user.ID > 0 {
		return userResponse, errors.New("Email already used")
	}

	if input.Password != input.ConfirmPassword {
		return userResponse, errors.New("Password does not match")
	}

	password, err := helpers.HashPassword(input.Password)
	if err != nil {
		return userResponse, err
	}

	if input.Email == "" || input.FullName == "" || input.Password == "" || input.ConfirmPassword == "" || input.Role == "admin" {
		return userResponse, errors.New("Failed to create user")
	}

	user.FullName = input.FullName
	user.Email = input.Email
	user.Password = password
	user.PhoneNumber = input.PhoneNumber
	user.ProfilePicture = "https://icon-library.com/images/default-user-icon/default-user-icon-13.jpg"
	user.Citizen = "Indonesia"
	user.Role = "user"

	user, err = u.userRepo.UserCreate(user)
	if err != nil {
		return userResponse, err
	}

	userResponse.ID = user.ID
	userResponse.FullName = user.FullName
	userResponse.Email = user.Email
	userResponse.PhoneNumber = user.PhoneNumber
	userResponse.BirthDate = helpers.FormatDateToYMD(user.BirthDate)
	userResponse.ProfilePicture = user.ProfilePicture
	userResponse.Citizen = user.Citizen
	userResponse.Role = user.Role
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return userResponse, err
}

// UserUpdatePassword godoc
// @Summary      Update Password
// @Description  User update an password
// @Tags         User - Account
// @Accept       json
// @Produce      json
// @Param        request body dtos.UserUpdatePasswordInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.UserStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/update-password [put]
// @Security BearerAuth
func (u *userUsecase) UserUpdatePassword(userId uint, input dtos.UserUpdatePasswordInput) (dtos.UserInformationResponse, error) {
	var (
		user         models.User
		userResponse dtos.UserInformationResponse
	)

	user, err := u.userRepo.UserGetById(userId)
	if err != nil {
		return userResponse, errors.New("User not found")
	}

	if input.OldPassword == "" || input.NewPassword == "" || input.ConfirmPassword == "" {
		return userResponse, errors.New("Failed to update user password")
	}

	password := helpers.ComparePassword(input.OldPassword, user.Password)
	if !password {
		return userResponse, errors.New("Old password is incorrect")
	}

	if input.OldPassword == input.NewPassword {
		return userResponse, errors.New("New password must be different")
	}

	if input.NewPassword != input.ConfirmPassword {
		return userResponse, errors.New("Confirm password does not match")
	}

	newPassword, err := helpers.HashPassword(input.ConfirmPassword)
	if err != nil {
		return userResponse, err
	}

	user.Password = newPassword

	user, err = u.userRepo.UserUpdate(user)
	if err != nil {
		return userResponse, err
	}

	userResponse.ID = user.ID
	userResponse.FullName = user.FullName
	userResponse.Email = user.Email
	userResponse.PhoneNumber = user.PhoneNumber
	userResponse.BirthDate = helpers.FormatDateToYMD(user.BirthDate)
	userResponse.ProfilePicture = user.ProfilePicture
	userResponse.Citizen = user.Citizen
	userResponse.Role = user.Role
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return userResponse, err
}

// UserUpdateProfile godoc
// @Summary      Update Profile
// @Description  User update an profile
// @Tags         User - Account
// @Accept       json
// @Produce      json
// @Param        request body dtos.UserUpdateProfileInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.UserStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/update-profile [put]
// @Security BearerAuth
func (u *userUsecase) UserUpdateProfile(userId uint, input dtos.UserUpdateProfileInput) (dtos.UserInformationResponse, error) {
	var (
		user         models.User
		userResponse dtos.UserInformationResponse
	)

	user, err := u.userRepo.UserGetById(userId)
	if err != nil {
		return userResponse, errors.New("User not found")
	}

	dateNow := "2006-01-02"
	birthDateParse, err := time.Parse(dateNow, input.BirthDate)
	if err != nil {
		return userResponse, errors.New("Failed to parse birth date")
	}

	user.FullName = input.FullName
	user.PhoneNumber = input.PhoneNumber
	user.BirthDate = &birthDateParse
	user.Citizen = input.Citizen

	user, err = u.userRepo.UserUpdate(user)
	if err != nil {
		return userResponse, err
	}

	userResponse.ID = user.ID
	userResponse.FullName = user.FullName
	userResponse.Email = user.Email
	userResponse.PhoneNumber = user.PhoneNumber
	userResponse.BirthDate = helpers.FormatDateToYMD(user.BirthDate)
	userResponse.ProfilePicture = user.ProfilePicture
	userResponse.Citizen = user.Citizen
	userResponse.Role = user.Role
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return userResponse, err
}

// UserCredential godoc
// @Summary      Get Credentials
// @Description  User get credentials
// @Tags         User - Account
// @Accept       json
// @Produce      json
// @Success      200 {object} dtos.UserStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user [get]
// @Security BearerAuth
func (u *userUsecase) UserCredential(userId uint) (dtos.UserInformationResponse, error) {
	var (
		user         models.User
		userResponse dtos.UserInformationResponse
	)

	user, err := u.userRepo.UserGetById(userId)
	if err != nil {
		return userResponse, errors.New("User not found")
	}

	userResponse.ID = user.ID
	userResponse.FullName = user.FullName
	userResponse.Email = user.Email
	userResponse.PhoneNumber = user.PhoneNumber
	userResponse.BirthDate = helpers.FormatDateToYMD(user.BirthDate)
	userResponse.ProfilePicture = user.ProfilePicture
	userResponse.Citizen = user.Citizen
	userResponse.Role = user.Role
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return userResponse, err
}

// UserUpdatePhotoProfile godoc
// @Summary      Update Photo Profile
// @Description  User update an photo profile
// @Tags         User - Account
// @Accept       json
// @Produce      json
// @Param        file formData file false "Photo file"
// @Success      200 {object} dtos.UserStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/update-photo-profile [put]
// @Security BearerAuth
func (u *userUsecase) UserUpdatePhotoProfile(userId uint, input dtos.UserUpdatePhotoProfileInput) (dtos.UserInformationResponse, error) {
	var (
		user         models.User
		userResponse dtos.UserInformationResponse
	)

	user, err := u.userRepo.UserGetById(userId)
	if err != nil {
		return userResponse, errors.New("User not found")
	}

	user.ProfilePicture = input.ProfilePicture

	user, err = u.userRepo.UserUpdate(user)
	if err != nil {
		return userResponse, err
	}

	userResponse.ID = user.ID
	userResponse.FullName = user.FullName
	userResponse.Email = user.Email
	userResponse.PhoneNumber = user.PhoneNumber
	userResponse.BirthDate = helpers.FormatDateToYMD(user.BirthDate)
	userResponse.ProfilePicture = user.ProfilePicture
	userResponse.Citizen = user.Citizen
	userResponse.Role = user.Role
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return userResponse, err
}

// UserDeletePhotoProfile godoc
// @Summary      Update Information
// @Description  User update an information
// @Tags         User - Account
// @Accept       json
// @Produce      json
// @Success      200 {object} dtos.UserStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/delete-photo-profile [delete]
// @Security BearerAuth
func (u *userUsecase) UserDeletePhotoProfile(userId uint) (dtos.UserInformationResponse, error) {
	var (
		user         models.User
		userResponse dtos.UserInformationResponse
	)

	user, err := u.userRepo.UserGetById(userId)
	if err != nil {
		return userResponse, errors.New("User not found")
	}

	user.ProfilePicture = "https://icon-library.com/images/default-user-icon/default-user-icon-13.jpg"

	user, err = u.userRepo.UserUpdate(user)
	if err != nil {
		return userResponse, err
	}

	userResponse.ID = user.ID
	userResponse.FullName = user.FullName
	userResponse.Email = user.Email
	userResponse.PhoneNumber = user.PhoneNumber
	userResponse.BirthDate = helpers.FormatDateToYMD(user.BirthDate)
	userResponse.ProfilePicture = user.ProfilePicture
	userResponse.Citizen = user.Citizen
	userResponse.Role = user.Role
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return userResponse, err
}
