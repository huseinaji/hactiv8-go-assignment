package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ProductID uint `gorm:"not null" json:"product_id" form:"product_id" valid:"required~ProductID is required"`
	UserID    uint `gorm:"not null" json:"user_id" form:"user_id" valid:"required~UserID is required"`
	Quantity  uint `gorm:"not null" json:"qty" form:"qty" valid:"required~qty is required"`
	Payments  []Payment
}

func (o *Order) BeforeSave(tx *gorm.DB) (err error) {
	_, errSave := govalidator.ValidateStruct(o)
	var product Product
	if errSave != nil {
		err = errSave
		return
	}

	tx.Model(&product).Where("ID = ? ", o.ProductID).Find(&product)
	err = nil
	return
}
