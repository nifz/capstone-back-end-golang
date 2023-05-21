package models

import (
	"gorm.io/gorm"
)

type TravelerDetails struct {
	gorm.Model
	UserId         uint   `form:"user_id" json:"user_id"`
	User           User   `gorm:"foreignKey:UserId"`
	Full_name      string `form:"full_name" json:"full_name"`
	Id_card_number string `form:"id_card_number" json:"id_card_number"`
}
