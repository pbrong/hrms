package model

import (
	"gorm.io/gorm"
	"time"
)

type Staff struct {
	gorm.Model
	StaffId       string    `gorm:"column:staff_id" json:"staff_id"`
	StaffName     string    `gorm:"column:staff_name" json:"staff_name"`
	LeaderStaffId string    `gorm:"column:leader_staff_id" json:"leader_staff_id"`
	LeaderName    string    `gorm:"column:leader_name" json:"leader_name"`
	Birthday      time.Time `gorm:"column:birthday" json:"birthday"`
	IdentityNum   string    `gorm:"column:identity_num" json:"identity_num"`
	Sex           int64     `gorm:"column:sex" json:"sex"`
	Nation        string    `gorm:"column:nation" json:"nation"`
	School        string    `gorm:"column:school" json:"school"`
	Major         string    `gorm:"column:major" json:"major"`
	EduLevel      string    `gorm:"column:edu_level" json:"edu_level"`
	BaseSalary    int64     `gorm:"column:base_salary" json:"base_salary"`
	CardNum       string    `gorm:"column:card_num" json:"card_num"`
	RankId        string    `gorm:"column:rank_id" json:"rank_id"`
	DepId         string    `gorm:"column:dep_id" json:"dep_id"`
	Email         string    `gorm:"column:email" json:"email"`
	Phone         int64     `gorm:"column:phone" json:"phone"`
	EntryDate     time.Time `gorm:"column:entry_date" json:"entry_date"`
}

type StaffVO struct {
	Staff
	DepName      string `json:"dep_name"`
	RankName     string `json:"rank_name"`
	UserTypeName string `json:"user_type_name"`
}

type StaffCreateDTO struct {
	StaffName     string `json:"staff_name" binding:"required"`
	LeaderStaffId string `gorm:"column:leader_staff_id" json:"leader_staff_id"`
	LeaderName    string `gorm:"column:leader_name" json:"leader_name"`
	BirthdayStr   string `json:"birthday_str" binding:"required"`
	IdentityNum   string `json:"identity_num" binding:"required"`
	SexStr        string `json:"sex_str" binding:"required""`
	Nation        string `json:"nation" binding:"required"`
	School        string `json:"school" binding:"required"`
	Major         string `json:"major" binding:"required"`
	EduLevel      string `json:"edu_level" binding:"required"`
	BaseSalary    int64  `json:"base_salary" binding:"required"`
	CardNum       string `json:"card_num" binding:"required"`
	RankId        string `json:"rank_id" binding:"required"`
	DepId         string `json:"dep_id" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Phone         int64  `gorm:"column:phone" json:"phone" binding:"required"`
	EntryDateStr  string `json:"entry_date_str" binding:"required"`
}

type StaffEditDTO struct {
	StaffId       string `json:"staff_id" binding:"required"`
	LeaderStaffId string `gorm:"column:leader_staff_id" json:"leader_staff_id"`
	LeaderName    string `gorm:"column:leader_name" json:"leader_name"`
	StaffName     string `json:"staff_name"`
	BirthdayStr   string `json:"birthday_str"`
	IdentityNum   string `json:"identity_num"`
	SexStr        string `json:"sex_str"`
	Sex           int64  `gorm:"column:sex" json:"sex"`
	Nation        string `json:"nation"`
	School        string `json:"school"`
	Major         string `json:"major"`
	EduLevel      string `json:"edu_level"`
	BaseSalary    int64  `json:"base_salary"`
	CardNum       string `json:"card_num"`
	RankId        string `json:"rank_id"`
	DepId         string `json:"dep_id"`
	Email         string `json:"email"`
	Phone         int64  `gorm:"column:phone" json:"phone"`
	EntryDateStr  string `json:"entry_date_str"`
}

func (s Staff) TableName() string {
	return "staff"
}
