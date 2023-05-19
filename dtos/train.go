package dtos

import "time"

type TrainInput struct {
	StationOriginID      uint    `json:"station_origin_id" form:"station_origin_id"`
	StationDestinationID uint    `json:"station_destination_id" form:"station_destination_id"`
	DepartureTime        string  `json:"departure_time" form:"departure_time"`
	ArriveTime           string  `json:"arrive_time" form:"departure_time"`
	Name                 string  `json:"name" form:"name"`
	Route                string  `json:"route" form:"route"`
	Status               *string `json:"status" form:"status"`
}

type TrainResponse struct {
	TrainID              uint         `json:"train_id" form:"train_id"`
	StationOriginID      uint         `json:"station_origin_id" form:"station_origin_id"`
	StationOrigin        StationInput `json:"station_origin" form:"station_origin"`
	StationDestinationID uint         `json:"station_destination_id" form:"station_destination_id"`
	StationDestination   StationInput `json:"station_destination" form:"station_destination"`
	DepartureTime        string       `json:"departure_time" form:"departure_time"`
	ArriveTime           string       `json:"arrive_time" form:"departure_time"`
	Name                 string       `json:"name" form:"name"`
	Route                string       `json:"route" form:"route"`
	Status               *string      `json:"status" form:"status"`
	UpdateAt             time.Time    `json:"update_at"`
}
