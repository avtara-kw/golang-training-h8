package models

import (
	"sesi10/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string `gorm:"not null,default:''" json:"full_name" form:"full_name" valid:"required~Your fullname is required"`
	Email    string `gorm:"not null,default:''" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null,default:''" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`

	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, errCreate := govalidator.ValidateStruct(u)

	u.Password, _ = helper.HashPass(u.Password)
	return errCreate
}
