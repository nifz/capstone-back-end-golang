package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       string    `form:"full_name" json:"full_name"`
	Email          string    `gorm:"unique" form:"email" json:"email"`
	Password       string    `form:"password" json:"password"`
	PhoneNumber    string    `form:"phone_number" json:"phone_number"`
	Gender         string    `form:"gender" json:"gender"`
	BirthDate      time.Time `form:"birth_date" json:"birth_date"`
	ProfilePicture string    `form:"profile_picture" json:"profile_picture"`
	Role           string    `form:"role" json:"role"`
}
