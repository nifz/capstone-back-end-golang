package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type TicketOrderRepository interface {
	GetAllTicketOrders(page, limit int) ([]models.TicketOrder, int, error)
	GetTicketOrderByID(id, userID uint) (models.TicketOrder, error)
	CreateTicketOrder(ticketOrder models.TicketOrder) (models.TicketOrder, error)
	UpdateTicketOrder(ticketOrder models.TicketOrder) (models.TicketOrder, error)
}

type ticketOrderRepository struct {
	db *gorm.DB
}

func NewTicketOrderRepository(db *gorm.DB) TicketOrderRepository {
	return &ticketOrderRepository{db}
}

func (r *ticketOrderRepository) GetAllTicketOrders(page, limit int) ([]models.TicketOrder, int, error) {
	var (
		ticketOrders []models.TicketOrder
		count        int64
	)
	err := r.db.Find(&ticketOrders).Count(&count).Error
	if err != nil {
		return ticketOrders, int(count), err
	}

	offset := (page - 1) * limit

	err = r.db.Limit(limit).Offset(offset).Find(&ticketOrders).Error

	return ticketOrders, int(count), err
}

func (r *ticketOrderRepository) GetTicketOrderByID(id, userID uint) (models.TicketOrder, error) {
	var ticketOrder models.TicketOrder
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&ticketOrder).Error
	return ticketOrder, err
}

func (r *ticketOrderRepository) CreateTicketOrder(ticketOrder models.TicketOrder) (models.TicketOrder, error) {
	err := r.db.Create(&ticketOrder).Error
	return ticketOrder, err
}

func (r *ticketOrderRepository) UpdateTicketOrder(ticketOrder models.TicketOrder) (models.TicketOrder, error) {
	err := r.db.Save(ticketOrder).Error
	return ticketOrder, err
}
