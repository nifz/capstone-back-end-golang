package models

import "gorm.io/gorm"

type Train struct {
	gorm.Model
	StationOriginID      uint    `json:"station_origin_id" form:"station_origin_id"`
	StationOrigin        Station `gorm:"foreignKey:StationOriginID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"station_origin" form:"station_origin"`
	StationDestinationID uint    `json:"station_destination_id" form:"station_destination_id"`
	StationDestination   Station `gorm:"foreignKey:StationDestinationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"station_destination" form:"station_destination"`
	DepartureTime        string  `json:"departure_time" form:"departure_time"`
	ArriveTime           string  `json:"arrive_time" form:"arrive_time"`
	Name                 string  `json:"name" form:"name"`
	Route                string  `json:"route" form:"route"`
	Status               *string `gorm:"type:ENUM('available','unavailable');null" json:"status" form:"status"`
}
