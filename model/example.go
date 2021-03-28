package model

import (
	"gorm.io/gorm"
)

type Example struct {
	gorm.Model
	ExampleId string `gorm:"column:example_id" json:"example_id"`
	Name      string `gorm:"column:name" json:"name"`
	Describe  string `gorm:"column:describe" json:"describe"`
	Date      string `gorm:"column:date" json:"date"`
	Limit     int64  `gorm:"column:limit" json:"limit"`
	Content   string `gorm:"column:content" json:"content"`
}

type ExampleCreateDTO struct {
	Name     string `gorm:"column:name" json:"name"`
	Date     string `gorm:"column:date" json:"date"`
	Describe string `gorm:"column:describe" json:"describe"`
	Limit    int64  `gorm:"column:limit" json:"limit"`
	Content  string `gorm:"column:content" json:"content"`
}

type ExampleEditDTO struct {
	Id       int64  `binding:"required"`
	Name     string `gorm:"column:name" json:"name"`
	Date     string `gorm:"column:date" json:"date"`
	Describe string `gorm:"column:describe" json:"describe"`
	Limit    int64  `gorm:"column:limit" json:"limit"`
}

type ExampleItem struct {
	Num   int
	Title string
	Items []string
	Ans   string
}

type ExampleScore struct {
	gorm.Model
	ExampleId string `gorm:"column:example_id" json:"example_id"`
	StaffId   string `gorm:"column:staff_id" json:"staff_id"`
	StaffName string `gorm:"column:staff_name" json:"staff_name"`
	Name      string `gorm:"column:name" json:"name"`
	Date      string `gorm:"column:date" json:"date"`
	Content   string `gorm:"column:content" json:"content"`
	Commit    string `gorm:"column:commit" json:"commit"`
	Score     int64  `gorm:"column:score" json:"score"`
}

type ExampleScoreCreateDTO struct {
	ExampleId string `gorm:"column:example_id" json:"example_id"  binding:"required"`
	StaffId   string `gorm:"column:staff_id" json:"staff_id" binding:"required"`
	StaffName string `gorm:"column:staff_name" json:"staff_name"`
	Name      string `gorm:"column:name" json:"name"`
	Date      string `gorm:"column:date" json:"date"`
	Content   string `gorm:"column:content" json:"content"  binding:"required"`
	Commit    string `gorm:"column:commit" json:"commit" binding:"required"`
}
