package dtos

import "time"

type TrainStationDTO struct {
	ID        uint      `form:"id" json:"id"`
	Name      string    `form:"name" json:"name" binding:"required"`
	Initial   string    `form:"initial" json:"initial" binding:"required"`
	CreatedAt time.Time `form:"created_at" json:"created_at"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at"`
}

type TrainStationDTOsResponse struct {
	Message string            `form:"message" json:"message"`
	Data    []TrainStationDTO `form:"data" json:"data"`
}
