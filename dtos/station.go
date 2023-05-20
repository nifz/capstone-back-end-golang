package dtos

import "time"

type StationInput struct {
	Origin  string `form:"origin" json:"origin"`
	Name    string `form:"name" json:"name"`
	Initial string `form:"initial" json:"initial"`
}

type StationResponse struct {
	StationID uint      `form:"station_id" json:"station_id"`
	Origin    string    `form:"origin" json:"origin"`
	Name      string    `form:"name" json:"name"`
	Initial   string    `form:"initial" json:"initial"`
	UpdateAt  time.Time `json:"update_at"`
}
