package models

import (
	"gorm.io/gorm"
)

type Station struct {
	gorm.Model
	Name    string `form:"name" json:"name"`
	Initial string `form:"initial" json:"initial"`
}
