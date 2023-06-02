package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"time"

	"github.com/google/uuid"
)

type TicketOrderUsecase interface {
	GetAllTicketOrders(page, limit int) ([]dtos.TicketOrderResponse, int, error)
	// GetTicketOrderByID(id uint) (dtos.TicketOrderResponse, error)
	CreateTicketOrder(userID uint, ticketOrderInput dtos.TicketOrderInput) (dtos.TicketOrderResponse, error)
	UpdateTicketOrder(userID, ticketOrderID uint, status string) (dtos.TicketOrderResponse, error)
}

type ticketOrderUsecase struct {
	ticketOrderRepo          repositories.TicketOrderRepository
	ticketTravelerDetailRepo repositories.TicketTravelerDetailRepository
	travelerDetailRepo       repositories.TravelerDetailRepository
	trainCarriageRepo        repositories.TrainCarriageRepository
	trainRepo                repositories.TrainRepository
	trainSeatRepo            repositories.TrainSeatRepository
	stationRepo              repositories.StationRepository
	trainStationRepo         repositories.TrainStationRepository
}

func NewTicketOrderUsecase(ticketOrderRepo repositories.TicketOrderRepository, ticketTravelerDetailRepo repositories.TicketTravelerDetailRepository, travelerDetailRepo repositories.TravelerDetailRepository, trainCarriageRepo repositories.TrainCarriageRepository, trainRepo repositories.TrainRepository, trainSeatRepo repositories.TrainSeatRepository, stationRepo repositories.StationRepository, trainStationRepo repositories.TrainStationRepository) TicketOrderUsecase {
	return &ticketOrderUsecase{ticketOrderRepo, ticketTravelerDetailRepo, travelerDetailRepo, trainCarriageRepo, trainRepo, trainSeatRepo, stationRepo, trainStationRepo}
}

func (u *ticketOrderUsecase) GetAllTicketOrders(page, limit int) ([]dtos.TicketOrderResponse, int, error) {
	var ticketOrderResponses []dtos.TicketOrderResponse

	ticketOrder, count, err := u.ticketOrderRepo.GetAllTicketOrders(page, limit)
	if err != nil {
		return nil, 0, err
	}
	for _, ticketOrder := range ticketOrder {
		ticketOrderResponse := dtos.TicketOrderResponse{
			TicketOrderID: int(ticketOrder.ID),
		}
		ticketOrderResponses = append(ticketOrderResponses, ticketOrderResponse)
	}

	return ticketOrderResponses, count, nil
}

// CreateTicketOrder godoc
// @Summary      Order ticket KA
// @Description  Order ticket KA
// @Tags         User - Train
// @Accept       json
// @Produce      json
// @Param        request body dtos.TicketOrderInput true "Payload Body [RAW]"
// @Success      201 {object} dtos.TicketOrderCreeatedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/train/order [post]
// @Security BearerAuth
func (u *ticketOrderUsecase) CreateTicketOrder(userID uint, ticketOrderInput dtos.TicketOrderInput) (dtos.TicketOrderResponse, error) {
	var ticketOrderResponse dtos.TicketOrderResponse
	createTicketOrder := models.TicketOrder{
		UserID:           userID,
		QuantityAdult:    ticketOrderInput.QuantityAdult,
		QuantityInfant:   ticketOrderInput.QuantityInfant,
		Price:            ticketOrderInput.Price,
		WithReturn:       ticketOrderInput.WithReturn,
		PaymentID:        ticketOrderInput.PaymentID,
		TotalAmount:      ticketOrderInput.QuantityAdult * ticketOrderInput.Price,
		NameOrder:        ticketOrderInput.NameOrder,
		EmailOrder:       ticketOrderInput.EmailOrder,
		PhoneNumberOrder: ticketOrderInput.PhoneNumberOrder,
		TicketOrderCode:  "ticket-order-" + uuid.New().String(),
		Status:           "pending",
	}

	createTicketOrder, err := u.ticketOrderRepo.CreateTicketOrder(createTicketOrder)
	if err != nil {
		return ticketOrderResponse, err
	}

	var travelerDetailResponses []dtos.TravelerDetailResponse
	var ticketTravelerDetailDepartureResponses []dtos.TicketTravelerDetailResponse
	var ticketTravelerDetailReturnResponses []dtos.TicketTravelerDetailResponse

	for _, travelerDetail := range ticketOrderInput.TravelerDetail {
		createTravelerDetail := models.TravelerDetail{
			UserID:        userID,
			TicketOrderID: createTicketOrder.ID,
			Title:         travelerDetail.Title,
			FullName:      travelerDetail.FullName,
			IDCardNumber:  &travelerDetail.IDCardNumber,
		}
		createTravelerDetail, err := u.travelerDetailRepo.CreateTravelerDetail(createTravelerDetail)
		if err != nil {
			return ticketOrderResponse, err
		}

		travelerDetailResponse := dtos.TravelerDetailResponse{
			ID:           int(createTravelerDetail.ID),
			Title:        createTravelerDetail.Title,
			FullName:     createTravelerDetail.FullName,
			IDCardNumber: *createTravelerDetail.IDCardNumber,
		}
		travelerDetailResponses = append(travelerDetailResponses, travelerDetailResponse)

	}

	for _, ticketTravelerDetailDeparture := range ticketOrderInput.TicketTravelerDetailDeparture {
		// dateDepartureParse, err := helpers.FormatStringToDate(ticketTravelerDetailDeparture.Date)
		// if err != nil {
		// 	return ticketOrderResponse, errors.New("Failed to parsing date")
		// }

		travelerDetail, err := u.travelerDetailRepo.GetTravelerDetailByID(createTicketOrder.ID)
		if err != nil {
			return ticketOrderResponse, err
		}

		getTrain, err := u.trainRepo.GetTrainByID(uint(ticketTravelerDetailDeparture.TrainID))
		if err != nil {
			return ticketOrderResponse, err
		}
		getTrainCarriage, err := u.trainCarriageRepo.GetTrainCarriageByID(uint(ticketTravelerDetailDeparture.TrainCarriageID))
		if err != nil {
			return ticketOrderResponse, err
		}
		getTrainSeat, err := u.trainSeatRepo.GetTrainSeatByID(uint(ticketTravelerDetailDeparture.TrainSeatID))
		if err != nil {
			return ticketOrderResponse, err
		}
		getStationOrigin, err := u.stationRepo.GetStationByID(uint(ticketTravelerDetailDeparture.StationOriginID))
		if err != nil {
			return ticketOrderResponse, err
		}
		getStationDestination, err := u.stationRepo.GetStationByID(uint(ticketTravelerDetailDeparture.StationDestinationID))
		if err != nil {
			return ticketOrderResponse, err
		}

		trainStationOrigin, err := u.trainStationRepo.GetTrainStationByTrainIDStationID(getTrain.ID, getStationOrigin.ID)
		if err != nil {
			return ticketOrderResponse, err
		}
		trainStationDestination, err := u.trainStationRepo.GetTrainStationByTrainIDStationID(getTrain.ID, getStationDestination.ID)
		if err != nil {
			return ticketOrderResponse, err
		}

		createTicketTravelerDetail := models.TicketTravelerDetail{
			TicketOrderID:        createTicketOrder.ID,
			TravelerDetailID:     travelerDetail.ID,
			TrainID:              uint(getTrain.ID),
			TrainCarriageID:      uint(getTrainCarriage.ID),
			TrainSeatID:          uint(getTrainSeat.ID),
			StationOriginID:      uint(getStationOrigin.ID),
			DepartureTime:        trainStationOrigin.ArriveTime,
			StationDestinationID: uint(getStationDestination.ID),
			ArrivalTime:          trainStationDestination.ArriveTime,
			DateOfDeparture:      time.Now(),
			BoardingTicketCode:   "boarding-ticket-" + uuid.New().String(),
		}
		createTicketTravelerDetail, err = u.ticketTravelerDetailRepo.CreateTicketTravelerDetail(createTicketTravelerDetail)
		if err != nil {
			return ticketOrderResponse, err
		}

		ticketTravelerDetailResponse := dtos.TicketTravelerDetailResponse{
			TicketTravelerDetailID: int(createTicketTravelerDetail.ID),
			Train: dtos.TrainResponsesSimply{
				TrainID:         getTrain.ID,
				CodeTrain:       getTrain.CodeTrain,
				Name:            getTrain.Name,
				Class:           getTrainCarriage.Class,
				TrainCarriageID: getTrainCarriage.ID,
				TrainCarriage:   getTrainCarriage.Name,
				TrainSeatID:     getTrainSeat.ID,
				TrainSeat:       getTrainSeat.Name,
			},
			StationOrigin: dtos.StationResponseSimply{
				StationID:  getStationOrigin.ID,
				Origin:     getStationOrigin.Origin,
				Name:       getStationOrigin.Name,
				Initial:    getStationOrigin.Initial,
				ArriveTime: createTicketTravelerDetail.DepartureTime,
			},
			StationDestination: dtos.StationResponseSimply{
				StationID:  getStationDestination.ID,
				Origin:     getStationDestination.Origin,
				Name:       getStationDestination.Name,
				Initial:    getStationDestination.Initial,
				ArriveTime: createTicketTravelerDetail.ArrivalTime,
			},
			Date:               createTicketTravelerDetail.DateOfDeparture,
			BoardingTicketCode: createTicketTravelerDetail.BoardingTicketCode,
		}
		ticketTravelerDetailDepartureResponses = append(ticketTravelerDetailDepartureResponses, ticketTravelerDetailResponse)
	}

	if createTicketOrder.WithReturn {
		for _, ticketTravelerDetailReturn := range ticketOrderInput.TicketTravelerDetailReturn {
			// dateReturn, err := helpers.FormatStringToDate(ticketTravelerDetailReturn.Date)
			// if err != nil {
			// 	return ticketOrderResponse, errors.New("Failed to parsing date")
			// }

			travelerDetail, err := u.travelerDetailRepo.GetTravelerDetailByID(createTicketOrder.ID)
			if err != nil {
				return ticketOrderResponse, err
			}

			getTrain, err := u.trainRepo.GetTrainByID(uint(ticketTravelerDetailReturn.TrainID))
			if err != nil {
				return ticketOrderResponse, err
			}
			getTrainCarriage, err := u.trainCarriageRepo.GetTrainCarriageByID(uint(ticketTravelerDetailReturn.TrainCarriageID))
			if err != nil {
				return ticketOrderResponse, err
			}
			getTrainSeat, err := u.trainSeatRepo.GetTrainSeatByID(uint(ticketTravelerDetailReturn.TrainSeatID))
			if err != nil {
				return ticketOrderResponse, err
			}
			getStationOrigin, err := u.stationRepo.GetStationByID(uint(ticketTravelerDetailReturn.StationOriginID))
			if err != nil {
				return ticketOrderResponse, err
			}
			getStationDestination, err := u.stationRepo.GetStationByID(uint(ticketTravelerDetailReturn.StationDestinationID))
			if err != nil {
				return ticketOrderResponse, err
			}

			trainStationOrigin, err := u.trainStationRepo.GetTrainStationByTrainIDStationID(getTrain.ID, getStationOrigin.ID)
			if err != nil {
				return ticketOrderResponse, err
			}
			trainStationDestination, err := u.trainStationRepo.GetTrainStationByTrainIDStationID(getTrain.ID, getStationDestination.ID)
			if err != nil {
				return ticketOrderResponse, err
			}

			createTicketTravelerDetail := models.TicketTravelerDetail{
				TicketOrderID:        createTicketOrder.ID,
				TravelerDetailID:     travelerDetail.ID,
				TrainID:              uint(getTrain.ID),
				TrainCarriageID:      uint(getTrainCarriage.ID),
				TrainSeatID:          uint(getTrainSeat.ID),
				StationOriginID:      uint(getStationOrigin.ID),
				DepartureTime:        trainStationOrigin.ArriveTime,
				StationDestinationID: uint(getStationDestination.ID),
				ArrivalTime:          trainStationDestination.ArriveTime,
				DateOfDeparture:      time.Now(),
				BoardingTicketCode:   "boarding-ticket-" + uuid.New().String(),
			}
			createTicketTravelerDetail, err = u.ticketTravelerDetailRepo.CreateTicketTravelerDetail(createTicketTravelerDetail)
			if err != nil {
				return ticketOrderResponse, err
			}

			ticketTravelerDetailResponse := dtos.TicketTravelerDetailResponse{
				TicketTravelerDetailID: int(createTicketTravelerDetail.ID),
				Train: dtos.TrainResponsesSimply{
					TrainID:         getTrain.ID,
					CodeTrain:       getTrain.CodeTrain,
					Name:            getTrain.Name,
					Class:           getTrainCarriage.Class,
					TrainCarriageID: getTrainCarriage.ID,
					TrainCarriage:   getTrainCarriage.Name,
					TrainSeatID:     getTrainSeat.ID,
					TrainSeat:       getTrainSeat.Name,
				},
				StationOrigin: dtos.StationResponseSimply{
					StationID:  getStationOrigin.ID,
					Origin:     getStationOrigin.Origin,
					Name:       getStationOrigin.Name,
					Initial:    getStationOrigin.Initial,
					ArriveTime: createTicketTravelerDetail.DepartureTime,
				},
				StationDestination: dtos.StationResponseSimply{
					StationID:  getStationDestination.ID,
					Origin:     getStationDestination.Origin,
					Name:       getStationDestination.Name,
					Initial:    getStationDestination.Initial,
					ArriveTime: createTicketTravelerDetail.ArrivalTime,
				},
				Date:               createTicketTravelerDetail.DateOfDeparture,
				BoardingTicketCode: createTicketTravelerDetail.BoardingTicketCode,
			}
			ticketTravelerDetailReturnResponses = append(ticketTravelerDetailReturnResponses, ticketTravelerDetailResponse)
		}
	}
	getOrderTicket, err := u.ticketOrderRepo.GetTicketOrderByID(createTicketOrder.ID, userID)
	if err != nil {
		return ticketOrderResponse, err
	}

	ticketOrderResponse = dtos.TicketOrderResponse{
		TicketOrderID:                 int(getOrderTicket.ID),
		QuantityAdult:                 getOrderTicket.QuantityAdult,
		QuantityInfant:                getOrderTicket.QuantityInfant,
		Price:                         getOrderTicket.Price,
		TotalAmount:                   getOrderTicket.TotalAmount,
		WithReturn:                    getOrderTicket.WithReturn,
		PaymentID:                     getOrderTicket.PaymentID,
		NameOrder:                     getOrderTicket.NameOrder,
		EmailOrder:                    getOrderTicket.EmailOrder,
		PhoneNumberOrder:              getOrderTicket.PhoneNumberOrder,
		TicketOrderCode:               getOrderTicket.TicketOrderCode,
		Status:                        getOrderTicket.Status,
		TravelerDetail:                travelerDetailResponses,
		TicketTravelerDetailDeparture: ticketTravelerDetailDepartureResponses,
		TicketTravelerDetailReturn:    ticketTravelerDetailReturnResponses,
	}

	return ticketOrderResponse, nil
}

// UpdateTicketOrder godoc
// @Summary      Update Order ticket KA
// @Description  Update Order ticket KA
// @Tags         User - Train
// @Accept       json
// @Produce      json
// @Param ticket_order_id query int true "Ticket Order ID"
// @Param status query string true "Update Status Order ID"
// @Success      200 {object} dtos.TicketOrderStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /user/train/order [patch]
// @Security BearerAuth
func (u *ticketOrderUsecase) UpdateTicketOrder(userID, ticketOrderID uint, status string) (dtos.TicketOrderResponse, error) {
	var ticketOrderResponse dtos.TicketOrderResponse

	createTicketOrder, err := u.ticketOrderRepo.GetTicketOrderByID(ticketOrderID, userID)
	if err != nil {
		return ticketOrderResponse, err
	}

	createTicketOrder.Status = status

	createTicketOrder, err = u.ticketOrderRepo.UpdateTicketOrder(createTicketOrder)
	if err != nil {
		return ticketOrderResponse, err
	}

	getTravelerDetail, err := u.travelerDetailRepo.GetTravelerDetailByTicketOrderID(createTicketOrder.ID)
	if err != nil {
		return ticketOrderResponse, err
	}

	var travelerDetailResponses []dtos.TravelerDetailResponse
	var ticketTravelerDetailDepartureResponses []dtos.TicketTravelerDetailResponse
	var ticketTravelerDetailReturnResponses []dtos.TicketTravelerDetailResponse

	for _, travelerDetail := range getTravelerDetail {
		travelerDetailResponse := dtos.TravelerDetailResponse{
			ID:           int(travelerDetail.ID),
			Title:        travelerDetail.Title,
			FullName:     travelerDetail.FullName,
			IDCardNumber: *travelerDetail.IDCardNumber,
		}
		travelerDetailResponses = append(travelerDetailResponses, travelerDetailResponse)
	}

	getTicketTravelerDetail, err := u.ticketTravelerDetailRepo.GetTicketTravelerDetailByTrainID(createTicketOrder.ID)
	if err != nil {
		return ticketOrderResponse, err
	}

	for _, ticketTravelerDetailDeparture := range getTicketTravelerDetail {
		getTrain, err := u.trainRepo.GetTrainByID(uint(ticketTravelerDetailDeparture.TrainID))
		if err != nil {
			return ticketOrderResponse, err
		}
		getTrainCarriage, err := u.trainCarriageRepo.GetTrainCarriageByID(uint(ticketTravelerDetailDeparture.TrainCarriageID))
		if err != nil {
			return ticketOrderResponse, err
		}
		getTrainSeat, err := u.trainSeatRepo.GetTrainSeatByID(uint(ticketTravelerDetailDeparture.TrainSeatID))
		if err != nil {
			return ticketOrderResponse, err
		}
		getStationOrigin, err := u.stationRepo.GetStationByID(uint(ticketTravelerDetailDeparture.StationOriginID))
		if err != nil {
			return ticketOrderResponse, err
		}
		getStationDestination, err := u.stationRepo.GetStationByID(uint(ticketTravelerDetailDeparture.StationDestinationID))
		if err != nil {
			return ticketOrderResponse, err
		}

		ticketTravelerDetailResponse := dtos.TicketTravelerDetailResponse{
			TicketTravelerDetailID: int(ticketTravelerDetailDeparture.ID),
			Train: dtos.TrainResponsesSimply{
				TrainID:         getTrain.ID,
				CodeTrain:       getTrain.CodeTrain,
				Name:            getTrain.Name,
				Class:           getTrainCarriage.Class,
				TrainCarriageID: getTrainCarriage.ID,
				TrainCarriage:   getTrainCarriage.Name,
				TrainSeatID:     getTrainSeat.ID,
				TrainSeat:       getTrainSeat.Name,
			},
			StationOrigin: dtos.StationResponseSimply{
				StationID:  getStationOrigin.ID,
				Origin:     getStationOrigin.Origin,
				Name:       getStationOrigin.Name,
				Initial:    getStationOrigin.Initial,
				ArriveTime: ticketTravelerDetailDeparture.DepartureTime,
			},
			StationDestination: dtos.StationResponseSimply{
				StationID:  getStationDestination.ID,
				Origin:     getStationDestination.Origin,
				Name:       getStationDestination.Name,
				Initial:    getStationDestination.Initial,
				ArriveTime: ticketTravelerDetailDeparture.ArrivalTime,
			},
			Date:               ticketTravelerDetailDeparture.DateOfDeparture,
			BoardingTicketCode: ticketTravelerDetailDeparture.BoardingTicketCode,
		}
		ticketTravelerDetailDepartureResponses = append(ticketTravelerDetailDepartureResponses, ticketTravelerDetailResponse)
	}

	ticketOrderResponse = dtos.TicketOrderResponse{
		TicketOrderID:                 int(createTicketOrder.ID),
		QuantityAdult:                 createTicketOrder.QuantityAdult,
		QuantityInfant:                createTicketOrder.QuantityInfant,
		Price:                         createTicketOrder.Price,
		TotalAmount:                   createTicketOrder.TotalAmount,
		WithReturn:                    createTicketOrder.WithReturn,
		PaymentID:                     createTicketOrder.PaymentID,
		NameOrder:                     createTicketOrder.NameOrder,
		EmailOrder:                    createTicketOrder.EmailOrder,
		PhoneNumberOrder:              createTicketOrder.PhoneNumberOrder,
		TicketOrderCode:               createTicketOrder.TicketOrderCode,
		Status:                        createTicketOrder.Status,
		TravelerDetail:                travelerDetailResponses,
		TicketTravelerDetailDeparture: ticketTravelerDetailDepartureResponses,
		TicketTravelerDetailReturn:    ticketTravelerDetailReturnResponses,
	}

	return ticketOrderResponse, nil
}
