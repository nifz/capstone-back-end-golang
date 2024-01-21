package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"errors"
)

type HotelRatingsUsecase interface {
	// user
	CreateHotelRating(userId uint, hotelRatingInput dtos.HotelRatingInput) (dtos.HotelRatingResponse, error)
	GetHotelRatingsByIdOrders(id uint) (dtos.HotelRatingResponse, error)
	GetAllHotelRatingsByIdHotels(page, limit int, hotelId uint) ([]dtos.RatingInfo, int, error)
	// admin
	GetHotelRatingsByHotelID(star, page, limit int, id uint, filter string) (dtos.HotelRatingsByIdHotels, int, error)
}

type hotelRatingsUsecase struct {
	hotelRatingsRepository repositories.HotelRatingsRepository
	hotelRepository        repositories.HotelRepository
	userRepository         repositories.UserRepository
	hotelOrderRepository   repositories.HotelOrderRepository
	notificationRepository repositories.NotificationRepository
}

func NewHotelRatingsUsecase(hotelRatingsRepository repositories.HotelRatingsRepository, hotelRepository repositories.HotelRepository, userRepository repositories.UserRepository, hotelOrderRepository repositories.HotelOrderRepository, notificationRepository repositories.NotificationRepository) HotelRatingsUsecase {
	return &hotelRatingsUsecase{
		hotelRatingsRepository,
		hotelRepository,
		userRepository,
		hotelOrderRepository,
		notificationRepository,
	}
}

// Implementasi fungsi-fungsi dari interface ItemUsecase

// CreateHotelRating godoc
// @Summary      Create a new hotel rating
// @Description  Create a new hotel rating
// @Tags         User - Hotel Rating
// @Accept       json
// @Produce      json
// @Param        request body dtos.HotelRatingInput true "Payload Body [RAW]"
// @Success      201 {object} dtos.HotelRatingCreeatedResponses
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/hotel-ratings [post]
// @Security BearerAuth
func (u *hotelRatingsUsecase) CreateHotelRating(userId uint, hotelRatingInput dtos.HotelRatingInput) (dtos.HotelRatingResponse, error) {
	var hotelRatingResponse dtos.HotelRatingResponse

	_, err := u.hotelRatingsRepository.CheckExistHotelRating(hotelRatingInput.HotelOrderID, userId)
	if err == nil {
		return hotelRatingResponse, errors.New("You have already rated this hotel")
	}

	hotelOrder, err := u.hotelOrderRepository.GetHotelOrderByID(hotelRatingInput.HotelOrderID, userId)
	if err != nil {
		return hotelRatingResponse, errors.New("Hotel Order ID is not valid")
	}

	if hotelRatingInput.Rating < 0 || hotelRatingInput.Rating > 5 {
		return hotelRatingResponse, errors.New("Rating must be between 0 and 5")
	}

	if len(hotelRatingInput.Review) < 10 {
		return hotelRatingResponse, errors.New("Review must be at least 10 characters long")
	}

	hotelRating := models.HotelRating{
		HotelOrderID: hotelOrder.ID,
		HotelID:      hotelOrder.HotelID,
		UserID:       userId,
		Rating:       hotelRatingInput.Rating,
		Review:       hotelRatingInput.Review,
	}

	// Save hotel rating
	createdRating, err := u.hotelRatingsRepository.CreateHotelRating(hotelRating)
	if err != nil {
		return hotelRatingResponse, err
	}

	if createdRating.ID > 0 {
		createNotification := models.Notification{
			UserID:     createdRating.UserID,
			TemplateID: 5,
		}

		_, err = u.notificationRepository.CreateNotification(createNotification)
		if err != nil {
			return hotelRatingResponse, err
		}
	}

	// Fill the response with the created rating
	hotelRatingResponse.HotelOrderID = createdRating.HotelOrderID
	hotelRatingResponse.HotelID = createdRating.HotelID
	hotelRatingResponse.UserID = createdRating.UserID
	hotelRatingResponse.Rating = createdRating.Rating
	hotelRatingResponse.Review = createdRating.Review

	return hotelRatingResponse, nil
}

// GetHotelByID godoc
// @Summary      Get hotel by ID
// @Description  Get hotel by ID
// @Tags         Admin - Hotel
// @Accept       json
// @Produce      json
// @Param id path integer true "ID Hotel"
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param filter query string false "Filter order by review hotel from user" Enums(latest, oldest)
// @Param rating query int false "Filter rating hotel by user" Enums(1,2,3,4,5)
// @Success      200 {object} dtos.GetAllRatingByIdHotelStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /public/hotel/{id}/rating [get]
func (u *hotelRatingsUsecase) GetHotelRatingsByHotelID(star, page, limit int, id uint, filter string) (dtos.HotelRatingsByIdHotels, int, error) {
	var (
		hotelRatingsResponse dtos.HotelRatingsByIdHotels
	)

	ratingCounts, hotelRatings, count, err := u.hotelRatingsRepository.GetHotelRatingsByHotelID(id, filter)
	if err != nil {
		return hotelRatingsResponse, 0, errors.New("Hotel ID is not found")
	}

	hotelRatingsResponse.HotelID = id
	hotelRatingsResponse.TotalRating = len(hotelRatings)
	// hotelRatingsResponse.RatingCounts = ratingCounts

	sumRatings := 0
	for _, rating := range hotelRatings {
		sumRatings += rating.Rating
	}

	if len(hotelRatings) > 0 {
		hotelRatingsResponse.RataRataRating = float64(sumRatings) / float64(len(hotelRatings))
	} else {
		hotelRatingsResponse.RataRataRating = 0
	}

	// Set individual rating counts
	hotelRatingsResponse.Rating5 = ratingCounts[5]
	hotelRatingsResponse.Rating4 = ratingCounts[4]
	hotelRatingsResponse.Rating3 = ratingCounts[3]
	hotelRatingsResponse.Rating2 = ratingCounts[2]
	hotelRatingsResponse.Rating1 = ratingCounts[1]

	for _, rating := range hotelRatings {
		userDetail, err := u.userRepository.UserGetById2(rating.UserID)
		if err != nil {
			return hotelRatingsResponse, 0, errors.New("User ID is not valid")
		}
		if star == 0 && star != rating.Rating {
			continue
		}

		ratingInfo := dtos.RatingInfo{
			UserID:    rating.UserID,
			Username:  userDetail.FullName,
			UserImage: userDetail.ProfilePicture,
			Rating:    rating.Rating,
			Review:    rating.Review,
			CreatedAt: rating.CreatedAt,
		}
		hotelRatingsResponse.Ratings = append(hotelRatingsResponse.Ratings, ratingInfo)
	}

	// Apply offset and limit to trainResponses
	start := (page - 1) * limit
	end := start + limit

	// Ensure that `start` is within the range of trainResponses
	if start >= len(hotelRatingsResponse.Ratings) {
		return hotelRatingsResponse, 0, nil
	}

	// Ensure that `end` does not exceed the length of trainResponses
	if end > len(hotelRatingsResponse.Ratings) {
		end = len(hotelRatingsResponse.Ratings)
	}

	subsetHotelRatingResponses := hotelRatingsResponse.Ratings[start:end]
	hotelRatingsResponse.Ratings = subsetHotelRatingResponses
	return hotelRatingsResponse, count, nil
}

func (u *hotelRatingsUsecase) GetHotelRatingsByIdOrders(id uint) (dtos.HotelRatingResponse, error) {
	var hotelRatingsResponse dtos.HotelRatingResponse

	hotelRatings, err := u.hotelRatingsRepository.GetHotelRatingsByIdOrders(id)
	if err != nil {
		return hotelRatingsResponse, errors.New("Hotel Order ID is not found")
	}

	hotelRatingsResponse.HotelOrderID = hotelRatings.HotelOrderID
	hotelRatingsResponse.HotelID = hotelRatings.HotelID
	hotelRatingsResponse.UserID = hotelRatings.UserID
	hotelRatingsResponse.Rating = hotelRatings.Rating
	hotelRatingsResponse.Review = hotelRatings.Review

	return hotelRatingsResponse, nil
}

func (u *hotelRatingsUsecase) GetAllHotelRatingsByIdHotels(page, limit int, hotelId uint) ([]dtos.RatingInfo, int, error) {
	var hotelRatingsResponse []dtos.RatingInfo

	hotelRatings, count, err := u.hotelRatingsRepository.GetAllHotelRatingsByIdHotels(page, limit, hotelId)
	if err != nil {
		return hotelRatingsResponse, 0, errors.New("Hotel ID is not found")
	}

	for _, rating := range hotelRatings {
		userDetail, err := u.userRepository.UserGetById2(rating.UserID)
		if err != nil {
			return hotelRatingsResponse, 0, errors.New("User ID is not valid")
		}

		ratingInfo := dtos.RatingInfo{
			UserID:    rating.UserID,
			Username:  userDetail.FullName,
			UserImage: userDetail.ProfilePicture,
			Rating:    rating.Rating,
			Review:    rating.Review,
			CreatedAt: rating.CreatedAt,
		}
		hotelRatingsResponse = append(hotelRatingsResponse, ratingInfo)
	}

	return hotelRatingsResponse, count, nil

}
