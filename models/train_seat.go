package models

import "gorm.io/gorm"

type TrainSeat struct {
	gorm.Model
	Class string `gorm:"type:varchar(255)"`
	Name  string
}
