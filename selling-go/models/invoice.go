package models

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model
	OrderID      uint `gorm:"not null" json:"order_id" form:"order_id" valid:"required~OrderID name is required"`
	UserID       uint `gorm:"not null" json:"user_id" form:"user_id" valid:"required~UserID name is required"`
	Quantity     uint `gorm:"not null" json:"qty" form:"qty" valid:"required~Quantity name is required"`
	Rate         uint
	TotalPayment uint
	Status       string
	Payments     []Payment
}
