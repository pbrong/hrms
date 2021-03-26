package model

import "gorm.io/gorm"

type Recruitment struct {
	gorm.Model
	RecruitmentId string `gorm:"column:recruitment_id" json:"recruitment_id"`
	JobName       string `gorm:"column:job_name" json:"job_name"`
	JobType       string `gorm:"column:job_type" json:"job_type"`
	BaseLocation  string `gorm:"column:base_location" json:"base_location"`
	BaseSalary    string `gorm:"column:base_salary" json:"base_salary"`
	EduLevel      string `gorm:"column:edu_level" json:"edu_level"`
	Experience    string `gorm:"column:experience" json:"experience"`
	Describe      string `gorm:"column:describe" json:"describe"`
	Email         string `gorm:"column:email" json:"email"`
}

type RecruitmentCreateDTO struct {
	JobName      string `gorm:"column:job_name" json:"job_name"  binding:"required"`
	JobType      string `gorm:"column:job_type" json:"job_type" binding:"required"`
	BaseLocation string `gorm:"column:base_location" json:"base_location" binding:"required"`
	BaseSalary   string `gorm:"column:base_salary" json:"base_salary" binding:"required"`
	EduLevel     string `gorm:"column:edu_level" json:"edu_level" binding:"required"`
	Experience   string `gorm:"column:experience" json:"experience" binding:"required"`
	Describe     string `gorm:"column:describe" json:"describe" binding:"required"`
	Email        string `gorm:"column:email" json:"email" binding:"required"`
}

type RecruitmentEditDTO struct {
	Id           int64  `binding:"required"`
	JobName      string `gorm:"column:job_name" json:"job_name"`
	JobType      string `gorm:"column:job_type" json:"job_type"`
	BaseLocation string `gorm:"column:base_location" json:"base_location"`
	BaseSalary   string `gorm:"column:base_salary" json:"base_salary"`
	EduLevel     string `gorm:"column:edu_level" json:"edu_level"`
	Experience   string `gorm:"column:experience" json:"experience"`
	Describe     string `gorm:"column:describe" json:"describe"`
	Email        string `gorm:"column:email" json:"email"`
}
