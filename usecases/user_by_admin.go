package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"errors"
	"strings"
)

type UserByAdminUsecase interface {
	GetAllUserByAdmin(page, limit int) ([]dtos.UserInformationResponse, int, error)
	GetUserByAdminByID(id uint) (dtos.UserInformationResponse, error)
	CreateUserByAdmin(input dtos.UserRegisterInputByAdmin) (dtos.UserInformationResponse, error)
	UpdateUserByAdmin(id uint, input dtos.UserRegisterInputByAdmin) (dtos.UserInformationResponse, error)
	DeleteUserByAdmin(id uint) error
}

type userByAdminUsecase struct {
	userByAdminRepo repositories.UserRepository
}

func NewUserByAdminUsecase(UserByAdminRepo repositories.UserRepository) UserByAdminUsecase {
	return &userByAdminUsecase{UserByAdminRepo}
}

// GetAllUserByAdmin godoc
// @Summary      Get all user by admin
// @Description  Get all user by admin
// @Tags         Admin - user
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success      200 {object} dtos.GetAllUserByAdminStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/user [get]
// @Security BearerAuth
func (u *userByAdminUsecase) GetAllUserByAdmin(page, limit int) ([]dtos.UserInformationResponse, int, error) {
	users, count, err := u.userByAdminRepo.GetAllUser(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var userRegisterResponse []dtos.UserInformationResponse
	for _, user := range users {
		userResponse := dtos.UserInformationResponse{
			ID:             user.ID,
			FullName:       user.FullName,
			Email:          user.Email,
			PhoneNumber:    user.PhoneNumber,
			Gender:         user.Gender,
			BirthDate:      helpers.FormatDateToYMD(user.BirthDate),
			ProfilePicture: user.ProfilePicture,
			Citizen:        user.Citizen,
			Role:           user.Role,
			CreatedAt:      user.CreatedAt,
			UpdatedAt:      user.UpdatedAt,
		}
		userRegisterResponse = append(userRegisterResponse, userResponse)
	}

	return userRegisterResponse, count, nil
}

// GetUserByAdminByID godoc
// @Summary      Get user by admin by ID
// @Description  Get user by admin by ID
// @Tags         Admin - User
// @Accept       json
// @Produce      json
// @Param id path integer true "ID user"
// @Success      200 {object} dtos.UserStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/user/{id} [get]
// @Security BearerAuth
func (u *userByAdminUsecase) GetUserByAdminByID(id uint) (dtos.UserInformationResponse, error) {
	var userResponse dtos.UserInformationResponse
	user, err := u.userByAdminRepo.UserGetById(id)
	if err != nil {
		return userResponse, err
	}
	userResponses := dtos.UserInformationResponse{
		ID:             user.ID,
		FullName:       user.FullName,
		Email:          user.Email,
		PhoneNumber:    user.PhoneNumber,
		Gender:         user.Gender,
		BirthDate:      helpers.FormatDateToYMD(user.BirthDate),
		ProfilePicture: user.ProfilePicture,
		Citizen:        user.Citizen,
		Role:           user.Role,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
	return userResponses, nil
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
// @Router       /admin/user [post]
func (u *userByAdminUsecase) CreateUserByAdmin(input dtos.UserRegisterInputByAdmin) (dtos.UserInformationResponse, error) {
	var (
		user         models.User
		userResponse dtos.UserInformationResponse
	)

	input.Email = strings.ToLower(input.Email)

	user, _ = u.userByAdminRepo.UserGetByEmail(input.Email)
	if user.ID > 0 {
		return userResponse, errors.New("email already used")
	}

	if input.Password != input.ConfirmPassword {
		return userResponse, errors.New("password does not match")
	}

	password, err := helpers.HashPassword(input.Password)
	if err != nil {
		return userResponse, err
	}

	if input.Email == "" || input.FullName == "" || input.Password == "" || input.ConfirmPassword == "" || input.Role == "admin" || input.PhoneNumber == "" || input.BirthDate == "" {
		return userResponse, errors.New("failed to create user")
	}

	user.FullName = input.FullName
	user.Email = input.Email
	user.Password = password
	user.PhoneNumber = input.PhoneNumber
	user.ProfilePicture = "https://icon-library.com/images/default-user-icon/default-user-icon-13.jpg"
	user.Citizen = "Indonesia"
	user.Role = "user"

	user, err = u.userByAdminRepo.UserCreate(user)
	if err != nil {
		return userResponse, err
	}

	userResponse.ID = user.ID
	userResponse.FullName = user.FullName
	userResponse.Email = user.Email
	userResponse.PhoneNumber = user.PhoneNumber
	userResponse.Gender = user.Gender
	userResponse.BirthDate = helpers.FormatDateToYMD(user.BirthDate)
	userResponse.ProfilePicture = user.ProfilePicture
	userResponse.Citizen = user.Citizen
	userResponse.Role = user.Role
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return userResponse, err
}

// UpdateUser godoc
// @Summary      Update user
// @Description  Update user
// @Tags         Admin - User
// @Accept       json
// @Produce      json
// @Param id path integer true "ID User"
// @Param        request body dtos.UserRegisterInputByAdmin true "Payload Body [RAW]"
// @Success      200 {object} dtos.UserStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/user [put]
// @Security BearerAuth
func (u *userByAdminUsecase) UpdateUserByAdmin(id uint, input dtos.UserRegisterInputByAdmin) (dtos.UserInformationResponse, error) {
	var user models.User
	var userResponse dtos.UserInformationResponse

	user, err := u.userByAdminRepo.UserGetById(id)
	if err != nil {
		return userResponse, errors.New("user not found")
	}

	birthDateParse, err := helpers.FormatStringToDate(input.BirthDate)
	if err != nil {
		return userResponse, errors.New("failed to parse birth date")
	}

	if input.Password == "" || input.ConfirmPassword == "" {
		return userResponse, errors.New("failed to update user password")
	}

	if input.Password != input.ConfirmPassword {
		return userResponse, errors.New("confirm password does not match")
	}

	newPassword, err := helpers.HashPassword(input.ConfirmPassword)
	if err != nil {
		return userResponse, err
	}

	user.BirthDate = &birthDateParse
	user.Email = input.Email
	user.FullName = input.FullName
	user.PhoneNumber = input.PhoneNumber
	user.Password = newPassword

	user, err = u.userByAdminRepo.UserUpdate(user)

	if err != nil {
		return userResponse, err
	}

	userResponse.ID = user.ID
	userResponse.FullName = user.FullName
	userResponse.Email = user.Email
	userResponse.PhoneNumber = user.PhoneNumber
	userResponse.Role = user.Role
	userResponse.Gender = user.Gender
	userResponse.ProfilePicture = user.ProfilePicture
	userResponse.BirthDate = helpers.FormatDateToYMD(user.BirthDate)
	userResponse.Citizen = user.Citizen
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return userResponse, nil

}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Delete a user
// @Tags         Admin - User
// @Accept       json
// @Produce      json
// @Param id path integer true "ID user"
// @Success      200 {object} dtos.StatusOKDeletedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/user/{id} [delete]
// @Security BearerAuth
func (u *userByAdminUsecase) DeleteUserByAdmin(id uint) error {
	user, err := u.userByAdminRepo.UserGetById(id)

	if err != nil {
		return err
	}

	err = u.userByAdminRepo.UserDelete(user)
	return err
}
