package models

import (
	"gorm.io/gorm"
)

type Station struct {
	gorm.Model
	Origin  string
	Name    string
	Initial string
}
