package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"errors"
)

type HistorySeenHotelUsecase interface {
	GetAllHistorySeenHotels(page, limit int, userId uint) ([]dtos.HistorySeenHotelResponse, int, error)
	CreateHistorySeenHotel(userId uint, historySeenHotelInput dtos.HistorySeenHotelInput) (dtos.HistorySeenHotelResponse, error)
}

type historySeenHotelUsecase struct {
	historySeenHotelRepo repositories.HistorySeenHotelRepository
	hotelRepo            repositories.HotelRepository
	hotelImageRepo       repositories.HotelImageRepository
	hotelFacilitiesRepo  repositories.HotelFacilitiesRepository
	hotelPoliciesRepo    repositories.HotelPoliciesRepository
}

func NewHistorySeenHotelUsecase(historySeenHotelRepo repositories.HistorySeenHotelRepository, hotelRepo repositories.HotelRepository, hotelImageRepo repositories.HotelImageRepository, hotelFacilitiesRepo repositories.HotelFacilitiesRepository, hotelPoliciesRepo repositories.HotelPoliciesRepository) HistorySeenHotelUsecase {
	return &historySeenHotelUsecase{historySeenHotelRepo, hotelRepo, hotelImageRepo, hotelFacilitiesRepo, hotelPoliciesRepo}
}

// GetAllHistorySeenHotels godoc
// @Summary      Get all history seen hotel by user id
// @Description  Get all history seen hotel by user id
// @Tags         User - History Seen
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success      200 {object} dtos.GetAllHistorySeenHotelStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/history-seen-hotel [get]
// @Security BearerAuth
func (u *historySeenHotelUsecase) GetAllHistorySeenHotels(page, limit int, userId uint) ([]dtos.HistorySeenHotelResponse, int, error) {
	historySeenHotels, count, err := u.historySeenHotelRepo.GetAllHistorySeenHotel(page, limit, userId)
	if err != nil {
		return nil, 0, err
	}

	var historySeenHotelResponses []dtos.HistorySeenHotelResponse
	for _, historySeenHotel := range historySeenHotels {
		getHotel, err := u.hotelRepo.GetHotelByID(historySeenHotel.HotelID)
		if err != nil {
			continue
		}
		getHotelImage, err := u.hotelImageRepo.GetAllHotelImageByID(getHotel.ID)
		if err != nil {
			continue
		}
		var hotelImageResponses []dtos.HotelImageResponse
		for _, hotel := range getHotelImage {
			hotelImageResponse := dtos.HotelImageResponse{
				HotelID:  hotel.HotelID,
				ImageUrl: hotel.ImageUrl,
			}
			hotelImageResponses = append(hotelImageResponses, hotelImageResponse)
		}
		getHotelFacilities, err := u.hotelFacilitiesRepo.GetAllHotelFacilitiesByID(getHotel.ID)
		if err != nil {
			continue
		}
		var hotelFacilitiesResponses []dtos.HotelFacilitiesResponse
		for _, hotel := range getHotelFacilities {
			hotelFacilitiesResponse := dtos.HotelFacilitiesResponse{
				HotelID: hotel.HotelID,
				Name:    hotel.Name,
			}
			hotelFacilitiesResponses = append(hotelFacilitiesResponses, hotelFacilitiesResponse)
		}
		getHotelPolicies, err := u.hotelPoliciesRepo.GetHotelPoliciesByIDHotel(getHotel.ID)
		if err != nil {
			continue
		}
		historySeenHotelResponse := dtos.HistorySeenHotelResponse{
			ID: historySeenHotel.ID,
			Hotel: dtos.HotelByIDSimply{
				HotelID:         getHotel.ID,
				Name:            getHotel.Name,
				Class:           getHotel.Class,
				Description:     getHotel.Description,
				PhoneNumber:     getHotel.PhoneNumber,
				Email:           getHotel.Email,
				Address:         getHotel.Address,
				HotelImage:      hotelImageResponses,
				HotelFacilities: hotelFacilitiesResponses,
				HotelPolicy: dtos.HotelPoliciesResponse{
					HotelID:            getHotelPolicies.HotelID,
					IsCheckInCheckOut:  getHotelPolicies.IsCheckInCheckOut,
					TimeCheckIn:        getHotelPolicies.TimeCheckIn,
					TimeCheckOut:       getHotelPolicies.TimeCheckOut,
					IsPolicyCanceled:   getHotelPolicies.IsPolicyCanceled,
					PolicyMinimumAge:   getHotelPolicies.PolicyMinimumAge,
					IsPolicyMinimumAge: getHotelPolicies.IsPolicyMinimumAge,
					IsCheckInEarly:     getHotelPolicies.IsCheckInEarly,
					IsCheckOutOverdue:  getHotelPolicies.IsCheckOutOverdue,
					IsBreakfast:        getHotelPolicies.IsBreakfast,
					TimeBreakfastStart: getHotelPolicies.TimeBreakfastStart,
					TimeBreakfastEnd:   getHotelPolicies.TimeBreakfastEnd,
					IsSmoking:          getHotelPolicies.IsSmoking,
					IsPet:              getHotelPolicies.IsPet,
				},
			},
			CreatedAt: historySeenHotel.CreatedAt,
			UpdatedAt: historySeenHotel.UpdatedAt,
		}
		historySeenHotelResponses = append(historySeenHotelResponses, historySeenHotelResponse)
	}

	return historySeenHotelResponses, count, nil
}

func (u *historySeenHotelUsecase) CreateHistorySeenHotel(userId uint, historySeenHotelInput dtos.HistorySeenHotelInput) (dtos.HistorySeenHotelResponse, error) {
	var historySeenHotelResponses dtos.HistorySeenHotelResponse

	if historySeenHotelInput.HotelID < 1 {
		return historySeenHotelResponses, errors.New("Failed to create history seen hotel")
	}

	createHistorySeenHotel := models.HistorySeenHotel{
		UserID:  userId,
		HotelID: historySeenHotelInput.HotelID,
	}

	getHistorySeenHotel, _ := u.historySeenHotelRepo.GetHistorySeenHotelByID(historySeenHotelInput.HotelID, userId)
	if getHistorySeenHotel.ID > 0 {
		createHistorySeenHotel, _ = u.historySeenHotelRepo.UpdateHistorySeenHotel(getHistorySeenHotel)
	} else {
		createHistorySeenHotel, _ = u.historySeenHotelRepo.CreateHistorySeenHotel(createHistorySeenHotel)
	}

	getHotel, _ := u.hotelRepo.GetHotelByID(createHistorySeenHotel.HotelID)
	getHotelImage, _ := u.hotelImageRepo.GetAllHotelImageByID(getHotel.ID)
	var hotelImageResponses []dtos.HotelImageResponse
	for _, hotel := range getHotelImage {
		hotelImageResponse := dtos.HotelImageResponse{
			HotelID:  hotel.HotelID,
			ImageUrl: hotel.ImageUrl,
		}
		hotelImageResponses = append(hotelImageResponses, hotelImageResponse)
	}
	getHotelFacilities, _ := u.hotelFacilitiesRepo.GetAllHotelFacilitiesByID(getHotel.ID)
	var hotelFacilitiesResponses []dtos.HotelFacilitiesResponse
	for _, hotel := range getHotelFacilities {
		hotelFacilitiesResponse := dtos.HotelFacilitiesResponse{
			HotelID: hotel.HotelID,
			Name:    hotel.Name,
		}
		hotelFacilitiesResponses = append(hotelFacilitiesResponses, hotelFacilitiesResponse)
	}
	getHotelPolicies, _ := u.hotelPoliciesRepo.GetHotelPoliciesByIDHotel(getHotel.ID)
	historySeenHotelResponse := dtos.HistorySeenHotelResponse{
		ID: createHistorySeenHotel.ID,
		Hotel: dtos.HotelByIDSimply{
			HotelID:         getHotel.ID,
			Name:            getHotel.Name,
			Class:           getHotel.Class,
			Description:     getHotel.Description,
			PhoneNumber:     getHotel.PhoneNumber,
			Email:           getHotel.Email,
			Address:         getHotel.Address,
			HotelImage:      hotelImageResponses,
			HotelFacilities: hotelFacilitiesResponses,
			HotelPolicy: dtos.HotelPoliciesResponse{
				HotelID:            getHotelPolicies.HotelID,
				IsCheckInCheckOut:  getHotelPolicies.IsCheckInCheckOut,
				TimeCheckIn:        getHotelPolicies.TimeCheckIn,
				TimeCheckOut:       getHotelPolicies.TimeCheckOut,
				IsPolicyCanceled:   getHotelPolicies.IsPolicyCanceled,
				PolicyMinimumAge:   getHotelPolicies.PolicyMinimumAge,
				IsPolicyMinimumAge: getHotelPolicies.IsPolicyMinimumAge,
				IsCheckInEarly:     getHotelPolicies.IsCheckInEarly,
				IsCheckOutOverdue:  getHotelPolicies.IsCheckOutOverdue,
				IsBreakfast:        getHotelPolicies.IsBreakfast,
				TimeBreakfastStart: getHotelPolicies.TimeBreakfastStart,
				TimeBreakfastEnd:   getHotelPolicies.TimeBreakfastEnd,
				IsSmoking:          getHotelPolicies.IsSmoking,
				IsPet:              getHotelPolicies.IsPet,
			},
		},
		CreatedAt: createHistorySeenHotel.CreatedAt,
		UpdatedAt: createHistorySeenHotel.UpdatedAt,
	}
	return historySeenHotelResponse, nil
}
