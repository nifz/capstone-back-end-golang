package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"errors"
)

type HotelRatingsUsecase interface {
	// user
	CreateHotelRating(hotelRatingInput dtos.HotelRatingInput) (dtos.HotelRatingResponse, error)
	GetHotelRatingsByIdOrders(id uint) (dtos.HotelRatingResponse, error)
	GetAllHotelRatingsByIdHotels(page, limit int, hotel_id uint) ([]dtos.RatingInfo, int, error)
	// admin
	GetHotelRatingsByHotelID(page, limit int, id uint, filter string) (dtos.HotelRatingsByIdHotels, int, error)
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
func (u *hotelRatingsUsecase) CreateHotelRating(hotelRatingInput dtos.HotelRatingInput) (dtos.HotelRatingResponse, error) {
	var hotelRatingResponse dtos.HotelRatingResponse

	_, err := u.hotelRatingsRepository.CheckExistHotelRating(hotelRatingInput.HotelOrderID, hotelRatingInput.UserID)
	if err == nil {
		return hotelRatingResponse, errors.New("You have already rated this hotel")
	}

	_, err = u.hotelOrderRepository.GetHotelOrderByID(hotelRatingInput.HotelOrderID, hotelRatingInput.UserID)
	if err != nil {
		return hotelRatingResponse, errors.New("Hotel Order ID is not valid")
	}

	if hotelRatingInput.Rating < 0 || hotelRatingInput.Rating > 5 {
		return hotelRatingResponse, errors.New("Rating must be between 0 and 5")
	}

	_, err = u.hotelRepository.GetHotelByID(hotelRatingInput.HotelID)
	if err != nil {
		return hotelRatingResponse, errors.New("Hotel ID is not valid")
	}

	_, err = u.userRepository.UserGetById2(hotelRatingInput.UserID)
	if err != nil {
		return hotelRatingResponse, errors.New("User ID is not valid")
	}

	if len(hotelRatingInput.Review) < 10 {
		return hotelRatingResponse, errors.New("Review must be at least 10 characters long")
	}

	hotelRating := models.HotelRating{
		HotelOrderID: hotelRatingInput.HotelOrderID,
		HotelID:      hotelRatingInput.HotelID,
		UserID:       hotelRatingInput.UserID,
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

func (u *hotelRatingsUsecase) GetHotelRatingsByHotelID(page, limit int, id uint, filter string) (dtos.HotelRatingsByIdHotels, int, error) {
	var (
		hotelRatingsResponse dtos.HotelRatingsByIdHotels
	)

	ratingCounts, hotelRatings, count, err := u.hotelRatingsRepository.GetHotelRatingsByHotelID(page, limit, id, filter)
	if err != nil {
		return hotelRatingsResponse, 0, errors.New("Hotel ID is not found")
	}

	hotelRatingsResponse.HotelID = id
	hotelRatingsResponse.TotalRating = len(hotelRatings)
	// hotelRatingsResponse.RatingCounts = ratingCounts

	var totalRating int
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
		hotelRatingsResponse.Ratings = append(hotelRatingsResponse.Ratings, ratingInfo)
		totalRating += rating.Rating
	}

	if len(hotelRatings) > 0 {
		hotelRatingsResponse.RataRataRating = float64(totalRating) / float64(len(hotelRatings))
	} else {
		hotelRatingsResponse.RataRataRating = 0
	}
	// Set individual rating counts
	hotelRatingsResponse.Rating5 = ratingCounts[5]
	hotelRatingsResponse.Rating4 = ratingCounts[4]
	hotelRatingsResponse.Rating3 = ratingCounts[3]
	hotelRatingsResponse.Rating2 = ratingCounts[2]
	hotelRatingsResponse.Rating1 = ratingCounts[1]

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

func (u *hotelRatingsUsecase) GetAllHotelRatingsByIdHotels(page, limit int, hotel_id uint) ([]dtos.RatingInfo, int, error) {
	var hotelRatingsResponse []dtos.RatingInfo

	hotelRatings, count, err := u.hotelRatingsRepository.GetAllHotelRatingsByIdHotels(page, limit, hotel_id)
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
