package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type TicketTravelerDetailRepository interface {
	GetAllTicketTravelerDetails(page, limit int) ([]models.TicketTravelerDetail, int, error)
	GetTicketTravelerDetailByID(id uint) (models.TicketTravelerDetail, error)
	GetTicketTravelerDetailByTrainID(id uint) ([]models.TicketTravelerDetail, error)
	CreateTicketTravelerDetail(ticketTravelerDetail models.TicketTravelerDetail) (models.TicketTravelerDetail, error)
	UpdateTicketTravelerDetail(ticketTravelerDetail models.TicketTravelerDetail) (models.TicketTravelerDetail, error)
}

type ticketTravelerDetailRepository struct {
	db *gorm.DB
}

func NewTicketTravelerDetailRepository(db *gorm.DB) TicketTravelerDetailRepository {
	return &ticketTravelerDetailRepository{db}
}

func (r *ticketTravelerDetailRepository) GetAllTicketTravelerDetails(page, limit int) ([]models.TicketTravelerDetail, int, error) {
	var (
		ticketTravelerDetails []models.TicketTravelerDetail
		count                 int64
	)
	err := r.db.Find(&ticketTravelerDetails).Count(&count).Error
	if err != nil {
		return ticketTravelerDetails, int(count), err
	}

	offset := (page - 1) * limit

	err = r.db.Limit(limit).Offset(offset).Find(&ticketTravelerDetails).Error

	return ticketTravelerDetails, int(count), err
}

func (r *ticketTravelerDetailRepository) GetTicketTravelerDetailByID(id uint) (models.TicketTravelerDetail, error) {
	var ticketTravelerDetail models.TicketTravelerDetail
	err := r.db.Where("id = ?", id).First(&ticketTravelerDetail).Error
	return ticketTravelerDetail, err
}

func (r *ticketTravelerDetailRepository) GetTicketTravelerDetailByTrainID(id uint) ([]models.TicketTravelerDetail, error) {
	var ticketTravelerDetail []models.TicketTravelerDetail
	err := r.db.Where("id = ?", id).Find(&ticketTravelerDetail).Error
	return ticketTravelerDetail, err
}

func (r *ticketTravelerDetailRepository) CreateTicketTravelerDetail(ticketTravelerDetail models.TicketTravelerDetail) (models.TicketTravelerDetail, error) {
	err := r.db.Create(&ticketTravelerDetail).Error
	return ticketTravelerDetail, err
}

func (r *ticketTravelerDetailRepository) UpdateTicketTravelerDetail(ticketTravelerDetail models.TicketTravelerDetail) (models.TicketTravelerDetail, error) {
	err := r.db.Save(ticketTravelerDetail).Error
	return ticketTravelerDetail, err
}
