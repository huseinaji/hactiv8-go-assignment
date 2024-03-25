package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserID    uint `gorm:"not null" json:"user_id" form:"user_id" valid:"required~user_id is required"`
	OrderID   uint `gorm:"not null" json:"order_id" form:"order_id" valid:"required~order_id is required"`
	InvoiceID uint `gorm:"not null" json:"invoice_id" form:"invoice_id" valid:"required~invoice_id is required"`
	Amount    uint
}
