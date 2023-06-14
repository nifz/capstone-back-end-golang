package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"errors"
)

type HotelRatingsUsecase interface {
	CreateHotelRating(hotelRatingInput dtos.HotelRatingInput) (dtos.HotelRatingResponse, error)
	GetHotelRatingsByHotelID(id uint) (dtos.HotelRatingsByIdHotels, error)
}

type hotelRatingsUsecase struct {
	hotelRatingsRepository repositories.HotelRatingsRepository
	hotelRepository        repositories.HotelRepository
	userRepository         repositories.UserRepository
}

func NewHotelRatingsUsecase(hotelRatingsRepository repositories.HotelRatingsRepository, hotelRepository repositories.HotelRepository, userRepository repositories.UserRepository) HotelRatingsUsecase {
	return &hotelRatingsUsecase{
		hotelRatingsRepository,
		hotelRepository,
		userRepository,
	}
}

// Implementasi fungsi-fungsi dari interface ItemUsecase


func (u *hotelRatingsUsecase) CreateHotelRating(hotelRatingInput dtos.HotelRatingInput) (dtos.HotelRatingResponse, error) {
	var hotelRatingResponse dtos.HotelRatingResponse

	if hotelRatingInput.Rating < 0 || hotelRatingInput.Rating > 5 {
		return hotelRatingResponse, errors.New("Rating must be between 0 and 5")
	}

	_, err := u.hotelRepository.GetHotelByID(hotelRatingInput.HotelID)
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
		HotelID: hotelRatingInput.HotelID,
		UserID:  hotelRatingInput.UserID,
		Rating:  hotelRatingInput.Rating,
		Review:  hotelRatingInput.Review,
	}

	// Save hotel rating
	createdRating, err := u.hotelRatingsRepository.CreateHotelRating(hotelRating)
	if err != nil {
		return hotelRatingResponse, err
	}

	// Fill the response with the created rating
	hotelRatingResponse.HotelID = createdRating.HotelID
	hotelRatingResponse.UserID = createdRating.UserID
	hotelRatingResponse.Rating = createdRating.Rating
	hotelRatingResponse.Review = createdRating.Review

	return hotelRatingResponse, nil
}

func (u *hotelRatingsUsecase) GetHotelRatingsByHotelID(id uint) (dtos.HotelRatingsByIdHotels, error) {
	var hotelRatingsResponse dtos.HotelRatingsByIdHotels

	ratingCounts, hotelRatings, err := u.hotelRatingsRepository.GetHotelRatingsByHotelID(id)
	if err != nil {
		return hotelRatingsResponse, errors.New("Hotel ID is not found")
	}

	hotelRatingsResponse.HotelID = id
	hotelRatingsResponse.TotalRating = len(hotelRatings)
	// hotelRatingsResponse.RatingCounts = ratingCounts

	var totalRating int
	for _, rating := range hotelRatings {
		userDetail, err := u.userRepository.UserGetById2(rating.UserID)
		if err != nil {
			return hotelRatingsResponse, errors.New("User ID is not valid")
		}

		ratingInfo := dtos.RatingInfo{
			UserID:    rating.UserID,
			Username:  userDetail.FullName,
			UserImage: userDetail.ProfilePicture,
			Rating:    rating.Rating,
			Review:    rating.Review,
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

	return hotelRatingsResponse, nil
}
