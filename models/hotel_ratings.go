package models

import "gorm.io/gorm"

type HotelRating struct {
	gorm.Model
	HotelOrderID uint       `form:"hotel_order_id" json:"hotel_order_id"`
	HotelOrder   HotelOrder `gorm:"foreignKey:HotelOrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	HotelID      uint       `form:"hotel_id" json:"hotel_id"`
	Hotel        Hotel      `gorm:"foreignKey:HotelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID       uint       `form:"user_id" json:"user_id"`
	User         User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rating       int        `form:"rating" json:"rating"`
	Review       string     `form:"review" json:"review"`
}
