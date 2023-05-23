package dtos

type ReservationCreateInput struct {
	Name          string `form:"name" json:"name" binding:"required"`
	Province_name string `form:"province_name" json:"province_name" binding:"required"`
	Regency_name  string `form:"regency_name" json:"regency_name" binding:"required"`
	District_name string `form:"district_name" json:"district_name" binding:"required"`
	Village_name  string `form:"village_name" json:"village_name" binding:"required"`
	Postal_code   string `form:"postal_code" json:"postal_code" binding:"required"`
	Full_address  string `form:"full_address" json:"full_address" binding:"required"`
	Type          string `form:"type" json:"type" binding:"required"`
	Price         int32  `form:"price" json:"price" binding:"required"`
	Thumbnail     string `form:"thumbnail" json:"thumbnail" binding:"required"`
	Description   string `form:"description" json:"description" binding:"required"`
	Tags          string `form:"tags" json:"tags" binding:"required"`
	Status        string `form:"status" json:"status" binding:"required"`
	ImageFile     string `form:"image" json:"image"`
}

type ReservationCreateResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Province_name string `json:"province_name"`
	Regency_name  string `json:"regency_name"`
	District_name string `json:"district_name"`
	Village_name  string `json:"village_name"`
	Postal_code   string `json:"postal_code"`
	Full_address  string `json:"full_address"`
	Type          string `json:"type"`
	Price         int32  `json:"price"`
	Thumbnail     string `json:"thumbnail"`
	Description   string `json:"description"`
	Tags          string `json:"tags"`
	Status        string `json:"status"`
	Image         string `json:"image"`
}

type ReservationResponse struct {
	ReservationID uint   `json:"reservation_id"`
	Name          string `json:"name"`
	Province_name string `json:"province_name"`
	Regency_name  string `json:"regency_name"`
	District_name string `json:"district_name"`
	Village_name  string `json:"village_name"`
	Postal_code   string `json:"postal_code"`
	Full_address  string `json:"full_address"`
	Type          string `json:"type"`
	Price         int32  `json:"price"`
	Thumbnail     string `json:"thumbnail"`
	Description   string `json:"description"`
	Tags          string `json:"tags"`
	Status        string `json:"status"`
	Image         string `json:"image"`
}
