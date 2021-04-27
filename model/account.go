package model

import (
	"gorm.io/gorm"
)

type LoginDTO struct {
	UserNo       string `json:"staff_id" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
	BranchId     string `json:"branch_id" binding:"required"`
}

type Authority struct {
	gorm.Model
	AuthorityId  string `gorm:"column:authority_id" json:"authority_id"`
	StaffId      string `gorm:"column:staff_id" json:"staff_id"`
	UserPassword string `gorm:"column:user_password" json:"user_password"`
	//Aval         int64  `gorm:"column:aval" json:"aval"`
	UserType string `gorm:"column:user_type" json:"user_type"`
}

type PasswordQueryVO struct {
	Id        int64  `json:"id"`
	StaffId   string `json:"staff_id"`
	StaffName string `json:"staff_name"`
	Password  string `json:"password"`
}

type PasswordEditDTO struct {
	StaffId  string `json:"staff_id"`
	Password string `json:"password"`
}
