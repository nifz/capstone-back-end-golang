package models

import "gorm.io/gorm"

type Train struct {
	gorm.Model
	StationOriginID      uint
	StationOrigin        Station `gorm:"foreignKey:StationOriginID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StationDestinationID uint
	StationDestination   Station `gorm:"foreignKey:StationDestinationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DepartureTime        string
	ArriveTime           string
	Name                 string
	Route                string
	Status               string `gorm:"type:ENUM('available','unavailable'); default:'unavailable'"`
}
