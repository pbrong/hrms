package model

import (
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	gorm.Model
	NoticeId      string    `gorm:"column:notice_id" json:"notice_id"`
	NoticeTitle   string    `gorm:"column:notice_title" json:"notice_title"`
	NoticeContent string    `gorm:"column:notice_content" json:"notice_content"`
	Type          string    `gorm:"column:type" json:"type"`
	Date          time.Time `gorm:"column:date" json:"date"`
}

type NotificationEditDTO struct {
	ID            int64  `gorm:"column:id" json:"id" binding:"required"`
	NoticeId      string `gorm:"column:notice_id" json:"notice_id"`
	NoticeTitle   string `gorm:"column:notice_title" json:"notice_title"`
	NoticeContent string `gorm:"column:notice_content" json:"notice_content"`
	Type          string `gorm:"column:type" json:"type"`
	Date          string `gorm:"column:date" json:"date"`
}

type NotificationDTO struct {
	NoticeTitle   string `gorm:"column:notice_title" json:"notice_title" binding:"required"`
	NoticeContent string `gorm:"column:notice_content" json:"notice_content" binding:"required"`
	Type          string `gorm:"column:type" json:"type" binding:"required"`
	Date          string `gorm:"column:date" json:"date" binding:"required"`
}
