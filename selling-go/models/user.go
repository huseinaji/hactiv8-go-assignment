package models

import (
	"selling-go/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"not null" json:"user_name" form:"user_name" valid:"required~Your user name is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have minimum length of 6 characters"`
	Orders   []Order
	Invoices []Invoice
	Payments []Payment
	Role     string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	User := User{}

	u.Password = helpers.HashPass(u.Password)

	//cek jka belum ada admin maka dicreate user dengan role admin
	err = tx.Where("role = ?", "admin").Take(&User).Error

	if err != nil {
		tx.Model(u).Update("role", "admin")
	} else {
		tx.Model(u).Update("role", "customer")
	}

	err = nil
	return
}
