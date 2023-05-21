package dtos

import "time"

type TrainPeronInput struct {
	TrainID uint   `json:"train_id" form:"train_id" example:"1"`
	Class   string `json:"class" form:"class" example:"Ekonomi"`
	Name    string `json:"name" form:"name" example:"Gerbong 1"`
	Price   int    `json:"price" form:"Price" example:"50000"`
	Status  string `json:"status" form:"status" example:"available"`
}

type TrainPeronResponse struct {
	TrainPeronID uint                `json:"train_peron_id" example:"1"`
	Train        TrainResponse       `json:"train"`
	Class        string              `json:"class" example:"Ekonomi"`
	Name         string              `json:"name" example:"Gerbong 1"`
	Seat         []TrainSeatResponse `json:"seat"`
	Price        int                 `json:"price" example:"50000"`
	Status       string              `json:"status" example:"available"`
	CreatedAt    time.Time           `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt    time.Time           `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}
