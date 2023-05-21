package dtos

import "back-end-golang/helpers"

type UserStatusOKResponse struct {
	StatusCode int                     `json:"status_code" example:"200"`
	Message    string                  `json:"message" example:"Successfully get user credentials"`
	Data       UserInformationResponse `json:"data"`
}

type UserCreeatedResponse struct {
	StatusCode int                     `json:"status_code" example:"201"`
	Message    string                  `json:"message" example:"Successfully registered"`
	Data       UserInformationResponse `json:"data"`
}

type GetAllStationStatusOKResponse struct {
	StatusCode int             `json:"status_code" example:"200"`
	Message    string          `json:"message" example:"Successfully get station"`
	Data       StationResponse `json:"data"`
	Meta       helpers.Meta    `json:"meta"`
}

type StationStatusOKResponse struct {
	StatusCode int             `json:"status_code" example:"200"`
	Message    string          `json:"message" example:"Successfully get station"`
	Data       StationResponse `json:"data"`
}

type StationCreeatedResponse struct {
	StatusCode int             `json:"status_code" example:"201"`
	Message    string          `json:"message" example:"Successfully created station"`
	Data       StationResponse `json:"data"`
}

type GetAllTrainStatusOKResponse struct {
	StatusCode int           `json:"status_code" example:"200"`
	Message    string        `json:"message" example:"Successfully get train"`
	Data       TrainResponse `json:"data"`
	Meta       helpers.Meta  `json:"meta"`
}

type TrainStatusOKResponse struct {
	StatusCode int           `json:"status_code" example:"200"`
	Message    string        `json:"message" example:"Successfully get train"`
	Data       TrainResponse `json:"data"`
}

type TrainCreeatedResponse struct {
	StatusCode int           `json:"status_code" example:"201"`
	Message    string        `json:"message" example:"Successfully created train"`
	Data       TrainResponse `json:"data"`
}

type GetAllTrainPeronStatusOKResponse struct {
	StatusCode int                `json:"status_code" example:"200"`
	Message    string             `json:"message" example:"Successfully get train peron"`
	Data       TrainPeronResponse `json:"data"`
	Meta       helpers.Meta       `json:"meta"`
}

type TrainPeronStatusOKResponse struct {
	StatusCode int                `json:"status_code" example:"200"`
	Message    string             `json:"message" example:"Successfully get train peron"`
	Data       TrainPeronResponse `json:"data"`
}

type TrainPeronCreeatedResponse struct {
	StatusCode int                `json:"status_code" example:"201"`
	Message    string             `json:"message" example:"Successfully created train peron"`
	Data       TrainPeronResponse `json:"data"`
}

type StatusOKDeletedResponse struct {
	StatusCode int         `json:"status_code" example:"200"`
	Message    string      `json:"message" example:"Successfully deleted"`
	Errors     interface{} `json:"errors"`
}

type BadRequestResponse struct {
	StatusCode int         `json:"status_code" example:"400"`
	Message    string      `json:"message" example:"Bad Request"`
	Errors     interface{} `json:"errors"`
}

type UnauthorizedResponse struct {
	StatusCode int         `json:"status_code" example:"401"`
	Message    string      `json:"message" example:"Unauthorized"`
	Errors     interface{} `json:"errors"`
}

type ForbiddenResponse struct {
	StatusCode int         `json:"status_code" example:"403"`
	Message    string      `json:"message" example:"Forbidden"`
	Errors     interface{} `json:"errors"`
}

type NotFoundResponse struct {
	StatusCode int         `json:"status_code" example:"404"`
	Message    string      `json:"message" example:"Not Found"`
	Errors     interface{} `json:"errors"`
}

type InternalServerErrorResponse struct {
	StatusCode int         `json:"status_code" example:"500"`
	Message    string      `json:"message" example:"Internal Server Error"`
	Errors     interface{} `json:"errors"`
}
