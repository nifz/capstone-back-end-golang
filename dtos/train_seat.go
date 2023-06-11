package dtos

type TrainSeatResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name" example:"A1"`
}

type TrainSeatAvailableResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name" example:"A1"`
	Available bool   `json:"available" example:"true"`
}

type TrainSeatResponseSimply struct {
	ID    int    `json:"id" example:"1"`
	Class string `json:"class" example:"Ekonomi"`
	Name  string `json:"name" example:"A1"`
}
