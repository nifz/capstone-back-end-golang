package dtos

import "time"

type TrainInput struct {
	CodeTrain string              `json:"code_train" form:"code_train" example:"TRAIN001"`
	Name      string              `json:"name" form:"name" example:"Bengawan"`
	Route     []TrainStationInput `json:"route"`
	Status    string              `json:"status" form:"status" example:"available"`
}

type TrainResponse struct {
	TrainID   uint                   `json:"train_id" example:"1"`
	CodeTrain string                 `json:"code_train" example:"TRAIN001"`
	Name      string                 `json:"name" example:"Bengawan"`
	Class     string                 `json:"class" example:"Ekonomi"`
	Price     int                    `json:"price" example:"50000"`
	Route     []TrainStationResponse `json:"route"`
	Status    string                 `json:"status" example:"available"`
	CreatedAt time.Time              `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt time.Time              `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}
