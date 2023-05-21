package dtos

import "time"

type TrainInput struct {
	StationOriginID      uint   `json:"station_origin_id" form:"station_origin_id" example:"1"`
	StationDestinationID uint   `json:"station_destination_id" form:"station_destination_id" example:"2"`
	DepartureTime        string `json:"departure_time" form:"departure_time" example:"05:00"`
	ArriveTime           string `json:"arrive_time" form:"arrive_time" example:"10:00"`
	Name                 string `json:"name" form:"name" example:"Bengawan"`
	Route                string `json:"route" form:"route" example:"Jakarta, Bekasi, Cikarang, Karawang, Purwokerto"`
	Status               string `json:"status" form:"status" example:"available"`
}

type TrainResponse struct {
	TrainID              uint         `json:"train_id" example:"1"`
	StationOriginID      uint         `json:"station_origin_id" example:"1"`
	StationOrigin        StationInput `json:"station_origin"`
	StationDestinationID uint         `json:"station_destination_id" example:"2"`
	StationDestination   StationInput `json:"station_destination"`
	DepartureTime        string       `json:"departure_time" example:"05:00"`
	ArriveTime           string       `json:"arrive_time" example:"10:00"`
	Name                 string       `json:"name" example:"Bengawan"`
	Route                string       `json:"route" example:"Jakarta, Bekasi, Cikarang, Karawang, Purwokerto"`
	Status               string       `json:"status" example:"available"`
	CreatedAt            time.Time    `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt            time.Time    `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}
