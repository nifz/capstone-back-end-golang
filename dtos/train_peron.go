package dtos

import "time"

type TrainPeronInput struct {
	TrainID uint    `json:"train_id" form:"train_id"`
	Class   string  `json:"class" form:"class"`
	Name    string  `json:"name" form:"name"`
	Price   int     `json:"price" form:"Price"`
	Status  *string `json:"status" form:"status"`
}

type TrainPeronResponse struct {
	TrainPeronID uint       `json:"train_peron_id" form:"train_peron_id"`
	TrainID      uint       `json:"train_id" form:"train_id"`
	Train        TrainInput `json:"train" form:"train"`
	Class        string     `json:"class" form:"class"`
	Name         string     `json:"name" form:"name"`
	Price        int        `json:"price" form:"Price"`
	Status       *string    `json:"status" form:"status"`
	UpdateAt     time.Time  `json:"update_at"`
}
