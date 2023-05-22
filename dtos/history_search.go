package dtos

type HistorySearchCreateInput struct {
	UserID uint   `form:"user_id" json:"user_id" binding:"required"`
	Name   string `form:"name" json:"name" binding:"required"`
}

type HistorySearchCreateResponse struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

type HistorySearchUpdateInput struct {
	UserID uint   `form:"user_id" json:"user_id" binding:"required"`
	Name   string `form:"name" json:"name" binding:"required"`
}

type HistorySearchUpdateResponse struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}
