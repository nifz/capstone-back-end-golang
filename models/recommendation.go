package models

import "gorm.io/gorm"

type Recomendation struct {
	gorm.Model
	Tag string `json:"tag" form:"tag"`
}
