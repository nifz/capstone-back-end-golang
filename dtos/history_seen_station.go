package dtos

import "time"

type HistorySeenStationInput struct {
	StationOriginID      uint `json:"station_origin_id" form:"station_origin_id" example:"1"`
	StationDestinationID uint `json:"station_destination_id" form:"station_destination_id" example:"2"`
}

type HistorySeenStationResponse struct {
	ID                 uint                   `json:"id" example:"1"`
	StationOrigin      StationResponseSimply2 `json:"station_origin"`
	StationDestination StationResponseSimply2 `json:"station_destination"`
	CreatedAt          time.Time              `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt          time.Time              `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}
