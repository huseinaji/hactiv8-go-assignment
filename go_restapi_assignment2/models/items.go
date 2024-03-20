package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ItemCode    string `gorm:"not null;type:varchar(100)"`
	Description string `gorm:"not null;type:varchar(100)"`
	Quantity    int
	OrderID     uint
}
