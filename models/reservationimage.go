package models

import (
	"gorm.io/gorm"
)

type ReservationImages struct {
	gorm.Model
	ReservationId uint         `form:"reservation_id" json:"reservation_id"`
	Reservations  Reservations `gorm:"foreignKey:ReservationId"`
	Image         string       `form:"image" json:"image"`
}