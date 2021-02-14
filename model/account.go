package model

import (
	"gorm.io/gorm"
)

type LoginDTO struct {
	UserNo int64 `json:"user_no" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
}

type Login struct {
	gorm.Model
	UserId int64 `gorm:"column:user_id" json:"user_id"`
	UserNo int64 `gorm:"column:user_no" json:"user_no"`
	UserPassword string `gorm:"column:user_password" json:"user_password"`
	Aval int64 `gorm:"column:aval" json:"aval"`
	UserType string `gorm:"column:user_type" json:"user_type"`
}

func (l Login) TableName() string {
	return "login"
}
