package model

import "gorm.io/gorm"

type Rank struct {
	gorm.Model
	RankId   string `gorm:"column:rank_id" json:"rank_id"`
	RankName string `gorm:"column:rank_name" json:"rank_name"`
}

type RankCreateDTO struct {
	RankName string `json:"rank_name" binding:"required"`
}

type RankEditDTO struct {
	RankId   string `json:"rank_id" binding:"required"`
	RankName string `json:"rank_name" binding:"required"`
}

func (d Rank) TableName() string {
	return "rank"
}
