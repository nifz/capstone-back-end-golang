package models

import "gorm.io/gorm"

type TrainPeron struct {
	gorm.Model
	TrainID uint    `json:"train_id" form:"train_id"`
	Train   Train   `gorm:"foreignKey:TrainID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"train" form:"train"`
	Class   string  `json:"class" form:"class"`
	Name    string  `json:"name" form:"name"`
	Price   int     `json:"price" form:"Price"`
	Status  *string `gorm:"type:ENUM('available','unavailable');null" json:"status" form:"status"`
}
