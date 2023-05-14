package dtos

import "time"

type UserRegisterInput struct {
	FullName       string `form:"full_name" json:"full_name" binding:"required"`
	Email          string `form:"email" json:"email" binding:"required,email"`
	Password       string `form:"password" json:"password" binding:"required"`
	PhoneNumber    string `form:"phone_number" json:"phone_number" binding:"required,min=11"`
	Gender         string `form:"gender" json:"gender" binding:"required"`
	BirthDate      string `form:"birth_date" json:"birth_date" binding:"required"`
	ProfilePicture string `form:"profile_picture" json:"profile_picture" binding:"required"`
	Role           string `form:"role" json:"role"`
}

type UserLoginInput struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

type UserRegisterResponse struct {
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	Gender         string    `json:"gender"`
	BirthDate      time.Time `json:"birth_date"`
	ProfilePicture string    `json:"profile_picture"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
