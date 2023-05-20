package models

import "gorm.io/gorm"

type TrainPeron struct {
	gorm.Model
	TrainID uint
	Train   Train  `gorm:"foreignKey:TrainID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Class   string `gorm:"type:varchar(255)"`
	Name    string
	Price   int
	Status  string `gorm:"type:ENUM('available','unavailable');default:'unavailable'"`
}
