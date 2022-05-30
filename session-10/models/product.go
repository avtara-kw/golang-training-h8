package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `json:"title" form:"title" valid:"required~Title is required"`
	Description string `json:"description" form:"description" valid:"required~Title is required"`
	UserID      uint

	User *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	return err
}

func (p *Product) BeforeUpdate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
