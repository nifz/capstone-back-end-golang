package models

import "gorm.io/gorm"

type HotelRating struct {
	gorm.Model
	HotelID uint   `form:"hotel_id" json:"hotel_id"`
	Hotel   Hotel  `gorm:"foreignKey:HotelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID  uint   `form:"user_id" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rating  int    `form:"rating" json:"rating"`
	Review  string `form:"review" json:"review"`
}
