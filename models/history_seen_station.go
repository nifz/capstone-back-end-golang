package models

import "gorm.io/gorm"

type HistorySeenStation struct {
	gorm.Model
	UserID               uint
	User                 User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StationOriginID      uint
	StationOrigin        Station `gorm:"foreignKey:StationOriginID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StationDestinationID uint
	StationDestination   Station `gorm:"foreignKey:StationDestinationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
