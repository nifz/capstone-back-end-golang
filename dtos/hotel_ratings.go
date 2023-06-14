package dtos

type HotelRatingInput struct {
	HotelID uint   `form:"hotel_id" json:"hotel_id"`
	UserID  uint   `form:"user_id" json:"user_id"`
	Rating  int    `form:"rating" json:"rating"`
	Review  string `form:"review" json:"review"`
}

type HotelRatingResponse struct {
	HotelID uint   `form:"hotel_id" json:"hotel_id"`
	UserID  uint   `form:"user_id" json:"user_id"`
	Rating  int    `form:"rating" json:"rating"`
	Review  string `form:"review" json:"review"`
}

type HotelRatingsByIdHotels struct {
	HotelID        uint    `json:"hotel_id"`
	TotalRating    int     `json:"total_rating"`
	RataRataRating float64 `json:"rata_rata_rating"`
	// RatingCounts   map[int]int  `json:"rating_counts"`
	Rating5 int          `json:"rating_5"`
	Rating4 int          `json:"rating_4"`
	Rating3 int          `json:"rating_3"`
	Rating2 int          `json:"rating_2"`
	Rating1 int          `json:"rating_1"`
	Ratings []RatingInfo `json:"ratings"`
}

type RatingInfo struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	UserImage string `json:"user_image"`
	Rating    int    `json:"rating"`
	Review    string `json:"review"`
}
