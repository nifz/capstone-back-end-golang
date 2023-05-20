package models

import (
	"gorm.io/gorm"
)

type Station struct {
	gorm.Model
	Origin  string `form:"origin" json:"origin"`
	Name    string `form:"name" json:"name"`
	Initial string `form:"initial" json:"initial"`
}
