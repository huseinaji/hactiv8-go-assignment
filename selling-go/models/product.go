package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name" form:"name" valid:"required~Product name is required"`
	Quantity uint   `gorm:"not null" json:"qty" form:"qty" valid:"required~qty is required"`
	Rate     uint   `gorm:"not null" json:"rate" form:"rate" valid:"required~rate is required"`
	Orders   []Order
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
