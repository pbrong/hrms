package model

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type DepartmentCreateDTO struct {
	PreDepId string `json:"pre_dep_id" binding:"required"`
	DepName  string `json:"dep_name" binding:"required"`
}

type DepartmentEditDTO struct {
	DepId    string `json:"dep_id" binding:"required"`
	PreDepId string `json:"pre_dep_id" binding:"required"`
	DepName  string `json:"dep_name" binding:"required"`
}

type Department struct {
	gorm.Model
	DepId    string `gorm:"column:dep_id" db:"column:dep_id" json:"dep_id" form:"dep_id"`
	PreDepId string `gorm:"column:pre_dep_id" db:"column:pre_dep_id" json:"pre_dep_id" form:"pre_dep_id"`
	DepName  string `gorm:"column:dep_name" db:"column:dep_name" json:"dep_name" form:"dep_name"`
}

func CreateDepId() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("dep_%v", rand.Uint32())
}

func (d Department) TableName() string {
	return "department"
}
