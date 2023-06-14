package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type HotelOrderRepository interface {
	GetHotelOrders(page, limit int, userID uint, status string) ([]models.HotelOrder, int, error)
	GetHotelOrderByStatusAndID(id, userID uint, status string) (models.HotelOrder, error)
	GetHotelOrderByID(id, userID uint) (models.HotelOrder, error)
	CreateHotelOrder(hotelOrder models.HotelOrder) (models.HotelOrder, error)
	UpdateHotelOrder(hotelOrder models.HotelOrder) (models.HotelOrder, error)
	DeleteHotelOrder(hotelOrder models.HotelOrder) (models.HotelOrder, error)
}

type hotelOrderRepository struct {
	db *gorm.DB
}

func NewHotelOrderRepository(db *gorm.DB) HotelOrderRepository {
	return &hotelOrderRepository{db}
}

func (r *hotelOrderRepository) GetHotelOrders(page, limit int, userID uint, status string) ([]models.HotelOrder, int, error) {
	var (
		hotelOrders []models.HotelOrder
		count       int64
		err         error
	)
	if userID == 1 {
		if status == "" {
			err = r.db.Find(&hotelOrders).Count(&count).Error
		} else {
			err = r.db.Where("status = ?", status).Find(&hotelOrders).Count(&count).Error
		}
	} else {
		if status == "" {
			err = r.db.Where("user_id = ?", userID).Find(&hotelOrders).Count(&count).Error
		} else {
			err = r.db.Where("user_id = ? AND status = ?", userID, status).Find(&hotelOrders).Count(&count).Error
		}
	}
	if err != nil {
		return hotelOrders, int(count), err
	}

	offset := (page - 1) * limit

	if userID == 1 {
		if status == "" {
			err = r.db.Find(&hotelOrders).Count(&count).Error
		} else {
			err = r.db.Where("status = ?", status).Limit(limit).Offset(offset).Find(&hotelOrders).Error
		}
	} else {
		if status == "" {
			err = r.db.Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&hotelOrders).Error
		} else {
			err = r.db.Where("user_id = ? AND status = ?", userID, status).Limit(limit).Offset(offset).Find(&hotelOrders).Error
		}
	}

	return hotelOrders, int(count), err
}

func (r *hotelOrderRepository) GetHotelOrderByStatusAndID(id, userID uint, status string) (models.HotelOrder, error) {
	var hotelOrder models.HotelOrder
	if userID == 1 {
		err := r.db.Where("id = ? AND status = ?", id, status).First(&hotelOrder).Error
		return hotelOrder, err
	}
	err := r.db.Where("id = ? AND user_id = ? AND status = ?", id, userID, status).First(&hotelOrder).Error
	return hotelOrder, err
}

func (r *hotelOrderRepository) GetHotelOrderByID(id, userID uint) (models.HotelOrder, error) {
	var hotelOrder models.HotelOrder
	if userID == 1 {
		err := r.db.Where("id = ?", id).First(&hotelOrder).Error
		return hotelOrder, err
	}
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&hotelOrder).Error
	return hotelOrder, err
}

func (r *hotelOrderRepository) CreateHotelOrder(hotelOrder models.HotelOrder) (models.HotelOrder, error) {
	err := r.db.Create(&hotelOrder).Error
	return hotelOrder, err
}

func (r *hotelOrderRepository) UpdateHotelOrder(hotelOrder models.HotelOrder) (models.HotelOrder, error) {
	err := r.db.Save(hotelOrder).Error
	return hotelOrder, err
}

func (r *hotelOrderRepository) DeleteHotelOrder(hotelOrder models.HotelOrder) (models.HotelOrder, error) {
	err := r.db.Unscoped().Delete(&hotelOrder).Error
	return hotelOrder, err
}
