package dtos

import "time"

type HistorySeenHotelInput struct {
	HotelID uint `json:"hotel_id" form:"hotel_id" example:"1"`
}

type HistorySeenHotelResponse struct {
	ID        uint            `json:"id" example:"1"`
	Hotel     HotelByIDSimply `json:"hotel"`
	CreatedAt time.Time       `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt time.Time       `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}
