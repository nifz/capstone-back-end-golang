package models

import (
	"gorm.io/gorm"
)

type statusCheckout string

const (
	Process       statusCheckout = "process"
	Pending       statusCheckout = "pending"
	Success       statusCheckout = "success"
	Canceled      statusCheckout = "canceled"
	Return        statusCheckout = "return"
	RefundSuccess statusCheckout = "refund_success"
)

type ReservationCheckouts struct {
	gorm.Model
	ReservationId     uint            `form:"reservation_id" json:"reservation_id"`
	Reservations      Reservations    `gorm:"foreignKey:ReservationId"`
	Checkin_date      string          `form:"checkin_date" json:"checkin_date"`
	Checkout_date     string          `form:"checkout_date" json:"checkout_date"`
	Qunatity          int32           `form:"qunatity" json:"qunatity"`
	TravelerDetailsId uint            `form:"traveler_details_id" json:"traveler_details_id"`
	TravelerDetails   TravelerDetails `gorm:"foreignKey:TravelerDetailsId"`
	StatusCheckout    statusCheckout  `form:"status_checkout" json:"status_checkout" gorm:"type:ENUM('process', 'pending', 'success', 'canceled', 'return', 'refund_success')"`
}
