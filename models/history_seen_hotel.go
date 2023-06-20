package models

import "gorm.io/gorm"

type HistorySeenHotel struct {
	gorm.Model
	UserID  uint
	User    User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	HotelID uint
	Hotel   Hotel `gorm:"foreignKey:HotelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
