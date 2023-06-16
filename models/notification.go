package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	UserID        uint            `json:"user_id" form:"user_id"`
	User          User            `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TemplateID    uint            `json:"template_id" form:"template_id"`
	Template      TemplateMessage `gorm:"foreignKey:TemplateID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	HotelOrderID  uint            `json:"hotel_order_id" form:"hotel_order_id"`
	TicketOrderID uint            `json:"ticket_order_id" form:"ticket_order_id"`
}
