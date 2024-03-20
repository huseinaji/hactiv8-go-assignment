package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string `gorm:"not null;type:varchar(100)"`
	OrderedAt    string `gorm:"not null;type:varchar(100)"`
	Items        []Item
}
