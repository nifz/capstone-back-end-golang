package dtos

import "time"

type ArticleInput struct {
	Title       string `json:"title" form:"title" `
	Image       string `json:"image" form:"image"`
	Description string `json:"description" form:"description" `
	Label       string `json:"label" form:"label" `
}

type ArticleResponse struct {
	ArticleID   uint      `json:"article_id"`
	Title       string    `json:"title" `
	Image       string    `json:"image" `
	Description string    `json:"description" `
	Label       string    `json:"label" `
	Slug        string    `json:"slug" `
	CreatedAt   time.Time `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}
