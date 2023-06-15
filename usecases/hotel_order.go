package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"errors"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

type HotelOrderUsecase interface {
	GetHotelOrders(page, limit int, userID uint, status string) ([]dtos.HotelOrderResponse, int, error)
	GetHotelOrdersByAdmin(page, limit, ratingClass int, search, dateStart, dateEnd, orderBy, status string) ([]dtos.HotelOrderResponse, int, error)
	GetHotelOrdersDetailByAdmin(hotelOrderId uint) (dtos.HotelOrderResponse, error)
	GetHotelOrderByID(userID, hotelOrderId uint) (dtos.HotelOrderResponse, error)
	CreateHotelOrder(userID uint, hotelOrderInput dtos.HotelOrderInput) (dtos.HotelOrderResponse, error)
	UpdateHotelOrder(userID, hotelOrderID uint, status string) (dtos.HotelOrderResponse, error)
}

type hotelOrderUsecase struct {
	hotelOrderRepo          repositories.HotelOrderRepository
	hotelRepo               repositories.HotelRepository
	hotelImageRepo          repositories.HotelImageRepository
	hotelFacilitiesRepo     repositories.HotelFacilitiesRepository
	hotelPoliciesRepo       repositories.HotelPoliciesRepository
	hotelRoomRepo           repositories.HotelRoomRepository
	hotelRoomImageRepo      repositories.HotelRoomImageRepository
	hotelRoomFacilitiesRepo repositories.HotelRoomFacilitiesRepository
	travelerDetailRepo      repositories.TravelerDetailRepository
	paymentRepo             repositories.PaymentRepository
	userRepo                repositories.UserRepository
	notificationRepo        repositories.NotificationRepository
}

func NewHotelOrderUsecase(hotelOrderRepo repositories.HotelOrderRepository, hotelRepo repositories.HotelRepository, hotelImageRepo repositories.HotelImageRepository, hotelFacilitiesRepo repositories.HotelFacilitiesRepository, hotelPoliciesRepo repositories.HotelPoliciesRepository, hotelRoomRepo repositories.HotelRoomRepository, hotelRoomImageRepo repositories.HotelRoomImageRepository, hotelRoomFacilitiesRepo repositories.HotelRoomFacilitiesRepository, travelerDetailRepo repositories.TravelerDetailRepository, paymentRepo repositories.PaymentRepository, userRepo repositories.UserRepository, notificationRepo repositories.NotificationRepository) HotelOrderUsecase {
	return &hotelOrderUsecase{hotelOrderRepo, hotelRepo, hotelImageRepo, hotelFacilitiesRepo, hotelPoliciesRepo, hotelRoomRepo, hotelRoomImageRepo, hotelRoomFacilitiesRepo, travelerDetailRepo, paymentRepo, userRepo, notificationRepo}
}

// GetHotelOrders godoc
// @Summary      Get Hotel Order User
// @Description  Get Hotel Order User
// @Tags         User - Order
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param status query string false "Filter by status order"
// @Success      200 {object} dtos.GetAllHotelOrderStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/order/hotel [get]
// @Security BearerAuth
func (u *hotelOrderUsecase) GetHotelOrders(page, limit int, userID uint, status string) ([]dtos.HotelOrderResponse, int, error) {
	var hotelOrderResponses []dtos.HotelOrderResponse

	hotelOrders, _, err := u.hotelOrderRepo.GetHotelOrders(page, limit, userID, status)
	if err != nil {
		return hotelOrderResponses, 0, err
	}

	for _, hotelOrder := range hotelOrders {
		getHotel, err := u.hotelRepo.GetHotelByID(hotelOrder.HotelID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		getHotelImage, err := u.hotelImageRepo.GetAllHotelImageByID(hotelOrder.HotelID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		var hotelImageResponses []dtos.HotelImageResponse
		for _, hotelImage := range getHotelImage {
			hotelImageResponse := dtos.HotelImageResponse{
				HotelID:  hotelImage.HotelID,
				ImageUrl: hotelImage.ImageUrl,
			}
			hotelImageResponses = append(hotelImageResponses, hotelImageResponse)
		}
		getHotelFacilities, err := u.hotelFacilitiesRepo.GetAllHotelFacilitiesByID(hotelOrder.HotelID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		getHotelPolicies, err := u.hotelPoliciesRepo.GetHotelPoliciesByIDHotel(hotelOrder.HotelID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		var hotelFacilitiesResponses []dtos.HotelFacilitiesResponse
		for _, hotelFacilities := range getHotelFacilities {
			hotelFacilitiesResponse := dtos.HotelFacilitiesResponse{
				HotelID: hotelFacilities.HotelID,
				Name:    hotelFacilities.Name,
			}
			hotelFacilitiesResponses = append(hotelFacilitiesResponses, hotelFacilitiesResponse)
		}
		getHotelRoom, err := u.hotelRoomRepo.GetHotelRoomByID(hotelOrder.HotelRoomID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		getHotelRoomImage, err := u.hotelRoomImageRepo.GetAllHotelRoomImageByID(getHotelRoom.ID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		var hotelRoomImageResponses []dtos.HotelRoomImageResponse
		for _, hotelRoomImage := range getHotelRoomImage {
			hotelRoomImageResponse := dtos.HotelRoomImageResponse{
				HotelID:     hotelRoomImage.HotelID,
				HotelRoomID: hotelRoomImage.ID,
				ImageUrl:    hotelRoomImage.ImageUrl,
			}
			hotelRoomImageResponses = append(hotelRoomImageResponses, hotelRoomImageResponse)
		}
		getHotelRoomFacilities, err := u.hotelRoomFacilitiesRepo.GetAllHotelRoomFacilitiesByHotelRoomID(getHotelRoom.ID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		var hotelRoomFacilitiesResponses []dtos.HotelRoomFacilitiesResponse
		for _, hotelRoomFacilities := range getHotelRoomFacilities {
			hotelRoomFacilitiesResponse := dtos.HotelRoomFacilitiesResponse{
				HotelID:     hotelRoomFacilities.HotelID,
				HotelRoomID: hotelRoomFacilities.ID,
				Name:        hotelRoomFacilities.Name,
			}
			hotelRoomFacilitiesResponses = append(hotelRoomFacilitiesResponses, hotelRoomFacilitiesResponse)
		}
		getPayment, err := u.paymentRepo.GetPaymentByID(uint(hotelOrder.PaymentID))
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		getTravelerDetail, err := u.travelerDetailRepo.GetTravelerDetailByHotelOrderID(hotelOrder.ID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}

		var travelerDetailResponses []dtos.TravelerDetailResponse

		for _, travelerDetail := range getTravelerDetail {
			travelerDetailResponse := dtos.TravelerDetailResponse{
				ID:           int(travelerDetail.ID),
				Title:        travelerDetail.Title,
				FullName:     travelerDetail.FullName,
				IDCardNumber: *travelerDetail.IDCardNumber,
			}
			travelerDetailResponses = append(travelerDetailResponses, travelerDetailResponse)
		}

		hotelOrderResponse := dtos.HotelOrderResponse{
			HotelOrderID:     int(hotelOrder.ID),
			QuantityAdult:    hotelOrder.QuantityAdult,
			QuantityInfant:   hotelOrder.QuantityInfant,
			NumberOfNight:    hotelOrder.NumberOfNight,
			DateStart:        helpers.FormatDateToYMD(&hotelOrder.DateStart),
			DateEnd:          helpers.FormatDateToYMD(&hotelOrder.DateEnd),
			Price:            hotelOrder.Price,
			TotalAmount:      hotelOrder.TotalAmount,
			NameOrder:        hotelOrder.NameOrder,
			EmailOrder:       hotelOrder.EmailOrder,
			PhoneNumberOrder: hotelOrder.PhoneNumberOrder,
			SpecialRequest:   hotelOrder.SpecialRequest,
			HotelOrderCode:   hotelOrder.HotelOrderCode,
			Status:           hotelOrder.Status,
			Hotel: dtos.HotelByIDResponses{
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
				HotelRoom: dtos.HotelRoomHotelIDResponse{
					HotelRoomID:       getHotelRoom.ID,
					HotelID:           getHotelRoom.HotelID,
					Name:              getHotelRoom.Name,
					SizeOfRoom:        getHotelRoom.SizeOfRoom,
					QuantityOfRoom:    getHotelRoom.QuantityOfRoom,
					Description:       getHotelRoom.Description,
					NormalPrice:       getHotelRoom.NormalPrice,
					Discount:          getHotelRoom.Discount,
					DiscountPrice:     getHotelRoom.DiscountPrice,
					NumberOfGuest:     getHotelRoom.NumberOfGuest,
					MattressSize:      getHotelRoom.MattressSize,
					NumberOfMattress:  getHotelRoom.NumberOfMattress,
					HotelRoomImage:    hotelRoomImageResponses,
					HotelRoomFacility: hotelRoomFacilitiesResponses,
				},
			},
			Payment: dtos.PaymentResponses{
				ID:            int(getPayment.ID),
				Type:          getPayment.Type,
				ImageUrl:      getPayment.ImageUrl,
				Name:          getPayment.Name,
				AccountName:   getPayment.AccountName,
				AccountNumber: getPayment.AccountNumber,
			},
			TravelerDetail: travelerDetailResponses,
			CreatedAt:      hotelOrder.CreatedAt,
			UpdatedAt:      hotelOrder.UpdatedAt,
		}
		hotelOrderResponses = append(hotelOrderResponses, hotelOrderResponse)
	}

	// Apply offset and limit to hotelOrderResponses
	// start := (page - 1) * limit
	// end := start + limit

	// Ensure that `start` is within the range of hotelOrderResponses
	// if start >= len(hotelOrderResponses) {
	// 	return nil, 0, nil
	// }

	// Ensure that `end` does not exceed the length of hotelOrderResponses
	// if end > len(hotelOrderResponses) {
	// 	end = len(hotelOrderResponses)
	// }

	// subsetHotelOrderResponses := hotelOrderResponses[start:end]

	return hotelOrderResponses, len(hotelOrderResponses), nil
}

// GetHotelOrdersByAdmin godoc
// @Summary      Get Hotel Order User
// @Description  Get Hotel Order User
// @Tags         Admin - Order
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param rating_class query int false "Hotel rating class"
// @Param search query string false "search hotel name"
// @Param date_start query string false "Date start"
// @Param date_end query string false "Date end"
// @Param order_by query string false "Order by name"
// @Param status query string false "Filter by status order"
// @Success      200 {object} dtos.GetAllHotelOrderStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/order/hotel [get]
// @Security BearerAuth
func (u *hotelOrderUsecase) GetHotelOrdersByAdmin(page, limit, ratingClass int, search, dateStart, dateEnd, orderBy, status string) ([]dtos.HotelOrderResponse, int, error) {
	var hotelOrderResponses []dtos.HotelOrderResponse

	hotelOrders, _, err := u.hotelOrderRepo.GetHotelOrders(page, limit, 1, status)
	if err != nil {
		return hotelOrderResponses, 0, err
	}

	// Parse dateStart and dateEnd strings into time.Time objects
	var startDate, endDate time.Time
	if dateStart != "" {
		startDate, err = time.Parse("2006-01-02", dateStart)
		if err != nil {
			return hotelOrderResponses, 0, errors.New("invalid dateStart format")
		}
		startDate = startDate.AddDate(0, 0, -1) // Subtract 1 day from startDate
	}
	if dateEnd != "" {
		endDate, err = time.Parse("2006-01-02", dateEnd)
		if err != nil {
			return hotelOrderResponses, 0, errors.New("invalid dateEnd format")
		}
	}

	for _, hotelOrder := range hotelOrders {
		// Filter hotel orders based on dateStart and dateEnd
		if !startDate.IsZero() && hotelOrder.DateStart.Before(startDate) {
			continue // Skip hotel order if its dateStart is before the specified startDate
		}
		if !endDate.IsZero() && hotelOrder.DateStart.After(endDate) {
			continue // Skip hotel order if its dateEnd is after the specified endDate
		}

		getHotel, err := u.hotelRepo.GetHotelByID(hotelOrder.HotelID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}

		// Check if the search query matches the hotel name, address, or traveler detail name
		if search != "" &&
			!strings.Contains(strings.ToLower(getHotel.Name), strings.ToLower(search)) &&
			!strings.Contains(strings.ToLower(getHotel.Address), strings.ToLower(search)) &&
			!hasMatchingTravelerDetail(hotelOrder.ID, search, u.travelerDetailRepo) {
			continue // Skip hotel order if it doesn't match the search query
		}

		getHotelImage, err := u.hotelImageRepo.GetAllHotelImageByID(hotelOrder.HotelID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		var hotelImageResponses []dtos.HotelImageResponse
		for _, hotelImage := range getHotelImage {
			hotelImageResponse := dtos.HotelImageResponse{
				HotelID:  hotelImage.HotelID,
				ImageUrl: hotelImage.ImageUrl,
			}
			hotelImageResponses = append(hotelImageResponses, hotelImageResponse)
		}
		getHotelFacilities, err := u.hotelFacilitiesRepo.GetAllHotelFacilitiesByID(hotelOrder.HotelID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		getHotelPolicies, err := u.hotelPoliciesRepo.GetHotelPoliciesByIDHotel(hotelOrder.HotelID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		var hotelFacilitiesResponses []dtos.HotelFacilitiesResponse
		for _, hotelFacilities := range getHotelFacilities {
			hotelFacilitiesResponse := dtos.HotelFacilitiesResponse{
				HotelID: hotelFacilities.HotelID,
				Name:    hotelFacilities.Name,
			}
			hotelFacilitiesResponses = append(hotelFacilitiesResponses, hotelFacilitiesResponse)
		}
		getHotelRoom, err := u.hotelRoomRepo.GetHotelRoomByID(hotelOrder.HotelRoomID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		getHotelRoomImage, err := u.hotelRoomImageRepo.GetAllHotelRoomImageByID(getHotelRoom.ID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		var hotelRoomImageResponses []dtos.HotelRoomImageResponse
		for _, hotelRoomImage := range getHotelRoomImage {
			hotelRoomImageResponse := dtos.HotelRoomImageResponse{
				HotelID:     hotelRoomImage.HotelID,
				HotelRoomID: hotelRoomImage.ID,
				ImageUrl:    hotelRoomImage.ImageUrl,
			}
			hotelRoomImageResponses = append(hotelRoomImageResponses, hotelRoomImageResponse)
		}
		getHotelRoomFacilities, err := u.hotelRoomFacilitiesRepo.GetAllHotelRoomFacilitiesByHotelRoomID(getHotelRoom.ID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		var hotelRoomFacilitiesResponses []dtos.HotelRoomFacilitiesResponse
		for _, hotelRoomFacilities := range getHotelRoomFacilities {
			hotelRoomFacilitiesResponse := dtos.HotelRoomFacilitiesResponse{
				HotelID:     hotelRoomFacilities.HotelID,
				HotelRoomID: hotelRoomFacilities.ID,
				Name:        hotelRoomFacilities.Name,
			}
			hotelRoomFacilitiesResponses = append(hotelRoomFacilitiesResponses, hotelRoomFacilitiesResponse)
		}
		getPayment, err := u.paymentRepo.GetPaymentByID(uint(hotelOrder.PaymentID))
		if err != nil {
			return hotelOrderResponses, 0, err
		}
		getTravelerDetail, err := u.travelerDetailRepo.GetTravelerDetailByHotelOrderID(hotelOrder.ID)
		if err != nil {
			return hotelOrderResponses, 0, err
		}

		var travelerDetailResponses []dtos.TravelerDetailResponse

		for _, travelerDetail := range getTravelerDetail {
			travelerDetailResponse := dtos.TravelerDetailResponse{
				ID:           int(travelerDetail.ID),
				Title:        travelerDetail.Title,
				FullName:     travelerDetail.FullName,
				IDCardNumber: *travelerDetail.IDCardNumber,
			}
			travelerDetailResponses = append(travelerDetailResponses, travelerDetailResponse)
		}

		hotelOrderResponse := dtos.HotelOrderResponse{
			HotelOrderID:     int(hotelOrder.ID),
			QuantityAdult:    hotelOrder.QuantityAdult,
			QuantityInfant:   hotelOrder.QuantityInfant,
			NumberOfNight:    hotelOrder.NumberOfNight,
			DateStart:        helpers.FormatDateToYMD(&hotelOrder.DateStart),
			DateEnd:          helpers.FormatDateToYMD(&hotelOrder.DateEnd),
			Price:            hotelOrder.Price,
			TotalAmount:      hotelOrder.TotalAmount,
			NameOrder:        hotelOrder.NameOrder,
			EmailOrder:       hotelOrder.EmailOrder,
			PhoneNumberOrder: hotelOrder.PhoneNumberOrder,
			SpecialRequest:   hotelOrder.SpecialRequest,
			HotelOrderCode:   hotelOrder.HotelOrderCode,
			Status:           hotelOrder.Status,
			Hotel: dtos.HotelByIDResponses{
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
				HotelRoom: dtos.HotelRoomHotelIDResponse{
					HotelRoomID:       getHotelRoom.ID,
					HotelID:           getHotelRoom.HotelID,
					Name:              getHotelRoom.Name,
					SizeOfRoom:        getHotelRoom.SizeOfRoom,
					QuantityOfRoom:    getHotelRoom.QuantityOfRoom,
					Description:       getHotelRoom.Description,
					NormalPrice:       getHotelRoom.NormalPrice,
					Discount:          getHotelRoom.Discount,
					DiscountPrice:     getHotelRoom.DiscountPrice,
					NumberOfGuest:     getHotelRoom.NumberOfGuest,
					MattressSize:      getHotelRoom.MattressSize,
					NumberOfMattress:  getHotelRoom.NumberOfMattress,
					HotelRoomImage:    hotelRoomImageResponses,
					HotelRoomFacility: hotelRoomFacilitiesResponses,
				},
			},
			Payment: dtos.PaymentResponses{
				ID:            int(getPayment.ID),
				Type:          getPayment.Type,
				ImageUrl:      getPayment.ImageUrl,
				Name:          getPayment.Name,
				AccountName:   getPayment.AccountName,
				AccountNumber: getPayment.AccountNumber,
			},
			TravelerDetail: travelerDetailResponses,
			CreatedAt:      hotelOrder.CreatedAt,
			UpdatedAt:      hotelOrder.UpdatedAt,
		}

		if ratingClass > 0 && getHotel.Class < ratingClass {
			continue // Skip the hotel if its rating class is below the specified ratingClass
		}
		hotelOrderResponses = append(hotelOrderResponses, hotelOrderResponse)
	}

	// Sort hotelOrderResponses based on the orderBy parameter
	switch orderBy {
	case "latest":
		// Sort hotelOrderResponses by descending order of CreatedAt
		sort.SliceStable(hotelOrderResponses, func(i, j int) bool {
			return hotelOrderResponses[i].CreatedAt.After(hotelOrderResponses[j].CreatedAt)
		})
	case "oldest":
		// Sort hotelOrderResponses by ascending order of CreatedAt
		sort.SliceStable(hotelOrderResponses, func(i, j int) bool {
			return hotelOrderResponses[i].CreatedAt.Before(hotelOrderResponses[j].CreatedAt)
		})
	case "highest_price":
		// Sort hotelOrderResponses by descending order of Price
		sort.SliceStable(hotelOrderResponses, func(i, j int) bool {
			return hotelOrderResponses[i].Price > hotelOrderResponses[j].Price
		})
	case "lowest_price":
		// Sort hotelOrderResponses by ascending order of Price
		sort.SliceStable(hotelOrderResponses, func(i, j int) bool {
			return hotelOrderResponses[i].Price < hotelOrderResponses[j].Price
		})
	}
	return hotelOrderResponses, len(hotelOrderResponses), nil
}

// GetHotelOrdersByAdmin godoc
// @Summary      Get Hotel Order User
// @Description  Get Hotel Order User
// @Tags         Admin - Order
// @Accept       json
// @Produce      json
// @Param hotel_order_id query int true "Hotel Order ID"
// @Success      200 {object} dtos.HotelOrderStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/order/hotel/detail [get]
// @Security BearerAuth
func (u *hotelOrderUsecase) GetHotelOrdersDetailByAdmin(hotelOrderId uint) (dtos.HotelOrderResponse, error) {
	var hotelOrderResponses dtos.HotelOrderResponse

	hotelOrder, err := u.hotelOrderRepo.GetHotelOrderByID(hotelOrderId, 1)
	if err != nil {
		return hotelOrderResponses, err
	}

	getHotel, err := u.hotelRepo.GetHotelByID(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponses, err
	}
	getHotelImage, err := u.hotelImageRepo.GetAllHotelImageByID(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponses, err
	}
	var hotelImageResponses []dtos.HotelImageResponse
	for _, hotelImage := range getHotelImage {
		hotelImageResponse := dtos.HotelImageResponse{
			HotelID:  hotelImage.HotelID,
			ImageUrl: hotelImage.ImageUrl,
		}
		hotelImageResponses = append(hotelImageResponses, hotelImageResponse)
	}
	getHotelFacilities, err := u.hotelFacilitiesRepo.GetAllHotelFacilitiesByID(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponses, err
	}
	getHotelPolicies, err := u.hotelPoliciesRepo.GetHotelPoliciesByIDHotel(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponses, err
	}
	var hotelFacilitiesResponses []dtos.HotelFacilitiesResponse
	for _, hotelFacilities := range getHotelFacilities {
		hotelFacilitiesResponse := dtos.HotelFacilitiesResponse{
			HotelID: hotelFacilities.HotelID,
			Name:    hotelFacilities.Name,
		}
		hotelFacilitiesResponses = append(hotelFacilitiesResponses, hotelFacilitiesResponse)
	}
	getHotelRoom, err := u.hotelRoomRepo.GetHotelRoomByID(hotelOrder.HotelRoomID)
	if err != nil {
		return hotelOrderResponses, err
	}
	getHotelRoomImage, err := u.hotelRoomImageRepo.GetAllHotelRoomImageByID(getHotelRoom.ID)
	if err != nil {
		return hotelOrderResponses, err
	}
	var hotelRoomImageResponses []dtos.HotelRoomImageResponse
	for _, hotelRoomImage := range getHotelRoomImage {
		hotelRoomImageResponse := dtos.HotelRoomImageResponse{
			HotelID:     hotelRoomImage.HotelID,
			HotelRoomID: hotelRoomImage.ID,
			ImageUrl:    hotelRoomImage.ImageUrl,
		}
		hotelRoomImageResponses = append(hotelRoomImageResponses, hotelRoomImageResponse)
	}
	getHotelRoomFacilities, err := u.hotelRoomFacilitiesRepo.GetAllHotelRoomFacilitiesByHotelRoomID(getHotelRoom.ID)
	if err != nil {
		return hotelOrderResponses, err
	}
	var hotelRoomFacilitiesResponses []dtos.HotelRoomFacilitiesResponse
	for _, hotelRoomFacilities := range getHotelRoomFacilities {
		hotelRoomFacilitiesResponse := dtos.HotelRoomFacilitiesResponse{
			HotelID:     hotelRoomFacilities.HotelID,
			HotelRoomID: hotelRoomFacilities.ID,
			Name:        hotelRoomFacilities.Name,
		}
		hotelRoomFacilitiesResponses = append(hotelRoomFacilitiesResponses, hotelRoomFacilitiesResponse)
	}
	getPayment, err := u.paymentRepo.GetPaymentByID(uint(hotelOrder.PaymentID))
	if err != nil {
		return hotelOrderResponses, err
	}
	getTravelerDetail, err := u.travelerDetailRepo.GetTravelerDetailByHotelOrderID(hotelOrder.ID)
	if err != nil {
		return hotelOrderResponses, err
	}

	var travelerDetailResponses []dtos.TravelerDetailResponse

	for _, travelerDetail := range getTravelerDetail {
		travelerDetailResponse := dtos.TravelerDetailResponse{
			ID:           int(travelerDetail.ID),
			Title:        travelerDetail.Title,
			FullName:     travelerDetail.FullName,
			IDCardNumber: *travelerDetail.IDCardNumber,
		}
		travelerDetailResponses = append(travelerDetailResponses, travelerDetailResponse)
	}

	hotelOrderResponses = dtos.HotelOrderResponse{
		HotelOrderID:     int(hotelOrder.ID),
		QuantityAdult:    hotelOrder.QuantityAdult,
		QuantityInfant:   hotelOrder.QuantityInfant,
		NumberOfNight:    hotelOrder.NumberOfNight,
		DateStart:        helpers.FormatDateToYMD(&hotelOrder.DateStart),
		DateEnd:          helpers.FormatDateToYMD(&hotelOrder.DateEnd),
		Price:            hotelOrder.Price,
		TotalAmount:      hotelOrder.TotalAmount,
		NameOrder:        hotelOrder.NameOrder,
		EmailOrder:       hotelOrder.EmailOrder,
		PhoneNumberOrder: hotelOrder.PhoneNumberOrder,
		SpecialRequest:   hotelOrder.SpecialRequest,
		HotelOrderCode:   hotelOrder.HotelOrderCode,
		Status:           hotelOrder.Status,
		Hotel: dtos.HotelByIDResponses{
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
			HotelRoom: dtos.HotelRoomHotelIDResponse{
				HotelRoomID:       getHotelRoom.ID,
				HotelID:           getHotelRoom.HotelID,
				Name:              getHotelRoom.Name,
				SizeOfRoom:        getHotelRoom.SizeOfRoom,
				QuantityOfRoom:    getHotelRoom.QuantityOfRoom,
				Description:       getHotelRoom.Description,
				NormalPrice:       getHotelRoom.NormalPrice,
				Discount:          getHotelRoom.Discount,
				DiscountPrice:     getHotelRoom.DiscountPrice,
				NumberOfGuest:     getHotelRoom.NumberOfGuest,
				MattressSize:      getHotelRoom.MattressSize,
				NumberOfMattress:  getHotelRoom.NumberOfMattress,
				HotelRoomImage:    hotelRoomImageResponses,
				HotelRoomFacility: hotelRoomFacilitiesResponses,
			},
		},
		Payment: dtos.PaymentResponses{
			ID:            int(getPayment.ID),
			Type:          getPayment.Type,
			ImageUrl:      getPayment.ImageUrl,
			Name:          getPayment.Name,
			AccountName:   getPayment.AccountName,
			AccountNumber: getPayment.AccountNumber,
		},
		TravelerDetail: travelerDetailResponses,
		CreatedAt:      hotelOrder.CreatedAt,
		UpdatedAt:      hotelOrder.UpdatedAt,
	}
	return hotelOrderResponses, nil
}

// GetHotelOrderByID godoc
// @Summary      Get Hotel Order User by ID
// @Description  Get Hotel Order User by ID
// @Tags         User - Order
// @Accept       json
// @Produce      json
// @Param hotel_order_id query int true "Hotel Order ID"
// @Success      200 {object} dtos.HotelOrderStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/order/hotel/detail [get]
// @Security BearerAuth
func (u *hotelOrderUsecase) GetHotelOrderByID(userID, hotelOrderId uint) (dtos.HotelOrderResponse, error) {
	var hotelOrderResponses dtos.HotelOrderResponse

	hotelOrder, err := u.hotelOrderRepo.GetHotelOrderByID(hotelOrderId, userID)
	if err != nil {
		return hotelOrderResponses, err
	}

	getHotel, err := u.hotelRepo.GetHotelByID(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponses, err
	}
	getHotelImage, err := u.hotelImageRepo.GetAllHotelImageByID(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponses, err
	}
	var hotelImageResponses []dtos.HotelImageResponse
	for _, hotelImage := range getHotelImage {
		hotelImageResponse := dtos.HotelImageResponse{
			HotelID:  hotelImage.HotelID,
			ImageUrl: hotelImage.ImageUrl,
		}
		hotelImageResponses = append(hotelImageResponses, hotelImageResponse)
	}
	getHotelFacilities, err := u.hotelFacilitiesRepo.GetAllHotelFacilitiesByID(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponses, err
	}
	getHotelPolicies, err := u.hotelPoliciesRepo.GetHotelPoliciesByIDHotel(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponses, err
	}
	var hotelFacilitiesResponses []dtos.HotelFacilitiesResponse
	for _, hotelFacilities := range getHotelFacilities {
		hotelFacilitiesResponse := dtos.HotelFacilitiesResponse{
			HotelID: hotelFacilities.HotelID,
			Name:    hotelFacilities.Name,
		}
		hotelFacilitiesResponses = append(hotelFacilitiesResponses, hotelFacilitiesResponse)
	}
	getHotelRoom, err := u.hotelRoomRepo.GetHotelRoomByID(hotelOrder.HotelRoomID)
	if err != nil {
		return hotelOrderResponses, err
	}
	getHotelRoomImage, err := u.hotelRoomImageRepo.GetAllHotelRoomImageByID(getHotelRoom.ID)
	if err != nil {
		return hotelOrderResponses, err
	}
	var hotelRoomImageResponses []dtos.HotelRoomImageResponse
	for _, hotelRoomImage := range getHotelRoomImage {
		hotelRoomImageResponse := dtos.HotelRoomImageResponse{
			HotelID:     hotelRoomImage.HotelID,
			HotelRoomID: hotelRoomImage.ID,
			ImageUrl:    hotelRoomImage.ImageUrl,
		}
		hotelRoomImageResponses = append(hotelRoomImageResponses, hotelRoomImageResponse)
	}
	getHotelRoomFacilities, err := u.hotelRoomFacilitiesRepo.GetAllHotelRoomFacilitiesByHotelRoomID(getHotelRoom.ID)
	if err != nil {
		return hotelOrderResponses, err
	}
	var hotelRoomFacilitiesResponses []dtos.HotelRoomFacilitiesResponse
	for _, hotelRoomFacilities := range getHotelRoomFacilities {
		hotelRoomFacilitiesResponse := dtos.HotelRoomFacilitiesResponse{
			HotelID:     hotelRoomFacilities.HotelID,
			HotelRoomID: hotelRoomFacilities.ID,
			Name:        hotelRoomFacilities.Name,
		}
		hotelRoomFacilitiesResponses = append(hotelRoomFacilitiesResponses, hotelRoomFacilitiesResponse)
	}
	getPayment, err := u.paymentRepo.GetPaymentByID(uint(hotelOrder.PaymentID))
	if err != nil {
		return hotelOrderResponses, err
	}
	getTravelerDetail, err := u.travelerDetailRepo.GetTravelerDetailByHotelOrderID(hotelOrder.ID)
	if err != nil {
		return hotelOrderResponses, err
	}

	var travelerDetailResponses []dtos.TravelerDetailResponse

	for _, travelerDetail := range getTravelerDetail {
		travelerDetailResponse := dtos.TravelerDetailResponse{
			ID:           int(travelerDetail.ID),
			Title:        travelerDetail.Title,
			FullName:     travelerDetail.FullName,
			IDCardNumber: *travelerDetail.IDCardNumber,
		}
		travelerDetailResponses = append(travelerDetailResponses, travelerDetailResponse)
	}

	hotelOrderResponses = dtos.HotelOrderResponse{
		HotelOrderID:     int(hotelOrder.ID),
		QuantityAdult:    hotelOrder.QuantityAdult,
		QuantityInfant:   hotelOrder.QuantityInfant,
		NumberOfNight:    hotelOrder.NumberOfNight,
		DateStart:        helpers.FormatDateToYMD(&hotelOrder.DateStart),
		DateEnd:          helpers.FormatDateToYMD(&hotelOrder.DateEnd),
		Price:            hotelOrder.Price,
		TotalAmount:      hotelOrder.TotalAmount,
		NameOrder:        hotelOrder.NameOrder,
		EmailOrder:       hotelOrder.EmailOrder,
		PhoneNumberOrder: hotelOrder.PhoneNumberOrder,
		SpecialRequest:   hotelOrder.SpecialRequest,
		HotelOrderCode:   hotelOrder.HotelOrderCode,
		Status:           hotelOrder.Status,
		Hotel: dtos.HotelByIDResponses{
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
			HotelRoom: dtos.HotelRoomHotelIDResponse{
				HotelRoomID:       getHotelRoom.ID,
				HotelID:           getHotelRoom.HotelID,
				Name:              getHotelRoom.Name,
				SizeOfRoom:        getHotelRoom.SizeOfRoom,
				QuantityOfRoom:    getHotelRoom.QuantityOfRoom,
				Description:       getHotelRoom.Description,
				NormalPrice:       getHotelRoom.NormalPrice,
				Discount:          getHotelRoom.Discount,
				DiscountPrice:     getHotelRoom.DiscountPrice,
				NumberOfGuest:     getHotelRoom.NumberOfGuest,
				MattressSize:      getHotelRoom.MattressSize,
				NumberOfMattress:  getHotelRoom.NumberOfMattress,
				HotelRoomImage:    hotelRoomImageResponses,
				HotelRoomFacility: hotelRoomFacilitiesResponses,
			},
		},
		Payment: dtos.PaymentResponses{
			ID:            int(getPayment.ID),
			Type:          getPayment.Type,
			ImageUrl:      getPayment.ImageUrl,
			Name:          getPayment.Name,
			AccountName:   getPayment.AccountName,
			AccountNumber: getPayment.AccountNumber,
		},
		TravelerDetail: travelerDetailResponses,
		CreatedAt:      hotelOrder.CreatedAt,
		UpdatedAt:      hotelOrder.UpdatedAt,
	}

	return hotelOrderResponses, nil
}

// CreateHotelOrder godoc
// @Summary      Order Hotel
// @Description  Order Hotel
// @Tags         User - Hotel
// @Accept       json
// @Produce      json
// @Param        request body dtos.HotelOrderInput true "Payload Body [RAW]"
// @Success      201 {object} dtos.HotelOrderCreeatedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/hotel/order [post]
// @Security BearerAuth
func (u *hotelOrderUsecase) CreateHotelOrder(userID uint, hotelOrderInput dtos.HotelOrderInput) (dtos.HotelOrderResponse, error) {
	var hotelOrderResponse dtos.HotelOrderResponse
	sumHotelPrice := 0
	if hotelOrderInput.HotelRoomID < 1 || hotelOrderInput.QuantityAdult < 1 || hotelOrderInput.DateStart == "" || hotelOrderInput.DateEnd == "" || hotelOrderInput.PaymentID < 1 || hotelOrderInput.NameOrder == "" || hotelOrderInput.EmailOrder == "" || hotelOrderInput.PhoneNumberOrder == "" || hotelOrderInput.TravelerDetail == nil {
		return hotelOrderResponse, errors.New("Failed to create hotel order")
	}
	getHotelRooms, err := u.hotelRoomRepo.GetHotelRoomByID(uint(hotelOrderInput.HotelRoomID))
	if err != nil {
		return hotelOrderResponse, err
	}
	getHotels, err := u.hotelRepo.GetHotelByID(getHotelRooms.HotelID)
	if err != nil {
		return hotelOrderResponse, err
	}

	dateNow := "2006-01-02"
	dateStartParse, err := time.Parse(dateNow, *&hotelOrderInput.DateStart)
	if err != nil {
		return hotelOrderResponse, errors.New("Failed to parse date start")
	}
	dateEndParse, err := time.Parse(dateNow, *&hotelOrderInput.DateEnd)
	if err != nil {
		return hotelOrderResponse, errors.New("Failed to parse date end")
	}

	days := 1

	if hotelOrderInput.DateStart < hotelOrderInput.DateEnd {
		duration := dateEndParse.Sub(dateStartParse)
		days = int(duration.Hours() / 24)
	} else if hotelOrderInput.DateStart == hotelOrderInput.DateEnd {
		days = 1
	} else {
		return hotelOrderResponse, errors.New("Failed to date start cannot be larger than date end")
	}

	createHotelOrder := models.HotelOrder{
		UserID:           userID,
		HotelID:          getHotels.ID,
		HotelRoomID:      uint(hotelOrderInput.HotelRoomID),
		QuantityAdult:    hotelOrderInput.QuantityAdult,
		QuantityInfant:   hotelOrderInput.QuantityInfant,
		NumberOfNight:    days,
		DateStart:        dateStartParse,
		DateEnd:          dateEndParse,
		Price:            0,
		PaymentID:        hotelOrderInput.PaymentID,
		TotalAmount:      0,
		NameOrder:        hotelOrderInput.NameOrder,
		EmailOrder:       hotelOrderInput.EmailOrder,
		PhoneNumberOrder: hotelOrderInput.PhoneNumberOrder,
		SpecialRequest:   hotelOrderInput.SpecialRequest,
		HotelOrderCode:   "hotel-order-" + uuid.New().String(),
		Status:           "unpaid",
	}

	getPayment, err := u.paymentRepo.GetPaymentByID2(uint(createHotelOrder.PaymentID))
	if err != nil {
		return hotelOrderResponse, errors.New("failed to get payment id")
	}

	createHotelOrder, err = u.hotelOrderRepo.CreateHotelOrder(createHotelOrder)
	if err != nil {
		return hotelOrderResponse, err
	}

	if createHotelOrder.ID > 0 && createHotelOrder.Status == "unpaid" {
		createNotification := models.Notification{
			UserID:     userID,
			TemplateID: 7,
		}

		_, err = u.notificationRepo.CreateNotification(createNotification)
		if err != nil {
			return hotelOrderResponse, err
		}
	}

	var travelerDetailResponses []dtos.TravelerDetailResponse
	for _, travelerDetail := range hotelOrderInput.TravelerDetail {
		travelerDetailResponse := models.TravelerDetail{
			UserID:       userID,
			HotelOrderID: &createHotelOrder.ID,
			Title:        travelerDetail.Title,
			FullName:     travelerDetail.FullName,
			IDCardNumber: &travelerDetail.IDCardNumber,
		}
		createTravelerDetail, err := u.travelerDetailRepo.CreateTravelerDetail(travelerDetailResponse)
		if err != nil {
			return hotelOrderResponse, err
		}
		travelerDetailResponseses := dtos.TravelerDetailResponse{
			ID:           int(createTravelerDetail.ID),
			Title:        createTravelerDetail.Title,
			FullName:     createTravelerDetail.FullName,
			IDCardNumber: *createTravelerDetail.IDCardNumber,
		}
		travelerDetailResponses = append(travelerDetailResponses, travelerDetailResponseses)
	}

	hotelOrder, err := u.hotelOrderRepo.GetHotelOrderByID(createHotelOrder.HotelID, userID)
	if err != nil {
		return hotelOrderResponse, err
	}

	getHotel, err := u.hotelRepo.GetHotelByID(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponse, err
	}
	getHotelImage, err := u.hotelImageRepo.GetAllHotelImageByID(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponse, err
	}
	var hotelImageResponses []dtos.HotelImageResponse
	for _, hotelImage := range getHotelImage {
		hotelImageResponse := dtos.HotelImageResponse{
			HotelID:  hotelImage.HotelID,
			ImageUrl: hotelImage.ImageUrl,
		}
		hotelImageResponses = append(hotelImageResponses, hotelImageResponse)
	}
	getHotelFacilities, err := u.hotelFacilitiesRepo.GetAllHotelFacilitiesByID(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponse, err
	}
	getHotelPolicies, err := u.hotelPoliciesRepo.GetHotelPoliciesByIDHotel(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponse, err
	}
	var hotelFacilitiesResponses []dtos.HotelFacilitiesResponse
	for _, hotelFacilities := range getHotelFacilities {
		hotelFacilitiesResponse := dtos.HotelFacilitiesResponse{
			HotelID: hotelFacilities.HotelID,
			Name:    hotelFacilities.Name,
		}
		hotelFacilitiesResponses = append(hotelFacilitiesResponses, hotelFacilitiesResponse)
	}
	getHotelRoom, err := u.hotelRoomRepo.GetHotelRoomByID(hotelOrder.HotelRoomID)
	if err != nil {
		return hotelOrderResponse, err
	}

	sumHotelPrice = getHotelRoom.DiscountPrice

	createHotelOrder.Price = sumHotelPrice
	createHotelOrder.TotalAmount = sumHotelPrice * createHotelOrder.NumberOfNight

	hotelOrder, err = u.hotelOrderRepo.UpdateHotelOrder(createHotelOrder)
	if err != nil {
		return hotelOrderResponse, err
	}

	getHotelRoomImage, err := u.hotelRoomImageRepo.GetAllHotelRoomImageByID(getHotelRoom.ID)
	if err != nil {
		return hotelOrderResponse, err
	}
	var hotelRoomImageResponses []dtos.HotelRoomImageResponse
	for _, hotelRoomImage := range getHotelRoomImage {
		hotelRoomImageResponse := dtos.HotelRoomImageResponse{
			HotelID:     hotelRoomImage.HotelID,
			HotelRoomID: hotelRoomImage.ID,
			ImageUrl:    hotelRoomImage.ImageUrl,
		}
		hotelRoomImageResponses = append(hotelRoomImageResponses, hotelRoomImageResponse)
	}
	getHotelRoomFacilities, err := u.hotelRoomFacilitiesRepo.GetAllHotelRoomFacilitiesByHotelRoomID(getHotelRoom.ID)
	if err != nil {
		return hotelOrderResponse, err
	}
	var hotelRoomFacilitiesResponses []dtos.HotelRoomFacilitiesResponse
	for _, hotelRoomFacilities := range getHotelRoomFacilities {
		hotelRoomFacilitiesResponse := dtos.HotelRoomFacilitiesResponse{
			HotelID:     hotelRoomFacilities.HotelID,
			HotelRoomID: hotelRoomFacilities.ID,
			Name:        hotelRoomFacilities.Name,
		}
		hotelRoomFacilitiesResponses = append(hotelRoomFacilitiesResponses, hotelRoomFacilitiesResponse)
	}

	hotelOrderResponse = dtos.HotelOrderResponse{
		HotelOrderID:     int(hotelOrder.ID),
		QuantityAdult:    hotelOrder.QuantityAdult,
		QuantityInfant:   hotelOrder.QuantityInfant,
		NumberOfNight:    hotelOrder.NumberOfNight,
		DateStart:        helpers.FormatDateToYMD(&hotelOrder.DateStart),
		DateEnd:          helpers.FormatDateToYMD(&hotelOrder.DateEnd),
		Price:            hotelOrder.Price,
		TotalAmount:      hotelOrder.TotalAmount,
		NameOrder:        hotelOrder.NameOrder,
		EmailOrder:       hotelOrder.EmailOrder,
		PhoneNumberOrder: hotelOrder.PhoneNumberOrder,
		SpecialRequest:   hotelOrder.SpecialRequest,
		HotelOrderCode:   hotelOrder.HotelOrderCode,
		Status:           hotelOrder.Status,
		Hotel: dtos.HotelByIDResponses{
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
			HotelRoom: dtos.HotelRoomHotelIDResponse{
				HotelRoomID:       getHotelRoom.ID,
				HotelID:           getHotelRoom.HotelID,
				Name:              getHotelRoom.Name,
				SizeOfRoom:        getHotelRoom.SizeOfRoom,
				QuantityOfRoom:    getHotelRoom.QuantityOfRoom,
				Description:       getHotelRoom.Description,
				NormalPrice:       getHotelRoom.NormalPrice,
				Discount:          getHotelRoom.Discount,
				DiscountPrice:     getHotelRoom.DiscountPrice,
				NumberOfGuest:     getHotelRoom.NumberOfGuest,
				MattressSize:      getHotelRoom.MattressSize,
				NumberOfMattress:  getHotelRoom.NumberOfMattress,
				HotelRoomImage:    hotelRoomImageResponses,
				HotelRoomFacility: hotelRoomFacilitiesResponses,
			},
		},
		Payment: dtos.PaymentResponses{
			ID:            int(getPayment.ID),
			Type:          getPayment.Type,
			ImageUrl:      getPayment.ImageUrl,
			Name:          getPayment.Name,
			AccountName:   getPayment.AccountName,
			AccountNumber: getPayment.AccountNumber,
		},
		TravelerDetail: travelerDetailResponses,
		CreatedAt:      hotelOrder.CreatedAt,
		UpdatedAt:      hotelOrder.UpdatedAt,
	}

	return hotelOrderResponse, nil
}

// UpdateHotelOrder godoc
// @Summary      Update Order Hotel
// @Description  Update Order Hotel
// @Tags         User - Hotel
// @Accept       json
// @Produce      json
// @Param hotel_order_id query int true "Hotel Order ID"
// @Param status query string true "Update Status Order ID"
// @Success      200 {object} dtos.HotelOrderStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/hotel/order [patch]
// @Security BearerAuth
func (u *hotelOrderUsecase) UpdateHotelOrder(userID, hotelOrderID uint, status string) (dtos.HotelOrderResponse, error) {
	var hotelOrderResponses dtos.HotelOrderResponse

	hotelOrder, err := u.hotelOrderRepo.GetHotelOrderByID(hotelOrderID, userID)
	if err != nil {
		return hotelOrderResponses, err
	}
	hotelOrder.Status = status
	hotelOrder, err = u.hotelOrderRepo.UpdateHotelOrder(hotelOrder)
	if err != nil {
		return hotelOrderResponses, err
	}

	if hotelOrder.ID > 0 && hotelOrder.Status == "paid" {
		createNotification := models.Notification{
			UserID:     userID,
			TemplateID: 5,
		}

		_, err = u.notificationRepo.CreateNotification(createNotification)
		if err != nil {
			return hotelOrderResponses, err
		}
	}

	if hotelOrder.ID > 0 && hotelOrder.Status == "canceled" {
		createNotification := models.Notification{
			UserID:     userID,
			TemplateID: 8,
		}

		_, err = u.notificationRepo.CreateNotification(createNotification)
		if err != nil {
			return hotelOrderResponses, err
		}
	}

	getHotel, err := u.hotelRepo.GetHotelByID(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponses, err
	}
	getHotelImage, err := u.hotelImageRepo.GetAllHotelImageByID(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponses, err
	}
	var hotelImageResponses []dtos.HotelImageResponse
	for _, hotelImage := range getHotelImage {
		hotelImageResponse := dtos.HotelImageResponse{
			HotelID:  hotelImage.HotelID,
			ImageUrl: hotelImage.ImageUrl,
		}
		hotelImageResponses = append(hotelImageResponses, hotelImageResponse)
	}
	getHotelFacilities, err := u.hotelFacilitiesRepo.GetAllHotelFacilitiesByID(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponses, err
	}
	getHotelPolicies, err := u.hotelPoliciesRepo.GetHotelPoliciesByIDHotel(hotelOrder.HotelID)
	if err != nil {
		return hotelOrderResponses, err
	}
	var hotelFacilitiesResponses []dtos.HotelFacilitiesResponse
	for _, hotelFacilities := range getHotelFacilities {
		hotelFacilitiesResponse := dtos.HotelFacilitiesResponse{
			HotelID: hotelFacilities.HotelID,
			Name:    hotelFacilities.Name,
		}
		hotelFacilitiesResponses = append(hotelFacilitiesResponses, hotelFacilitiesResponse)
	}
	getHotelRoom, err := u.hotelRoomRepo.GetHotelRoomByID(hotelOrder.HotelRoomID)
	if err != nil {
		return hotelOrderResponses, err
	}
	getHotelRoomImage, err := u.hotelRoomImageRepo.GetAllHotelRoomImageByID(getHotelRoom.ID)
	if err != nil {
		return hotelOrderResponses, err
	}
	var hotelRoomImageResponses []dtos.HotelRoomImageResponse
	for _, hotelRoomImage := range getHotelRoomImage {
		hotelRoomImageResponse := dtos.HotelRoomImageResponse{
			HotelID:     hotelRoomImage.HotelID,
			HotelRoomID: hotelRoomImage.ID,
			ImageUrl:    hotelRoomImage.ImageUrl,
		}
		hotelRoomImageResponses = append(hotelRoomImageResponses, hotelRoomImageResponse)
	}
	getHotelRoomFacilities, err := u.hotelRoomFacilitiesRepo.GetAllHotelRoomFacilitiesByHotelRoomID(getHotelRoom.ID)
	if err != nil {
		return hotelOrderResponses, err
	}
	var hotelRoomFacilitiesResponses []dtos.HotelRoomFacilitiesResponse
	for _, hotelRoomFacilities := range getHotelRoomFacilities {
		hotelRoomFacilitiesResponse := dtos.HotelRoomFacilitiesResponse{
			HotelID:     hotelRoomFacilities.HotelID,
			HotelRoomID: hotelRoomFacilities.ID,
			Name:        hotelRoomFacilities.Name,
		}
		hotelRoomFacilitiesResponses = append(hotelRoomFacilitiesResponses, hotelRoomFacilitiesResponse)
	}
	getPayment, err := u.paymentRepo.GetPaymentByID(uint(hotelOrder.PaymentID))
	if err != nil {
		return hotelOrderResponses, err
	}
	getTravelerDetail, err := u.travelerDetailRepo.GetTravelerDetailByHotelOrderID(hotelOrder.ID)
	if err != nil {
		return hotelOrderResponses, err
	}

	var travelerDetailResponses []dtos.TravelerDetailResponse

	for _, travelerDetail := range getTravelerDetail {
		travelerDetailResponse := dtos.TravelerDetailResponse{
			ID:           int(travelerDetail.ID),
			Title:        travelerDetail.Title,
			FullName:     travelerDetail.FullName,
			IDCardNumber: *travelerDetail.IDCardNumber,
		}
		travelerDetailResponses = append(travelerDetailResponses, travelerDetailResponse)
	}

	hotelOrderResponses = dtos.HotelOrderResponse{
		HotelOrderID:     int(hotelOrder.ID),
		QuantityAdult:    hotelOrder.QuantityAdult,
		QuantityInfant:   hotelOrder.QuantityInfant,
		NumberOfNight:    hotelOrder.NumberOfNight,
		DateStart:        helpers.FormatDateToYMD(&hotelOrder.DateStart),
		DateEnd:          helpers.FormatDateToYMD(&hotelOrder.DateEnd),
		Price:            hotelOrder.Price,
		TotalAmount:      hotelOrder.TotalAmount,
		NameOrder:        hotelOrder.NameOrder,
		EmailOrder:       hotelOrder.EmailOrder,
		PhoneNumberOrder: hotelOrder.PhoneNumberOrder,
		SpecialRequest:   hotelOrder.SpecialRequest,
		HotelOrderCode:   hotelOrder.HotelOrderCode,
		Status:           hotelOrder.Status,
		Hotel: dtos.HotelByIDResponses{
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
			HotelRoom: dtos.HotelRoomHotelIDResponse{
				HotelRoomID:       getHotelRoom.ID,
				HotelID:           getHotelRoom.HotelID,
				Name:              getHotelRoom.Name,
				SizeOfRoom:        getHotelRoom.SizeOfRoom,
				QuantityOfRoom:    getHotelRoom.QuantityOfRoom,
				Description:       getHotelRoom.Description,
				NormalPrice:       getHotelRoom.NormalPrice,
				Discount:          getHotelRoom.Discount,
				DiscountPrice:     getHotelRoom.DiscountPrice,
				NumberOfGuest:     getHotelRoom.NumberOfGuest,
				MattressSize:      getHotelRoom.MattressSize,
				NumberOfMattress:  getHotelRoom.NumberOfMattress,
				HotelRoomImage:    hotelRoomImageResponses,
				HotelRoomFacility: hotelRoomFacilitiesResponses,
			},
		},
		Payment: dtos.PaymentResponses{
			ID:            int(getPayment.ID),
			Type:          getPayment.Type,
			ImageUrl:      getPayment.ImageUrl,
			Name:          getPayment.Name,
			AccountName:   getPayment.AccountName,
			AccountNumber: getPayment.AccountNumber,
		},
		TravelerDetail: travelerDetailResponses,
		CreatedAt:      hotelOrder.CreatedAt,
		UpdatedAt:      hotelOrder.UpdatedAt,
	}

	return hotelOrderResponses, nil
}

func hasMatchingTravelerDetail(hotelOrderID uint, search string, travelerDetailRepo repositories.TravelerDetailRepository) bool {
	travelerDetails, err := travelerDetailRepo.GetTravelerDetailByHotelOrderID(hotelOrderID)
	if err != nil {
		return false
	}

	for _, travelerDetail := range travelerDetails {
		if strings.Contains(strings.ToLower(travelerDetail.FullName), strings.ToLower(search)) {
			return true
		}
	}

	return false
}
