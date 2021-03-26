package model

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	CandidateId string `gorm:"column:candidate_id" json:"candidate_id"`
	StaffId     string `gorm:"column:staff_id" json:"staff_id"`
	Name        string `gorm:"column:name" json:"name"`
	JobName     string `gorm:"column:job_name" json:"job_name"`
	EduLevel    string `gorm:"column:edu_level" json:"edu_level"`
	Major       string `gorm:"column:major" json:"major"`
	Experience  string `gorm:"column:experience" json:"experience"`
	Describe    string `gorm:"column:describe" json:"describe"`
	Email       string `gorm:"column:email" json:"email"`
	Evaluation  string `gorm:"column:evaluation" json:"evaluation"`
	Status      int64  `gorm:"column:status" json:"status"`
}

type CandidateCreateDTO struct {
	StaffId    string `gorm:"column:staff_id" json:"staff_id"`
	Name       string `gorm:"column:name" json:"name"`
	JobName    string `gorm:"column:job_name" json:"job_name"`
	EduLevel   string `gorm:"column:edu_level" json:"edu_level"`
	Major      string `gorm:"column:major" json:"major"`
	Experience string `gorm:"column:experience" json:"experience"`
	Describe   string `gorm:"column:describe" json:"describe"`
	Email      string `gorm:"column:email" json:"email"`
}

type CandidateEditDTO struct {
	Id         int64
	StaffId    string `gorm:"column:staff_id" json:"staff_id"`
	Name       string `gorm:"column:name" json:"name"`
	JobName    string `gorm:"column:job_name" json:"job_name"`
	EduLevel   string `gorm:"column:edu_level" json:"edu_level"`
	Major      string `gorm:"column:major" json:"major"`
	Experience string `gorm:"column:experience" json:"experience"`
	Describe   string `gorm:"column:describe" json:"describe"`
	Email      string `gorm:"column:email" json:"email"`
	Evaluation string `gorm:"column:evaluation" json:"evaluation"`
}
