package models

import "gorm.io/gorm"

type Recommendation struct {
	gorm.Model
	Tag string `json:"tag" form:"tag"`
}
