package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
	"errors"
)

type HistorySearchUseCase interface {
	HistorySearchGetById(id uint) (models.HistorySearch, error)
	HistorySearchCreate(input dtos.HistorySearchCreateInput) (dtos.HistorySearchCreateResponse, error)
	HistorySearchUpdate(userId uint, input dtos.HistorySearchUpdateInput) (dtos.HistorySearchUpdateResponse, error)
	HistorySearchDelete(id uint) (models.HistorySearch, error)
}

type historySearchUsecase struct {
	historySearchRepository repositories.HistorySearchRepository
	userRepo                repositories.UserRepository
}

func NewHistorySearchUsecase(historySearchRepository repositories.HistorySearchRepository, userRepo repositories.UserRepository) HistorySearchUseCase {
	return &historySearchUsecase{historySearchRepository, userRepo}
}

func (u *historySearchUsecase) HistorySearchGetById(id uint) (models.HistorySearch, error) {
	var history models.HistorySearch
	history, err := u.historySearchRepository.HistorySearchGetById(id)
	if err != nil {
		return history, err
	}

	return history, nil
}

func (u *historySearchUsecase) HistorySearchCreate(input dtos.HistorySearchCreateInput) (dtos.HistorySearchCreateResponse, error) {
	var (
		history         models.HistorySearch
		historyResponse dtos.HistorySearchCreateResponse
	)

	history, err := u.historySearchRepository.HistorySearchCreate(history)
	if err != nil {
		return historyResponse, err
	}

	historyResponse.UserID = history.UserID
	historyResponse.Name = history.Name

	return historyResponse, nil
}

func (u *historySearchUsecase) HistorySearchUpdate(userId uint, input dtos.HistorySearchUpdateInput) (dtos.HistorySearchUpdateResponse, error) {
	var (
		history         models.HistorySearch
		historyResponse dtos.HistorySearchUpdateResponse
		user            models.User
	)

	user, err := u.userRepo.UserGetById(userId)
	if err != nil {
		return historyResponse, errors.New("User not found")
	}

	history, err = u.historySearchRepository.HistorySearchUpdate(history)
	if err != nil {
		return historyResponse, err
	}

	historyResponse.UserID = user.ID
	historyResponse.Name = history.Name

	return historyResponse, nil
}

func (u *historySearchUsecase) HistorySearchDelete(id uint) (models.HistorySearch, error) {
	var history models.HistorySearch

	history, err := u.historySearchRepository.HistorySearchGetById(id)
	if err != nil {
		return history, err
	}

	history, err = u.historySearchRepository.HistorySearchUpdate(history)
	if err != nil {
		return history, err
	}

	return history, nil
}
