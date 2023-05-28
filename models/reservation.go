package models

import (
	"gorm.io/gorm"
)

type TypeReservation string

const (
	Hotels    TypeReservation = "hotels"
	Villa     TypeReservation = "villa"
	GustHouse TypeReservation = "guest_house"
)

type Status string

const (
	Avaliable   Status = "available"
	Unavaliable Status = "unavailable"
)

type Reservations struct {
	gorm.Model
	Name          string          `form:"name" json:"name"`
	Province_name string          `form:"province_name" json:"province_name"`
	Regency_name  string          `form:"regency_name" json:"regency_name"`
	District_name string          `form:"district_name" json:"district_name"`
	Village_name  string          `form:"village_name" json:"village_name"`
	Postal_code   string          `form:"postal_code" json:"postal_code"`
	Full_address  string          `form:"full_address" json:"full_address"`
	Type          TypeReservation `form:"type" json:"type" gorm:"type:ENUM('hotels', 'villa', 'guest_house')"`
	Price         int32           `form:"price" json:"price"`
	Thumbnail     string          `form:"thumbnail" json:"thumbnail"`
	Description   string          `form:"description" json:"description"`
	Tags          string          `form:"tags" json:"tags"`
	Status        Status          `form:"status" json:"status" gorm:"type:ENUM('available', 'unavailable')"`
	// Created_at    string          `form:"created_at" json:"created_at" gorm:"autoCreateTime"`
	// Updated_at    string          `form:"updated_at" json:"updated_at" gorm:"autoUpdateTime"`
}
