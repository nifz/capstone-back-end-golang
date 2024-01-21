package models

import (
	"time"

	"gorm.io/gorm"
)

type HotelOrder struct {
	gorm.Model
	UserID           uint      `form:"user_id" json:"user_id"`
	User             User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	HotelID          uint      `form:"hotel_id" json:"hotel_id"`
	Hotel            Hotel     `gorm:"foreignKey:HotelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	HotelRoomID      uint      `form:"hotel_room_id" json:"hotel_room_id"`
	HotelRoom        HotelRoom `gorm:"foreignKey:HotelRoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	QuantityAdult    int
	QuantityInfant   int
	NumberOfNight    int
	DateStart        time.Time `gorm:"type:DATE"`
	DateEnd          time.Time `gorm:"type:DATE"`
	Price            int
	PaymentID        int
	TotalAmount      int
	NameOrder        string
	EmailOrder       string
	PhoneNumberOrder string
	SpecialRequest   string
	HotelOrderCode   string
	PaymentURL       string
	IsCheckIn        bool   `gorm:"default:false"`
	IsCheckOut       bool   `gorm:"default:false"`
	Status           string `gorm:"type:ENUM('unpaid', 'paid', 'done', 'canceled', 'refund')"`
}
