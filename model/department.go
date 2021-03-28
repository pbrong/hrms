package model

import (
	"gorm.io/gorm"
)

type DepartmentCreateDTO struct {
	DepDescribe string `json:"dep_describe" binding:"required"`
	DepName     string `json:"dep_name" binding:"required"`
}

type DepartmentEditDTO struct {
	DepId       string `json:"dep_id" binding:"required"`
	DepDescribe string `json:"dep_describe" binding:"required"`
	DepName     string `json:"dep_name" binding:"required"`
}

type Department struct {
	gorm.Model
	DepId       string `gorm:"column:dep_id" db:"column:dep_id" json:"dep_id"`
	DepDescribe string `gorm:"column:dep_describe" json:"dep_describe"`
	DepName     string `gorm:"column:dep_name" db:"column:dep_name" json:"dep_name"`
}

func (d *Department) AfterFind(tx *gorm.DB) (err error) {

	return nil
}
